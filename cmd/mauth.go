package cmd

import (
	"fmt"
	_ "goprueba/docs" // docs is generated by Swag CLI, you have to import it.

	"os"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//http -v --json POST localhost:8000/login username=admin password=admin

func Mauth() {

	port := os.Getenv("AUTH_PORT")

	fmt.Printf("port auth: %v\n", port)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong auth",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	go r.Run(":" + port)

}
