package client

import (
	"sync"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/common"
	log "github.com/sirupsen/logrus"
)

func Start() {
	var wg sync.WaitGroup
	channelSnapshots := make(chan common.Snapshot)

	wg.Add(1)
	go startCoinspotQueryClient(&wg, channelSnapshots)
	wg.Add(1)
	go startProcessor(&wg, channelSnapshots)

	wg.Wait()
}

func startProcessor(wg *sync.WaitGroup, channelSnapshots chan common.Snapshot) {
	log.Info("Processor started.")
	defer wg.Done()
	defer log.Info("Advisor stopped.")
	var snapshots []*common.Snapshot

	for channelSnapshot := range channelSnapshots {
		log.Debugf("Added snapshot to sliding window: %+v", channelSnapshot)
		snapshots = append(snapshots, &channelSnapshot)
		snapshots = cleanupSnapshots(snapshots)
	}
}

func cleanupSnapshots(snapshots []*common.Snapshot) []*common.Snapshot {
	// Get the current time
	now := time.Now()

	// Calculate the threshold time
	thresholdTime := now.Add(-common.WindowSize)

	// Iterate over the snapshots in reverse order
	for i := len(snapshots) - 1; i >= 0; i-- {
		// If the snapshot's time is before the threshold time,
		// remove it from the list of snapshots
		if snapshots[i].Time.Before(thresholdTime) {
			snapshots = append(snapshots[:i], snapshots[i+1:]...)
		} else {
			// Since the snapshots are ordered by Time in ascending order,
			// we can stop iterating as soon as we encounter a snapshot
			// that is not before the threshold time
			break
		}
	}

	return snapshots
}
