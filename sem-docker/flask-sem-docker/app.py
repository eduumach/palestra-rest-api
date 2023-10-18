from flask import Flask, request
from flask import Response
import json


app = Flask(__name__)


def create_response(data, status_code=200):
    response = Response()
    response.data = json.dumps(data)  # Convert data to JSON
    response.status_code = status_code
    response.headers.set('Backend-Type', 'Flask')
    # Set the content type to JSON
    response.headers['Content-Type'] = 'application/json'
    return response

produtos = [{'id': 1, 'nome': 'Produto 1', 'descricao': 'Descrição 1', 'preco': 100.00, 'vendido': False},
            {'id': 2, 'nome': 'Produto 2', 'descricao': 'Descrição 2', 'preco': 200.00, 'vendido': False},
            {'id': 3, 'nome': 'Produto 3', 'descricao': 'Descrição 3', 'preco': 300.00, 'vendido': False},
            {'id': 4, 'nome': 'Produto 4', 'descricao': 'Descrição 4', 'preco': 400.00, 'vendido': False},
            {'id': 5, 'nome': 'Produto 5', 'descricao': 'Descrição 5', 'preco': 500.00, 'vendido': False}]

# Rota para listar todos os produtos (método GET)
@app.route('/produtos', methods=['GET'])
def listar_produtos():
    return create_response(produtos)


# Rota para obter detalhes de um produto por ID (método GET)
@app.route('/produtos/<int:id>', methods=['GET'])
def obter_produto(id):
    produto = next((p for p in produtos if p["id"] == id), None)
    if produto is None:
        return jsonify({"mensagem": "Produto não encontrado"}), 404
    return create_response(produto)


# Rota para criar um novo produto (método POST)
@app.route('/produtos', methods=['POST'])
def criar_produto():
    novo_produto = request.json
    nome = novo_produto['nome']
    descricao = novo_produto['descricao']
    preco = novo_produto['preco']

    produtos.append({'id': len(produtos) + 1, 'nome': nome, 'descricao': descricao, 'preco': preco, 'vendido': False})

    return create_response(novo_produto, status_code=201)


# Rota para atualizar um produto por ID (método PUT)
@app.route('/produtos/<int:id>', methods=['PUT'])
def atualizar_produto(id):
    produto_atualizado = request.json
    nome = produto_atualizado['nome']
    descricao = produto_atualizado['descricao']
    preco = produto_atualizado['preco']
    
    produto = next((p for p in produtos if p["id"] == id), None)
    if produto is None:
        return jsonify({"mensagem": "Produto não encontrado"}), 404
    dados_atualizados = {'id': id,'nome': nome, 'descricao': descricao, 'preco': preco}
    produto.update(dados_atualizados)

    return create_response(produto_atualizado)


@app.route('/produtos/<int:id>', methods=['PATCH'])
def atualizar_produto_parcial(id):
    produto_atualizado = request.json
    vendido = produto_atualizado['vendido']
    
    produto = next((p for p in produtos if p["id"] == id), None)
    if produto is None:
        return jsonify({"mensagem": "Produto não encontrado"}), 404
    dados_atualizados['vendido'] = vendido
    produto.update(dados_atualizados)
    return create_response(produto_atualizado)


# Rota para excluir um produto por ID (método DELETE)
@app.route('/produtos/<int:id>', methods=['DELETE'])
def excluir_produto(id):
    produto = next((p for p in produtos if p["id"] == id), None)
    if produto is None:
        return jsonify({"mensagem": "Produto não encontrado"}), 404
    produtos.remove(produto)
    return create_response(produto)


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)
