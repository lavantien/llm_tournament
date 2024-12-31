package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Use gorilla/mux router for more flexible routing
	r := mux.NewRouter()

	// Set up HTTP handlers
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/load", loadDataHandler(db)).Methods("POST")
	r.HandleFunc("/clear", clearDataHandler(db)).Methods("POST")
	r.HandleFunc("/generate", generateScoresHandler(db)).Methods("POST")

	// Model Manager routes
	r.HandleFunc("/models", getModelsHandler(db)).Methods("GET")
	r.HandleFunc("/models/{name}", getModelHandler(db)).Methods("GET")
	r.HandleFunc("/models", createModelHandler(db)).Methods("POST")
	r.HandleFunc("/models/{name}", updateModelHandler(db)).Methods("PUT")
	r.HandleFunc("/models/{name}", deleteModelHandler(db)).Methods("DELETE")

	// Start the server
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
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

// Bot struct to hold data for a single bot
type Bot struct {
	Name          string  `json:"name"`
	Path          string  `json:"path"`
	Size          float64 `json:"size"`
	Param         float64 `json:"param"`
	Quant         string  `json:"quant"`
	GPULayers     int     `json:"gpuLayers"`
	GPULayersUsed int     `json:"gpuLayersUsed"`
	Ctx           int     `json:"ctx"`
	CtxUsed       int     `json:"ctxUsed"`
	KingOf        string  `json:"kingOf"`
}

// getModelsHandler returns a list of all models
func getModelsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM bots")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query bots: %v", err), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var bots []Bot
		for rows.Next() {
			var bot Bot
			if err := rows.Scan(&bot.Name, &bot.Path, &bot.Size, &bot.Param, &bot.Quant, &bot.GPULayers, &bot.GPULayersUsed, &bot.Ctx, &bot.CtxUsed, &bot.KingOf); err != nil {
				http.Error(w, fmt.Sprintf("Failed to scan bot: %v", err), http.StatusInternalServerError)
				return
			}
			bots = append(bots, bot)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bots)
	}
}

// getModelHandler returns a single model by name
func getModelHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		var bot Bot
		err := db.QueryRow("SELECT * FROM bots WHERE name = ?", name).Scan(&bot.Name, &bot.Path, &bot.Size, &bot.Param, &bot.Quant, &bot.GPULayers, &bot.GPULayersUsed, &bot.Ctx, &bot.CtxUsed, &bot.KingOf)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Model not found", http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf("Failed to query bot: %v", err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bot)
	}
}

// createModelHandler creates a new model
func createModelHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bot Bot
		if err := json.NewDecoder(r.Body).Decode(&bot); err != nil {
			http.Error(w, fmt.Sprintf("Failed to decode bot: %v", err), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO bots(name, path, size, param, quant, gpuLayers, gpuLayersUsed, ctx, ctxUsed) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", bot.Name, bot.Path, bot.Size, bot.Param, bot.Quant, bot.GPULayers, bot.GPULayersUsed, bot.Ctx, bot.CtxUsed)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert bot: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// updateModelHandler updates an existing model
func updateModelHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		var bot Bot
		if err := json.NewDecoder(r.Body).Decode(&bot); err != nil {
			http.Error(w, fmt.Sprintf("Failed to decode bot: %v", err), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE bots SET path = ?, size = ?, param = ?, quant = ?, gpuLayers = ?, gpuLayersUsed = ?, ctx = ?, ctxUsed = ? WHERE name = ?", bot.Path, bot.Size, bot.Param, bot.Quant, bot.GPULayers, bot.GPULayersUsed, bot.Ctx, bot.CtxUsed, name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to update bot: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// deleteModelHandler deletes a model by name
func deleteModelHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		_, err := db.Exec("DELETE FROM bots WHERE name = ?", name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete bot: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
