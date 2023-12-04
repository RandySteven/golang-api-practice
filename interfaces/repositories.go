package interfaces

import (
	"context"
	"test-api/entities/models"

	"gorm.io/gorm"
)

type (
	Repository[T any] interface {
		Save(ctx context.Context, tx *gorm.DB, entity *T) (*T, error)
		Find(ctx context.Context, tx *gorm.DB) ([]T, error)
		GetById(ctx context.Context, tx *gorm.DB, id uint) (*T, error)
		Update(ctx context.Context, tx *gorm.DB, entity *T) (*T, error)
	}

	UserRepository interface {
		Repository[models.User]
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (*models.User, error)
	}

	ProductRepository interface {
		Repository[models.Product]
	}

	TransactionRepository interface {
		Repository[models.Transaction]
	}

	TransactionDetailRepository interface {
		Repository[models.TransactionDetail]
		GetTransactionsBasedTransactionID(ctx context.Context, tx *gorm.DB, transactionID uint) ([]models.TransactionDetail, error)
	}
)
