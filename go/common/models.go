package common

import "time"

type Coin struct {
	Rate float64
}

type Wallet struct {
	Balance    float64
	AudBalance float64
	Rate       float64
}

type Snapshot struct {
	Time   time.Time
	Coins  map[string]Coin
	Wallet map[string]Wallet
}
