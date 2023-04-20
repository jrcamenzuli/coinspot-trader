package web

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	log "github.com/sirupsen/logrus"
)

func startWebServer(wg *sync.WaitGroup, channelSnapshots chan coinspot.Snapshot) {
	defer wg.Done()
	fs := http.FileServer(http.Dir("subscriber/frontend"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleWebSocket)

	log.Println("The web server is listening on :8081...")
	go handleWebsocketMessages()
	go handleCoinspotMessages(channelSnapshots)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleCoinspotMessages(channelSnapshots chan coinspot.Snapshot) {
	for channelSnapshot := range channelSnapshots {
		jsonBytes, err := json.Marshal(channelSnapshot)
		log.Info(channelSnapshot)
		if err != nil {
			log.Error("Error converting snapshots to JSON:", err)
			return
		}
		for conn := range connections {
			if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
				delete(connections, conn)
				log.Error(err)
				break
			}
		}
	}
}

func handleWebsocketMessages() {
	for {
		message := <-broadcast
		for conn := range connections {
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				delete(connections, conn)
				log.Println(err)
				break
			}
		}
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	connections[conn] = true

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(connections, conn)
			break
		}

		if messageType == websocket.TextMessage {
			log.Println("Received message: ", string(message))
			broadcast <- message
		} else {
			log.Println("Received non-text message")
		}
	}
}
