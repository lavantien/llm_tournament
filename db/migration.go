package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

// func ensureDBDir() {
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		log.Fatalf("Failed to get user home directory: %v", err)
// 	}
// 	dbDir := filepath.Join(homeDir, ".llp")
// 	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
// 		err := os.MkdirAll(dbDir, 0755)
// 		if err != nil {
// 			log.Fatalf("Failed to create database directory: %v", err)
// 		}
// 	}
// }

func Migrate() (*sql.DB, error) {
	// ensureDBDir() // Removed this line

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
		"Booting":              "You are an extremely versatile and intelligent person, with advanced thinking skills and the ability to form and explain your cohesive chain of thoughts.",
		"Programming":          "You are a hobbyist programmer with deep passion for programming and problem solving. You love doing recreational programming and creative programming work.",
		"General":              "You are an extremely intelligent and wise person. You're calm and compose, and excel at reasoning and contemplation, from mundane aspects of life to the fundamentals of reality.",
		"AGI Probing":          "You're an expert philosopher and a long-term practicing Buddhist. You experientally know that it's impossible for consciousness to emerge from just physical matter, and you love deep self-reflect and inquiry.",
		"Creative Writing":     "You're a mystical writer, who are well known for various myths and legends. Your style of writing blends between reality and mythological exposition that captivate the reader and transfer them to an alternate dimension, where they can experience a realistic dream that can directly foster their morality.",
		"Software Engineering": "You are a senior software engineer who have mastered Golang, Rust, Python, and JavaScript. You are skilled at crafting complex concurrent backends and robust distributed systems. You can break down step-by-step your chain of thoughts. You can write high quality and modular code. You can write extensive unit tests and generate comprehensive test cases that will cover all the edge cases. You can write extensive and comprehensive explanation of the theory, the analysis of how the code work, and the data flow between input and output. And you can suggest improvements and enhancements.",
	}

	systemPromptIDs := make(map[string]int)
	for categoryName, promptContent := range systemPrompts {
		result, err := db.Exec("INSERT INTO system_prompts (content) VALUES (?)", promptContent)
		if err != nil {
			return fmt.Errorf("failed to insert system prompt for %s: %w", categoryName, err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get last insert id for system prompt %s: %w", categoryName, err)
		}
		systemPromptIDs[categoryName] = int(id)
	}

	for categoryName, systemPromptID := range systemPromptIDs {
		_, err = db.Exec("INSERT INTO categories (name, system_prompt_id) VALUES (?, ?)", categoryName, systemPromptID)
		if err != nil {
			return fmt.Errorf("failed to insert category %s: %w", categoryName, err)
		}
	}

	prompts := map[string][]string{
		"Booting": {
			"What is the meaning of life?",
			"What is the nature of consciousness?",
		},
		"Programming": {
			"Write a function in Go that calculates the factorial of a number.",
			"Explain the concept of closures in JavaScript.",
		},
		"General": {
			"What are the main causes of climate change?",
			"Explain the theory of relativity.",
		},
		"AGI Probing": {
			"Can a machine truly understand human emotions?",
			"What are the ethical implications of artificial intelligence?",
		},
		"Creative Writing": {
			"Write a short story about a journey to another world.",
			"Compose a poem about the beauty of nature.",
		},
		"Software Engineering": {
			"Explain the SOLID principles of object-oriented design.",
			"Describe the benefits of using microservices architecture.",
		},
	}

	for categoryName, promptList := range prompts {
		categoryID, err := getCategoryID(db, categoryName)
		if err != nil {
			return fmt.Errorf("failed to get category ID for %s: %w", categoryName, err)
		}
		for _, promptContent := range promptList {
			var systemPromptID int
			if categoryName == "Booting" {
				systemPromptID = systemPromptIDs["Booting"]
			} else if categoryName == "Programming" {
				systemPromptID = systemPromptIDs["Programming"]
			} else if categoryName == "General" {
				systemPromptID = systemPromptIDs["General"]
			} else if categoryName == "AGI Probing" {
				systemPromptID = systemPromptIDs["AGI Probing"]
			} else if categoryName == "Creative Writing" {
				systemPromptID = systemPromptIDs["Creative Writing"]
			} else if categoryName == "Software Engineering" {
				systemPromptID = systemPromptIDs["Software Engineering"]
			}
			_, err := db.Exec("INSERT INTO prompts (category_id, content, system_prompt_id) VALUES (?, ?, ?)", categoryID, promptContent, systemPromptID)
			if err != nil {
				return fmt.Errorf("failed to insert prompt for category %s: %w", categoryName, err)
			}
		}
	}

	return nil
}

// func getSystemPromptID(db *sql.DB, categoryName string) (int, error) {
// 	var id int
// 	err := db.QueryRow("SELECT id FROM system_prompts WHERE content = ?", categoryName).Scan(&id)
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to get system prompt ID for %s: %w", categoryName, err)
// 	}
// 	return id, nil
// }

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
