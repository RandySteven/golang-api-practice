package utils

import (
	"test-api/entities/models"
	"test-api/entities/payload/res"
)

func TransactionDetailResponse(model *models.TransactionDetail) (res *res.TransactionDetailResponse) {
	res.ProductID = model.ProductID
	res.ProductName = model.Product.Name
	res.ProductPrice = model.Product.Price
	res.Quantity = model.Quantity
	return
}
