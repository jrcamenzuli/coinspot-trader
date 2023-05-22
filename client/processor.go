package client

import (
	"sync"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/common"
	"github.com/jrcamenzuli/coinspot-trader/utils"
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
	defer log.Info("Processor stopped.")
	snapshots := []common.Snapshot{}

	// preload snapshots
	tLast := time.Now()
	firstTime := true
	for channelSnapshot := range channelSnapshots {
		snapshots = append(snapshots, channelSnapshot)
		now := time.Now()
		if firstTime {
			tLast = now
			firstTime = false
		}
		if now.Sub(tLast) > 1*time.Second {
			break
		}
		tLast = now
	}

	log.Infof("Preloaded %d snapshots.", len(snapshots))

	for channelSnapshot := range channelSnapshots {
		snapshots = append(snapshots, channelSnapshot)
		snapshots = filterByAge(snapshots, 24*time.Hour)

		ages := []time.Duration{
			24 * time.Hour,
			12 * time.Hour,
			6 * time.Hour,
			1 * time.Hour,
			30 * time.Minute,
			10 * time.Minute,
			5 * time.Minute,
			1 * time.Minute}
		for _, age := range ages {
			s := filterByAge(snapshots, age)
			log.Debugf("There are %d snapshots in %+v.", len(s), age)
			slope := averageSlope(s, "BTC")
			log.Infof("The BTC rate is changing at %f (AUD/s) over the last %+v", slope, age)
		}
	}
}

func filterByAge(snapshots []common.Snapshot, age time.Duration) []common.Snapshot {
	now := time.Now().UTC()
	keep := 0
	threshold := now.Add(-age)
	for i := len(snapshots) - 1; i >= 0; i-- {
		if snapshots[i].Time.After(threshold) {
			keep++
		} else {
			break
		}
	}
	if keep < len(snapshots) {
		snapshots = snapshots[len(snapshots)-keep:]
	}
	return snapshots
}

func averageSlope(snapshots []common.Snapshot, symbol string) float64 {
	var points []utils.Point
	for _, snapshot := range snapshots {
		points = append(points, utils.Point{X: float64(snapshot.Time.UnixNano()) / 1e9, Y: snapshot.Coins[symbol].Rate})
	}
	slope := utils.AverageSlope(points)
	return slope
}
