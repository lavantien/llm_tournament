package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type promptmanModel struct {
	db               *sql.DB
	categories       []Category
	prompts          []Prompt
	selectedRow      int
	focusIndex       int
	categoryInput    textinput.Model
	promptInput      textinput.Model
	width            int
	height           int
	viewport         viewport.Model
	creating         bool
	editing          bool
	deleting         bool
	searchQuery      textinput.Model
	selectedCategory int
}

func initialPromptmanModel(db *sql.DB) promptmanModel {
	categoryInput := textinput.New()
	categoryInput.Placeholder = "Category Name"
	categoryInput.CharLimit = 50
	categoryInput.Width = 30

	promptInput := textinput.New()
	promptInput.Placeholder = "Prompt Content"
	promptInput.Width = 50
	// Removed Height as it's not a valid field
	// promptInput.Height = 5

	searchQuery := textinput.New()
	searchQuery.Placeholder = "Search Prompts"
	searchQuery.Width = 30

	return promptmanModel{
		db:               db,
		selectedRow:      -1,
		focusIndex:       0,
		categoryInput:    categoryInput,
		promptInput:      promptInput,
		searchQuery:      searchQuery,
		viewport:         viewport.New(80, 20),
		selectedCategory: -1,
	}
}

func (m promptmanModel) Init() tea.Cmd {
	return m.updateCategories()
}

func (m promptmanModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.creating || m.editing {
				m.focusIndex = (m.focusIndex + 1) % 2
			}
		case "shift+tab":
			if m.creating || m.editing {
				m.focusIndex = (m.focusIndex - 1 + 2) % 2
			}
		case "enter":
			if m.creating {
				if m.focusIndex == 1 {
					return m, m.createPrompt()
				}
			} else if m.editing {
				if m.focusIndex == 1 {
					return m, m.updatePrompt()
				}
			}
		case "ctrl+n":
			m.creating = true
			m.editing = false
			m.focusIndex = 0
			m.promptInput.SetValue("")
		case "ctrl+e":
			if m.selectedRow != -1 {
				m.editing = true
				m.creating = false
				m.focusIndex = 0
				m.loadSelectedPrompt()
			}
		case "ctrl+d":
			if m.selectedRow != -1 {
				m.deleting = true
				return m, nil
			}
		case "y":
			if m.deleting {
				return m, m.deletePrompt()
			}
		case "n", "esc":
			m.deleting = false
			m.creating = false
			m.editing = false
		case "up", "k":
			if m.selectedRow > 0 {
				m.selectedRow--
			}
		case "down", "j":
			if m.selectedRow < len(m.prompts)-1 {
				m.selectedRow++
			}
		case "left", "h":
			if m.selectedCategory > 0 {
				m.selectedCategory--
				return m, m.updatePrompts()
			}
		case "right", "l":
			if m.selectedCategory < len(m.categories)-1 {
				m.selectedCategory++
				return m, m.updatePrompts()
			}
		case "pgdown":
			m.viewport.LineDown(5)
		case "pgup":
			m.viewport.LineUp(5)
		}
	case modelUpdateMsg:
		return m, nil
	case error:
		log.Printf("Error: %v", msg)
		return m, nil
	}

	if m.creating || m.editing {
		switch m.focusIndex {
		case 0:
			m.promptInput, cmd = m.promptInput.Update(msg)
		case 1:
			m.categoryInput, cmd = m.categoryInput.Update(msg)
		}
	} else {
		m.searchQuery, cmd = m.searchQuery.Update(msg)
		if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "enter" {
			return m, m.updatePrompts()
		}
	}
	return m, cmd
}

func (m promptmanModel) View() string {
	var sb strings.Builder

	if m.creating || m.editing {
		sb.WriteString("Creating/Editing Prompt:\n\n")

		promptView := m.promptInput.View()
		if m.focusIndex == 0 {
			promptView = focusedStyle.Render(promptView)
		}
		sb.WriteString(fmt.Sprintf("Prompt: %s\n", promptView))

		categoryView := m.categoryInput.View()
		if m.focusIndex == 1 {
			categoryView = focusedStyle.Render(categoryView)
		}
		sb.WriteString(fmt.Sprintf("Category: %s\n", categoryView))

		submitButton := "Submit"
		if m.focusIndex == 1 {
			submitButton = focusedStyle.Render(submitButton)
		}
		sb.WriteString(fmt.Sprintf("\n%s\n", submitButton))
	} else if m.deleting {
		sb.WriteString("Are you sure you want to delete this prompt? (y/n)\n")
	} else {
		sb.WriteString("Prompt Manager:\n\n")
		sb.WriteString(fmt.Sprintf("Search: %s\n\n", m.searchQuery.View()))
		sb.WriteString("Categories:\n")
		for i, category := range m.categories {
			if i == m.selectedCategory {
				sb.WriteString(fmt.Sprintf("> %s\n", category.Name))
			} else {
				sb.WriteString(fmt.Sprintf("  %s\n", category.Name))
			}
		}
		sb.WriteString("\nPrompts:\n")
		for i, prompt := range m.prompts {
			if i == m.selectedRow {
				sb.WriteString(fmt.Sprintf("> %s\n", truncateString(prompt.Content, 50)))
			} else {
				sb.WriteString(fmt.Sprintf("  %s\n", truncateString(prompt.Content, 50)))
			}
		}
		m.viewport.SetContent(sb.String())
		return baseStyle.Render(m.viewport.View())
	}

	return baseStyle.Render(sb.String())
}

