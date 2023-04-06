package trader

import (
	"time"

	"github.com/jrcamenzuli/coinspot-trader/utils"
	log "github.com/sirupsen/logrus"
)

type Strategy1 struct{}

func averageSlope(snapshots []Snapshot) float64 {
	points := make([]utils.Point, len(snapshots))
	for i, snapshot := range snapshots {
		points[i] = utils.Point{X: float64(snapshot.time.UnixNano() / int64(time.Second)), Y: snapshot.coin.rate}
	}
	slope := utils.AverageSlope(points)
	return slope
}

func (s *Strategy1) Run(snapshots []Snapshot) error {
	n := len(snapshots)
	if n >= 10 {
		slope := averageSlope(snapshots[n-10:])
		log.Infof("slope(10): $%f/s", slope)
	}
	if n >= 100 {
		slope := averageSlope(snapshots[n-100:])
		log.Infof("slope(100): $%f/s", slope)
	}
	if n >= 1000 {
		slope := averageSlope(snapshots[n-1000:])
		log.Infof("slope(1000): $%f/s", slope)
	}
	return nil
}
