package main

import (
	"./docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	MATRIX "./controller"
)

func main() {

	//swagger info
	docs.SwaggerInfo.Title = "League Backend Challenge"
	docs.SwaggerInfo.Description = "Matrix API Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/echo", MATRIX.Echo)
		v1.POST("/invert", MATRIX.InvertMatrix)
		v1.POST("/flatten", MATRIX.FlattenMatrix)
		v1.POST("/sum", MATRIX.SumOfMatrix)
		v1.POST("/multiply", MATRIX.MultiplyMatrix)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
