package processor

import (
	"sync"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
)

func Start() {
	var wg sync.WaitGroup
	channelSnapshots := make(chan coinspot.Snapshot)

	wg.Add(1)
	go startCoinspotQueryClient(&wg, channelSnapshots)
	wg.Add(1)
	go startAdvisor(&wg, channelSnapshots)

	wg.Wait()
}
