package domain

type Order struct {
	ID         string
	Symbol     string
	Side       string
	Quantity   int
	StrategyID string
	Price      float64
	Status     string
}
