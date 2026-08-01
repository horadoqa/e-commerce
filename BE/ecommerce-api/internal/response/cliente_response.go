package response

import (
	"time"

	"github.com/google/uuid"
)

type ClienteResponse struct {
	ID uuid.UUID `json:"id"`

	Nome string `json:"nome"`

	CPF string `json:"cpf"`

	Email string `json:"email"`

	Telefone string `json:"telefone"`

	DataNascimento time.Time `json:"dataNascimento"`

	Endereco string `json:"endereco"`

	Numero string `json:"numero"`

	Complemento string `json:"complemento"`

	Bairro string `json:"bairro"`

	Cidade string `json:"cidade"`

	Estado string `json:"estado"`

	CEP string `json:"cep"`

	Ativo bool `json:"ativo"`

	CreatedAt time.Time `json:"createdAt"`

	UpdatedAt time.Time `json:"updatedAt"`
}
