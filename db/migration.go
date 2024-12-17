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

	bots := []struct {
		Name              string
		Arch              string
		CompatibilityType string
		Quantization      string
		MaxContextLength  int
	}{
		{
			Name:              "Llama-3.3-70B-Instruct-IQ2_M.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "IQ2_M",
			MaxContextLength:  32768,
		},
		{
			Name:              "Codestral-22B-v0.1-Q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Qwen2.5-Coder-32B-Instruct.i1-Q5_K_M.gguf",
			Arch:              "gwen2",
			CompatibilityType: "gguf",
			Quantization:      "Q5_K_M",
			MaxContextLength:  32768,
		},
		{
			Name:              "c4ai-command-r-08-2024-Q5_K_M.gguf",
			Arch:              "command-r",
			CompatibilityType: "gguf",
			Quantization:      "Q5_K_M",
			MaxContextLength:  32768,
		},
		{
			Name:              "gemma-2-27b-it-Q6_K.gguf",
			Arch:              "gemma2",
			CompatibilityType: "gguf",
			Quantization:      "Q6_K",
			MaxContextLength:  32768,
		},
		{
			Name:              "GritLM-8x7B.i1-IQ3_M.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "IQ3_M",
			MaxContextLength:  32768,
		},
		{
			Name:              "internlm2_5-20b-chat.Q8_0.gguf",
			Arch:              "internlm2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "aya-23-35B.i1-IQ4_XS.gguf",
			Arch:              "command-r",
			CompatibilityType: "gguf",
			Quantization:      "IQ4_XS",
			MaxContextLength:  32768,
		},
		{
			Name:              "Yi-1.5-34B-Chat-16K.IQ4_XS.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "IQ4_XS",
			MaxContextLength:  32768,
		},
		{
			Name:              "WizardCoder-33B-V1.1.i1-IQ4_XS.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "IQ4_XS",
			MaxContextLength:  32768,
		},
		{
			Name:              "DeepSeek-Coder-V2-Lite-Instruct-Q8_0.gguf",
			Arch:              "deepseek2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Qwen2.5-Coder-14B-Instruct-Q8_0.gguf",
			Arch:              "gwen2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Virtuoso-Small-Q8_0.gguf",
			Arch:              "qwen2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "phi-4-Q8_0.gguf",
			Arch:              "phi3",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "TowerInstruct-13B-v0.1.Q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "vicuna-13b-v1.5-16k-Q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Mistral-Nemo-Instruct-2407-Q8_0.gguf",
			Arch:              "phi3",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "stablelm-2-12b-chat-Q8_0.gguf",
			Arch:              "stablelm",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Fimbulvetr-11B-Ultra-Quality-plus-Q8_0-imat.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "phi3.1-medium-Q6_K.gguf",
			Arch:              "phi3",
			CompatibilityType: "gguf",
			Quantization:      "Q6_K",
			MaxContextLength:  32768,
		},
		{
			Name:              "madlad400-10b-mt-q8_0.gguf",
			Arch:              "gwen2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "gemma-2-9b-it-abliterated-Q8_0.gguf",
			Arch:              "gemma2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Yi-1.5-9B-Chat-16K-abliterated.Q8_0.gguf",
			Arch:              "phi3",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "aya-23-8B.Q8_0.gguf",
			Arch:              "phi3",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Poppy_Porpoise-v0.7-L3-8B.Q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "OpenCoder-8B-Instruct-Q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Qwen2.5-7B-Instruct_Q8_0.gguf",
			Arch:              "gwen2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "marco-o1-q8_0.gguf",
			Arch:              "qwen2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "llava-v1.5-7b-Q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "falcon-mamba-7b-instruct-Q8_0.gguf",
			Arch:              "mamba",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Mistral-7B-Instruct-v0.3.Q8_0.gguf",
			Arch:              "phi3",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "llava-v1.6-mistral-7b.Q8_0.gguf",
			Arch:              "t5",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Phi-3.5-mini-instruct.Q8_0.gguf",
			Arch:              "phi3",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Qwen2.5-3B-Instruct-abliterated-Evol-CoT_SLERP.Q8_0.gguf",
			Arch:              "gwen2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "qwen2.5-coder-3b-instruct-q8_0.gguf",
			Arch:              "gwen2",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Llama-Doctor-3.2-3B-Instruct.Q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "Hermes-3-Llama-3.2-3B.Q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "stable-code-instruct-3b-Q8_0.gguf",
			Arch:              "stablelm",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "whisper-large-v3-candle-q8_0.gguf",
			Arch:              "phi3",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "gemma-2-2b-it-IQ4_XS.gguf",
			Arch:              "gemma2",
			CompatibilityType: "gguf",
			Quantization:      "IQ4_XS",
			MaxContextLength:  32768,
		},
		{
			Name:              "llama-3.2-1b-instruct-q8_0.gguf",
			Arch:              "llama",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
		{
			Name:              "flan-t5-large-grammar-synthesis-Q8_0.gguf",
			Arch:              "t5",
			CompatibilityType: "gguf",
			Quantization:      "Q8_0",
			MaxContextLength:  32768,
		},
	}

	for _, bot := range bots {
		_, err := db.Exec(
			"INSERT INTO bots (name, arch, compatibility_type, quantization, max_context_length) VALUES (?, ?, ?, ?, ?)",
			bot.Name, bot.Arch, bot.CompatibilityType, bot.Quantization, bot.MaxContextLength,
		)
		if err != nil {
			return fmt.Errorf("failed to insert bot %s: %w", bot.Name, err)
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
