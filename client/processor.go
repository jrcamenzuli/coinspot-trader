package client

import (
	"math"
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

type Slope struct {
	age   time.Duration
	slope float64 // AUD/s
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
		// remove snapshots older than 24 hours
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
		maxAge := ages[0]
		for _, age := range ages {
			if age > maxAge {
				maxAge = age
			}
		}
		symbol := "BTC"
		slopes := []Slope{}
		for _, age := range ages {
			s := filterByAge(snapshots, age)
			log.Debugf("There are %d snapshots in %+v.", len(s), age)
			if isInvalid(s, age) {
				log.Infof("Time window %+v is invalid", age)
				slopes = append(slopes, Slope{age: age, slope: 0.0})
				continue
			}
			slope := averageSlope(s, symbol)
			slopes = append(slopes, Slope{age: age, slope: slope})
			// log.Infof("The %s rate is changing at %f (AUD/s) over the last %+v", symbol, slope, age)
		}
		chance := 0.0
		sumOfAges := 0.0
		for _, slope := range slopes {
			if slope.slope > 0.0 {
				chance += slope.age.Seconds()
			}
			sumOfAges += slope.age.Seconds()
		}
		chance /= sumOfAges
		coin, ok := snapshots[len(snapshots)-1].Coins[symbol]
		currentRate := math.NaN()
		if ok {
			currentRate = coin.Rate
		}
		log.Infof("The chance the rate for %s will increase over the next %+v is %.0f%% and is now %f", symbol, maxAge, chance*100, currentRate)
	}
}
