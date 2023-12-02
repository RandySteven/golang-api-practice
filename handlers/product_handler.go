package handlers

import (
	"context"
	"net/http"
	"test-api/entities/payload/req"
	"test-api/entities/payload/res"
	"test-api/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	usecase interfaces.ProductUsecase
}

// CreateProduct implements interfaces.ProductHandler.
func (handler *ProductHandler) CreateProduct(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		request   *req.ProductRequest
	)
	if err := c.ShouldBind(&request); err != nil {
		return
	}
	productRes, err := handler.usecase.CreateProduct(ctx, request)
	if err != nil {
		return
	}
	resp := res.Response{
		Message: "Success created product",
		Data:    productRes,
	}
	c.JSON(http.StatusCreated, resp)
}

// GetAllProducts implements interfaces.ProductHandler.
func (handler *ProductHandler) GetAllProducts(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	)
	productsRes, err := handler.usecase.GetAllProducts(ctx)
	if err != nil {
		return
	}
	resp := res.Response{
		Message: "Success get all products",
		Data:    productsRes,
	}
	c.JSON(http.StatusOK, resp)
}

func NewProductHandler(usecase interfaces.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

var _ interfaces.ProductHandler = &ProductHandler{}
