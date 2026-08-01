# **Swagger/OpenAPI**

Como estamos usando **Go + Gin**, a melhor opção é usar o **Swagger/OpenAPI** com a biblioteca **swaggo**. Ela gera automaticamente a documentação a partir de comentários no código e disponibiliza uma interface Swagger UI.

## 1. Instale as dependências

```bash
go install github.com/swaggo/swag/cmd/swag@latest

go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

Verifique se o binário `swag` está disponível:

```bash
swag --version

swag version v1.16.4
```

Se não estiver, adicione o diretório de binários do Go ao `PATH`.

---

## 2. Adicione as anotações no `main.go`

Antes da função `main()`, adicione:

```go
// @title Ecommerce API
// @version 1.0
// @description API REST para gerenciamento de um e-commerce.
// @termsOfService http://swagger.io/terms/

// @contact.name Ricardo Fahham
// @contact.email ricardo@email.com

// @license.name MIT

// @host localhost:8080
// @BasePath /
```

---

## 3. Gere a documentação

Na raiz do projeto execute:

```bash
swag init -g cmd/api/main.go
```

Será criada uma pasta:

```
docs/
    docs.go
    swagger.json
    swagger.yaml
```

---

## 4. Registrar o Swagger no Gin

No `main.go` adicione os imports:

```go
import (
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    _ "github.com/horadoqa/ecommerce-api/docs"
)
```

Depois das demais rotas:

```go
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

Agora você poderá acessar:

```
http://localhost:8080/swagger/index.html
```

---

# 5. Documentando um endpoint

Por exemplo, em `cliente_handler.go`:

```go
// Listar godoc
//
// @Summary Lista todos os clientes
// @Description Retorna todos os clientes ativos
// @Tags Clientes
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /clientes [get]
func (h *ClienteHandler) Listar(c *gin.Context) {
    ...
}
```

---

## Buscar cliente

```go
// Buscar godoc
//
// @Summary Busca um cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path string true "UUID do cliente"
// @Success 200 {object} models.Cliente
// @Failure 404 {object} response.ErrorResponse
// @Router /clientes/{id} [get]
```

---

## Criar cliente

```go
// Criar godoc
//
// @Summary Cria um cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param cliente body dto.ClienteRequest true "Dados do cliente"
// @Success 201 {object} models.Cliente
// @Failure 400 {object} response.ErrorResponse
// @Router /clientes [post]
```

---

## Atualizar cliente

```go
// Atualizar godoc
//
// @Summary Atualiza um cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path string true "UUID"
// @Param cliente body dto.ClienteRequest true "Cliente"
// @Success 200 {object} models.Cliente
// @Router /clientes/{id} [put]
```

---

## PATCH

```go
// AtualizarParcial godoc
//
// @Summary Atualiza parcialmente um cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path string true "UUID"
// @Param cliente body map[string]interface{} true "Campos"
// @Success 200 {object} response.SuccessResponse
// @Router /clientes/{id} [patch]
```

---

## DELETE

```go
// Excluir godoc
//
// @Summary Desativa um cliente
// @Tags Clientes
// @Produce json
// @Param id path string true "UUID"
// @Success 204
// @Router /clientes/{id} [delete]
```

---

# Resultado

Depois de executar novamente:

```bash
swag init -g cmd/api/main.go
go run ./cmd/api
```

Acesse:

```
http://localhost:8080/swagger/index.html
```

Você terá uma interface semelhante à do Swagger Editor, com:

* Todos os endpoints organizados por grupo ("Clientes", futuramente "Produtos", "Pedidos");
* Descrição de cada operação;
* Campos esperados no corpo da requisição;
* Códigos de resposta;
* Botão **Try it out**, que permite testar os endpoints diretamente pelo navegador.

Para uma API de e-commerce, essa abordagem facilita tanto o desenvolvimento quanto a integração com clientes e outros sistemas.
