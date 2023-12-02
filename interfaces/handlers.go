package interfaces

import "github.com/gin-gonic/gin"

type (
	ProductHandler interface {
		CreateProduct(c *gin.Context)
		GetAllProducts(c *gin.Context)
	}

	UserHandler interface {
		RegisterUser(c *gin.Context)
		LoginUser(c *gin.Context)
		GetAllUsers(c *gin.Context)
		GetUserById(c *gin.Context)
	}
)
