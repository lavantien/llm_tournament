import sqlite3

class Database:
    def __init__(self, db_name='llm_benchmark.db'):
        self.conn = sqlite3.connect(db_name)
        self.create_tables()

    def create_tables(self):
        with self.conn:
            self.conn.execute('''
                CREATE TABLE IF NOT EXISTS models (
                    id INTEGER PRIMARY KEY,
                    name TEXT NOT NULL
                )
            ''')
            self.conn.execute('''
                CREATE TABLE IF NOT EXISTS profiles (
                    id INTEGER PRIMARY KEY,
                    name TEXT NOT NULL
                )
            ''')
            self.conn.execute('''
                CREATE TABLE IF NOT EXISTS prompts (
                    id INTEGER PRIMARY KEY,
                    content TEXT NOT NULL,
                    solution TEXT NOT NULL,
                    profile TEXT NOT NULL,
                    category TEXT NOT NULL
                )
            ''')
            self.conn.execute('''
                CREATE TABLE IF NOT EXISTS scores (
                    id INTEGER PRIMARY KEY,
                    model_id INTEGER,
                    prompt_id INTEGER,
                    score INTEGER,
                    FOREIGN KEY(model_id) REFERENCES models(id),
                    FOREIGN KEY(prompt_id) REFERENCES prompts(id)
                )
            ''')

    def execute_query(self, query, params=()):
        with self.conn:
            cursor = self.conn.execute(query, params)
            return cursor.fetchall()

    def execute_non_query(self, query, params=()):
        with self.conn:
            self.conn.execute(query, params)
