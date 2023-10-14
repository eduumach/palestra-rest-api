CREATE TABLE IF NOT EXISTS produtos
(
    id        SERIAL PRIMARY KEY,
    name      TEXT,
    descricao TEXT,
    preco     FLOAT
);


INSERT INTO produtos (id, name, descricao, preco)
values (1, 'Banana', 'Fruta Banana', 5.99);

INSERT INTO produtos (id, name, descricao, preco)
values (2, 'Mamao', 'Fruta Mamao', 5.99);