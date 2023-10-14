from flask import Flask, request, jsonify
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

# Dados de exemplo (simulando um banco de dados)
produtos = [
    {"id": 1, "nome": "Produto 1", "descricao": "Descrição do Produto 1", "preco": 19.99},
    {"id": 2, "nome": "Produto 2", "descricao": "Descrição do Produto 2", "preco": 29.99},
]

# Rota para listar todos os produtos (método GET)
@app.route('/produtos', methods=['GET'])
def listar_produtos():
    return jsonify(produtos)

# Rota para obter detalhes de um produto por ID (método GET)
@app.route('/produtos/<int:id>', methods=['GET'])
def obter_produto(id):
    produto = next((p for p in produtos if p["id"] == id), None)
    if produto is None:
        return jsonify({"mensagem": "Produto não encontrado"}), 404
    return jsonify(produto)

# Rota para criar um novo produto (método POST)
@app.route('/produtos', methods=['POST'])
def criar_produto():
    novo_produto = request.json
    novo_produto["id"] = len(produtos) + 1
    produtos.append(novo_produto)
    return jsonify({"mensagem": "Produto criado com sucesso"}), 201

# Rota para atualizar um produto por ID (método PUT)
@app.route('/produtos/<int:id>', methods=['PUT'])
def atualizar_produto(id):
    produto = next((p for p in produtos if p["id"] == id), None)
    if produto is None:
        return jsonify({"mensagem": "Produto não encontrado"}), 404
    dados_atualizados = request.json
    produto.update(dados_atualizados)
    return jsonify({"mensagem": "Produto atualizado com sucesso"})

# Rota para excluir um produto por ID (método DELETE)
@app.route('/produtos/<int:id>', methods=['DELETE'])
def excluir_produto(id):
    produto = next((p for p in produtos if p["id"] == id), None)
    if produto is None:
        return jsonify({"mensagem": "Produto não encontrado"}), 404
    produtos.remove(produto)
    return jsonify({"mensagem": "Produto excluído com sucesso"})

if __name__ == '__main__':
    app.run(debug=True)
