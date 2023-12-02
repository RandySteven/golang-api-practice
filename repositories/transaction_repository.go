package repositories

import (
	"context"
	"test-api/entities/models"
	"test-api/enums"
	"test-api/interfaces"
	"test-api/utils"

	"gorm.io/gorm"
)

type transactionRepository struct {
}

// Find implements interfaces.TransactionRepository.
func (*transactionRepository) Find(ctx context.Context, tx *gorm.DB) ([]models.Transaction, error) {
	panic("unimplemented")
}

// GetById implements interfaces.TransactionRepository.
func (*transactionRepository) GetById(ctx context.Context, tx *gorm.DB, id uint) (*models.Transaction, error) {
	panic("unimplemented")
}

// Save implements interfaces.TransactionRepository.
func (*transactionRepository) Save(ctx context.Context, tx *gorm.DB, entity *models.Transaction) (*models.Transaction, error) {
	return utils.SaveQuery[models.Transaction](ctx, tx, entity, enums.Create)
}

// Update implements interfaces.TransactionRepository.
func (*transactionRepository) Update(ctx context.Context, tx *gorm.DB, entity *models.Transaction) (*models.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionRepository() *transactionRepository {
	return &transactionRepository{}
}

var _ interfaces.TransactionRepository = &transactionRepository{}
