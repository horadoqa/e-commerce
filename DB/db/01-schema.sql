-- ==========================================
-- Configuração inicial do banco ecommerce
-- ==========================================

-- Extensão para geração de UUID
CREATE EXTENSION IF NOT EXISTS "pgcrypto";


-- ==========================================
-- Schema principal da aplicação
-- ==========================================

CREATE SCHEMA IF NOT EXISTS ecommerce;


-- Define o schema padrão para a aplicação
ALTER DATABASE ecommerce
SET search_path TO ecommerce, public;


-- ==========================================
-- Função genérica para atualização de data
-- ==========================================

CREATE OR REPLACE FUNCTION atualizar_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;