package server

type BuyOrder struct {
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
	Total  float64 `json:"total"`
	Coin   string  `json:"coin"`
	Market string  `json:"market"`
}

type SellOrder struct {
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
	Total  float64 `json:"total"`
	Coin   string  `json:"coin"`
	Market string  `json:"market"`
}

type listOpenOrdersRequest struct {
	Nonce    string `json:"nonce"`
	CoinType string `json:"cointype"`
}

type ListOpenOrdersResponse struct {
	Status     string       `json:"status"`
	BuyOrders  []*BuyOrder  `json:"buyorders"`
	SellOrders []*SellOrder `json:"sellorders"`
	Message    string       `json:"message"`
}

type HistoricalOrder struct {
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
	Total  float64 `json:"total"`
	Coin   string  `json:"coin"`
	Market string  `json:"market"`
}

type listOrderHistoryRequest struct {
	Nonce    string `json:"nonce"`
	CoinType string `json:"cointype"`
}

type ListOrderHistoryResponse struct {
	Status  string             `json:"status"`
	Orders  []*HistoricalOrder `json:"orders"`
	Message string             `json:"message"`
}

type depositCoinsRequest struct {
	Nonce    string `json:"nonce"`
	CoinType string `json:"cointype"`
}

type DepositCoinsResponse struct {
	Status  string `json:"status"`
	Address string `json:"address"`
	Message string `json:"message"`
}

type quickBuyQuoteRequest struct {
	Nonce    string  `json:"nonce"`
	CoinType string  `json:"cointype"`
	Amount   float64 `json:"amount"`
}

type QuickBuyQuoteResponse struct {
	Status    string  `json:"status"`
	Quote     float64 `json:"quote"`
	TimeFrame float64 `json:"timeframe"`
	Message   string  `json:"message"`
}

type quickSellQuoteRequest struct {
	Nonce    string  `json:"nonce"`
	CoinType string  `json:"cointype"`
	Amount   float64 `json:"amount"`
}

type QuickSellQuoteResponse struct {
	Status    string  `json:"status"`
	Quote     float64 `json:"quote"`
	TimeFrame float64 `json:"timeframe"`
	Message   string  `json:"message"`
}

type listBalancesRequest struct {
	Nonce string `json:"nonce"`
}

type BalanceResponse struct {
	Balance    float64 `json:"balance"`
	AudBalance float64 `json:"audbalance"`
	Rate       float64 `json:"rate"`
}

type ListBalancesResponse1 struct {
	Status   string                       `json:"status"`
	Message  string                       `json:"message"`
	Balances []map[string]BalanceResponse `json:"balances"`
}

type ListBalancesResponse struct {
	Status   string                     `json:"status"`
	Message  string                     `json:"message"`
	Balances map[string]BalanceResponse `json:"balances"`
}

type listMyOrdersRequest struct {
	Nonce string `json:"nonce"`
}

type ListMyOrdersResponse struct {
	Status     string       `json:"status"`
	BuyOrders  []*BuyOrder  `json:"buyorders"`
	SellOrders []*SellOrder `json:"sellorders"`
	Message    string       `json:"message"`
}

type placeBuyOrderRequest struct {
	Nonce    string  `json:"nonce"`
	CoinType string  `json:"cointype"`
	Amount   float64 `json:"amount"`
	Rate     float64 `json:"rate"`
}

type PlaceBuyOrderResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type placeSellOrderRequest struct {
	Nonce    string  `json:"nonce"`
	CoinType string  `json:"cointype"`
	Amount   float64 `json:"amount"`
	Rate     float64 `json:"rate"`
}

type PlaceSellOrderResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type cancelBuyOrderRequest struct {
	Nonce string `json:"nonce"`
	Id    string `json:"id"`
}

type CancelBuyOrderResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type cancelSellOrderRequest struct {
	Nonce string `json:"nonce"`
	Id    string `json:"id"`
}

type CancelSellOrderResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
