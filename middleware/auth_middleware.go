package middleware

import (
	"net/http"
	"strings"
	"test-api/auth"
	"test-api/entities/payload/res"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func validateToken(c *gin.Context) *auth.JWTClaim {
	tokenAuthorization := c.GetHeader("Authorization")
	tokenStr := strings.Trim(tokenAuthorization, "Bearer ")
	claims := &auth.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(j *jwt.Token) (interface{}, error) {
		return auth.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		return nil
	}
	return claims
}

func AuthMiddleware(c *gin.Context) {
	claims := validateToken(c)
	if claims == nil {
		resp := res.Response{
			Errors: []string{
				"Unauthorized",
			},
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
	}

	c.Set("x-user-id", claims.ID)
	c.Set("x-user-name", claims.Name)
	c.Set("x-user-email", claims.Email)
	c.Next()
}
