package app

import "github.com/gin-gonic/gin"

func (h *Handlers) CallRouters(r *gin.RouterGroup) {

	userRouter := r.Group("users")
	userRouter.GET("", h.UserHandler.GetAllUsers)
	userRouter.GET("/:id", h.UserHandler.GetUserById)

	productRouter := r.Group("products")
	productRouter.GET("", h.ProductHandler.GetAllProducts)
	productRouter.POST("", h.ProductHandler.CreateProduct)
}
