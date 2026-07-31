package dto

import "time"

type ClienteRequest struct {
	Nome string `json:"nome" binding:"required"`

	CPF string `json:"cpf" binding:"required,len=11"`

	Email string `json:"email" binding:"required,email"`

	Telefone string `json:"telefone"`

	DataNascimento time.Time `json:"dataNascimento"`

	Endereco string `json:"endereco"`

	Numero string `json:"numero"`

	Complemento string `json:"complemento"`

	Bairro string `json:"bairro"`

	Cidade string `json:"cidade"`

	Estado string `json:"estado"`

	CEP string `json:"cep"`
}
