# Cenários

Abaixo estão cenários de testes para uma **API de e-commerce** usando **Gherkin (BDD)**. Eles cobrem funcionalidades, regras de negócio, segurança, validações e integrações.

---

# Autenticação de usuários

```gherkin
Feature: Autenticação de usuários na API

  Scenario: Usuário realiza login com credenciais válidas
    Given que existe um usuário cadastrado
    And o usuário possui email "cliente@email.com"
    When a API recebe uma requisição de login com email e senha corretos
    Then a API deve retornar status 200
    And deve retornar um token de acesso válido

  Scenario: Usuário tenta realizar login com senha inválida
    Given que existe um usuário cadastrado
    When a API recebe uma requisição de login com senha incorreta
    Then a API deve retornar status 401
    And deve informar que as credenciais são inválidas

  Scenario: Usuário tenta acessar recurso sem autenticação
    Given que o usuário não possui um token válido
    When ele solicita seus pedidos
    Then a API deve retornar status 401
    And deve negar o acesso
```

---

# Cadastro de usuários

```gherkin
Feature: Cadastro de clientes

  Scenario: Criar usuário com dados válidos
    Given que não existe usuário com o email informado
    When a API recebe uma requisição de cadastro válida
    Then a API deve retornar status 201
    And deve criar o usuário no banco de dados

  Scenario: Impedir cadastro com email duplicado
    Given que existe um usuário cadastrado com o email informado
    When a API recebe uma tentativa de cadastro usando o mesmo email
    Then a API deve retornar status 409
    And deve informar que o usuário já existe

  Scenario: Impedir cadastro sem campos obrigatórios
    Given que o cadastro não possui o campo email
    When a API recebe a requisição
    Then a API deve retornar status 400
    And deve informar os campos obrigatórios ausentes
```

---

# Produtos

```gherkin
Feature: Gerenciamento de produtos

  Scenario: Consultar lista de produtos
    Given que existem produtos cadastrados
    When a API recebe uma requisição GET /produtos
    Then deve retornar status 200
    And deve retornar uma lista de produtos

  Scenario: Consultar produto existente
    Given que existe um produto com ID 100
    When a API recebe uma requisição GET /produtos/100
    Then deve retornar status 200
    And deve retornar os dados do produto

  Scenario: Consultar produto inexistente
    Given que não existe produto com ID 999
    When a API recebe uma requisição GET /produtos/999
    Then deve retornar status 404
    And deve informar que o produto não foi encontrado
```

---

# Cadastro e atualização de produtos

```gherkin
Feature: Administração de produtos

  Scenario: Administrador cria um produto
    Given que o usuário autenticado possui perfil administrador
    When ele envia uma requisição para criar um produto válido
    Then a API deve retornar status 201
    And o produto deve ser salvo no banco de dados

  Scenario: Cliente tenta criar produto
    Given que o usuário autenticado possui perfil cliente
    When ele tenta criar um produto
    Then a API deve retornar status 403
    And deve negar a operação

  Scenario: Impedir produto com preço inválido
    Given que o preço informado é menor ou igual a zero
    When a API recebe uma solicitação de criação de produto
    Then deve retornar status 400
    And deve informar que o preço é inválido
```

---

# Estoque

```gherkin
Feature: Controle de estoque

  Scenario: Produto possui estoque disponível
    Given que um produto possui estoque igual a 10 unidades
    When um cliente consulta o produto
    Then a API deve informar que o produto está disponível

  Scenario: Impedir compra acima do estoque disponível
    Given que um produto possui estoque igual a 3 unidades
    When o cliente tenta comprar 5 unidades
    Then a API deve retornar status 400
    And deve informar que o estoque é insuficiente

  Scenario: Atualizar estoque após pedido confirmado
    Given que existe um produto com estoque igual a 10
    When um pedido de 2 unidades é confirmado
    Then o estoque deve ser atualizado para 8 unidades
```

---

# Carrinho de compras

