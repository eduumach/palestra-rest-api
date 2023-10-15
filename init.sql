CREATE TABLE IF NOT EXISTS produtos
(
    id        SERIAL PRIMARY KEY,
    nome      TEXT,
    descricao TEXT,
    preco     FLOAT,
    vendido   BOOLEAN DEFAULT FALSE
);


INSERT INTO produtos (id, nome, descricao, preco, vendido)
values (1, 'Banana', 'Fruta Banana', 5.99, true);

INSERT INTO produtos (id, nome, descricao, preco, vendido)
values (2, 'Mamao', 'Fruta Mamao', 5.99. false);