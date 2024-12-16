CREATE TABLE IF NOT EXISTS system_prompts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  content TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE,
  system_prompt_id INTEGER,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (system_prompt_id) REFERENCES system_prompts (id)
);

CREATE TABLE IF NOT EXISTS prompts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  category_id INTEGER NOT NULL,
  content TEXT NOT NULL,
  system_prompt_id INTEGER,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES categories (id),
  FOREIGN KEY (system_prompt_id) REFERENCES system_prompts (id)
);

CREATE TABLE IF NOT EXISTS scores (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  bot_id INTEGER NOT NULL,
  prompt_id INTEGER NOT NULL,
  score INTEGER NOT NULL,
  speed REAL NOT NULL,
  output_path TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (bot_id) REFERENCES bots (id),
  FOREIGN KEY (prompt_id) REFERENCES prompts (id),
  UNIQUE (bot_id, prompt_id)
);

CREATE VIRTUAL TABLE IF NOT EXISTS bots_fts USING fts5 (
  name,
  arch,
  compatibility_type,
  quantization,
  content = 'bots'
);

CREATE TRIGGER IF NOT EXISTS bots_ai AFTER INSERT ON bots BEGIN
INSERT INTO
  bots_fts (
    rowid,
    name,
    arch,
    compatibility_type,
    quantization
  )
VALUES
  (
    new.id,
    new.name,
    new.arch,
    new.compatibility_type,
    new.quantization
  );

END;

CREATE TRIGGER IF NOT EXISTS bots_ad AFTER DELETE ON bots BEGIN
INSERT INTO
  bots_fts (bots_fts, rowid)
VALUES
  ('delete', old.id);

END;

CREATE TRIGGER IF NOT EXISTS bots_au AFTER
UPDATE ON bots BEGIN
INSERT INTO
  bots_fts (
    bots_fts,
    rowid,
    name,
    arch,
    compatibility_type,
    quantization
  )
VALUES
  (
    'update',
    old.id,
    new.name,
    new.arch,
    new.compatibility_type,
    new.quantization
  );

END;

CREATE VIRTUAL TABLE IF NOT EXISTS prompts_fts USING fts5 (content, content = 'prompts');

CREATE TRIGGER IF NOT EXISTS prompts_ai AFTER INSERT ON prompts BEGIN
INSERT INTO
  prompts_fts (rowid, content)
VALUES
  (new.id, new.content);

END;

CREATE TRIGGER IF NOT EXISTS prompts_ad AFTER DELETE ON prompts BEGIN
INSERT INTO
  prompts_fts (prompts_fts, rowid)
VALUES
  ('delete', old.id);

END;

CREATE TRIGGER IF NOT EXISTS prompts_au AFTER
UPDATE ON prompts BEGIN
INSERT INTO
  prompts_fts (prompts_fts, rowid, content)
VALUES
  ('update', old.id, new.content);

END;
