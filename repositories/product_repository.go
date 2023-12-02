package repositories

import (
	"context"
	"test-api/entities/models"
	"test-api/enums"
	"test-api/interfaces"
	"test-api/utils"

	"gorm.io/gorm"
)

type productRepository struct {
}

// Update implements interfaces.ProductRepository.
func (*productRepository) Update(ctx context.Context, tx *gorm.DB, entity *models.Product) (*models.Product, error) {
	return utils.SaveQuery[models.Product](ctx, tx, entity, enums.Update)
}

// GetById implements interfaces.ProductRepository.
func (*productRepository) GetById(ctx context.Context, tx *gorm.DB, id uint) (*models.Product, error) {
	return utils.GetById[models.Product](ctx, tx, id)
}

// Find implements interfaces.productRepository.
func (repo *productRepository) Find(ctx context.Context, tx *gorm.DB) ([]models.Product, error) {
	return utils.SelectQuery[models.Product](ctx, tx)
}

// Save implements interfaces.productRepository.
func (repo *productRepository) Save(ctx context.Context, tx *gorm.DB, product *models.Product) (*models.Product, error) {
	return utils.SaveQuery[models.Product](ctx, tx, product, enums.Create)
}

var _ interfaces.ProductRepository = &productRepository{}

func NewProductRepository() *productRepository {
	return &productRepository{}
}
