package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestRecalculateLeaderboard(t *testing.T) {
	// Create a temporary test database
	tempDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open temporary database: %v", err)
	}
	defer tempDB.Close()

	// Initialize the database schema
	initDB()

	// Insert some mock data for testing
	_, err = tempDB.Exec("INSERT INTO bots(name, path) VALUES('bot1', 'path1'), ('bot2', 'path2')")
	if err != nil {
		t.Fatalf("Failed to insert mock bots: %v", err)
	}
	_, err = tempDB.Exec("INSERT INTO profiles(name, systemPrompt) VALUES('profile1', 'prompt1'), ('profile2', 'prompt2')")
	if err != nil {
		t.Fatalf("Failed to insert mock profiles: %v", err)
	}
	_, err = tempDB.Exec("INSERT INTO prompts(number, content, solution, profile) VALUES(1, 'content1', 'solution1', 'profile1'), (2, 'content2', 'solution2', 'profile2')")
	if err != nil {
		t.Fatalf("Failed to insert mock prompts: %v", err)
	}
	_, err = tempDB.Exec("INSERT INTO scores(attempt, elo, botId, promptId, profile) VALUES(1, 100, 'bot1', 1, 'profile1'), (2, 50, 'bot2', 2, 'profile2')")
	if err != nil {
		t.Fatalf("Failed to insert mock scores: %v", err)
	}

	// Test recalculateLeaderboard
	err = recalculateLeaderboard(tempDB)
	if err != nil {
		t.Errorf("recalculateLeaderboard failed: %v", err)
	}

	// Verify the recalculated leaderboard data
	var bot1TotalElo, bot2TotalElo float64
	err = tempDB.QueryRow("SELECT totalElo FROM bots WHERE name = 'bot1'").Scan(&bot1TotalElo)
	if err != nil {
		t.Fatalf("Failed to query bot1 total Elo: %v", err)
	}
	if bot1TotalElo != 100 {
		t.Errorf("Expected bot1 total Elo to be 100, got %f", bot1TotalElo)
	}

	err = tempDB.QueryRow("SELECT totalElo FROM bots WHERE name = 'bot2'").Scan(&bot2TotalElo)
	if err != nil {
		t.Fatalf("Failed to query bot2 total Elo: %v", err)
	}
	if bot2TotalElo != 50 {
		t.Errorf("Expected bot2 total Elo to be 50, got %f", bot2TotalElo)
	}
}
