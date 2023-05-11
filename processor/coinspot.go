package processor

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	log "github.com/sirupsen/logrus"
)

func startCoinspotQueryClient(wg *sync.WaitGroup, channelSnapshots chan coinspot.Snapshot) {
	log.Info("Coinspot Query Client started.")
	defer wg.Done()
	defer log.Info("Coinspot Query Client stopped.")

	lastSnapshotTime := time.Unix(0, 0)

	for {
		response, err := http.Get("http://192.168.0.40:10000/query?t=" + strconv.Itoa(int(lastSnapshotTime.Unix())))
		if err != nil {
			log.Error("Error:", err)
			continue
		}

		var receivedSnapshots []*coinspot.Snapshot

		if err := json.NewDecoder(response.Body).Decode(&receivedSnapshots); err != nil {
			log.Error("Error:", err)
			continue
		}

		for i := 0; i < len(receivedSnapshots); i++ {
			if receivedSnapshots[i].Time.Before(lastSnapshotTime) || receivedSnapshots[i].Time.Equal(lastSnapshotTime) {
				continue
			}
			channelSnapshots <- *receivedSnapshots[i]
			lastSnapshotTime = receivedSnapshots[i].Time
		}

		response.Body.Close()

		time.Sleep(5 * time.Second)
	}
}
