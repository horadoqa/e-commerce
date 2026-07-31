Como você já possui o banco modelado, o ideal é construir a API seguindo uma arquitetura em camadas. Isso facilita a manutenção e o crescimento do sistema.

## Estrutura do projeto

Uma estrutura bastante utilizada em projetos Go é:

```text
ecommerce-api/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── config/
│   ├── database.go
│   └── env.go
│
├── internal/
│   ├── models/
│   │   ├── cliente.go
│   │   ├── produto.go
│   │   ├── pedido.go
│   │   └── pedido_item.go
│   │
│   ├── repository/
│   │   ├── cliente_repository.go
│   │   ├── produto_repository.go
│   │   ├── pedido_repository.go
│   │   └── interfaces.go
│   │
│   ├── service/
│   │   ├── cliente_service.go
│   │   ├── produto_service.go
│   │   └── pedido_service.go
│   │
│   ├── handler/
│   │   ├── cliente_handler.go
│   │   ├── produto_handler.go
│   │   └── pedido_handler.go
│   │
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── logger.go
│   │   └── cors.go
│   │
│   └── routes/
│       └── routes.go
│
├── pkg/
│   ├── response/
│   └── utils/
│
├── .env
├── go.mod
└── README.md
```

Essa organização separa responsabilidades e facilita testes.

---

# Criando o projeto

```bash
mkdir ecommerce-api

cd ecommerce-api

go mod init github.com/seuusuario/ecommerce-api
```

---

## Instalando as dependências

### Gin

```bash
go get github.com/gin-gonic/gin
```

### GORM

```bash
go get gorm.io/gorm
```

### Driver PostgreSQL

```bash
go get gorm.io/driver/postgres
```

### Variáveis de ambiente

```bash
go get github.com/joho/godotenv
```

### UUID

```bash
go get github.com/google/uuid
```

---

## Arquivo `.env`

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=senha
DB_NAME=ecommerce
DB_SSLMODE=disable

SERVER_PORT=8080
```

---

# Configuração do banco

`config/database.go`

```go
package config

import (
    "fmt"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSLMODE"),
    )

    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
```

---

# main.go

```go
package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "ecommerce-api/config"
)

func main() {

    godotenv.Load()

    db, err := config.ConnectDatabase()

    if err != nil {
        log.Fatal(err)
    }

    _ = db

    router := gin.Default()

    router.Run(":8080")
}
```

---

# Model Cliente

```go
package models

import (
    "time"

    "github.com/google/uuid"
)

type Cliente struct {
    ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

    Nome            string
    CPF             string
    Email           string
    Telefone        string

    DataNascimento  time.Time

    Endereco        string
    Numero          string
    Complemento     string

    Bairro          string
    Cidade          string
    Estado          string
    CEP             string

    Ativo           bool

    CreatedAt       time.Time
    UpdatedAt       time.Time
}

func (Cliente) TableName() string {
    return "clientes"
}
```

---

# Model Produto

```go
type Produto struct {
    ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

    SKU string

    Nome string

    Descricao string

    Categoria string

    Preco float64

    Custo float64

    QuantidadeEstoque int

    EstoqueMinimo int

    Peso float64

    Ativo bool

    CreatedAt time.Time

    UpdatedAt time.Time
}

func (Produto) TableName() string {
    return "produtos"
}
```

---

# Repository

Exemplo para Produto.

```go
type ProdutoRepository struct {
    DB *gorm.DB
}

func (r ProdutoRepository) Listar() ([]models.Produto, error) {

    var produtos []models.Produto

    err := r.DB.Find(&produtos).Error

    return produtos, err
}
```

---

# Service

```go
type ProdutoService struct {
    Repository repository.ProdutoRepository
}

func (s ProdutoService) Listar() ([]models.Produto, error) {

    return s.Repository.Listar()
}
```

---

# Handler

```go
type ProdutoHandler struct {
    Service service.ProdutoService
}

func (h ProdutoHandler) Listar(c *gin.Context) {

    produtos, err := h.Service.Listar()

    if err != nil {
        c.JSON(500, gin.H{
            "erro": err.Error(),
        })
        return
    }

    c.JSON(200, produtos)
}
```

---

# Rotas

```go
func SetupRoutes(router *gin.Engine, handler ProdutoHandler) {

    api := router.Group("/api")

    api.GET("/produtos", handler.Listar)
}
```

---

# Próximas APIs

Depois você pode implementar os endpoints principais:

## Clientes

```
GET     /clientes
GET     /clientes/:id
POST    /clientes
PUT     /clientes/:id
DELETE  /clientes/:id
```

## Produtos

```
GET     /produtos
GET     /produtos/:id
POST    /produtos
PUT     /produtos/:id
DELETE  /produtos/:id
```

## Pedidos

```
GET     /pedidos
GET     /pedidos/:id
POST    /pedidos
PUT     /pedidos/:id/status
```

## Itens

```
POST    /pedidos/:id/itens
DELETE  /pedidos/:id/itens/:itemId
```

## Recomendações adicionais

Como seu banco já existe, eu faria alguns ajustes em relação ao uso do GORM:

* **Não utilize `AutoMigrate()`**. Como o schema já está criado, deixe o GORM apenas mapear as tabelas.
* Considere usar o tipo `decimal.Decimal` (pacote `github.com/shopspring/decimal`) para campos monetários (`preco`, `custo`, `valor_total`), evitando problemas de precisão que podem ocorrer com `float64`.
* Utilize transações (`db.Transaction(...)`) ao criar pedidos, para garantir que o pedido, seus itens e a atualização do estoque sejam gravados de forma atômica.
* Centralize validações na camada de serviço (estoque disponível, CPF e e-mail únicos, regras de negócio) e deixe os handlers responsáveis apenas por receber a requisição e enviar a resposta.

Para um e-commerce desse porte, essa arquitetura (**Gin + GORM + PostgreSQL + Repository + Service + Handler**) é uma base sólida, organizada e preparada para crescer com autenticação, cache, filas e integrações futuras.
