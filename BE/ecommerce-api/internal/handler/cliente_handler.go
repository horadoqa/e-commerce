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

func (h *ClienteHandler) Listar(c *gin.Context) {

	clientes, err := h.Service.Listar()

	if err != nil {
		response.InternalServerError(c)
		return
	}

	response.OK(c, clientes)
}

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

func (h *ClienteHandler) AtualizarParcial(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		response.BadRequest(c, "UUID inválido")
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

func (h *ClienteHandler) Excluir(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "UUID inválido.")
		return
	}

	if err := h.Service.Excluir(id); err != nil {
		response.InternalServerError(c)
		return
	}

	response.NoContent(c)
}
