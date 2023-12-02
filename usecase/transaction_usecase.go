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

type transactionUsecase struct {
	userRepo              interfaces.UserRepository
	productRepo           interfaces.ProductRepository
	transactionRepo       interfaces.TransactionRepository
	transactionDetailRepo interfaces.TransactionDetailRepository
	db                    *gorm.DB
}

// PurchaseProduct implements interfaces.TransactionUsecase.
func (usecase *transactionUsecase) PurchaseProduct(ctx context.Context, req *req.TransactionRequest) (*res.TransactionResponse, error) {
	tx := usecase.db.Begin()
	var res *res.TransactionResponse

	user, err := usecase.userRepo.GetById(ctx, tx, req.UserID)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		UserID: user.ID,
	}

	transaction, err = usecase.transactionRepo.Save(ctx, tx, transaction)

	for _, request := range req.TransactionDetailsRequests {
		product, err := usecase.productRepo.GetById(ctx, tx, request.ProductID)
		if err != nil {
			return nil, err
		}

		if product.Stock == 0 || product.Stock < request.Quantity {
			return nil, err
		}

		product.Stock -= request.Quantity
		product, err = usecase.productRepo.Update(ctx, tx, product)
		if err != nil {
			return nil, err
		}

		transactionDetail := &models.TransactionDetail{
			TransactionID: transaction.ID,
			ProductID:     product.ID,
			Quantity:      request.Quantity,
		}

		transactionDetail, err = usecase.transactionDetailRepo.Save(ctx, tx, transactionDetail)
		if err != nil {
			return nil, err
		}
	}

	tx.Commit()
	return res, nil
}

func NewTransactionUsecase(db *gorm.DB) *transactionUsecase {
	return &transactionUsecase{
		userRepo:              repositories.NewUserRepository(),
		productRepo:           repositories.NewProductRepository(),
		transactionRepo:       repositories.NewTransactionRepository(),
		transactionDetailRepo: repositories.NewTransactionDetailRepository(),
		db:                    db,
	}
}

var _ interfaces.TransactionUsecase = &transactionUsecase{}
