package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/horadoqa/ecommerce-api/internal/models"
)

type ClienteRepository struct {
	DB *gorm.DB
}

func (r *ClienteRepository) FindAll() ([]models.Cliente, error) {

	var clientes []models.Cliente

	err := r.DB.Order("nome").Find(&clientes).Error

	return clientes, err
}

func (r *ClienteRepository) FindByID(id uuid.UUID) (*models.Cliente, error) {

	var cliente models.Cliente

	err := r.DB.First(&cliente, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &cliente, nil
}

func (r *ClienteRepository) Create(cliente *models.Cliente) error {

	return r.DB.Create(cliente).Error

}

// Atualizar cliente
func (r *ClienteRepository) PartialUpdate(cliente *models.Cliente) error {

	return r.DB.Save(cliente).Error
}

func (r *ClienteRepository) Update(cliente *models.Cliente) error {

	return r.DB.Save(cliente).Error
}

// func (r *ClienteRepository) Delete(id uuid.UUID) error {

// 	return r.DB.Delete(&models.Cliente{}, "id = ?", id).Error

// }

func (r *ClienteRepository) Delete(id uuid.UUID) error {

	return r.DB.
		Model(&models.Cliente{}).
		Where("id = ?", id).
		Update("ativo", false).
		Error
}
