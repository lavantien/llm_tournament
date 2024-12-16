package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"

	"llm_tournament/db"
)

func setupTestDB() (*sql.DB, func()) {
	tempDir, err := os.MkdirTemp("", "llp_test")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	dbPath := filepath.Join(tempDir, "test.db")

	testDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open test database: %v", err)
	}

	schema, err := os.ReadFile("db/schema.sql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}

	_, err = testDB.Exec(string(schema))
	if err != nil {
		log.Fatalf("Failed to execute schema: %v", err)
	}

	if err := seedTestData(testDB); err != nil {
		log.Fatalf("Failed to seed test data: %v", err)
	}

	cleanup := func() {
		testDB.Close()
		os.RemoveAll(tempDir)
	}

	return testDB, cleanup
}

func seedTestData(db *sql.DB) error {
	categories := []string{"Test Category 1", "Test Category 2"}
	for _, category := range categories {
		_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", category)
		if err != nil {
			return fmt.Errorf("failed to insert category %s: %w", category, err)
		}
	}

	var categoryID1 int
	err := db.QueryRow("SELECT id FROM categories WHERE name = ?", "Test Category 1").Scan(&categoryID1)
	if err != nil {
		return fmt.Errorf("failed to get category ID: %w", err)
	}

	var categoryID2 int
	err = db.QueryRow("SELECT id FROM categories WHERE name = ?", "Test Category 2").Scan(&categoryID2)
	if err != nil {
		return fmt.Errorf("failed to get category ID: %w", err)
	}

	prompts := []struct {
		categoryID int
		content    string
	}{
		{categoryID1, "Test Prompt 1"},
		{categoryID1, "Test Prompt 2"},
		{categoryID2, "Test Prompt 3"},
	}
	for _, prompt := range prompts {
		_, err := db.Exec("INSERT INTO prompts (category_id, content) VALUES (?, ?)", prompt.categoryID, prompt.content)
		if err != nil {
			return fmt.Errorf("failed to insert prompt: %w", err)
		}
	}

	_, err = db.Exec("INSERT INTO bots (name, arch, compatibility_type, quantization, max_context_length) VALUES (?, ?, ?, ?, ?)", "Test Bot 1", "test_arch", "test_compat", "test_quant", 1024)
	if err != nil {
		return fmt.Errorf("failed to insert bot: %w", err)
	}

	_, err = db.Exec("INSERT INTO bots (name, arch, compatibility_type, quantization, max_context_length) VALUES (?, ?, ?, ?, ?)", "Test Bot 2", "test_arch", "test_compat", "test_quant", 2048)
	if err != nil {
		return fmt.Errorf("failed to insert bot: %w", err)
	}

	return nil
}

func TestIntegration(t *testing.T) {
	testDB, cleanup := setupTestDB()
	defer cleanup()

	m := initialModel(testDB)
	if err := m.Init(); err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	// Test Leaderboard
	t.Run("Leaderboard", func(t *testing.T) {
		log.Println("TestIntegration: Starting Leaderboard test")
		m.activeTab = 0
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyEnter}); err != nil {
			t.Fatalf("Failed to select row: %v", err)
		}
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlI}); err != nil {
			t.Fatalf("Failed to open ingressor: %v", err)
		}
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyEnter})
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlE}); err != nil {
			t.Fatalf("Failed to open egressor: %v", err)
		}
		m.Update(tea.KeyMsg{Type: tea.KeyEsc})
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlX}); err != nil {
			t.Fatalf("Failed to export data: %v", err)
		}
		log.Println("TestIntegration: Finished Leaderboard test")
	})

	// Test Bot Manager
	t.Run("Bot Manager", func(t *testing.T) {
		log.Println("TestIntegration: Starting Bot Manager test")
		m.activeTab = 1
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlN}); err != nil {
			t.Fatalf("Failed to create bot: %v", err)
		}
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyEnter})
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlE}); err != nil {
			t.Fatalf("Failed to edit bot: %v", err)
		}
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyEnter})
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlD}); err != nil {
			t.Fatalf("Failed to delete bot: %v", err)
		}
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'y'}}); err != nil {
			t.Fatalf("Failed to confirm delete bot: %v", err)
		}
		log.Println("TestIntegration: Finished Bot Manager test")
	})

	// Test Prompt Manager
	t.Run("Prompt Manager", func(t *testing.T) {
		log.Println("TestIntegration: Starting Prompt Manager test")
		m.activeTab = 2
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlN}); err != nil {
			t.Fatalf("Failed to create prompt: %v", err)
		}
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyEnter})
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlE}); err != nil {
			t.Fatalf("Failed to edit prompt: %v", err)
		}
		m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m.Update(tea.KeyMsg{Type: tea.KeyEnter})
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlD}); err != nil {
			t.Fatalf("Failed to delete prompt: %v", err)
		}
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'y'}}); err != nil {
			t.Fatalf("Failed to confirm delete prompt: %v", err)
		}
		log.Println("TestIntegration: Finished Prompt Manager test")
	})

	// Test Conducer
	t.Run("Conducer", func(t *testing.T) {
		log.Println("TestIntegration: Starting Conducer test")
		m.activeTab = 3
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyDown}); err != nil {
			t.Fatalf("Failed to select bot: %v", err)
		}
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyTab}); err != nil {
			t.Fatalf("Failed to select category: %v", err)
		}
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyDown}); err != nil {
			t.Fatalf("Failed to select category: %v", err)
		}
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyTab}); err != nil {
			t.Fatalf("Failed to select prompt: %v", err)
		}
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyDown}); err != nil {
			t.Fatalf("Failed to select prompt: %v", err)
		}
		if _, err := m.Update(tea.KeyMsg{Type: tea.KeyCtrlT}); err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		time.Sleep(2 * time.Second)
		log.Println("TestIntegration: Finished Conducer test")
	})
}

func TestDBMigration(t *testing.T) {
	_, cleanup := setupTestDB()
	defer cleanup()

	db, err := db.GetDB()
	if err != nil {
		t.Fatalf("Failed to get database: %v", err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query categories count: %v", err)
	}
	if count != 2 {
		t.Fatalf("Expected 2 categories, got %d", count)
	}

	err = db.QueryRow("SELECT COUNT(*) FROM prompts").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query prompts count: %v", err)
	}
	if count != 3 {
		t.Fatalf("Expected 3 prompts, got %d", count)
	}

	err = db.QueryRow("SELECT COUNT(*) FROM bots").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query bots count: %v", err)
	}
	if count != 2 {
		t.Fatalf("Expected 2 bots, got %d", count)
	}
}

func TestGetDB(t *testing.T) {
	_, cleanup := setupTestDB()
	defer cleanup()

	db, err := db.GetDB()
	if err != nil {
		t.Fatalf("Failed to get database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
