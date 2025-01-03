package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./llm-tournament.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	initDB(db)

	if len(os.Args) > 1 && os.Args[1] == "load" {
		err := loadData(db)
		if err != nil {
			log.Fatalf("Failed to load  %v", err)
		}
		log.Println("Data loaded successfully.")
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/load_data", loadDataHandler).Methods("POST")
	r.HandleFunc("/clear_data", clearDataHandler).Methods("POST")
	r.HandleFunc("/generate_mock_scores", generateMockScoresHandler).Methods("POST")
	r.HandleFunc("/models", getModelsHandler).Methods("GET")
	r.HandleFunc("/models", createModelHandler).Methods("POST")
	r.HandleFunc("/models/{name}", getModelHandler).Methods("GET")
	r.HandleFunc("/models/{name}", updateModelHandler).Methods("PUT")
	r.HandleFunc("/models/{name}", deleteModelHandler).Methods("DELETE")
	r.HandleFunc("/profiles", getProfilesHandler).Methods("GET")
	r.HandleFunc("/profiles", createProfileHandler).Methods("POST")
	r.HandleFunc("/profiles/{name}", getProfileHandler).Methods("GET")
	r.HandleFunc("/profiles/{name}", updateProfileHandler).Methods("PUT")
	r.HandleFunc("/profiles/{name}", deleteProfileHandler).Methods("DELETE")
	r.HandleFunc("/prompts", getPromptsHandler).Methods("GET")
	r.HandleFunc("/prompts", createPromptHandler).Methods("POST")
	r.HandleFunc("/prompts/{number}", getPromptHandler).Methods("GET")
	r.HandleFunc("/prompts/{number}", updatePromptHandler).Methods("PUT")
	r.HandleFunc("/prompts/{number}", deletePromptHandler).Methods("DELETE")
	r.HandleFunc("/model_manager", modelManagerHandler)
	r.HandleFunc("/profile_manager", profileManagerHandler)
	r.HandleFunc("/prompt_manager", promptManagerHandler)
	r.HandleFunc("/leaderboard", leaderboardHandler)
	r.HandleFunc("/leaderboard_data", getLeaderboardDataHandler)
	r.HandleFunc("/stats", statsHandler)
	r.HandleFunc("/stats_data", getStatsDataHandler)
	r.HandleFunc("/conclude_stats", concludeStatsHandler).Methods("POST")
	r.HandleFunc("/update_score", updateScoreHandler).Methods("POST")
	r.HandleFunc("/refresh_leaderboard", refreshLeaderboardDataHandler).Methods("GET")

	r.HandleFunc("/nav", navHandler)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDB(db *sql.DB) {
	createTables := `
    CREATE TABLE IF NOT EXISTS bots (
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
        totalElo REAL DEFAULT 0,
        FOREIGN KEY (kingOf) REFERENCES profiles(name)
    );

    CREATE TABLE IF NOT EXISTS profiles (
        name TEXT PRIMARY KEY,
        systemPrompt TEXT,
        repeatPenalty REAL,
        topK INTEGER,
        topP REAL,
        minP REAL,
        topA REAL
    );

    CREATE TABLE IF NOT EXISTS prompts (
        number INTEGER PRIMARY KEY,
        content TEXT,
        solution TEXT,
        profile TEXT,
        FOREIGN KEY (profile) REFERENCES profiles(name)
    );

    CREATE TABLE IF NOT EXISTS scores (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        attempt INTEGER,
        elo REAL,
        botId TEXT,
        promptId INTEGER,
        profile TEXT,
        FOREIGN KEY (botId) REFERENCES bots(name),
        FOREIGN KEY (promptId) REFERENCES prompts(number),
        FOREIGN KEY (profile) REFERENCES profiles(name)
    );

    CREATE TABLE IF NOT EXISTS profile_bot (
        profile_name TEXT,
        bot_name TEXT,
        PRIMARY KEY (profile_name, bot_name),
        FOREIGN KEY (profile_name) REFERENCES profiles(name),
        FOREIGN KEY (bot_name) REFERENCES bots(name)
    );
    `

	_, err := db.Exec(createTables)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html", "templates/nav.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func navHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/nav.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadDataHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Loading data")
	err := loadData(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load  %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Data loaded successfully.")
}

func clearDataHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Clearing data")
	_, err := db.Exec("DELETE FROM scores")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to clear scores: %v", err), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("DELETE FROM prompts")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to clear prompts: %v", err), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("DELETE FROM bots")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to clear bots: %v", err), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("DELETE FROM profiles")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to clear profiles: %v", err), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("UPDATE sqlite_sequence SET seq = 0 WHERE name = 'scores'")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to reset sequence: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Data cleared successfully.")
}

func generateMockScoresHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Generating mock scores")
	err := generateMockScores(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate mock scores: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Mock scores generated successfully.")
}

type Bot struct {
	Name          string         `json:"name"`
	Path          string         `json:"path"`
	Size          float64        `json:"size"`
	Param         float64        `json:"param"`
	Quant         string         `json:"quant"`
	GPULayers     int            `json:"gpuLayers"`
	GPULayersUsed int            `json:"gpuLayersUsed"`
	Ctx           int            `json:"ctx"`
	CtxUsed       int            `json:"ctxUsed"`
	KingOf        sql.NullString `json:"kingOf"`
}

func getModelsHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting models")
	rows, err := db.Query("SELECT name, path, size, param, quant, gpuLayers, gpuLayersUsed, ctx, ctxUsed, kingOf FROM bots")
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

	for i := range bots {
		if !bots[i].KingOf.Valid {
			bots[i].KingOf.String = ""
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bots); err != nil {
		fmt.Println("Failed to get bots", err)
	}
}

func createModelHandler(w http.ResponseWriter, r *http.Request) {
	var bot Bot
	if err := json.NewDecoder(r.Body).Decode(&bot); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode bot: %v", err), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO bots(name, path, size, param, quant, gpuLayers, gpuLayersUsed, ctx, ctxUsed) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		bot.Name, bot.Path, bot.Size, bot.Param, bot.Quant, bot.GPULayers, bot.GPULayersUsed, bot.Ctx, bot.CtxUsed)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert bot: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getModelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var bot Bot
	err := db.QueryRow("SELECT name, path, size, param, quant, gpuLayers, gpuLayersUsed, ctx, ctxUsed, kingOf FROM bots WHERE name = ?", name).Scan(&bot.Name, &bot.Path, &bot.Size, &bot.Param, &bot.Quant, &bot.GPULayers, &bot.GPULayersUsed, &bot.Ctx, &bot.CtxUsed, &bot.KingOf)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Bot not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Failed to query bot: %v", err), http.StatusInternalServerError)
		}
		return
	}

	if !bot.KingOf.Valid {
		bot.KingOf.String = ""
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bot); err != nil {
		fmt.Println("Failed to get bot", err)
	}
}

func updateModelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	originalName := vars["name"]

	var bot Bot
	if err := json.NewDecoder(r.Body).Decode(&bot); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode bot: %v", err), http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to begin transaction: %v", err), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	if originalName != bot.Name {
		_, err := tx.Exec("UPDATE bots SET name = ? WHERE name = ?", bot.Name, originalName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to update bot name: %v", err), http.StatusInternalServerError)
			return
		}
	}

	_, err = tx.Exec("UPDATE bots SET path = ?, size = ?, param = ?, quant = ?, gpuLayers = ?, gpuLayersUsed = ?, ctx = ?, ctxUsed = ?, kingOf = ? WHERE name = ?",
		bot.Path, bot.Size, bot.Param, bot.Quant, bot.GPULayers, bot.GPULayersUsed, bot.Ctx, bot.CtxUsed, bot.KingOf, bot.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update bot: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, fmt.Sprintf("Failed to commit transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteModelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	_, err := db.Exec("DELETE FROM bots WHERE name = ?", name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete bot: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type Profile struct {
	Name          string  `json:"name"`
	SystemPrompt  string  `json:"systemPrompt"`
	RepeatPenalty float64 `json:"repeatPenalty"`
	TopK          int     `json:"topK"`
	TopP          float64 `json:"topP"`
	MinP          float64 `json:"minP"`
	TopA          float64 `json:"topA"`
}

func getProfilesHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting profiles")
	rows, err := db.Query("SELECT name, systemPrompt, repeatPenalty, topK, topP, minP, topA FROM profiles")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to query profiles: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var profiles []Profile
	for rows.Next() {
		var profile Profile
		if err := rows.Scan(&profile.Name, &profile.SystemPrompt, &profile.RepeatPenalty, &profile.TopK, &profile.TopP, &profile.MinP, &profile.TopA); err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan profile: %v", err), http.StatusInternalServerError)
			return
		}
		profiles = append(profiles, profile)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(profiles); err != nil {
		fmt.Println("Failed to get profiles", err)
	}
}

func createProfileHandler(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode profile: %v", err), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO profiles(name, systemPrompt, repeatPenalty, topK, topP, minP, topA) VALUES(?, ?, ?, ?, ?, ?, ?)",
		profile.Name, profile.SystemPrompt, profile.RepeatPenalty, profile.TopK, profile.TopP, profile.MinP, profile.TopA)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert profile: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var profile Profile
	err := db.QueryRow("SELECT name, systemPrompt, repeatPenalty, topK, topP, minP, topA FROM profiles WHERE name = ?", name).Scan(&profile.Name, &profile.SystemPrompt, &profile.RepeatPenalty, &profile.TopK, &profile.TopP, &profile.MinP, &profile.TopA)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Profile not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Failed to query profile: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		fmt.Println("Failed to get profile", err)
	}
}

