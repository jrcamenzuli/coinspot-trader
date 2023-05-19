package main

import (
	"flag"
	"fmt"

	"github.com/jrcamenzuli/coinspot-trader/client"
	"github.com/jrcamenzuli/coinspot-trader/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	modePtr := flag.String("mode", "", "mode: coinspot or web")
	flag.Parse()

	switch *modePtr {
	case "coinspot":
		log.Infof("Starting Server...")
		server.Start()
	case "processor":
		log.Infof("Starting Client...")
		client.Start()
	default:
		fmt.Println("Invalid mode:", *modePtr)
		flag.PrintDefaults()
	}
}
