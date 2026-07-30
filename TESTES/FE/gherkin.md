# Cenários

Utilizando **Gherkin** (linguagem usada em BDD — *Behavior Driven Development*). 

Abaixo estão cenários de testes para o **Front End de um e-commerce** usando a estrutura:

* **Feature**: funcionalidade
* **Scenario**: cenário
* **Given** (*Dado*): contexto inicial
* **When** (*Quando*): ação do usuário
* **Then** (*Então*): resultado esperado

```gherkin
Feature: Navegação no catálogo de produtos

  Scenario: Usuário visualiza a lista de produtos
    Given que o usuário acessa a página inicial do e-commerce
    When a página de produtos é carregada
    Then o sistema deve exibir uma lista de produtos disponíveis
    And cada produto deve apresentar nome, imagem e preço
```

---

## Busca de produtos

```gherkin
Feature: Pesquisa de produtos

  Scenario: Usuário pesquisa um produto existente
    Given que o usuário está na página de produtos
    When ele informa "Notebook" no campo de busca
    And clica no botão de pesquisar
    Then o sistema deve exibir produtos relacionados a "Notebook"

  Scenario: Usuário pesquisa um produto inexistente
    Given que o usuário está na página de produtos
    When ele informa "ProdutoXYZ123" no campo de busca
    And clica no botão de pesquisar
    Then o sistema deve informar que nenhum produto foi encontrado
```

---

## Visualização de detalhes do produto

```gherkin
Feature: Detalhes do produto

  Scenario: Usuário acessa detalhes de um produto
    Given que o usuário está visualizando a lista de produtos
    When ele seleciona um produto
    Then o sistema deve exibir:
      | Informação |
      | Nome       |
      | Imagem     |
      | Preço      |
      | Descrição  |
      | Estoque    |

  Scenario: Produto indisponível não pode ser comprado
    Given que o usuário acessa um produto sem estoque
    When a página do produto é carregada
    Then o botão "Comprar" deve estar desabilitado
    And deve exibir a mensagem "Produto indisponível"
```

---

# Carrinho de compras

```gherkin
Feature: Gerenciamento do carrinho

  Scenario: Adicionar produto ao carrinho
    Given que o usuário está na página de detalhes de um produto disponível
    When ele clica no botão "Adicionar ao carrinho"
    Then o produto deve ser adicionado ao carrinho
    And a quantidade de itens do carrinho deve ser atualizada

  Scenario: Alterar quantidade de produto no carrinho
    Given que o usuário possui um produto no carrinho
    When ele altera a quantidade de 1 para 3 unidades
    Then o sistema deve atualizar o subtotal do produto
    And atualizar o valor total da compra

  Scenario: Remover produto do carrinho
    Given que o usuário possui produtos no carrinho
    When ele remove um produto
    Then o produto deve desaparecer do carrinho
    And o valor total deve ser recalculado
```

---

# Cadastro de usuário

```gherkin
Feature: Cadastro de cliente

  Scenario: Usuário realiza cadastro com dados válidos
    Given que o usuário está na tela de cadastro
    When ele informa nome, email e senha válidos
    And confirma o cadastro
    Then o sistema deve criar a conta do usuário
    And exibir uma mensagem de sucesso

  Scenario: Usuário tenta cadastrar email já existente
    Given que existe um usuário cadastrado com o email "cliente@email.com"
    When o usuário tenta criar uma conta usando esse email
    Then o sistema deve impedir o cadastro
    And exibir a mensagem "Email já cadastrado"
```

---

# Login

```gherkin
Feature: Autenticação do usuário

  Scenario: Usuário realiza login com sucesso
    Given que o usuário possui uma conta cadastrada
    When ele informa email e senha corretos
    And clica no botão "Entrar"
    Then o sistema deve autenticar o usuário
    And redirecionar para a área logada

  Scenario: Usuário informa senha incorreta
    Given que o usuário possui uma conta cadastrada
    When ele informa uma senha inválida
    And clica no botão "Entrar"
    Then o sistema deve negar o acesso
    And exibir uma mensagem de erro
```

---

# Checkout

```gherkin
Feature: Finalização da compra

  Scenario: Usuário inicia checkout com produtos no carrinho
    Given que o usuário possui produtos no carrinho
    When ele clica em "Finalizar compra"
    Then o sistema deve abrir a tela de checkout

  Scenario: Usuário tenta finalizar compra sem endereço
    Given que o usuário possui produtos no carrinho
    And não possui endereço cadastrado
    When ele tenta finalizar a compra
    Then o sistema deve solicitar o cadastro do endereço

  Scenario: Usuário confirma pedido
    Given que o usuário está no checkout
    And possui endereço e forma de pagamento válidos
    When ele confirma a compra
    Then o sistema deve enviar o pedido para processamento
    And exibir a confirmação do pedido
```

---

# Tratamento de erros da API

```gherkin
Feature: Tratamento de falhas de comunicação

  Scenario: API de produtos está indisponível
    Given que o usuário acessa a página de produtos
    And a API retorna erro 500
    When a página tenta carregar os produtos
    Then o sistema deve exibir uma mensagem de erro
    And disponibilizar uma opção para tentar novamente

  Scenario: Sessão do usuário expirada
    Given que o usuário está autenticado
    When sua sessão expira
    And ele tenta acessar uma área restrita
    Then o sistema deve redirecionar para o login
```

---

# Acessibilidade

```gherkin
Feature: Acessibilidade da interface

  Scenario: Usuário navega utilizando teclado
    Given que o usuário acessa o site
    When ele utiliza somente a tecla TAB para navegar
    Then todos os elementos interativos devem receber foco
    And o usuário deve conseguir acessar as funcionalidades

  Scenario: Imagens possuem descrição
    Given que existem imagens de produtos na página
    When a página é analisada por um leitor de tela
    Then cada imagem deve possuir texto alternativo
```

---

# Responsividade

```gherkin
Feature: Layout responsivo

  Scenario: Usuário acessa pelo celular
    Given que o usuário acessa o e-commerce em um dispositivo móvel
    When a página é carregada
    Then os elementos devem se adaptar ao tamanho da tela
    And os produtos devem permanecer visíveis e utilizáveis
```

---

## Matriz resumida de cenários FE

| Área           | Cenários principais              |
| -------------- | -------------------------------- |
| Catálogo       | Listagem, busca, filtros         |
| Produto        | Detalhes, estoque, imagens       |
| Carrinho       | Adicionar, alterar, remover      |
| Cadastro       | Dados válidos e inválidos        |
| Login          | Sucesso, erro, sessão expirada   |
| Checkout       | Endereço, pagamento, confirmação |
| Erros          | API fora, mensagens de erro      |
| Acessibilidade | Teclado, leitor de tela          |
| Responsividade | Desktop, tablet, mobile          |

Esses cenários podem ser utilizados diretamente em ferramentas BDD como **Cucumber**, **Cypress com Cucumber**, **Playwright BDD** ou integrados ao pipeline de testes do Front End.
