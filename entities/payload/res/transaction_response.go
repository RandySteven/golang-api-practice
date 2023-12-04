package res

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionResponse struct {
	ID              uint                        `json:"id"`
	UserID          uint                        `json:"user_id"`
	TransactionDate time.Time                   `json:"transaction_date"`
	Details         []TransactionDetailResponse `json:"transaction_details"`
}

type TransactionDetailResponse struct {
	ProductID    uint            `json:"product_id"`
	ProductName  string          `json:"product_name"`
	ProductPrice decimal.Decimal `json:"product_price"`
	Quantity     uint            `json:"quantity"`
}
