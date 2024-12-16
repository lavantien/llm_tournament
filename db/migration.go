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

func getDBPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}
	return filepath.Join(homeDir, ".llp", dbName)
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
	ensureDBDir()
	dbPath := getDBPath()

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

	categories := []string{"Booting", "Programming", "General", "AGI Probing", "Creative Writing", "Software Engineering"}
	for _, category := range categories {
		_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", category)
		if err != nil {
			return fmt.Errorf("failed to insert category %s: %w", category, err)
		}
	}

	if err := seedPrompts(db); err != nil {
		return fmt.Errorf("failed to seed prompts: %w", err)
	}

	return nil
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

func GetCurrentTimestamp() time.Time {
	return time.Now()
}
