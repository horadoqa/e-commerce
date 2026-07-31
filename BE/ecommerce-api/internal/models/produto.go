package models

import (
	"time"

	"github.com/google/uuid"
)

type Produto struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	SKU string `gorm:"column:sku"`

	Nome string `gorm:"column:nome"`

	Descricao string `gorm:"column:descricao"`

	Categoria string `gorm:"column:categoria"`

	Preco float64 `gorm:"column:preco"`

	Custo float64 `gorm:"column:custo"`

	QuantidadeEstoque int `gorm:"column:quantidade_estoque"`

	EstoqueMinimo int `gorm:"column:estoque_minimo"`

	Peso float64 `gorm:"column:peso"`

	Ativo bool `gorm:"column:ativo"`

	CreatedAt time.Time `gorm:"column:created_at"`

	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Produto) TableName() string {
	return "produtos"
}
