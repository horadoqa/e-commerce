package routes

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/horadoqa/ecommerce-api/internal/handler"
)

func SetupRoutes(
	router *gin.Engine,
	clienteHandler *handler.ClienteHandler,
) {

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas da aplicação
	ClienteRoutes(router, clienteHandler)
}
