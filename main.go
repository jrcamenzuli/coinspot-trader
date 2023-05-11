package main

import (
	"flag"
	"fmt"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	"github.com/jrcamenzuli/coinspot-trader/processor"
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
		log.Infof("Starting Processor...")
		processor.Start()
	default:
		fmt.Println("Invalid mode:", *modePtr)
		flag.PrintDefaults()
	}
}
