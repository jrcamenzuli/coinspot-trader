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

	// Initialize a counter for removed elements
	var removed int

	// Iterate over the snapshots in reverse order
	for i := len(snapshots) - 1; i >= 0; i-- {
		// If the snapshot's time is before the threshold time,
		if snapshots[i].Time.Before(thresholdTime) {
			// increment the counter and move on to the next snapshot
			removed++
			continue
		}
		// if the snapshot's time is not before the threshold time,
		// and some elements were already removed,
		if removed > 0 {
			// move the snapshot to its new index
			snapshots[i+removed] = snapshots[i]
		}
	}
	// If some elements were removed, resize the slice accordingly
	if removed > 0 {
		snapshots = snapshots[:len(snapshots)-removed]
	}

	return snapshots
}
