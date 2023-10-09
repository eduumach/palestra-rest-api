const axios = require('axios');

// Listar todos os produtos
listarProdutos();

// Obter detalhes de um produto por ID
obterDetalhesDoProduto(1);

// Criar um novo produto
const novoProduto = {
    nome: "Novo Produto",
    descricao: "Descrição do Novo Produto",
    preco: 39.99
};
criarProduto(novoProduto);

// Atualizar um produto existente (Substituir o ID pelo ID do produto real)
const produtoAtualizado = {
    id: 1,
    nome: "Produto Atualizado",
    descricao: "Descrição do Produto Atualizado",
    preco: 49.99
};
atualizarProduto(produtoAtualizado);

// Excluir um produto por ID (Substituir o ID pelo ID do produto real)
excluirProduto(2);

function listarProdutos() {
    axios.get('http://localhost:5000/produtos')
        .then(response => {
            console.log("Lista de Produtos:");
            console.log(response.data);
        })
        .catch(error => {
            console.error("Erro ao listar produtos:", error);
        });
}

function obterDetalhesDoProduto(id) {
    axios.get(`http://localhost:5000/produtos/${id}`)
        .then(response => {
            console.log("Detalhes do Produto:");
            console.log(response.data);
        })
        .catch(error => {
            console.error("Erro ao obter detalhes do produto:", error);
        });
}

function criarProduto(produto) {
    axios.post('http://localhost:5000/produtos', produto)
        .then(() => {
            console.log("Novo Produto criado com sucesso!");
        })
        .catch(error => {
            console.error("Erro ao criar novo produto:", error);
        });
}

function atualizarProduto(produto) {
    axios.put(`http://localhost:5000/produtos/${produto.id}`, produto)
        .then(() => {
            console.log("Produto Atualizado com sucesso!");
        })
        .catch(error => {
            console.error("Erro ao atualizar produto:", error);
        });
}

function excluirProduto(id) {
    axios.delete(`http://localhost:5000/produtos/${id}`)
        .then(() => {
            console.log("Produto Excluído com sucesso!");
        })
        .catch(error => {
            console.error("Erro ao excluir produto:", error);
        });
}
