import os
import sqlite3

class DBManager:
    def __init__(self, db_path):
        self.db_path = db_path
        self.create_database_if_not_exists()

    def create_database_if_not_exists(self):
        if not os.path.exists(self.db_path):
            print(f"Database file does not exist. Creating new file at {self.db_path}")
            open(self.db_path, 'w').close()
            self.conn = sqlite3.connect(self.db_path)
            self.create_tables()
        else:
            try:
                self.conn = sqlite3.connect(self.db_path)
                # Check if the database is valid by querying the sqlite_master table
                self.conn.execute("SELECT name FROM sqlite_master WHERE type='table'")
            except sqlite3.DatabaseError:
                print(f"Database file is not valid. Recreating the database at {self.db_path}")
                self.conn.close()
                os.remove(self.db_path)
                open(self.db_path, 'w').close()
                self.conn = sqlite3.connect(self.db_path)
                self.create_tables()
            else:
                print(f"Database file already exists at {self.db_path}")

    def create_tables(self):
        with self.conn:
            print("Creating tables...")
            self.conn.execute('''CREATE TABLE IF NOT EXISTS prompt_suites (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                description TEXT
            )''')
            self.conn.execute('''CREATE TABLE IF NOT EXISTS prompts (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                suite_id INTEGER,
                content TEXT NOT NULL,
                category TEXT,
                FOREIGN KEY (suite_id) REFERENCES prompt_suites (id)
            )''')
            self.conn.execute('''CREATE TABLE IF NOT EXISTS llms (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                version TEXT,
                description TEXT
            )''')
            self.conn.execute('''CREATE TABLE IF NOT EXISTS attempts (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                llm_id INTEGER,
                prompt_id INTEGER,
                score REAL,
                output TEXT,
                timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
                FOREIGN KEY (llm_id) REFERENCES llms (id),
                FOREIGN KEY (prompt_id) REFERENCES prompts (id)
            )''')
            print("Tables created successfully.")

    def insert_prompt_suite(self, name, description):
        with self.conn:
            self.conn.execute('INSERT INTO prompt_suites (name, description) VALUES (?, ?)', (name, description))

    def insert_prompt(self, suite_id, content, category):
        with self.conn:
            self.conn.execute('INSERT INTO prompts (suite_id, content, category) VALUES (?, ?, ?)', (suite_id, content, category))

    def insert_llm(self, name, version, description):
        with self.conn:
            self.conn.execute('INSERT INTO llms (name, version, description) VALUES (?, ?, ?)', (name, version, description))

    def insert_attempt(self, llm_id, prompt_id, score, output):
        with self.conn:
            self.conn.execute('INSERT INTO attempts (llm_id, prompt_id, score, output) VALUES (?, ?, ?, ?)', (llm_id, prompt_id, score, output))

    def fetch_all_prompt_suites(self):
        with self.conn:
            return self.conn.execute('SELECT * FROM prompt_suites').fetchall()

    def fetch_all_prompts(self):
        with self.conn:
            return self.conn.execute('SELECT * FROM prompts').fetchall()

    def fetch_all_llms(self):
        with self.conn:
            return self.conn.execute('SELECT * FROM llms').fetchall()

    def fetch_all_attempts(self):
        with self.conn:
            return self.conn.execute('SELECT * FROM attempts').fetchall()

    def fetch_llm_by_id(self, llm_id):
        with self.conn:
            return self.conn.execute('SELECT * FROM llms WHERE id = ?', (llm_id,)).fetchone()
