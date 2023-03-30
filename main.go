package main

import (
	"os"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	"github.com/jrcamenzuli/coinspot-trader/trader"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("started coinspot-trader")

	// Read key and secret from environment variables
	key := os.Getenv("COINSPOT_KEY")
	secret := os.Getenv("COINSPOT_SECRET")

	// Check if key and secret are not empty
	if key == "" || secret == "" {
		log.Fatal("COINSPOT_KEY and COINSPOT_SECRET environment variables are not set")
	}

	api := coinspot.NewCoinSpotApi(key, secret)
	trader := new(trader.Trader)
	trader.Api = api

	trader.Run()
}
