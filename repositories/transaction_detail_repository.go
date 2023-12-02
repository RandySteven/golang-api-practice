package repositories

import (
	"context"
	"test-api/entities/models"
	"test-api/enums"
	"test-api/interfaces"
	"test-api/utils"

	"gorm.io/gorm"
)

type transactionDetailRepository struct{}

// Find implements interfaces.TransactionDetailRepository.
func (*transactionDetailRepository) Find(ctx context.Context, tx *gorm.DB) ([]models.TransactionDetail, error) {
	panic("unimplemented")
}

// GetById implements interfaces.TransactionDetailRepository.
func (*transactionDetailRepository) GetById(ctx context.Context, tx *gorm.DB, id uint) (*models.TransactionDetail, error) {
	panic("unimplemented")
}

// Save implements interfaces.TransactionDetailRepository.
func (*transactionDetailRepository) Save(ctx context.Context, tx *gorm.DB, entity *models.TransactionDetail) (*models.TransactionDetail, error) {
	return utils.SaveQuery[models.TransactionDetail](ctx, tx, entity, enums.Create)
}

// Update implements interfaces.TransactionDetailRepository.
func (*transactionDetailRepository) Update(ctx context.Context, tx *gorm.DB, entity *models.TransactionDetail) (*models.TransactionDetail, error) {
	panic("unimplemented")
}

func NewTransactionDetailRepository() *transactionDetailRepository {
	return &transactionDetailRepository{}
}

var _ interfaces.TransactionDetailRepository = &transactionDetailRepository{}
