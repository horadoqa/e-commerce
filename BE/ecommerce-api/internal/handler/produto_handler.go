package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/horadoqa/ecommerce-api/internal/service"
)

type ProdutoHandler struct {
	Service *service.ProdutoService
}

func (h *ProdutoHandler) Listar(c *gin.Context) {

	c.JSON(200, gin.H{
		"mensagem": "Produtos funcionando",
	})
}
