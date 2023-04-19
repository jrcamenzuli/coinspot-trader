package subscriber

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jrcamenzuli/coinspot-trader/publisher"
)

func startQueryClient(wg *sync.WaitGroup, channelSnapshots chan publisher.Snapshot) {
	defer wg.Done()

	lastSnapshotTime := time.Unix(0, 0)

	for {
		response, err := http.Get("http://192.168.0.40:10000/query?n=1000")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		var receivedSnapshots []*publisher.Snapshot

		if err := json.NewDecoder(response.Body).Decode(&receivedSnapshots); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		for i := 0; i < len(receivedSnapshots); i++ {
			if receivedSnapshots[i].Time.Before(lastSnapshotTime) {
				continue
			}
			channelSnapshots <- *receivedSnapshots[i]
			lastSnapshotTime = receivedSnapshots[i].Time
		}

		response.Body.Close()

		time.Sleep(5 * time.Second)
	}
}
