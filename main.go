package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/mattn/go-sqlite3"
)

type AppModel struct {
	currentView string
	table       table.Model
	keys        KeyMap
}

var db *sql.DB

// KeyMap defines application keybindings.
type KeyMap struct {
	Quit   key.Binding
	Tab    key.Binding
	Select key.Binding
}

// Initialize keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "quit"),
		),
		Tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch views"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
	}
}

func NewModel() AppModel {
	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "Model Name", Width: 20},
		{Title: "Version", Width: 10},
		{Title: "Status", Width: 15},
	}
	rows := []table.Row{
		{"1", "Qwen2.5-Coder", "7B", "Active"},
		{"2", "Gemma-2", "27B", "Inactive"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	t.SetStyles(table.DefaultStyles())

	return AppModel{
		currentView: "table",
		table:       t,
		keys:        DefaultKeyMap(),
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Tab):
			if m.currentView == "table" {
				m.currentView = "stats"
			} else {
				m.currentView = "table"
			}
			return m, nil

		case key.Matches(msg, m.keys.Select):
			log.Println("Selected row")
		}
	}

	if m.currentView == "table" {
		tableModel, cmd := m.table.Update(msg)
		m.table = tableModel
		return m, cmd
	}

	return m, nil
}

func (m AppModel) View() string {
	var body string

	switch m.currentView {
	case "table":
		body = m.table.View()
	case "stats":
		body = "Statistics View - Work in Progress"
	}

	footer := lipgloss.NewStyle().Background(lipgloss.Color("236")).Foreground(lipgloss.Color("230")).Padding(0, 1).Render("[q] Quit [Tab] Switch Views [Enter] Select")
	return lipgloss.JoinVertical(lipgloss.Center, body, footer)
}

func setupDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "llm.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS llms (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		version TEXT NOT NULL,
		status TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func exportData() {
	rows, err := db.Query("SELECT * FROM llms")
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer rows.Close()

	var data []map[string]interface{}
	for rows.Next() {
		var id int
		var name, version, status string
		if err := rows.Scan(&id, &name, &version, &status); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		data = append(data, map[string]interface{}{
			"id":      id,
			"name":    name,
			"version": version,
			"status":  status,
		})
	}

	file, err := os.Create("llms_export.json")
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}

	log.Println("Data exported to llms_export.json")
}

func main() {
	setupDatabase()
	defer db.Close()

	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		os.Exit(1)
	}
}
