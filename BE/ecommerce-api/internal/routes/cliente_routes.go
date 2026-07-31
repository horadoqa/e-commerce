package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/horadoqa/ecommerce-api/internal/handler"
)

func ClienteRoutes(
	router *gin.Engine,
	clienteHandler *handler.ClienteHandler,
) {

	clientes := router.Group("/clientes")
	{
		clientes.GET("", clienteHandler.Listar)
		clientes.GET("/:id", clienteHandler.Buscar)
		clientes.POST("", clienteHandler.Criar)
		clientes.PUT("/:id", clienteHandler.Atualizar)
		clientes.PATCH("/:id", clienteHandler.AtualizarParcial)
		clientes.DELETE("/:id", clienteHandler.Excluir)
	}
}