# Ecommerce API

API REST para gerenciamento de um sistema de e-commerce desenvolvida em **Go**, utilizando **Gin Framework**, **GORM** e **PostgreSQL**.

Atualmente implementado:

- Cadastro de clientes
- Consulta de clientes
- Atualização completa (PUT)
- Atualização parcial (PATCH)
- Soft Delete de clientes
- Padronização de respostas da API
- Healthcheck

---

# Tecnologias utilizadas

- Go 1.25+
- Gin Framework
- GORM
- PostgreSQL
- UUID como identificador
- Godotenv para gerenciamento de variáveis de ambiente

---

# Estrutura do projeto

```
ecommerce-api
│
├── cmd
│   └── api
│       └── main.go
│
├── config
│   ├── database.go
│   └── env.go
│
├── internal
│   │
│   ├── dto
│   │
│   ├── handler
│   │
│   ├── middleware
│   │
│   ├── models
│   │
│   ├── repository
│   │
│   ├── routes
│   │
│   ├── service
│   │
│   └── response
│
├── go.mod
└── README.md
```

---

# Executando a aplicação

Entre no diretório do projeto:

```bash
cd ecommerce-api
```

---

## Verificar formatação do projeto

Execute:

```bash
go fmt ./...
```

Se não houver erros, o projeto está formatado corretamente.

---

## Subir a API

Execute:

```bash
go run ./cmd/api
```

A aplicação estará disponível em:

```
http://localhost:8080
```

Mensagem esperada:

```
[GIN-debug] Listening and serving HTTP on :8080
```

Isso indica que a API iniciou corretamente.

---

# Healthcheck

Para validar se a API está online:

```bash
curl http://localhost:8080/healthcheck
```

Resposta:

```json
{
    "message": "WORKING !!!"
}
```

---

# Rota principal

Teste:

```bash
curl http://localhost:8080/
```

Resposta:

```json
{
    "api": "ecommerce-api",
    "status": "online"
}
```

---

# Clientes

## Listar clientes

Endpoint:

```
GET /clientes
```

Exemplo:

```bash
curl http://localhost:8080/clientes
```

Resposta:

```json
{
    "success": true,
    "message": "Operação realizada com sucesso.",
    "data": []
}
```

---

# Criar cliente

Endpoint:

```
POST /clientes
```

Exemplo:

```bash
curl -X POST http://localhost:8080/clientes \
-H "Content-Type: application/json" \
-d '{
    "nome":"João Silva",
    "cpf":"12345678906",
    "email":"joao@email.com",
    "telefone":"21999999999",
    "dataNascimento":"1990-01-01T00:00:00Z",
    "endereco":"Rua Principal",
    "numero":"100",
    "bairro":"Centro",
    "cidade":"Petrópolis",
    "estado":"RJ",
    "cep":"25600000"
}'
```

---

# Buscar cliente por ID

Endpoint:

```
GET /clientes/:id
```

Exemplo:

```bash
curl http://localhost:8080/clientes/fd6e5ac5-7155-4899-a1d7-7e78e201c166
```

---

# Atualizar cliente (PUT)

O PUT realiza uma atualização completa do registro.

Endpoint:

```
PUT /clientes/:id
```

Exemplo:

```bash
curl -X PUT http://localhost:8080/clientes/fd6e5ac5-7155-4899-a1d7-7e78e201c166 \
-H "Content-Type: application/json" \
-d '{
    "nome":"João Silva Atualizado",
    "cpf":"12345678906",
    "email":"joao.atualizado@email.com",
    "telefone":"21988887777",
    "dataNascimento":"1990-05-20T00:00:00Z",
    "endereco":"Rua Nova",
    "numero":"456",
    "bairro":"Centro",
    "cidade":"Petrópolis",
    "estado":"RJ",
    "cep":"25600000"
}'
```

---

# Atualizar parcialmente (PATCH)

O PATCH altera somente os campos enviados.

Endpoint:

```
PATCH /clientes/:id
```

Exemplo:

Alterando telefone:

```bash
curl -X PATCH http://localhost:8080/clientes/fd6e5ac5-7155-4899-a1d7-7e78e201c166 \
-H "Content-Type: application/json" \
-d '{
    "telefone":"21999990000"
}'
```

Alterando múltiplos campos:

```bash
curl -X PATCH http://localhost:8080/clientes/fd6e5ac5-7155-4899-a1d7-7e78e201c166 \
-H "Content-Type: application/json" \
-d '{
    "telefone":"21999990000",
    "cidade":"Rio de Janeiro"
}'
```

---

# Remover cliente (Soft Delete)

A aplicação não remove fisicamente clientes.

O DELETE apenas altera:

```sql
ativo = false
```

Isso mantém o histórico dos pedidos relacionados ao cliente.

Endpoint:

```
DELETE /clientes/:id
```

Exemplo:

```bash
curl -X DELETE http://localhost:8080/clientes/fd6e5ac5-7155-4899-a1d7-7e78e201c166
```

---

# Rotas disponíveis

| Método | Endpoint | Descrição |
|---|---|---|
| GET | `/` | Status da API |
| GET | `/healthcheck` | Healthcheck |
| GET | `/clientes` | Lista clientes |
| GET | `/clientes/:id` | Busca cliente |
| POST | `/clientes` | Cria cliente |
| PUT | `/clientes/:id` | Atualiza cliente completo |
| PATCH | `/clientes/:id` | Atualiza parcialmente |
| DELETE | `/clientes/:id` | Soft delete |

---

# Variáveis de ambiente

Criar arquivo:

```
.env
```

Exemplo:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecommerce
```

---

# Banco de dados

O projeto utiliza PostgreSQL.

As tabelas principais:

- clientes
- produtos
- pedidos
- pedido_itens

Os IDs utilizam UUID:

```sql
UUID PRIMARY KEY DEFAULT gen_random_uuid()
```

---

# Ambiente de produção

## Gin Release Mode

Em produção utilize:

```bash
export GIN_MODE=release
```

ou:

```go
gin.SetMode(gin.ReleaseMode)
```

---

## Trusted Proxies

Para ambientes com proxy/load balancer:

```go
router.SetTrustedProxies(nil)
```

Isso evita confiar automaticamente em proxies desconhecidos.

---

# Próximos passos

Implementações planejadas:

- [ ] CRUD Produtos
- [ ] CRUD Pedidos
- [ ] Controle de estoque
- [ ] Autenticação JWT
- [ ] Middleware de autorização
- [ ] Swagger/OpenAPI
- [ ] Docker
- [ ] Testes automatizados