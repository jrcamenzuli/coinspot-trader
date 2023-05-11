package processor

import (
	"sync"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	log "github.com/sirupsen/logrus"
)

func startAdvisor(wg *sync.WaitGroup, channelSnapshots chan coinspot.Snapshot) {
	log.Info("Processor started.")
	defer wg.Done()
	defer log.Info("Advisor stopped.")

	for channelSnapshot := range channelSnapshots {
		log.Infof("%+v\n", channelSnapshot.Wallet)
	}
}
