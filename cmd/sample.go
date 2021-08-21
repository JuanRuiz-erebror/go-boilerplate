package cmd

import (
	"fmt"
	"goprueba/Services/sample/controller"
	_ "goprueba/docs" // docs is generated by Swag CLI, you have to import it.
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//var ctx = context.TODO()

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func Sample() {

	r := gin.Default()
	port := os.Getenv("MAIN_PORT")
	fmt.Printf("port main: %v\n", port)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")

		accounts.GET(":id", c.ShowAccount)
		accounts.GET("", c.ListAccounts)
		accounts.POST("", c.AddAccount)
		accounts.DELETE(":id", c.DeleteAccount)
		accounts.PATCH(":id", c.UpdateAccount)
		accounts.POST(":id/images", c.UploadAccountImage)

		bras := v1.Group("/bras")
		bras.GET("", c.ListBras)

		bottles := v1.Group("/bottles")
		//bottles.Use(ginoauth2.Auth(zalando.UidCheck(USERS), zalando.OAuth2Endpoint))

		bottles.GET(":id", c.ShowBottle)
		bottles.GET("", c.ListBottles)
		bottles.GET("bb", c.ListBras)

		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Auth)
		}
		examples := v1.Group("/examples")
		{
			examples.GET("ping", c.PingExample)
			examples.GET("calc", c.CalcExample)
			examples.GET("groups/:group_id/accounts/:account_id", c.PathParamsExample)
			examples.GET("header", c.HeaderExample)
			examples.GET("securities", c.SecuritiesExample)
			examples.GET("attribute", c.AttributeExample)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + port)

}
