package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Set up HTTP handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/load", loadDataHandler(db))
	http.HandleFunc("/clear", clearDataHandler(db))
	http.HandleFunc("/generate", generateScoresHandler(db))

	// Start the server
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func initDB() (*sql.DB, error) {
	// Remove existing database file if it exists
	if _, err := os.Stat("llm-tournament.db"); err == nil {
		os.Remove("llm-tournament.db")
	}

	// Open a new database connection
	db, err := sql.Open("sqlite3", "llm-tournament.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Create tables
	_, err = db.Exec(`
        CREATE TABLE bots (
            name TEXT PRIMARY KEY,
            path TEXT,
            size REAL,
            param REAL,
            quant TEXT,
            gpuLayers INTEGER,
            gpuLayersUsed INTEGER,
            ctx INTEGER,
            ctxUsed INTEGER,
            kingOf TEXT,
            FOREIGN KEY (kingOf) REFERENCES profiles(name)
        );

        CREATE TABLE profiles (
            name TEXT PRIMARY KEY,
            systemPrompt TEXT,
            repeatPenalty REAL,
            topK INTEGER,
            topP REAL,
            minP REAL,
            topA REAL,
            bestBots TEXT
        );

        CREATE TABLE prompts (
            number INTEGER PRIMARY KEY,
            content TEXT,
            solution TEXT,
            profile TEXT,
            FOREIGN KEY (profile) REFERENCES profiles(name)
        );

        CREATE TABLE scores (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            attempt INTEGER,
            elo REAL,
            botId TEXT,
            promptId INTEGER,
            FOREIGN KEY (botId) REFERENCES bots(name),
            FOREIGN KEY (promptId) REFERENCES prompts(number)
        );
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}

	return db, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the LLM Tournament App!")
}

func loadDataHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := loadData(db); err != nil {
			http.Error(w, fmt.Sprintf("Failed to load  %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Data loaded successfully!")
	}
}

func clearDataHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := db.Exec("DELETE FROM bots"); err != nil {
			http.Error(w, fmt.Sprintf("Failed to clear bots  %v", err), http.StatusInternalServerError)
			return
		}
		if _, err := db.Exec("DELETE FROM profiles"); err != nil {
			http.Error(w, fmt.Sprintf("Failed to clear profiles  %v", err), http.StatusInternalServerError)
			return
		}
		if _, err := db.Exec("DELETE FROM prompts"); err != nil {
			http.Error(w, fmt.Sprintf("Failed to clear prompts  %v", err), http.StatusInternalServerError)
			return
		}
		if _, err := db.Exec("DELETE FROM scores"); err != nil {
			http.Error(w, fmt.Sprintf("Failed to clear scores  %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Data cleared successfully!")
	}
}

func generateScoresHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := generateMockScores(db); err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate scores: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Mock scores generated successfully!")
	}
}
