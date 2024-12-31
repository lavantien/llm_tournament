package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetStatsData(t *testing.T) {
	// Create a temporary test database
	tempDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open temporary database: %v", err)
	}
	defer tempDB.Close()

	// Initialize the database schema
	initDB()

	// Insert some mock data for testing
	_, err = tempDB.Exec("INSERT INTO bots(name, path) VALUES('bot1', 'path1'), ('bot2', 'path2'), ('bot3', 'path3')")
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
	_, err = tempDB.Exec("INSERT INTO scores(attempt, elo, botId, promptId, profile) VALUES(1, 100, 'bot1', 1, 'profile1'), (2, 50, 'bot2', 2, 'profile2'), (1, 200, 'bot3', 1, 'profile1')")
	if err != nil {
		t.Fatalf("Failed to insert mock scores: %v", err)
	}

	// Test getStatsData
	statsData, err := getStatsData(tempDB)
	if err != nil {
		t.Errorf("getStatsData failed: %v", err)
	}

	// Verify Lord of LLM
	if statsData.LordOfLLM != "bot3" {
		t.Errorf("Expected Lord of LLM to be bot3, got %s", statsData.LordOfLLM)
	}

	// Verify top bots per profile
	if len(statsData.Profiles["profile1"].TopBots) != 2 {
		t.Errorf("Expected 2 top bots for profile1, got %d", len(statsData.Profiles["profile1"].TopBots))
	}
	if statsData.Profiles["profile1"].TopBots[0].Name != "bot3" {
		t.Errorf("Expected top bot for profile1 to be bot3, got %s", statsData.Profiles["profile1"].TopBots[0].Name)
	}

	// Verify profile graph data
	if len(statsData.ProfileData["profile1"]) != 2 {
		t.Errorf("Expected 2 bots in profile graph data for profile1, got %d", len(statsData.ProfileData["profile1"]))
	}
	if statsData.ProfileData["profile1"][0].Name != "bot1" {
		t.Errorf("Expected first bot in profile graph data for profile1 to be bot1, got %s", statsData.ProfileData["profile1"][0].Name)
	}

	// Verify total Elo graph data
	if len(statsData.TotalElos) != 3 {
		t.Errorf("Expected 3 bots in total Elo graph data, got %d", len(statsData.TotalElos))
	}
	if statsData.TotalElos["bot3"] != 200 {
		t.Errorf("Expected total Elo for bot3 to be 200, got %f", statsData.TotalElos["bot3"])
	}
}

func TestConcludeStats(t *testing.T) {
	// Create a temporary test database
	tempDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open temporary database: %v", err)
	}
	defer tempDB.Close()

	// Initialize the database schema
	initDB()

	// Insert some mock data for testing
	_, err = tempDB.Exec("INSERT INTO bots(name, path) VALUES('bot1', 'path1'), ('bot2', 'path2'), ('bot3', 'path3')")
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
	_, err = tempDB.Exec("INSERT INTO scores(attempt, elo, botId, promptId, profile) VALUES(1, 100, 'bot1', 1, 'profile1'), (2, 50, 'bot2', 2, 'profile2'), (1, 200, 'bot3', 1, 'profile1')")
	if err != nil {
		t.Fatalf("Failed to insert mock scores: %v", err)
	}

	// Test concludeStats
	err = concludeStats(tempDB)
	if err != nil {
		t.Errorf("concludeStats failed: %v", err)
	}

	// Verify bestBots for each profile
	var count int
	err = tempDB.QueryRow("SELECT COUNT(*) FROM profile_bot WHERE profile_name = 'profile1'").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query profile_bot count: %v", err)
	}
	if count != 2 {
		t.Errorf("Expected 2 bestBots for profile1, got %d", count)
	}

	// Verify kingOf for each bot
	var kingOf string
	err = tempDB.QueryRow("SELECT kingOf FROM bots WHERE name = 'bot3'").Scan(&kingOf)
	if err != nil {
		t.Fatalf("Failed to query kingOf for bot3: %v", err)
	}
	if kingOf != "profile1" {
		t.Errorf("Expected kingOf for bot3 to be profile1, got %s", kingOf)
	}

	// Verify Lord of LLM
	err = tempDB.QueryRow("SELECT kingOf FROM bots WHERE name = 'bot3'").Scan(&kingOf)
	if err != nil {
		t.Fatalf("Failed to query kingOf for bot3: %v", err)
	}
	if kingOf != "profile1" {
		t.Errorf("Expected Lord of LLM to be bot3, got %s", kingOf)
	}
}
