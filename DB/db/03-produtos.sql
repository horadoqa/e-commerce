CREATE TABLE ecommerce.produtos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    sku VARCHAR(50) UNIQUE NOT NULL,

    nome VARCHAR(150) NOT NULL,

    descricao TEXT,

    categoria VARCHAR(100),

    preco NUMERIC(12,2) NOT NULL,

    custo NUMERIC(12,2),

    quantidade_estoque INTEGER NOT NULL DEFAULT 0,

    estoque_minimo INTEGER NOT NULL DEFAULT 0,

    peso NUMERIC(10,3),

    ativo BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- ==========================================
-- Índices
-- ==========================================

CREATE INDEX idx_produtos_nome
    ON ecommerce.produtos(nome);


CREATE INDEX idx_produtos_sku
    ON ecommerce.produtos(sku);


CREATE INDEX idx_produtos_categoria
    ON ecommerce.produtos(categoria);


CREATE INDEX idx_produtos_ativo
    ON ecommerce.produtos(ativo);


-- ==========================================
-- Trigger de atualização
-- ==========================================

CREATE TRIGGER trigger_produtos_updated_at
BEFORE UPDATE ON ecommerce.produtos
FOR EACH ROW
EXECUTE FUNCTION atualizar_updated_at();