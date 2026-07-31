package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/horadoqa/ecommerce-api/internal/models"
	"github.com/horadoqa/ecommerce-api/internal/repository"
)

type ClienteService struct {
	Repository *repository.ClienteRepository
}

func (s *ClienteService) Listar() ([]models.Cliente, error) {

	return s.Repository.FindAll()
}

func (s *ClienteService) Buscar(id uuid.UUID) (*models.Cliente, error) {

	return s.Repository.FindByID(id)
}

func (s *ClienteService) Criar(cliente *models.Cliente) error {

	return s.Repository.Create(cliente)
}

func (s *ClienteService) AtualizarParcial(id uuid.UUID, dados map[string]interface{}) error {

	cliente, err := s.Repository.FindByID(id)

	if err != nil {
		return err
	}

	// Validações de negócio

	if email, ok := dados["email"]; ok {

		if email == "" {
			return errors.New("email não pode ser vazio")
		}

		cliente.Email = email.(string)
	}

	if nome, ok := dados["nome"]; ok {

		if nome == "" {
			return errors.New("nome não pode ser vazio")
		}

		cliente.Nome = nome.(string)
	}

	if telefone, ok := dados["telefone"]; ok {
		cliente.Telefone = telefone.(string)
	}

	if endereco, ok := dados["endereco"]; ok {
		cliente.Endereco = endereco.(string)
	}

	if numero, ok := dados["numero"]; ok {
		cliente.Numero = numero.(string)
	}

	if complemento, ok := dados["complemento"]; ok {
		cliente.Complemento = complemento.(string)
	}

	if bairro, ok := dados["bairro"]; ok {
		cliente.Bairro = bairro.(string)
	}

	if cidade, ok := dados["cidade"]; ok {
		cliente.Cidade = cidade.(string)
	}

	if estado, ok := dados["estado"]; ok {
		cliente.Estado = estado.(string)
	}

	if cep, ok := dados["cep"]; ok {
		cliente.CEP = cep.(string)
	}

	if ativo, ok := dados["ativo"]; ok {
		cliente.Ativo = ativo.(bool)
	}

	return s.Repository.PartialUpdate(cliente)
}

func (s *ClienteService) Atualizar(id uuid.UUID, cliente *models.Cliente) error {

	existente, err := s.Repository.FindByID(id)

	if err != nil {
		return errors.New("cliente não encontrado")
	}

	cliente.ID = existente.ID

	return s.Repository.Update(cliente)
}

func (s *ClienteService) Excluir(id uuid.UUID) error {

	return s.Repository.Delete(id)
}
