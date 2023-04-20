package main

import (
	"flag"
	"fmt"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	log "github.com/sirupsen/logrus"
)

func main() {
	modePtr := flag.String("mode", "", "mode: publisher or subscriber")
	flag.Parse()

	switch *modePtr {
	case "publisher":
		log.Infof("Starting publisher...")
		coinspot.Start()
	case "subscriber":
		log.Infof("Starting subscriber...")
		coinspot.Start()
	default:
		fmt.Println("Invalid mode:", *modePtr)
		flag.PrintDefaults()
	}
}
