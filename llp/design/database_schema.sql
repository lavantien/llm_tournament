-- LLP (Lightweight LLM Benchmarking) Database Schema
-- Version: 1.0.0
-- Last Updated: December 2024
-- Enable foreign key support
PRAGMA foreign_keys = ON;

-- Models Table: Track individual LLM models
CREATE TABLE models (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE,
  file_path TEXT,
  context_length INTEGER DEFAULT 8192,
  gpu_layers INTEGER,
  model_type TEXT, -- e.g., Instruct, Chat, Code
  quantization_type TEXT, -- e.g., Q4_K_M, Q5_K_M
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  last_benchmarked_at DATETIME
);

-- Task Categories Table: Define benchmark task categories
CREATE TABLE task_categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE,
  description TEXT,
  max_total_points INTEGER NOT NULL
);

-- Prompts Table: Store detailed prompt information
CREATE TABLE prompts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  category_id INTEGER NOT NULL,
  title TEXT NOT NULL,
  prompt_text TEXT NOT NULL,
  complexity_level INTEGER,
  max_points INTEGER NOT NULL,
  system_prompt TEXT,
  FOREIGN KEY (category_id) REFERENCES task_categories (id)
);

-- Benchmark Results Table: Track individual benchmark runs
CREATE TABLE benchmark_results (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  model_id INTEGER NOT NULL,
  prompt_id INTEGER NOT NULL,
  execution_timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
  -- Performance Metrics
  speed_tokens_per_sec REAL,
  total_tokens INTEGER,
  time_to_first_token REAL,
  output_stop_reason TEXT,
  -- Scoring
  points_earned REAL,
  manual_adjustment REAL DEFAULT 0,
  -- Output Storage
  output_file_path TEXT,
  -- Evaluation Flags
  is_valid BOOLEAN DEFAULT 1,
  needs_review BOOLEAN DEFAULT 0,
  FOREIGN KEY (model_id) REFERENCES models (id),
  FOREIGN KEY (prompt_id) REFERENCES prompts (id)
);

-- Tournament Aggregate Results Table
CREATE TABLE tournament_standings (
  model_id INTEGER PRIMARY KEY,
  total_score REAL DEFAULT 0,
  booting_score REAL DEFAULT 0,
  programming_score REAL DEFAULT 0,
  general_score REAL DEFAULT 0,
  agi_score REAL DEFAULT 0,
  writing_score REAL DEFAULT 0,
  swe_score REAL DEFAULT 0,
  FOREIGN KEY (model_id) REFERENCES models (id)
);

-- Indexes for Performance Optimization
CREATE INDEX idx_model_name ON models (name);

CREATE INDEX idx_prompt_category ON prompts (category_id);

CREATE INDEX idx_benchmark_model ON benchmark_results (model_id);

CREATE INDEX idx_benchmark_prompt ON benchmark_results (prompt_id);

-- Initial Data Population
-- Task Categories
INSERT INTO
  task_categories (name, description, max_total_points)
VALUES
  (
    'Booting',
    'Initial model loading and performance',
    100
  ),
  ('Programming', 'Coding and technical tasks', 1200),
  (
    'General',
    'General knowledge and reasoning',
    1145
  ),
  (
    'AGI Probing',
    'Advanced cognitive evaluation',
    630
  ),
  (
    'Creative Writing',
    'Narrative and creative tasks',
    1425
  ),
  (
    'Software Engineering',
    'Software development tasks',
    500
  );
