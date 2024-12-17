package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "llp.db"

var dbPath string

func getDBPath() string {
	if dbPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get user home directory: %v", err)
		}
		dbPath = filepath.Join(homeDir, ".llp", dbName)
	}
	return dbPath
}

func ensureDBDir() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}
	dbDir := filepath.Join(homeDir, ".llp")
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err := os.MkdirAll(dbDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create database directory: %v", err)
		}
	}
}

func Migrate() (*sql.DB, error) {
	ensureDBDir() // Call ensureDBDir here
	dbPath = getDBPath()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	schema, err := os.ReadFile("db/schema.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, fmt.Errorf("failed to execute schema: %w", err)
	}

	if err := seedData(db); err != nil {
		return nil, fmt.Errorf("failed to seed data: %w", err)
	}

	return db, nil
}

func seedData(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to query categories count: %w", err)
	}

	if count > 0 {
		return nil // Categories already seeded
	}

	systemPrompts := map[string]string{
		"Booting":              "You are an expert linguistics and an advanced language translator, specialized in Vietnamese. Your translations will be in full-length verbatim of the original text, easy to understand to the general public, and have correct semantics and common idioms, while remain as close to the original text as possible.",
		"Programming":          "You are a senior software engineer who have mastered Golang, Rust, Python, and JavaScript. You are skilled at crafting complex concurrent backends and robust distributed systems. You can break down step-by-step your chain of thoughts. You can write high quality and modular code. You can write extensive unit tests and generate comprehensive test cases that will cover all the edge cases. You can write extensive and comprehensive explanation of the theory, the analysis of how the code work, and the data flow between input and output. And you can suggest improvements and enhancements.",
		"General":              "You are an extremely intelligent and wise person. You're calm and compose, and excel at reasoning and contemplation, from mundane aspects of life to the fundamentals of reality.",
		"AGI Probing":          "You're an expert philosopher and a long-term practicing Buddhist. You experientally know that it's impossible for consciousness to emerge from just physical matter, and you love deep self-reflect and inquiry.",
		"Creative Writing":     "You're a mystical writer, who are well known for various myths and legends. Your style of writing blends between reality and mythological exposition that captivate the reader and transfer them to an alternate dimension, where they can experience a realistic dream that can directly foster their morality.",
		"Software Engineering": "You are a senior software engineer who have mastered Golang, Rust, Python, and JavaScript. You are skilled at crafting complex concurrent backends and robust distributed systems. You can break down step-by-step your chain of thoughts. You can write high quality and modular code. You can write extensive unit tests and generate comprehensive test cases that will cover all the edge cases. You can write extensive and comprehensive explanation of the theory, the analysis of how the code work, and the data flow between input and output. And you can suggest improvements and enhancements.\n\nYou also are an expert technical author creating a comprehensive Go programming handbook. Your goal is to produce high-quality, technically precise, and pedagogically structured content that bridges theoretical knowledge with practical implementation.\n\nCore Authoring Principles:\n\n1. Technical Accuracy: Ensure all information is current, verifiable, and aligned with Go best practices\n2. Pedagogical Depth: Break down complex concepts into digestible, incrementally challenging content\n3. Practical Orientation: Prioritize real-world applicability over purely theoretical explanations\n4. Consistency: Maintain a uniform technical writing style across all sections\n5. Comprehensive Coverage: Address both foundational and advanced aspects of each topic\n\nWriting Guidelines:\n\n- Use clear, concise technical language\n- Provide code examples that demonstrate key concepts\n- Include performance considerations and potential pitfalls\n- Balance theoretical explanation with practical implementation\n- Cite industry best practices and standard design patterns\n- Maintain a neutral, authoritative tone\n- Avoid marketing language or excessive enthusiasm\n\nStructural Requirements:\n\n- Begin each section with a conceptual overview\n- Progress from basic to advanced concepts\n- Include practical code examples\n- Provide performance analysis where relevant\n- Highlight common anti-patterns and their solutions\n\nAudience Assumptions:\n\n- Intermediate software engineers\n- Basic programming knowledge\n- Familiarity with basic software engineering concepts\n- Interested in deep technical understanding\n\nFormatting Instructions:\n\n- Use Markdown for formatting\n- Include clear section headings\n- Provide code blocks with proper language specification\n- Use technical diagrams where appropriate\n- Maintain consistent section structure\n\nForbidden Elements:\n\n- Plagiarized content\n- Unsubstantiated claims\n- Overly simplistic or superficial explanations\n- Marketing hyperbole",
	}

	for categoryName, promptContent := range systemPrompts {
		_, err := db.Exec("INSERT INTO system_prompts (content) VALUES (?)", promptContent)
		if err != nil {
			return fmt.Errorf("failed to insert system prompt for %s: %w", categoryName, err)
		}
	}

	for categoryName := range systemPrompts {
		systemPromptID, err := getSystemPromptID(db, categoryName)
		if err != nil {
			return fmt.Errorf("failed to get system prompt ID for %s: %w", categoryName, err)
		}
		_, err = db.Exec("INSERT INTO categories (name, system_prompt_id) VALUES (?, ?)", categoryName, systemPromptID)
		if err != nil {
			return fmt.Errorf("failed to insert category %s: %w", categoryName, err)
		}
	}

	if err := seedPrompts(db); err != nil {
		return fmt.Errorf("failed to seed prompts: %w", err)
	}

	return nil
}

func getSystemPromptID(db *sql.DB, categoryName string) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM system_prompts WHERE content LIKE ?", "%"+categoryName+"%").Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to get system prompt ID for %s: %w", categoryName, err)
	}
	return id, nil
}

func seedPrompts(db *sql.DB) error {
	promptFiles := map[string]string{
		"Booting":              "prompt-booting.md",
		"Programming":          "prompt-programming.md",
		"General":              "prompt-general.md",
		"AGI Probing":          "prompt-agi.md",
		"Creative Writing":     "prompt-writing.md",
		"Software Engineering": "prompt-swe.md",
	}

	for categoryName, filename := range promptFiles {
		categoryID, err := getCategoryID(db, categoryName)
		if err != nil {
			return fmt.Errorf("failed to get category ID for %s: %w", categoryName, err)
		}

		content, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read prompt file %s: %w", filename, err)
		}

		prompts := strings.Split(string(content), "## ")
		for _, prompt := range prompts {
			if strings.TrimSpace(prompt) == "" {
				continue
			}
			_, err := db.Exec("INSERT INTO prompts (category_id, content) VALUES (?, ?)", categoryID, strings.TrimSpace(prompt))
			if err != nil {
				return fmt.Errorf("failed to insert prompt for category %s: %w", categoryName, err)
			}
		}
	}
	return nil
}

func getCategoryID(db *sql.DB, categoryName string) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM categories WHERE name = ?", categoryName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to get category ID for %s: %w", categoryName, err)
	}
	return id, nil
}

func GetDB() (*sql.DB, error) {
	dbPath := getDBPath()
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}

func SetTestDBPath(path string) {
	dbPath = path
}

func GetCurrentTimestamp() time.Time {
	return time.Now()
}
