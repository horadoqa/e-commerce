package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, status int, message string, data interface{}) {

	c.JSON(status, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func OK(c *gin.Context, data interface{}) {

	Success(c, http.StatusOK, "Operação realizada com sucesso.", data)
}

func Created(c *gin.Context, data interface{}) {

	Success(c, http.StatusCreated, "Registro criado com sucesso.", data)
}

func NoContent(c *gin.Context) {

	c.Status(http.StatusNoContent)
}
