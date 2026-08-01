package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/horadoqa/ecommerce-api/internal/dto"
	"github.com/horadoqa/ecommerce-api/internal/models"
	"github.com/horadoqa/ecommerce-api/internal/response"
	"github.com/horadoqa/ecommerce-api/internal/service"
)

type ClienteHandler struct {
	Service *service.ClienteService
}

// Listar godoc
//
// @Summary Lista todos os clientes
// @Description Retorna todos os clientes
// @Tags Clientes
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /clientes [get]
func (h *ClienteHandler) Listar(c *gin.Context) {

	clientes, err := h.Service.Listar()

	if err != nil {
		response.InternalServerError(c)
		return
	}

	response.OK(c, clientes)
}

// Buscar godoc
//
// @Summary Busca um cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path string true "UUID do cliente"
// @Success 200 {object} response.ClienteSuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /clientes/{id} [get]
func (h *ClienteHandler) Buscar(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		response.BadRequest(c, "UUID inválido.")
		return
	}

	cliente, err := h.Service.Buscar(id)

	if err != nil {
		response.NotFound(c, "Cliente não encontrado.")
		return
	}

	response.OK(c, cliente)
}

// Criar godoc
//
// @Summary Cria um cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param cliente body dto.ClienteRequest true "Dados do cliente"
// @Success 201 {object} response.ClienteSuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /clientes [post]
func (h *ClienteHandler) Criar(c *gin.Context) {

	var req dto.ClienteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	cliente := models.Cliente{
		Nome:           req.Nome,
		CPF:            req.CPF,
		Email:          req.Email,
		Telefone:       req.Telefone,
		DataNascimento: req.DataNascimento,
		Endereco:       req.Endereco,
		Numero:         req.Numero,
		Complemento:    req.Complemento,
		Bairro:         req.Bairro,
		Cidade:         req.Cidade,
		Estado:         req.Estado,
		CEP:            req.CEP,
		Ativo:          true,
	}

	if err := h.Service.Criar(&cliente); err != nil {

		response.InternalServerError(c)
		return
	}

	response.Created(c, cliente)
}

// AtualizarParcial godoc
//
// @Summary Atualiza parcialmente um cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path string true "UUID"
// @Param cliente body map[string]interface{} true "Campos para alteração"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /clientes/{id} [patch]
func (h *ClienteHandler) AtualizarParcial(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {

		response.BadRequest(c, "UUID inválido.")
		return
	}

	var dados map[string]interface{}

	if err := c.ShouldBindJSON(&dados); err != nil {

		response.BadRequest(c, err.Error())
		return
	}

	if err := h.Service.AtualizarParcial(id, dados); err != nil {

		response.NotFound(c, err.Error())
		return
	}

	response.OK(c, dados)
}

// Atualizar godoc
//
// @Summary Atualiza um cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path string true "UUID"
// @Param cliente body dto.ClienteRequest true "Dados completos do cliente"
// @Success 200 {object} response.ClienteSuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /clientes/{id} [put]
func (h *ClienteHandler) Atualizar(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {

		response.BadRequest(c, "UUID inválido.")
		return
	}

	var req dto.ClienteRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.BadRequest(c, err.Error())
		return
	}

	cliente := models.Cliente{

		Nome:           req.Nome,
		CPF:            req.CPF,
		Email:          req.Email,
		Telefone:       req.Telefone,
		DataNascimento: req.DataNascimento,
		Endereco:       req.Endereco,
		Numero:         req.Numero,
		Complemento:    req.Complemento,
		Bairro:         req.Bairro,
		Cidade:         req.Cidade,
		Estado:         req.Estado,
		CEP:            req.CEP,
		Ativo:          true,
	}

	if err := h.Service.Atualizar(id, &cliente); err != nil {

		response.NotFound(c, err.Error())
		return
	}

	response.OK(c, cliente)
}

// Excluir godoc
//
// @Summary Desativa um cliente
// @Description Realiza soft delete alterando ativo para false
// @Tags Clientes
// @Produce json
// @Param id path string true "UUID"
// @Success 204
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /clientes/{id} [delete]
func (h *ClienteHandler) Excluir(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {

		response.BadRequest(c, "UUID inválido.")
		return
	}

	if err := h.Service.Excluir(id); err != nil {

		response.NotFound(c, err.Error())
		return
	}

	response.NoContent(c)
}
