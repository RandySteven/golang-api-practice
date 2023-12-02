package req

import "github.com/shopspring/decimal"

type ProductRequest struct {
	Name       string          `json:"name"`
	Price      decimal.Decimal `json:"price"`
	Stock      uint            `json:"stock"`
	CategoryID uint            `json:"category_id"`
}
