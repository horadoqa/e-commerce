package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Error(c *gin.Context, status int, message string) {

	c.JSON(status, ErrorResponse{
		Success: false,
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {

	Error(
		c,
		http.StatusBadRequest,
		message,
	)
}

func NotFound(c *gin.Context, message string) {

	Error(
		c,
		http.StatusNotFound,
		message,
	)
}

func Conflict(c *gin.Context, message string) {

	Error(
		c,
		http.StatusConflict,
		message,
	)
}

func Unauthorized(c *gin.Context, message string) {

	Error(
		c,
		http.StatusUnauthorized,
		message,
	)
}

func InternalServerError(c *gin.Context) {

	Error(
		c,
		http.StatusInternalServerError,
		"Erro interno do servidor.",
	)
}
