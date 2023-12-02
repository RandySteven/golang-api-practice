package usecase

import (
	"context"
	"test-api/entities/models"
	"test-api/entities/payload/req"
	"test-api/entities/payload/res"
	"test-api/interfaces"
	"test-api/repositories"

	"gorm.io/gorm"
)

type productUsecase struct {
	productRepository interfaces.ProductRepository
	db                *gorm.DB
}

// CreateProduct implements interfaces.ProductUsecase.
func (usecase *productUsecase) CreateProduct(ctx context.Context, request *req.ProductRequest) (*res.ProductResponse, error) {
	tx := usecase.db.Begin()
	product := &models.Product{
		Name:       request.Name,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: request.CategoryID,
	}
	product, err := usecase.productRepository.Save(ctx, tx, product)
	if err != nil {
		return nil, err
	}
	productResponse := &res.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}
	tx.Commit()
	return productResponse, nil
}

// GetAllProducts implements interfaces.ProductUsecase.
func (usecase *productUsecase) GetAllProducts(ctx context.Context) ([]res.ProductResponse, error) {
	panic("unimplemented")
}

func NewProductUsecase(db *gorm.DB) *productUsecase {
	return &productUsecase{
		productRepository: repositories.NewProductRepository(),
		db:                db,
	}
}

var _ interfaces.ProductUsecase = &productUsecase{}
