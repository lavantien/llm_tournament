package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetLeaderboardData(t *testing.T) {
	// Create a temporary test database
	tempDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open temporary database: %v", err)
	}
	defer tempDB.Close()

	// Initialize the database schema
	initDB(tempDB)

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

	// Test getLeaderboardData
	data, err := getLeaderboardData(tempDB)
	if err != nil {
		t.Errorf("getLeaderboardData failed: %v", err)
	}

	// Verify the leaderboard data
	if len(data.Profiles) != 2 {
		t.Errorf("Expected 2 profiles, got %d", len(data.Profiles))
	}
	if len(data.Bots) != 2 {
		t.Errorf("Expected 2 bots, got %d", len(data.Bots))
	}
	if data.Bots[0].TotalElo != 100 {
		t.Errorf("Expected bot1 total Elo to be 100, got %f", data.Bots[0].TotalElo)
	}
	if data.Bots[1].TotalElo != 50 {
		t.Errorf("Expected bot2 total Elo to be 50, got %f", data.Bots[1].TotalElo)
	}
}
