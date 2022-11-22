package handlers

type TransactionRequest struct {
	Amount    uint   `json:"amount" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
}

type LocationSetRequest struct {
	Location string `json:"city" binding:"required"`
}
