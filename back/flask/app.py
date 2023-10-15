from flask import Flask, request, jsonify
from flask import Response
import json
import logging

from database import DataBase

app = Flask(__name__)

db = DataBase()


def create_response(data, status_code=200):
    response = Response()
    response.data = json.dumps(data)  # Convert data to JSON
    response.status_code = status_code
    response.headers.set('Backend-Type', 'Flask')
    # Set the content type to JSON
    response.headers['Content-Type'] = 'application/json'
    return response


# Rota para listar todos os produtos (método GET)
@app.route('/produtos', methods=['GET'])
def listar_produtos():
    result = db.make_query("SELECT * FROM produtos", fetch_all=True)

    return create_response(result)


# Rota para obter detalhes de um produto por ID (método GET)
@app.route('/produtos/<int:id>', methods=['GET'])
def obter_produto(id):
    result = db.make_query(
        f"SELECT * FROM produtos WHERE id={id}", fetch_all=True)
    return create_response(result[0])


# Rota para criar um novo produto (método POST)
@app.route('/produtos', methods=['POST'])
def criar_produto():
    novo_produto = request.json
    nome = novo_produto['nome']
    descricao = novo_produto['descricao']
    preco = novo_produto['preco']

    db.make_query(
        f"INSERT INTO produtos (nome, descricao, preco) VALUES ('{nome}', '{descricao}', {preco})")

    return create_response(novo_produto, status_code=201)


# Rota para atualizar um produto por ID (método PUT)
@app.route('/produtos/<int:id>', methods=['PUT'])
def atualizar_produto(id):
    produto_atualizado = request.json
    nome = produto_atualizado['nome']
    descricao = produto_atualizado['descricao']
    preco = produto_atualizado['preco']

    db.make_query(
        f"UPDATE produtos SET nome='{nome}', descricao='{descricao}', preco={preco} WHERE id={id}")

    return create_response(produto_atualizado)


@app.route('/produtos/<int:id>', methods=['PATCH'])
def atualizar_produto_parcial(id):
    produto_atualizado = request.json
    vendido = produto_atualizado['vendido']

    db.make_query(f"UPDATE produtos SET vendido={vendido} WHERE id={id}")

    return create_response(produto_atualizado)


# Rota para excluir um produto por ID (método DELETE)
@app.route('/produtos/<int:id>', methods=['DELETE'])
def excluir_produto(id):
    db.make_query(f"DELETE FROM produtos WHERE id={id}")

    return create_response('', status_code=204)


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)
