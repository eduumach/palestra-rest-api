import asyncio
import json

import tornado

from database import DataBase

db = DataBase()


class MainHandler(tornado.web.RequestHandler):
    def get(self):
        result = db.make_query("SELECT * FROM produtos", fetch_all=True)
        self.set_header("Backend-Type", "Tornado")
        self.set_header("Content-Type", "text/json; charset=UTF-8")
        self.write(json.dumps(result))


def make_app():
    return tornado.web.Application([
        (r"/produtos", MainHandler),
    ])


async def main():
    app = make_app()
    app.listen(8888)
    await asyncio.Event().wait()


if __name__ == "__main__":
    asyncio.run(main())
