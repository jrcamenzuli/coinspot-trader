package subscriber

import (
	"sync"

	"github.com/jrcamenzuli/coinspot-trader/publisher"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

var (
	connections = make(map[*websocket.Conn]bool)
	broadcast   = make(chan []byte)
	upgrader    = websocket.Upgrader{}
)

func Start() {
	var wg sync.WaitGroup
	channelSnapshots := make(chan publisher.Snapshot)

	wg.Add(1)
	go startWebServer(&wg, channelSnapshots)
	wg.Add(1)
	go startBroadcastClient(&wg, channelSnapshots)

	log.Info("Web server started.")
	log.Info("Query client started.")
	wg.Wait()
	log.Info("Web server stopped.")
	wg.Wait()
	log.Info("Query client stopped.")
}
