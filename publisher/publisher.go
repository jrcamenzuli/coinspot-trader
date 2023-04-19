package publisher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	log "github.com/sirupsen/logrus"
)

// write a publisher that uses "cloud.google.com/go/pubsub" and publishes messages to ws://localhost:8080/updates
const coinName = "BTC"
const windowSize = 1000
const snapshotInterval = 5 * time.Second

var (
	snapshotsMutex sync.Mutex
	snapshots      []*Snapshot
	api            coinspot.CoinspotApi
)

type Coin struct {
	Rate float64
}

type Snapshot struct {
	Time   time.Time
	Coin   Coin
	Wallet map[string]coinspot.BalanceResponse
}

func Start() {
	snapshots = []*Snapshot{}

	// Read key and secret from environment variables
	key := os.Getenv("COINSPOT_KEY")
	secret := os.Getenv("COINSPOT_SECRET")

	// Check if key and secret are not empty
	if key == "" || secret == "" {
		log.Fatal("COINSPOT_KEY and COINSPOT_SECRET environment variables are not set")
	}

	api = coinspot.NewCoinSpotApi(key, secret)

	http.HandleFunc("/query", get)
	go func() { http.ListenAndServe(":8080", nil) }()

	ticker := time.NewTicker(snapshotInterval)
	defer ticker.Stop()
	for {
		<-ticker.C
		err := loop()
		if err != nil {
			log.Error(err)
		}
	}
}

func getLastNSnapshots(snapshots []*Snapshot, n int) []*Snapshot {
	if n >= len(snapshots) {
		// return a copy of the entire slice
		return append([]*Snapshot{}, snapshots...)
	} else {
		// return a copy of the last n elements
		return append([]*Snapshot{}, snapshots[len(snapshots)-n:]...)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	defer snapshotsMutex.Unlock()

	n := r.FormValue("n")
	intN, err := strconv.Atoi(n)
	if err != nil {
		intN = windowSize
	}

	snapshotsMutex.Lock()

	snapshots := getLastNSnapshots(snapshots, intN)

	jsonBytes, err := json.MarshalIndent(snapshots, "", "  ")
	if err != nil {
		log.Error("Error converting snapshots to JSON:", err)
		return
	}

	fmt.Fprint(w, string(jsonBytes))

	log.Infof("Something requested the latest %d snapshots and got %d in return", intN, len(snapshots))
}

func appendSnapshot(snapshot *Snapshot) {
	defer snapshotsMutex.Unlock()
	snapshotsMutex.Lock()
	if len(snapshots) >= windowSize {
		snapshots = snapshots[1:]
	}
	snapshots = append(snapshots, snapshot)
}

func makeSnapshot() (*Snapshot, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	snapshot := Snapshot{
		Time: time.Now().UTC(),
	}

	// get wallet
	errCh1 := make(chan error, 1)
	go func() {
		defer wg.Done()
		resp, err := api.ListBalances()
		if err != nil {
			errCh1 <- err
			return
		}
		snapshot.Wallet = resp.Balances
	}()

	// get coin prices
	errCh2 := make(chan error, 1)
	go func() {
		defer wg.Done()
		resp, err := api.LatestCoinPrices(coinName)
		if err != nil {
			errCh2 <- err
			return
		}
		rate, err := strconv.ParseFloat(resp.Prices.Ask, 64)
		if err != nil {
			errCh2 <- err
			return
		}
		snapshot.Coin = Coin{Rate: rate}
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

func loop() error {
	snapshot, err := makeSnapshot()
	if err != nil {
		return err
	}
	appendSnapshot(snapshot)
	log.Infof("There are %d snapshots in the sliding window of size %d.", len(snapshots), windowSize)
	snapshots := make([]Snapshot, len(snapshots))
	for i, p := range snapshots {
		snapshots[i] = p
	}
	if err != nil {
		return err
	}
	return nil
}
