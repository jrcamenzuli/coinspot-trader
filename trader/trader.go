package trader

import (
	"time"

	"github.com/jrcamenzuli/coinspot-trader/coinspot"
	log "github.com/sirupsen/logrus"
)

type Trader struct {
	Api coinspot.CoinspotApi
}

func (trader *Trader) Run() {
	isFirstTime := true
	for {
		if !isFirstTime {
			time.Sleep(10 * time.Second)
		}
		isFirstTime = false

		//+
		balances, err := trader.Api.ListBalances()
		if err != nil {
			continue
		}
		for key, value := range balances.Balances {
			log.Infof("%s %+v", key, value)
		}
		//-

	}
}
