package res

import (
	"test-api/entities/models"

	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	Price    decimal.Decimal `json:"price"`
	Stock    uint            `json:"stock"`
	Category models.Category `json:"category"`
}
