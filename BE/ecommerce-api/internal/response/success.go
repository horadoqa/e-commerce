package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool `json:"success"`

	Message string `json:"message"`

	Data interface{} `json:"data"`
}

func OK(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, SuccessResponse{

		Success: true,

		Message: "Operação realizada com sucesso.",

		Data: data,
	})
}

func Created(c *gin.Context, data interface{}) {

	c.JSON(http.StatusCreated, SuccessResponse{

		Success: true,

		Message: "Registro criado com sucesso.",

		Data: data,
	})
}

func NoContent(c *gin.Context) {

	c.Status(http.StatusNoContent)
}
