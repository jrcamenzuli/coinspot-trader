package trader

import log "github.com/sirupsen/logrus"

type Strategy1 struct{}

func (s *Strategy1) Run(snapshots []Snapshot) error {
	log.Infof("%v", snapshots)
	return nil
}
