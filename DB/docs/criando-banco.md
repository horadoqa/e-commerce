# Criando o Banco de Dados

Se o objetivo é apenas subir um banco **PostgreSQL** com um banco chamado **`ecommerce`**, um `docker-compose.yml` é suficiente, sem necessidade de um `Dockerfile`.

```yaml
services:
  postgres:
    image: postgres:17
    container_name: ecommerce-postgres

    restart: unless-stopped

    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: ecommerce

    ports:
      - "5432:5432"

    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro

    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres -d ecommerce"]
      interval: 10s
      timeout: 5s
      retries: 5

volumes:
  postgres_data:
  
```

### Subir o banco

```bash
docker compose up -d
```

### Verificar os containers

```bash
docker compose ps
```

### Ver os logs

```bash
docker compose logs -f postgres
```

### Parar

```bash
docker compose down
```

### Remover tudo (inclusive os dados)

```bash
docker compose down -v
```

---

Após subir o container, a conexão será:

| Parâmetro | Valor       |
| --------- | ----------- |
| Host      | `localhost` |
| Porta     | `5432`      |
| Banco     | `ecommerce` |
| Usuário   | `postgres`  |
| Senha     | `postgres`  |

Se a intenção for um ambiente de desenvolvimento, também posso montar um `docker-compose.yml` com **PostgreSQL + pgAdmin**, permitindo administrar o banco pela interface web.
