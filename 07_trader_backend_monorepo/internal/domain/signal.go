package domain

type Signal struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"` // ETHUSDT - EURUSD
	Side       string `json:"side"`   //  BUY / SELL
	Quantity   int    `json:"quantity"`
	StrategyID string `json:"strategy_id"`
}
