import asyncio
import json

import tornado

from database import DataBase

db = DataBase()


class MainHandler(tornado.web.RequestHandler):
    def _create_response(self, data, status_code=200):
        self.set_header('Backend-Type', 'Tornado')
        self.set_header('Content-Type', 'application/json')
        self.set_status(status_code)
        self.write(json.dumps(data))
    
    def get(self, produto_id=None):
        if produto_id is None:
            result = db.make_query("SELECT * FROM produtos", fetch_all=True)
            self._create_response(result)
        else:
            result = db.make_query(f"SELECT * FROM produtos WHERE id={produto_id}", fetch_all=True)
            self._create_response(result[0])

    def post(self):
        data = json.loads(self.request.body)
        nome = data.get("nome")
        preco = data.get("preco")
        descricao = data.get("descricao")
        db.make_query(f"INSERT INTO produtos (nome, preco, descricao) VALUES ('{nome}', {preco}, '{descricao}')")
        self._create_response(data, status_code=201)

    def put(self, produto_id):
        data = json.loads(self.request.body)
        nome = data.get("nome")
        preco = data.get("preco")
        descricao = data.get("descricao")
        db.make_query(f"UPDATE produtos SET nome='{nome}', preco={preco}, descricao='{descricao}' WHERE id={produto_id}")
        self._create_response(data)
    
    def patch(self, produto_id):
        data = json.loads(self.request.body)
        vendido = data.get("vendido")
        db.make_query(f"UPDATE produtos SET vendido={vendido} WHERE id={produto_id}")
        self._create_response(data)

    def delete(self, produto_id):
        db.make_query(f"DELETE FROM produtos WHERE id={produto_id}")
        self._create_response({})


def make_app():
    return tornado.web.Application([
        (r"/produtos", MainHandler),
        (r"/produtos/(\d+)", MainHandler),
    ])


async def main():
    app = make_app()
    app.listen(8888)
    await asyncio.Event().wait()


if __name__ == "__main__":
    asyncio.run(main())
