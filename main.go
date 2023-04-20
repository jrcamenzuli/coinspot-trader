package main

import (
	"flag"
	"fmt"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	"github.com/jrcamenzuli/coinspot-trader/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	modePtr := flag.String("mode", "", "mode: coinspot or web")
	flag.Parse()

	switch *modePtr {
	case "coinspot":
		log.Infof("Starting Coinspot middle man...")
		coinspot.Start()
	case "web":
		log.Infof("Starting Web Server...")
		web.Start()
	default:
		fmt.Println("Invalid mode:", *modePtr)
		flag.PrintDefaults()
	}
}
