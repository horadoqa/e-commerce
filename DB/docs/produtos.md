# Pedidos

Consultando tabela:

```bash
SELECT * FROM ecommerce.produtos;
```

Inserindo dados:

```bash
INSERT INTO ecommerce.produtos (
    sku,
    nome,
    descricao,
    categoria,
    preco,
    custo,
    quantidade_estoque,
    estoque_minimo,
    peso
)
VALUES
(
    'NOTE-001',
    'Notebook Dell Inspiron',
    'Notebook 15 polegadas com processador Intel Core i5',
    'Informática',
    3500.00,
    2800.00,
    10,
    2,
    1.800
),
(
    'MOU-001',
    'Mouse Wireless',
    'Mouse sem fio USB',
    'Acessórios',
    120.00,
    50.00,
    50,
    10,
    0.150
),
(
    'TEC-001',
    'Teclado Mecânico',
    'Teclado mecânico RGB para jogos',
    'Acessórios',
    250.00,
    120.00,
    30,
    5,
    0.900
);

