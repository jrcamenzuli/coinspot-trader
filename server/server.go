package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/common"
	log "github.com/sirupsen/logrus"
)

var (
	snapshotsMutex sync.Mutex
	snapshots      []*common.Snapshot
	api            CoinspotApi
	coinNames      []string
)

func Start() {
	coinNames = []string{"BTC", "ETH"}

	snapshots = []*common.Snapshot{}

	// Read key and secret from environment variables
	key := os.Getenv("COINSPOT_KEY")
	secret := os.Getenv("COINSPOT_SECRET")

	// Check if key and secret are not empty
	if key == "" || secret == "" {
		log.Fatal("COINSPOT_KEY and COINSPOT_SECRET environment variables are not set")
	}

	api = NewCoinSpotApi(key, secret)

	http.HandleFunc("/query", inboundQuery)
	go func() { http.ListenAndServe(":8080", nil) }()

	ticker := time.NewTicker(common.SnapshotInterval)
	defer ticker.Stop()
	for {
		<-ticker.C
		err := loop()
		if err != nil {
			log.Error(err)
		}
	}
}

func getLastNSnapshots(snapshots []*common.Snapshot, n int) []*common.Snapshot {
	if n >= len(snapshots) {
		// return a copy of the entire slice
		return append([]*common.Snapshot{}, snapshots...)
	} else {
		// return a copy of the last n elements
		return append([]*common.Snapshot{}, snapshots[len(snapshots)-n:]...)
	}
}

func getSnapshotsFromTime(snapshots []*common.Snapshot, fromTime time.Time) []*common.Snapshot {
	for i, snapshot := range snapshots {
		if snapshot.Time.After(fromTime) || snapshot.Time.Equal(fromTime) {
			return snapshots[i:]
		}
	}
	// Return an empty slice if no snapshots are after the given fromTime
	return []*common.Snapshot{}
}

func inboundQuery(w http.ResponseWriter, r *http.Request) {
	defer snapshotsMutex.Unlock()

	t := r.FormValue("t")
	intT, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		intT = 0
	}
	fromTime := time.Unix(intT, 0)

	snapshotsMutex.Lock()

	snapshots := getSnapshotsFromTime(snapshots, fromTime)

	jsonBytes, err := json.MarshalIndent(snapshots, "", "  ")
	if err != nil {
		log.Error("Error converting snapshots to JSON:", err)
		return
	}

	fmt.Fprint(w, string(jsonBytes))

	log.Infof("Something requested snapshots from time %s and got %d in return", fromTime, len(snapshots))
}

func appendSnapshot(snapshot *common.Snapshot) {
	defer snapshotsMutex.Unlock()
	snapshotsMutex.Lock()

	i := 0
	fromTime := time.Now().UTC()
	fromTime = fromTime.Add(-common.SnapshotWindowSize)
	for ; i < len(snapshots); i++ {
		if snapshots[i].Time.After(fromTime) {
			break
		}
	}

	snapshots = snapshots[i:]
	snapshots = append(snapshots, snapshot)
}

func makeSnapshot() (*common.Snapshot, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	snapshot := common.Snapshot{
		Time:  time.Now().UTC(),
		Coins: make(map[string]common.Coin),
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
		for key, value := range resp.Balances {
			snapshot.Wallet[key] = common.Wallet(value)
		}
	}()

	// get coin prices
	errCh2 := make(chan error, 1)
	go func() {
		defer wg.Done()
		for _, coinName := range coinNames {
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
			snapshot.Coins[coinName] = common.Coin{Rate: rate}
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

func loop() error {
	snapshot, err := makeSnapshot()
	if err != nil {
		return err
	}
	appendSnapshot(snapshot)
	log.Infof("There are %d snapshots in the sliding window of size %s.", len(snapshots), common.SnapshotWindowSize)
	snapshots := make([]common.Snapshot, len(snapshots))
	for i, p := range snapshots {
		snapshots[i] = p
	}
	if err != nil {
		return err
	}
	return nil
}
