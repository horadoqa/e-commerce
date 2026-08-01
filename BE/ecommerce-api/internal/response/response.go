package response

import "github.com/horadoqa/ecommerce-api/internal/models"

type APIResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Data    models.Cliente `json:"data"`
	Error   interface{}    `json:"error,omitempty"`
}
