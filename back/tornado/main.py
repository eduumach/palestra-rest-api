import asyncio

import tornado


class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.set_header("Backend-Type", "Tornado")
        self.write("Hello, world")


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
