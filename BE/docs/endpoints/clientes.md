Essa é uma ótima primeira etapa. Eu implementaria seguindo a arquitetura **Handler → Service → Repository → Database**, que deixa o código organizado e facilita a evolução do projeto.

## Estrutura

```text
internal/
├── handler/
│   └── cliente_handler.go
├── service/
│   └── cliente_service.go
├── repository/
│   └── cliente_repository.go
├── models/
│   └── cliente.go
├── dto/
│   ├── cliente_request.go
│   └── cliente_response.go
└── routes/
    └── cliente_routes.go
```

Além disso, criaria DTOs para não expor diretamente o modelo do banco.

---

## Model

```go
package models

import (
	"time"

	"github.com/google/uuid"
)

type Cliente struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Nome string `gorm:"column:nome"`
	CPF string `gorm:"column:cpf"`

	Email string `gorm:"column:email"`
	Telefone string `gorm:"column:telefone"`

	DataNascimento time.Time `gorm:"column:data_nascimento"`

	Endereco string `gorm:"column:endereco"`
	Numero string `gorm:"column:numero"`
	Complemento string `gorm:"column:complemento"`

	Bairro string `gorm:"column:bairro"`
	Cidade string `gorm:"column:cidade"`
	Estado string `gorm:"column:estado"`
	CEP string `gorm:"column:cep"`

	Ativo bool `gorm:"column:ativo"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Cliente) TableName() string {
	return "clientes"
}
```

---

# DTO Request

```go
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
```

---

# Repository

```go
package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"ecommerce/internal/models"
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

func (r *ClienteRepository) Update(cliente *models.Cliente) error {

	return r.DB.Save(cliente).Error
}

func (r *ClienteRepository) Delete(id uuid.UUID) error {

	return r.DB.Delete(&models.Cliente{}, "id = ?", id).Error
}
```

---

# Service

```go
package service

import (
	"errors"

	"github.com/google/uuid"

	"ecommerce/internal/models"
	"ecommerce/internal/repository"
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
```

---

# Handler

```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ecommerce/internal/dto"
	"ecommerce/internal/models"
	"ecommerce/internal/service"
)

type ClienteHandler struct {
	Service *service.ClienteService
}

func (h *ClienteHandler) Listar(c *gin.Context) {

	clientes, err := h.Service.Listar()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clientes)
}

func (h *ClienteHandler) Buscar(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro":"UUID inválido"})
		return
	}

	cliente, err := h.Service.Buscar(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro":"Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

func (h *ClienteHandler) Criar(c *gin.Context) {

	var req dto.ClienteRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, err)

		return
	}

	cliente := models.Cliente{
		Nome: req.Nome,
		CPF: req.CPF,
		Email: req.Email,
		Telefone: req.Telefone,
		DataNascimento: req.DataNascimento,
		Endereco: req.Endereco,
		Numero: req.Numero,
		Complemento: req.Complemento,
		Bairro: req.Bairro,
		Cidade: req.Cidade,
		Estado: req.Estado,
		CEP: req.CEP,
		Ativo: true,
	}

	if err := h.Service.Criar(&cliente); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"erro":err.Error()})

		return
	}

	c.JSON(http.StatusCreated, cliente)
}

func (h *ClienteHandler) Atualizar(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var req dto.ClienteRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, err)

		return
	}

	cliente := models.Cliente{
		Nome: req.Nome,
		CPF: req.CPF,
		Email: req.Email,
		Telefone: req.Telefone,
		DataNascimento: req.DataNascimento,
		Endereco: req.Endereco,
		Numero: req.Numero,
		Complemento: req.Complemento,
		Bairro: req.Bairro,
		Cidade: req.Cidade,
		Estado: req.Estado,
		CEP: req.CEP,
		Ativo: true,
	}

	if err := h.Service.Atualizar(id, &cliente); err != nil {

		c.JSON(http.StatusNotFound, gin.H{"erro":err.Error()})

		return
	}

	c.JSON(http.StatusOK, cliente)
}

func (h *ClienteHandler) Excluir(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"erro":"UUID inválido"})

		return
	}

	if err := h.Service.Excluir(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"erro":err.Error()})

		return
	}

	c.Status(http.StatusNoContent)
}
```

---

# Rotas

```go
package routes

import (
	"github.com/gin-gonic/gin"

	"ecommerce/internal/handler"
)

func ClienteRoutes(router *gin.Engine, handler *handler.ClienteHandler) {

	clientes := router.Group("/clientes")
	{
		clientes.GET("", handler.Listar)
		clientes.GET("/:id", handler.Buscar)
		clientes.POST("", handler.Criar)
		clientes.PUT("/:id", handler.Atualizar)
		clientes.DELETE("/:id", handler.Excluir)
	}
}
```

## Próximos aprimoramentos

Embora esse CRUD funcione, para um e-commerce em produção eu faria alguns refinamentos antes de seguir para Produtos e Pedidos:

* Criar um pacote `response` para padronizar todas as respostas da API.
* Adicionar paginação e filtros no `GET /clientes` (`page`, `limit`, `nome`, `email`, `ativo`).
* Verificar CPF e e-mail duplicados na camada de serviço antes de inserir ou atualizar.
* Utilizar transações quando uma operação envolver mais de uma tabela.
* Implementar logs estruturados, tratamento centralizado de erros e testes unitários.
* Gerar documentação com Swagger/OpenAPI.

Essa base servirá para implementar os demais módulos (`Produtos`, `Pedidos` e `PedidoItens`) seguindo exatamente o mesmo padrão de arquitetura.
