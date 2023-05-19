package client

import (
	"sync"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/common"
	log "github.com/sirupsen/logrus"
)

func startProcessor(wg *sync.WaitGroup, channelSnapshots chan common.Snapshot) {
	log.Info("Processor started.")
	defer wg.Done()
	defer log.Info("Advisor stopped.")
	var snapshots []common.Snapshot

	for channelSnapshot := range channelSnapshots {
		log.Debugf("Added snapshot to sliding window: %+v", channelSnapshot)
		snapshots = append(snapshots, channelSnapshot)
		snapshots = cleanupSnapshots(snapshots)
	}
}

func cleanupSnapshots(snapshots []common.Snapshot) []common.Snapshot {
	ret := make([]common.Snapshot, 0)

	for _, s := range snapshots {
		if s.Time.After(time.Now().Add(-common.WindowSize)) {
			ret = append(ret, s)
		}
	}

	return ret
}
