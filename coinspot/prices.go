package coinspot

type CoinPrice struct {
	Bid  string `json:"bid"`
	Ask  string `json:"ask"`
	Last string `json:"last"`
}

type LatestPricesResponse struct {
	Status string                `json:"status"`
	Prices map[string]*CoinPrice `json:"prices"`
}