func (m *promptmanModel) updateCategories() tea.Cmd {
	return func() tea.Msg {
		categories, err := m.fetchCategories()
		if err != nil {
			return err
		}
		m.categories = categories
		if len(categories) > 0 {
			if m.selectedCategory == -1 {
				m.selectedCategory = 0
			}
			return m.updatePrompts()
		} else {
			m.selectedCategory = -1
			m.prompts = nil
			m.selectedRow = -1
			return nil
		}
	}
}

func (m *promptmanModel) fetchCategories() ([]Category, error) {
	rows, err := m.db.Query("SELECT id, name, created_at, updated_at FROM categories")
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (m *promptmanModel) updatePrompts() tea.Cmd {
	return func() tea.Msg {
		prompts, err := m.fetchPrompts()
		if err != nil {
			return err
		}
		m.prompts = prompts
		if len(prompts) > 0 {
			m.selectedRow = 0
		} else {
			m.selectedRow = -1
		}
		return nil
	}
}

func (m *promptmanModel) fetchPrompts() ([]Prompt, error) {
	var query string
	var rows *sql.Rows
	var err error

	search := m.searchQuery.Value()
	if search != "" {
		query = `
			SELECT p.id, p.category_id, p.content, p.created_at, p.updated_at
			FROM prompts p
			JOIN prompts_fts f ON p.id = f.rowid
			WHERE p.category_id = ? AND f.content MATCH ?
		`
		rows, err = m.db.Query(query, m.categories[m.selectedCategory].ID, search)
	} else {
		query = "SELECT id, category_id, content, created_at, updated_at FROM prompts WHERE category_id = ?"
		rows, err = m.db.Query(query, m.categories[m.selectedCategory].ID)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query prompts: %w", err)
	}
	defer rows.Close()

	var prompts []Prompt
	for rows.Next() {
		var prompt Prompt
		if err := rows.Scan(&prompt.ID, &prompt.CategoryID, &prompt.Content, &prompt.CreatedAt, &prompt.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan prompt: %w", err)
		}
		prompts = append(prompts, prompt)
	}
	return prompts, nil
}

func (m *promptmanModel) createPrompt() tea.Cmd {
	return func() tea.Msg {
		if m.selectedCategory == -1 {
			return fmt.Errorf("no category selected")
		}
		content := m.promptInput.Value()
		categoryID := m.categories[m.selectedCategory].ID

		_, err := m.db.Exec(
			"INSERT INTO prompts (category_id, content) VALUES (?, ?)",
			categoryID, content,
		)
		if err != nil {
			return fmt.Errorf("failed to insert prompt: %w", err)
		}

		m.creating = false
		m.focusIndex = 0
		m.promptInput.SetValue("")
		return m.updatePrompts()
	}
}

func (m *promptmanModel) updatePrompt() tea.Cmd {
	return func() tea.Msg {
		if m.selectedRow == -1 {
			return fmt.Errorf("no prompt selected")
		}
		prompt := m.prompts[m.selectedRow]
		content := m.promptInput.Value()
		categoryID := m.categories[m.selectedCategory].ID

		_, err := m.db.Exec(
			"UPDATE prompts SET content = ?, category_id = ? WHERE id = ?",
			content, categoryID, prompt.ID,
		)
		if err != nil {
			return fmt.Errorf("failed to update prompt: %w", err)
		}

		m.editing = false
		m.focusIndex = 0
		m.promptInput.SetValue("")
		return m.updatePrompts()
	}
}

func (m *promptmanModel) deletePrompt() tea.Cmd {
	return func() tea.Msg {
		if m.selectedRow == -1 {
			return fmt.Errorf("no prompt selected")
		}
		prompt := m.prompts[m.selectedRow]
		_, err := m.db.Exec("DELETE FROM prompts WHERE id = ?", prompt.ID)
		if err != nil {
			return fmt.Errorf("failed to delete prompt: %w", err)
		}
		m.deleting = false
		return m.updatePrompts()
	}
}

func (m *promptmanModel) loadSelectedPrompt() {
	if m.selectedRow == -1 {
		return
	}
	prompt := m.prompts[m.selectedRow]
	m.promptInput.SetValue(prompt.Content)
	m.categoryInput.SetValue(m.categories[m.selectedCategory].Name)
}

func (m *promptmanModel) nextField() tea.Cmd {
	m.focusIndex = (m.focusIndex + 1) % 2
	return nil
}

func (m *promptmanModel) prevField() tea.Cmd {
	m.focusIndex = (m.focusIndex - 1 + 2) % 2
	return nil
}

func (m *promptmanModel) nextRow() tea.Cmd {
	if m.selectedRow < len(m.prompts)-1 {
		m.selectedRow++
	}
	return nil
}

func (m *promptmanModel) prevRow() tea.Cmd {
	if m.selectedRow > 0 {
		m.selectedRow--
	}
	return nil
}

func (m *promptmanModel) selectRow() tea.Cmd {
	if m.selectedRow == -1 {
		return nil
	}
	return nil
}
