package req

type TransactionRequest struct {
	UserID                     uint                       `json:"user_id"`
	TransactionDetailsRequests []TransactionDetailRequest `json:"transaction_details"`
}

type TransactionDetailRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
