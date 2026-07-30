# Clientes

Consultando tabela:

```bash
SELECT * FROM ecommerce.clientes;
```

Inserindo dados na tabela:

```bash
INSERT INTO clientes (
    nome,
    cpf,
    email,
    telefone,
    data_nascimento,
    endereco,
    numero,
    complemento,
    bairro,
    cidade,
    estado,
    cep,
    ativo
)
VALUES (
    'João da Silva',
    '12345678901',
    'joao@email.com',
    '21999998888',
    '1990-05-20',
    'Rua das Flores',
    '123',
    'Apto 301',
    'Centro',
    'Petrópolis',
    'RJ',
    '25600000',
    TRUE
);
```
