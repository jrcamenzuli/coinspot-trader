package main

import (
	"flag"
	"fmt"

	"github.com/jrcamenzuli/coinspot-trader/publisher"
	"github.com/jrcamenzuli/coinspot-trader/subscriber"
	log "github.com/sirupsen/logrus"
)

func main() {
	modePtr := flag.String("mode", "", "mode: publisher or subscriber")
	flag.Parse()

	switch *modePtr {
	case "publisher":
		log.Infof("Starting publisher...")
		publisher.Start()
	case "subscriber":
		log.Infof("Starting subscriber...")
		subscriber.Start()
	default:
		fmt.Println("Invalid mode:", *modePtr)
		flag.PrintDefaults()
	}
}
