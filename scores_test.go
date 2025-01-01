package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

func TestUpdateScore(t *testing.T) {
	// Create a temporary test database
	tempDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		slog.Error("Failed to open temporary database", "error", err)
	}
	defer tempDB.Close()

	// Initialize the database schema
	initDB(tempDB)

	// Insert some mock data for testing
	_, err = tempDB.Exec("INSERT INTO bots(name, path) VALUES('bot1', 'path1')")
	if err != nil {
		slog.Error("Failed to insert mock bot", "error", err)
	}
	_, err = tempDB.Exec("INSERT INTO prompts(number, content, solution, profile) VALUES(1, 'content1', 'solution1', 'profile1')")
	if err != nil {
		slog.Error("Failed to insert mock prompt", "error", err)
	}

	// Test updating a score
	err = updateScore(tempDB, "bot1", "profile1", 1, 1, 100)
	if err != nil {
		slog.Error("updateScore failed", "error", err)
	}

	// Verify the updated score
	var elo float64
	err = tempDB.QueryRow("SELECT elo FROM scores WHERE botId = 'bot1' AND promptId = 1").Scan(&elo)
	if err != nil {
		slog.Error("Failed to query updated score", "error", err)
	}
	if elo != 100 {
		t.Errorf("Expected updated score to be 100, got %f", elo)
	}

	// Test updating an existing score
	err = updateScore(tempDB, "bot1", "profile1", 1, 2, 50)
	if err != nil {
		slog.Error("updateScore failed", "error", err)
	}

	// Verify the updated score
	err = tempDB.QueryRow("SELECT elo FROM scores WHERE botId = 'bot1' AND promptId = 1").Scan(&elo)
	if err != nil {
		slog.Error("Failed to query updated score", "error", err)
	}
	if elo != 50 {
		t.Errorf("Expected updated score to be 50, got %f", elo)
	}

	// Test recalculating leaderboard after updating score
	var totalElo float64
	err = tempDB.QueryRow("SELECT totalElo FROM bots WHERE name = 'bot1'").Scan(&totalElo)
	if err != nil {
		slog.Error("Failed to query bot total Elo after update", "error", err)
	}
	if totalElo != 50 {
		t.Errorf("Expected bot total Elo to be updated to 50, got %f", totalElo)
	}
}