func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode profile: %v", err), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE profiles SET systemPrompt = ?, repeatPenalty = ?, topK = ?, topP = ?, minP = ?, topA = ? WHERE name = ?",
		profile.SystemPrompt, profile.RepeatPenalty, profile.TopK, profile.TopP, profile.MinP, profile.TopA, name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update profile: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	_, err := db.Exec("DELETE FROM profiles WHERE name = ?", name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete profile: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type Prompt struct {
	Number   int    `json:"number"`
	Content  string `json:"content"`
	Solution string `json:"solution"`
	Profile  string `json:"profile"`
}

func getPromptsHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting prompts")
	rows, err := db.Query("SELECT number, content, solution, profile FROM prompts")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to query prompts: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var prompts []Prompt
	for rows.Next() {
		var prompt Prompt
		if err := rows.Scan(&prompt.Number, &prompt.Content, &prompt.Solution, &prompt.Profile); err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan prompt: %v", err), http.StatusInternalServerError)
			return
		}
		prompts = append(prompts, prompt)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prompts); err != nil {
		fmt.Println("Failed to get prompts", err)
	}
}

func createPromptHandler(w http.ResponseWriter, r *http.Request) {
	var prompt Prompt
	if err := json.NewDecoder(r.Body).Decode(&prompt); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode prompt: %v", err), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO prompts(number, content, solution, profile) VALUES(?, ?, ?, ?)",
		prompt.Number, prompt.Content, prompt.Solution, prompt.Profile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert prompt: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getPromptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])
	if err != nil {
		http.Error(w, "Invalid prompt number", http.StatusBadRequest)
		return
	}

	var prompt Prompt
	err = db.QueryRow("SELECT number, content, solution, profile FROM prompts WHERE number = ?", number).Scan(&prompt.Number, &prompt.Content, &prompt.Solution, &prompt.Profile)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Prompt not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Failed to query prompt: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prompt); err != nil {
		fmt.Println("Failed to get prompt", err)
	}
}

func updatePromptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])
	if err != nil {
		http.Error(w, "Invalid prompt number", http.StatusBadRequest)
		return
	}

	var prompt Prompt
	if err := json.NewDecoder(r.Body).Decode(&prompt); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode prompt: %v", err), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE prompts SET content = ?, solution = ?, profile = ? WHERE number = ?",
		prompt.Content, prompt.Solution, prompt.Profile, number)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update prompt: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deletePromptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])
	if err != nil {
		http.Error(w, "Invalid prompt number", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM prompts WHERE number = ?", number)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete prompt: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func modelManagerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/model_manager.html", "templates/nav.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "model_manager.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func profileManagerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/profile_manager.html", "templates/nav.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "profile_manager.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func promptManagerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/prompt_manager.html", "templates/nav.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "prompt_manager.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func leaderboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/leaderboard.html", "templates/nav.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "leaderboard.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type LeaderboardData struct {
	Profiles []string `json:"profiles"`
	Bots     []struct {
		Name     string             `json:"name"`
		Elos     map[string]float64 `json:"elos"`
		TotalElo float64            `json:"totalElo"`
	} `json:"bots"`
}

func getLeaderboardDataHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting leaderboard data")
	data, err := getLeaderboardData(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get leaderboard  %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Failed to get leaderboard data", err)
	}
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/stats.html", "templates/nav.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "stats.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getStatsDataHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting stats data")
	data, err := getStatsData(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get stats  %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Failed to get stats data", err)
	}
}

func concludeStatsHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Concluding stats")
	err := concludeStats(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to conclude stats: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Stats concluded successfully.")
}

func updateScoreHandler(w http.ResponseWriter, r *http.Request) {
	var scoreData struct {
		BotId    string  `json:"botId"`
		Profile  string  `json:"profile"`
		PromptId int     `json:"promptId"`
		Attempt  int     `json:"attempt"`
		Elo      float64 `json:"elo"`
	}

	if err := json.NewDecoder(r.Body).Decode(&scoreData); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode score  %v", err), http.StatusBadRequest)
		return
	}

	if err := updateScore(db, scoreData.BotId, scoreData.Profile, scoreData.PromptId, scoreData.Attempt, scoreData.Elo); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update score: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Score updated successfully.")
}

func refreshLeaderboardDataHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Refreshing leaderboard data")
	data, err := getLeaderboardData(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to refresh leaderboard  %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Failed to refresh leaderboard data", err)
	}
}
