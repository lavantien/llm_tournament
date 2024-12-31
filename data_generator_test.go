package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGenerateMockScores(t *testing.T) {
	// Create a temporary test database
	tempDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open temporary database: %v", err)
	}
	defer tempDB.Close()

	// Initialize the database schema
	// Assign the temporary database to the global db variable
	db = tempDB
	initDB()

	// Insert some mock bots and prompts for testing
	_, err = tempDB.Exec("INSERT INTO bots(name, path) VALUES('bot1', 'path1'), ('bot2', 'path2')")
	if err != nil {
		t.Fatalf("Failed to insert mock bots: %v", err)
	}
	_, err = tempDB.Exec("INSERT INTO prompts(number, content, solution, profile) VALUES(1, 'content1', 'solution1', 'profile1'), (2, 'content2', 'solution2', 'profile2')")
	if err != nil {
		t.Fatalf("Failed to insert mock prompts: %v", err)
	}

	// Test generateMockScores
	err = generateMockScores(tempDB)
	if err != nil {
		t.Errorf("generateMockScores failed: %v", err)
	}

	// Verify that scores were generated
	var count int
	err = tempDB.QueryRow("SELECT COUNT(*) FROM scores").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query scores count: %v", err)
	}
	if count == 0 {
		t.Errorf("Expected scores to be generated, but found none")
	}
}
