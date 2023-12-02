package app

import (
	"test-api/configs"
	"test-api/handlers"
	"test-api/interfaces"
)

type Handlers struct {
	interfaces.UserHandler
	interfaces.ProductHandler
}

func NewHandlers(repo configs.Repositories) (*Handlers, error) {

	return &Handlers{
		UserHandler:    handlers.NewUserHandler(repo.UserUsecase),
		ProductHandler: handlers.NewProductHandler(repo.ProductUsecase),
	}, nil
}
