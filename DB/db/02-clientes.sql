CREATE TABLE clientes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    nome VARCHAR(150) NOT NULL,
    cpf CHAR(11) UNIQUE NOT NULL,

    email VARCHAR(255) UNIQUE NOT NULL,
    telefone VARCHAR(20),

    data_nascimento DATE NOT NULL,

    endereco VARCHAR(255) NOT NULL,
    numero VARCHAR(20) NOT NULL,
    complemento VARCHAR(100),

    bairro VARCHAR(100) NOT NULL,
    cidade VARCHAR(100) NOT NULL,
    estado CHAR(2) NOT NULL,
    cep CHAR(8) NOT NULL,

    ativo BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);