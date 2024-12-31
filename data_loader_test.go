package main

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestLoadData(t *testing.T) {
	// Create a temporary test database
	tempDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open temporary database: %v", err)
	}
	defer tempDB.Close()

	// Initialize the database schema
	initDB()

	// Create dummy data files for testing
	createDummyDataFiles()
	defer cleanupDummyDataFiles()

	// Test loadBots
	err = loadBots(tempDB, "data/test_models.md")
	if err != nil {
		t.Errorf("loadBots failed: %v", err)
	}

	// Test loadProfiles
	err = loadProfiles(tempDB, "data/test_profiles.md")
	if err != nil {
		t.Errorf("loadProfiles failed: %v", err)
	}

	// Test loadPrompts
	err = loadPrompts(tempDB, "data/test_prompts.md")
	if err != nil {
		t.Errorf("loadPrompts failed: %v", err)
	}
}

func createDummyDataFiles() {
	os.Mkdir("data", 0755)

	modelsContent := `
- Path: "model1-path"
  Size: 1.0
  Param: 2.0
  Quant: "q4"
  GPU Layers: 10
  GPU Layers Used: 5
  Ctx: 1024
  Ctx Used: 512
- Path: "model2-path"
  Size: 2.0
  Param: 3.0
  Quant: "q5"
  GPU Layers: 20
  GPU Layers Used: 10
  Ctx: 2048
  Ctx Used: 1024
`
	os.WriteFile("data/test_models.md", []byte(modelsContent), 0644)

	profilesContent := `
## Profile1

### System prompt

This is a test system prompt.

- repeat_penalty: 1.1
- top_k: 40
- top_p: 0.95
- min_p: 0.05
- top_a: 0.8

## Profile2

### System prompt

Another test system prompt.

- repeat_penalty: 1.2
- top_k: 50
- top_p: 0.9
- min_p: 0.1
- top_a: 0.7
`
	os.WriteFile("data/test_profiles.md", []byte(profilesContent), 0644)

	promptsContent := `
## Default Profile

### 1

#### Content

Test prompt 1.

#### Solution

Test solution 1.

---

### 2

#### Content

Test prompt 2.

#### Solution

Test solution 2.
`
	os.WriteFile("data/test_prompts.md", []byte(promptsContent), 0644)
}

func cleanupDummyDataFiles() {
	os.RemoveAll("data")
}
