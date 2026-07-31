package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

func Error(c *gin.Context, status int, code string, message string) {

	c.JSON(status, APIResponse{
		Success: false,
		Error: APIError{
			Code:    code,
			Message: message,
		},
	})
}

func BadRequest(c *gin.Context, message string) {

	Error(c, http.StatusBadRequest, "BAD_REQUEST", message)
}

func NotFound(c *gin.Context, message string) {

	Error(c, http.StatusNotFound, "NOT_FOUND", message)
}

func Conflict(c *gin.Context, message string) {

	Error(c, http.StatusConflict, "CONFLICT", message)
}

func Unauthorized(c *gin.Context, message string) {

	Error(c, http.StatusUnauthorized, "UNAUTHORIZED", message)
}

func InternalServerError(c *gin.Context) {

	Error(
		c,
		http.StatusInternalServerError,
		"INTERNAL_SERVER_ERROR",
		"Erro interno do servidor.",
	)
}
