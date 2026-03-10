package strategies

type Strategy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateResponse struct {
	StrategyID string `json:"strategy_id"`
}

type GetByIDRequest struct {
	StrategyID string `json:"strategy_id"`
}

type GetByIDResponse struct {
	Strategy Strategy `json:"strategy"`
}
