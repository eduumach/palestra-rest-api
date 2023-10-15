"""
Singleton class encapsulating a psycopg2 connection.
"""
from typing import Dict

import psycopg2
from psycopg2.extras import RealDictCursor

psycopg2.extensions.register_type(psycopg2.extensions.UNICODE)
psycopg2.extensions.register_type(psycopg2.extensions.UNICODEARRAY)


class DataBase(object):
    """Borg pattern singleton"""
    __state = {}

    def __init__(self):
        self.__dict__ = self.__state
        if not hasattr(self, 'conn'):
            self.conn = psycopg2.connect(dbname='postgres',
                                         user='postgres',
                                         password='postgres',
                                         host='pgbouncer',
                                         port='6432'

                                         )
            self.conn.autocommit = True
            self.cur = self.conn.cursor(cursor_factory=RealDictCursor)

    def make_query(self, query: str, params: Dict = None, fetch_all: bool = False):
        try:
            self.cur.execute(query, params)
        except psycopg2.Error as e:
            print(f"An error Ocurred {e}")

        if fetch_all:
            values = self.cur.fetchall()

            return values


if __name__ == '__main__':
    db = DataBase()
    result = db.make_query("SELECT satellite_image_id, ST_AsGeoJSON(geometry) FROM satellite_images", fetch_all=True)

    print(result)
