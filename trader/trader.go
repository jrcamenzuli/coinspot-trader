package trader

import (
	"strconv"
	"sync"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	log "github.com/sirupsen/logrus"
)

type Coin struct {
	rate float64
}

type Snapshot struct {
	cryptos map[string]Coin
	wallet  map[string]coinspot.BalanceResponse
}

func (t *Trader) addSnapshot(snapshot *Snapshot) {
	if len(t.snapshots) >= t.windowSize {
		t.snapshots = t.snapshots[1:]
	}
	t.snapshots = append(t.snapshots, snapshot)
}

type Trader struct {
	isAlreadyStart bool
	api            coinspot.CoinspotApi
	windowSize     int
	snapshots      []*Snapshot
	tickers        []string
}

func (t *Trader) Start(api coinspot.CoinspotApi, tickers []string) {
	if t.isAlreadyStart {
		log.Error("already initialized")
		return
	}
	t.isAlreadyStart = true
	t.api = api
	t.snapshots = make([]*Snapshot, 0)
	t.tickers = tickers
	t.windowSize = 60

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		t.loop()
	}
}

func (trader *Trader) average(ticker string) float64 {
	average := 0.0
	for i := 0; i < len(trader.snapshots); i++ {
		average += trader.snapshots[i].cryptos[ticker].rate
	}
	average /= float64(len(trader.snapshots))
	return average
}

func (t *Trader) getSnapshot() (*Snapshot, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	snapshot := Snapshot{
		cryptos: make(map[string]Coin),
	}

	// get wallet
	errCh1 := make(chan error, 1)
	go func() {
		defer wg.Done()
		resp, err := t.api.ListBalances()
		snapshot.wallet = resp.Balances
		if err != nil {
			errCh1 <- err
			return
		}
	}()

	// get coin prices
	errCh2 := make(chan error, 1)
	go func() {
		defer wg.Done()
		for _, ticker := range t.tickers {
			resp, err := t.api.LatestCoinPrices(ticker)
			if err != nil {
				errCh2 <- err
				return
			}
			rate, err := strconv.ParseFloat(resp.Prices.Ask, 64)
			if err != nil {
				errCh2 <- err
				return
			}
			snapshot.cryptos[ticker] = Coin{rate: rate}
		}
	}()

	go func() {
		wg.Wait()
		close(errCh1)
		close(errCh2)
	}()

	for err := range errCh1 {
		return nil, err
	}

	for err := range errCh2 {
		return nil, err
	}

	return &snapshot, nil
}

func (t *Trader) loop() {
	time.Sleep(10 * time.Second)
	snapshot, err := t.getSnapshot()
	if err != nil {
		return
	}
	t.addSnapshot(snapshot)
	log.Infof("There are %d snapshots in the sliding window of size %d.", len(t.snapshots), t.windowSize)
	if len(t.snapshots) < t.windowSize {
		return
	}
}
