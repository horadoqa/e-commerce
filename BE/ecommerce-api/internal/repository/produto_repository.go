package repository

import (
	"gorm.io/gorm"

	"github.com/horadoqa/ecommerce-api/internal/models"
)

type ProdutoRepository struct {
	DB *gorm.DB
}

func (r *ProdutoRepository) FindAll() ([]models.Produto, error) {

	var produtos []models.Produto

	err := r.DB.Find(&produtos).Error

	return produtos, err
}
