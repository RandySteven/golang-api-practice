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

// GetTransactionsBasedTransactionID implements interfaces.TransactionDetailRepository.
func (*transactionDetailRepository) GetTransactionsBasedTransactionID(ctx context.Context, tx *gorm.DB, transactionID uint) ([]models.TransactionDetail, error) {
	var details []models.TransactionDetail
	err := tx.WithContext(ctx).
		Model(&models.TransactionDetail{}).
		Where("transaction_id = ? ", transactionID).
		Find(&details).
		Error
	if err != nil {
		return nil, err
	}
	return details, nil
}

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
