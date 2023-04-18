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
	time   time.Time
	coin   Coin
	wallet map[string]coinspot.BalanceResponse
}

func (t *Trader) addSnapshot(snapshot *Snapshot) {
	if len(t.snapshots) >= t.windowSize {
		t.snapshots = t.snapshots[1:]
	}
	t.snapshots = append(t.snapshots, snapshot)
}

type Trader struct {
	isAlreadyStarted bool
	api              coinspot.CoinspotApi
	windowSize       int
	snapshotInterval time.Duration
	snapshots        []*Snapshot
	coinName         string
	strategy         Strategy
}

func (t *Trader) Start(api coinspot.CoinspotApi, coinName string) {
	if t.isAlreadyStarted {
		log.Error("already initialized")
		return
	}
	t.isAlreadyStarted = true
	t.api = api
	t.snapshots = make([]*Snapshot, 0)
	t.coinName = coinName
	t.windowSize = 1000
	t.snapshotInterval = 1 * time.Second
	t.strategy = &Strategy1{}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		err := t.loop()
		if err != nil {
			log.Error(err)
		}
	}
}

func (t *Trader) getSnapshot() (*Snapshot, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	snapshot := Snapshot{
		time: time.Now().UTC(),
	}

	// get wallet
	errCh1 := make(chan error, 1)
	go func() {
		defer wg.Done()
		resp, err := t.api.ListBalances()
		if err != nil {
			errCh1 <- err
			return
		}
		snapshot.wallet = resp.Balances
	}()

	// get coin prices
	errCh2 := make(chan error, 1)
	go func() {
		defer wg.Done()
		resp, err := t.api.LatestCoinPrices(t.coinName)
		if err != nil {
			errCh2 <- err
			return
		}
		rate, err := strconv.ParseFloat(resp.Prices.Ask, 64)
		if err != nil {
			errCh2 <- err
			return
		}
		snapshot.coin = Coin{rate: rate}
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

func (t *Trader) loop() error {
	snapshot, err := t.getSnapshot()
	if err != nil {
		return err
	}
	t.addSnapshot(snapshot)
	log.Infof("There are %d snapshots in the sliding window of size %d.", len(t.snapshots), t.windowSize)
	snapshots := make([]Snapshot, len(t.snapshots))
	for i, p := range t.snapshots {
		snapshots[i] = *p
	}
	t.strategy.Run(snapshots)
	if err != nil {
		return err
	}
	return nil
}
