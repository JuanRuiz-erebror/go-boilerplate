package middlewares

import (
	"fmt"
	"goprueba/Services/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		tokenAuth, err := service.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		fmt.Printf("tokenAuth: %v\n", tokenAuth)
		userId, err := service.FetchAuth(tokenAuth)
		fmt.Printf("userId: %v\n", userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
