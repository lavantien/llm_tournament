# Detailed Design Document

The LLP Application is a terminal-based utility designed to manage and track Large Language Model (LLM) benchmarking results with minimal computational overhead.

## 1. System Architecture

### 1.1 Technology Stack

- **Programming Language**: Golang 1.23
- **UI Framework**: BubbleTea & BubbleZone
- **Database**: SQLite 3.47
- **Diagramming**: PlantUML

### 1.2 High-Level Component Structure

1. **Data Management Layer**

   - SQLite database handler
   - Data models for LLMs, tasks, and results
   - CRUD operations for benchmark data

2. **UI Rendering Layer**

   - Terminal User Interface (TUI) components
   - State management
   - User interaction handlers

3. **Benchmarking Ingestion Layer**
   - Result parsing
   - Performance metric extraction
   - Scoring calculation

## 2. Schema Design Document

Schema Design Highlights:

- Comprehensive tracking of LLM models and benchmark results
- Flexible scoring system with manual adjustment capabilities
- Performance optimization through strategic indexing
- Supports multiple task categories and detailed prompt management
- Captures rich metadata about benchmark executions

Database Initialization Approach:

- Use the schema script to create initial database structure
- Implement Go structs for type-safe database interactions

```go
package models

import "time"

type Model struct {
    ID              int64     `db:"id"`
    Name            string    `db:"name"`
    FilePath        string    `db:"file_path"`
    ContextLength   int       `db:"context_length"`
    GPULayers       int       `db:"gpu_layers"`
    ModelType       string    `db:"model_type"`
    Quantization    string    `db:"quantization_type"`
    CreatedAt       time.Time `db:"created_at"`
    LastBenchmarked time.Time `db:"last_benchmarked_at"`
}

type TaskCategory struct {
    ID              int64  `db:"id"`
    Name            string `db:"name"`
    Description     string `db:"description"`
    MaxTotalPoints  int    `db:"max_total_points"`
}

type Prompt struct {
    ID              int64  `db:"id"`
    CategoryID      int64  `db:"category_id"`
    Title           string `db:"title"`
    PromptText      string `db:"prompt_text"`
    ComplexityLevel int    `db:"complexity_level"`
    MaxPoints       int    `db:"max_points"`
    SystemPrompt    string `db:"system_prompt"`
}

type BenchmarkResult struct {
    ID                  int64     `db:"id"`
    ModelID             int64     `db:"model_id"`
    PromptID            int64     `db:"prompt_id"`
    ExecutionTimestamp  time.Time `db:"execution_timestamp"`

    SpeedTokensPerSec   float64   `db:"speed_tokens_per_sec"`
    TotalTokens         int       `db:"total_tokens"`
    TimeToFirstToken    float64   `db:"time_to_first_token"`
    OutputStopReason    string    `db:"output_stop_reason"`

    PointsEarned        float64   `db:"points_earned"`
    ManualAdjustment    float64   `db:"manual_adjustment"`

    OutputFilePath      string    `db:"output_file_path"`

    IsValid             bool      `db:"is_valid"`
    NeedsReview         bool      `db:"needs_review"`
}

type TournamentStanding struct {
    ModelID             int64   `db:"model_id"`
    TotalScore          float64 `db:"total_score"`
    BootingScore        float64 `db:"booting_score"`
    ProgrammingScore    float64 `db:"programming_score"`
    GeneralScore        float64 `db:"general_score"`
    AGIScore            float64 `db:"agi_score"`
    WritingScore        float64 `db:"writing_score"`
    SWEScore            float64 `db:"swe_score"`
}
```

## 3. Database Schema

### 3.1 Models Table

```sql
CREATE TABLE models (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    file_path TEXT,
    context_length INTEGER,
    gpu_layers INTEGER
);
```

### 3.2 Tasks Table

```sql
CREATE TABLE tasks (
    id INTEGER PRIMARY KEY,
    category TEXT NOT NULL,
    complexity_level INTEGER,
    max_points INTEGER,
    prompt TEXT
);
```

### 3.3 Results Table

```sql
CREATE TABLE results (
    id INTEGER PRIMARY KEY,
    model_id INTEGER,
    task_id INTEGER,
    speed REAL,
    tokens INTEGER,
    time_to_first_token REAL,
    points_earned INTEGER,
    output_path TEXT,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(model_id) REFERENCES models(id),
    FOREIGN KEY(task_id) REFERENCES tasks(id)
);
```

### Detailed Schema

- [`database_schema.sql`](./database_schema.sql)

## 4. Application Workflows

### 4.1 Result Ingestion Workflow

1. Select task category
2. Choose model
3. Input performance metrics
4. Paste output
5. Validate and store results
6. Calculate and update points

### 4.2 Export Workflow

1. Select export format (Markdown/CSV/JSON)
2. Choose export scope
3. Select metrics
4. Generate report
5. Save to specified location

## 5. Performance Considerations

- Minimal memory footprint
- Quick database operations
- Efficient UI rendering
- Non-blocking background processes

## 6. Error Handling Strategy

- Comprehensive input validation
- Graceful error messages
- Logging of critical errors
- Ability to recover or rollback operations

## 7. Security Considerations

- Sanitize all user inputs
- Restrict file system access
- Implement read-only modes for sensitive operations

## 8. Future Extensibility

- Plugin system for custom task types
- Advanced visualization capabilities
- Cloud/remote result synchronization

## 9. Development Workflow

1. Initialize project structure
2. Set up SQLite schema
3. Develop core data models
4. Implement BubbleTea UI components
5. Create benchmarking ingestion logic
6. Develop export and reporting features
7. Comprehensive testing

## 10. Testing Strategy

- Unit tests for each component
- Integration tests for workflow
- Performance benchmarking of the application itself
- Edge case and error handling verification

## 11. Deployment Considerations

- Single binary distribution
- Minimal external dependencies
- Cross-platform compatibility (Linux, macOS, Windows)
- Optional system-wide installation script

---

**Version**: 0.1.0
**Last Updated**: December 2024
**Status**: Initial Design
