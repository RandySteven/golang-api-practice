package interfaces

import (
	"context"
	"test-api/entities/payload/req"
	"test-api/entities/payload/res"
)

type (
	UserUsecase interface {
		RegisterUser(ctx context.Context, register *req.UserRegister) (*res.UserResponse, error)
		LoginUser(ctx context.Context, login *req.UserLogin) (*res.UserLoginResponse, error)
		GetAllUsers(ctx context.Context) ([]res.UserResponse, error)
		GetUserById(ctx context.Context, id uint) (*res.UserResponse, error)
	}

	ProductUsecase interface {
		CreateProduct(ctx context.Context, request *req.ProductRequest) (*res.ProductResponse, error)
		GetAllProducts(ctx context.Context) ([]res.ProductResponse, error)
	}

	TransactionUsecase interface {
		PurchaseProduct(ctx context.Context, request *req.TransactionRequest) (*res.TransactionResponse, error)
	}
)
