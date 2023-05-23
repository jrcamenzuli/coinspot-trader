package client

import (
	"time"

	"github.com/jrcamenzuli/coinspot-trader/common"
)

func isInvalid(snapshots []common.Snapshot, expectedAge time.Duration) bool {
	if len(snapshots) <= 0 {
		return true
	}
	minTime := snapshots[0].Time
	maxTime := snapshots[0].Time
	for _, snapshot := range snapshots {
		if snapshot.Time.After(maxTime) {
			maxTime = snapshot.Time
		}
		if snapshot.Time.Before(minTime) {
			minTime = snapshot.Time
		}
	}
	actualAge := maxTime.Sub(minTime)
	lowerLimit := expectedAge - 10*time.Second
	upperLimit := expectedAge + 10*time.Second
	if (actualAge < lowerLimit) || (actualAge > upperLimit) {
		return true
	}
	return false
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
	var points []common.Point
	for _, snapshot := range snapshots {
		coin, ok := snapshot.Coins[symbol]
		if ok {
			// Coin exists, access the rate
			points = append(points, common.Point{X: float64(snapshot.Time.UnixNano()) / 1e9, Y: coin.Rate})
		}
	}
	slope := common.AverageSlope(points)
	return slope
}
