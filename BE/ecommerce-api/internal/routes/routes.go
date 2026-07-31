package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/horadoqa/ecommerce-api/internal/handler"
)

func SetupRoutes(
	router *gin.Engine,
	clienteHandler *handler.ClienteHandler,
) {

	ClienteRoutes(router, clienteHandler)

}
