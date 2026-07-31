package models

import (
	"time"

	"github.com/google/uuid"
)

type Cliente struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Nome string `gorm:"column:nome"`
	CPF  string `gorm:"column:cpf"`

	Email    string `gorm:"column:email"`
	Telefone string `gorm:"column:telefone"`

	DataNascimento time.Time `gorm:"column:data_nascimento"`

	Endereco    string `gorm:"column:endereco"`
	Numero      string `gorm:"column:numero"`
	Complemento string `gorm:"column:complemento"`

	Bairro string `gorm:"column:bairro"`
	Cidade string `gorm:"column:cidade"`
	Estado string `gorm:"column:estado"`
	CEP    string `gorm:"column:cep"`

	Ativo bool `gorm:"column:ativo"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Cliente) TableName() string {
	return "clientes"
}
