package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

func TestRecalculateLeaderboard(t *testing.T) {
	// Create a temporary test database
	tempDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		slog.Error("Failed to open temporary database", "error", err)
	}
	defer tempDB.Close()

	// Initialize the database schema
	initDB(tempDB)

	// Insert some mock data for testing
	_, err = tempDB.Exec("INSERT INTO bots(name, path) VALUES('bot1', 'path1'), ('bot2', 'path2')")
	if err != nil {
		slog.Error("Failed to insert mock bots", "error", err)
	}
	_, err = tempDB.Exec("INSERT INTO profiles(name, systemPrompt) VALUES('profile1', 'prompt1'), ('profile2', 'prompt2')")
	if err != nil {
		slog.Error("Failed to insert mock profiles", "error", err)
	}
	_, err = tempDB.Exec("INSERT INTO prompts(number, content, solution, profile) VALUES(1, 'content1', 'solution1', 'profile1'), (2, 'content2', 'solution2', 'profile2')")
	if err != nil {
		slog.Error("Failed to insert mock prompts", "error", err)
	}
	_, err = tempDB.Exec("INSERT INTO scores(attempt, elo, botId, promptId, profile) VALUES(1, 100, 'bot1', 1, 'profile1'), (2, 50, 'bot2', 2, 'profile2')")
	if err != nil {
		slog.Error("Failed to insert mock scores", "error", err)
	}

	// Test recalculateLeaderboard
	err = recalculateLeaderboard(tempDB)
	if err != nil {
		slog.Error("recalculateLeaderboard failed", "error", err)
	}

	// Verify the recalculated leaderboard data
	var bot1TotalElo, bot2TotalElo float64
	err = tempDB.QueryRow("SELECT totalElo FROM bots WHERE name = 'bot1'").Scan(&bot1TotalElo)
	if err != nil {
		slog.Error("Failed to query bot1 total Elo", "error", err)
	}
	if bot1TotalElo != 100 {
		t.Errorf("Expected bot1 total Elo to be 100, got %f", bot1TotalElo)
	}

	err = tempDB.QueryRow("SELECT totalElo FROM bots WHERE name = 'bot2'").Scan(&bot2TotalElo)
	if err != nil {
		slog.Error("Failed to query bot2 total Elo", "error", err)
	}
	if bot2TotalElo != 50 {
		t.Errorf("Expected bot2 total Elo to be 50, got %f", bot2TotalElo)
	}
}
