package trader

import (
	"strconv"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	log "github.com/sirupsen/logrus"
)

type Coin struct {
	rate float64
}

type Snapshot struct {
	cryptos map[string]Coin
}

func (t *Trader) addSnapshot(snapshot Snapshot) {
	if len(t.snapshots) >= t.windowSize {
		t.snapshots = t.snapshots[1:]
	}
	t.snapshots = append(t.snapshots, snapshot)
}

type Trader struct {
	isAlreadyStart bool
	api            coinspot.CoinspotApi
	windowSize     int
	snapshots      []Snapshot
	tickers        []string
}

func (t *Trader) Start(api coinspot.CoinspotApi, tickers []string) {
	if t.isAlreadyStart {
		log.Error("already initialised")
		return
	}
	t.isAlreadyStart = true
	t.api = api
	t.snapshots = make([]Snapshot, 0)
	t.tickers = tickers
	t.windowSize = 60

	t.run()
}

func (trader *Trader) average(ticker string) float64 {
	average := 0.0
	for i := 0; i < len(trader.snapshots); i++ {
		average += trader.snapshots[i].cryptos[ticker].rate
	}
	average /= float64(len(trader.snapshots))
	return average
}

func (t *Trader) run() {
	isFirstTime := true
	for {
		if !isFirstTime {
			time.Sleep(10 * time.Second)
		}
		isFirstTime = false

		//+
		balances, err := t.api.ListBalances()
		if err != nil {
			continue
		}
		for key, value := range balances.Balances {
			log.Infof("%s %+v", key, value)
		}
		//-

		snapshot := Snapshot{cryptos: make(map[string]Coin)}
		for _, ticker := range t.tickers {
			resp, err := t.api.LatestCoinPrices(ticker)
			if err != nil {
				continue
			}
			rate, err := strconv.ParseFloat(resp.Prices.Ask, 64)
			if err != nil {
				continue
			}
			snapshot.cryptos[ticker] = Coin{rate: rate}
		}
		t.addSnapshot(snapshot)
	}
}