```gherkin
Feature: Carrinho de compras

  Scenario: Adicionar produto ao carrinho
    Given que o usuário está autenticado
    And existe um produto disponível
    When a API recebe uma solicitação para adicionar o produto ao carrinho
    Then deve retornar status 201
    And o produto deve aparecer no carrinho do usuário

  Scenario: Remover produto do carrinho
    Given que o usuário possui um produto no carrinho
    When a API recebe uma solicitação para remover o produto
    Then deve retornar status 200
    And o produto deve ser removido

  Scenario: Impedir quantidade inválida
    Given que o usuário adiciona quantidade igual a zero
    When a API recebe a solicitação
    Then deve retornar status 400
```

---

# Pedidos

```gherkin
Feature: Gerenciamento de pedidos

  Scenario: Criar pedido com sucesso
    Given que o usuário possui produtos no carrinho
    And existe estoque disponível
    And o pagamento foi aprovado
    When a API recebe uma solicitação de criação de pedido
    Then deve retornar status 201
    And deve criar o pedido
    And o status deve ser "PROCESSANDO"

  Scenario: Cliente consulta seus pedidos
    Given que o usuário possui pedidos cadastrados
    When ele solicita sua lista de pedidos
    Then a API deve retornar apenas seus pedidos

  Scenario: Cliente tenta acessar pedido de outro usuário
    Given que existe um pedido pertencente a outro cliente
    When o usuário tenta consultar esse pedido
    Then a API deve retornar status 403
    And deve negar o acesso
```

---

# Pagamentos

```gherkin
Feature: Processamento de pagamentos

  Scenario: Pagamento aprovado
    Given que um pedido foi criado
    When o gateway de pagamento retorna aprovação
    Then a API deve atualizar o pedido para "PAGO"

  Scenario: Pagamento recusado
    Given que um pedido foi criado
    When o gateway de pagamento retorna rejeição
    Then a API deve atualizar o pedido para "PAGAMENTO_RECUSADO"

  Scenario: Falha no gateway de pagamento
    Given que o serviço de pagamento está indisponível
    When a API tenta processar o pagamento
    Then deve retornar status 503
    And não deve confirmar o pedido
```

---

# Segurança da API

```gherkin
Feature: Segurança da API

  Scenario: Bloquear token expirado
    Given que o usuário possui um token expirado
    When ele acessa um endpoint protegido
    Then a API deve retornar status 401

  Scenario: Bloquear usuário sem permissão
    Given que o usuário possui perfil cliente
    When ele acessa uma rota administrativa
    Then a API deve retornar status 403

  Scenario: Limitar tentativas de login
    Given que um usuário realiza várias tentativas inválidas
    When ultrapassa o limite permitido
    Then a API deve bloquear temporariamente novas tentativas
```

---

# Validação de dados

```gherkin
Feature: Validação de entradas

  Scenario: Rejeitar dados com formato inválido
    Given que o email informado possui formato inválido
    When a API recebe a solicitação de cadastro
    Then deve retornar status 400

  Scenario: Rejeitar campos desconhecidos
    Given que a requisição possui campos não permitidos
    When a API processa a solicitação
    Then deve rejeitar ou ignorar os campos conforme configuração
```

---

# Performance

```gherkin
Feature: Performance da API

  Scenario: API responde dentro do tempo esperado
    Given que existem 10000 produtos cadastrados
    When um usuário consulta a lista de produtos
    Then a API deve responder em até 500 milissegundos

  Scenario: API suporta múltiplos usuários simultâneos
    Given que existem 500 usuários acessando a API ao mesmo tempo
    When todos consultam produtos
    Then a API deve continuar respondendo corretamente
```

---

# Resumo dos cenários de API

| Área         | Cenários                               |
| ------------ | -------------------------------------- |
| Autenticação | Login, token, sessão                   |
| Usuários     | Cadastro, duplicidade, validação       |
| Produtos     | Consulta, criação, atualização         |
| Estoque      | Disponibilidade, baixa, reposição      |
| Carrinho     | Adicionar, remover, validar quantidade |
| Pedidos      | Criar, consultar, permissões           |
| Pagamento    | Aprovação, rejeição, falha             |
| Segurança    | Acesso, tokens, abuso                  |
| Performance  | Carga e tempo de resposta              |

Esses cenários podem ser automatizados com ferramentas como **Cucumber, RestAssured (Java), Postman/Newman, Karate, Playwright API Testing, Jest/Supertest ou Pytest**, dependendo da tecnologia usada no Back End.
