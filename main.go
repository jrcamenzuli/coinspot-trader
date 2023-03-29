package main

import (
	"os"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
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
	balances, err := api.ListBalances()
	if err != nil {
		panic(err)
	}
	log.Infof("%+v", balances)
	log.Infof("%+v", balances.Balances[2])
	log.Infof("%+v", balances.Balances[2]["DOGE"])
	log.Infof("%+v", balances.Balances[2]["DOGE"].AudBalance)
}
