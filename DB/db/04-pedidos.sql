-- ==========================================
-- Tabela de Pedidos
-- ==========================================

CREATE TABLE pedidos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    cliente_id UUID NOT NULL,

    status VARCHAR(30) NOT NULL DEFAULT 'PENDENTE',

    valor_total NUMERIC(12,2) NOT NULL DEFAULT 0,

    observacao TEXT,

    data_pedido TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    data_atualizacao TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_pedido_cliente
        FOREIGN KEY (cliente_id)
        REFERENCES clientes(id)
        ON DELETE RESTRICT
);


-- ==========================================
-- Tabela de Itens do Pedido
-- ==========================================

CREATE TABLE pedido_itens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    pedido_id UUID NOT NULL,

    produto_id UUID NOT NULL,

    quantidade INTEGER NOT NULL DEFAULT 1,

    valor_unitario NUMERIC(12,2) NOT NULL,

    desconto NUMERIC(12,2) NOT NULL DEFAULT 0,

    valor_total NUMERIC(12,2) GENERATED ALWAYS AS
        ((quantidade * valor_unitario) - desconto)
        STORED,


    CONSTRAINT fk_item_pedido
        FOREIGN KEY (pedido_id)
        REFERENCES pedidos(id)
        ON DELETE CASCADE,


    CONSTRAINT fk_item_produto
        FOREIGN KEY (produto_id)
        REFERENCES produtos(id)
        ON DELETE RESTRICT
);


-- ==========================================
-- Índices
-- ==========================================

CREATE INDEX idx_pedidos_cliente
    ON pedidos(cliente_id);


CREATE INDEX idx_pedidos_status
    ON pedidos(status);


CREATE INDEX idx_pedido_itens_pedido
    ON pedido_itens(pedido_id);


CREATE INDEX idx_pedido_itens_produto
    ON pedido_itens(produto_id);


-- ==========================================
-- Trigger para atualizar data de alteração
-- ==========================================

CREATE OR REPLACE FUNCTION atualizar_data_pedido()
RETURNS TRIGGER AS $$
BEGIN
    NEW.data_atualizacao = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_atualizar_data_pedido
BEFORE UPDATE ON pedidos
FOR EACH ROW
EXECUTE FUNCTION atualizar_data_pedido();