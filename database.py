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
                    name TEXT NOT NULL,
                    path TEXT,
                    gpu_layers INTEGER,
                    ctx_size INTEGER,
                    batch_size INTEGER,
                    threads INTEGER,
                    keep INTEGER,
                    predict INTEGER,
                    flash_attn BOOLEAN,
                    mlock BOOLEAN,
                    cache_type_k TEXT,
                    cache_type_v TEXT
                )
            ''')
            self.conn.execute('''
                CREATE TABLE IF NOT EXISTS profiles (
                    id INTEGER PRIMARY KEY,
                    name TEXT NOT NULL,
                    system_prompt TEXT,
                    dry_multiplier REAL,
                    dry_base REAL,
                    dry_allowed_length INTEGER,
                    dry_penalty_last_n INTEGER,
                    repeat_penalty REAL,
                    repeat_last_n INTEGER,
                    top_k INTEGER,
                    top_p REAL,
                    min_p REAL,
                    top_a REAL,
                    xtc_threshold REAL,
                    xtc_probability REAL,
                    temperature REAL
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

    def populate_mock_data(self):
        with self.conn:
            # Insert mock data for models
            self.conn.executemany('''
                INSERT INTO models (name, path, gpu_layers, ctx_size, batch_size, threads, keep, predict, flash_attn, mlock, cache_type_k, cache_type_v)
                VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
            ''', [
                ('Model1', '/path/to/model1', 10, 2048, 4, 8, 1, 1, True, True, 'type_k1', 'type_v1'),
                ('Model2', '/path/to/model2', 15, 4096, 8, 16, 2, 2, False, False, 'type_k2', 'type_v2'),
            ])

            # Insert mock data for profiles
            self.conn.executemany('''
                INSERT INTO profiles (name, system_prompt, dry_multiplier, dry_base, dry_allowed_length, dry_penalty_last_n, repeat_penalty, repeat_last_n, top_k, top_p, min_p, top_a, xtc_threshold, xtc_probability, temperature)
                VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
            ''', [
                ('Profile1', 'System prompt 1', 1.0, 1.0, 100, 5, 1.0, 5, 50, 0.9, 0.1, 0.5, 0.5, 0.5, 0.7),
                ('Profile2', 'System prompt 2', 1.5, 1.5, 200, 10, 1.5, 10, 100, 0.8, 0.2, 0.6, 0.6, 0.6, 0.8),
            ])

            # Insert mock data for prompts
            self.conn.executemany('''
                INSERT INTO prompts (content, solution, profile, category)
                VALUES (?, ?, ?, ?)
            ''', [
                ('Prompt content 1', 'Solution 1', 'Profile1', 'Category1'),
                ('Prompt content 2', 'Solution 2', 'Profile2', 'Category2'),
            ])

            # Insert mock data for scores
            self.conn.executemany('''
                INSERT INTO scores (model_id, prompt_id, score)
                VALUES (?, ?, ?)
            ''', [
                (1, 1, 85),
                (1, 2, 90),
                (2, 1, 80),
                (2, 2, 88),
            ])
