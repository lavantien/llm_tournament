package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type conducerModel struct {
	db               *sql.DB
	botInput         textinput.Model
	categoryInput    textinput.Model
	promptInput      textinput.Model
	output           string
	focusIndex       int
	width            int
	height           int
	bots             []Bot
	categories       []Category
	prompts          []Prompt
	selectedBot      int
	selectedCategory int
	selectedPrompt   int
	sending          bool
}

func initialConducerModel(db *sql.DB) conducerModel {
	botInput := textinput.New()
	botInput.Placeholder = "Select Bot"
	botInput.Width = 30

	categoryInput := textinput.New()
	categoryInput.Placeholder = "Select Category"
	categoryInput.Width = 30

	promptInput := textinput.New()
	promptInput.Placeholder = "Select Prompt"
	promptInput.Width = 50
	// Removed Height as it's not a valid field
	// promptInput.Height = 5

	return conducerModel{
		db:               db,
		botInput:         botInput,
		categoryInput:    categoryInput,
		promptInput:      promptInput,
		focusIndex:       0,
		selectedBot:      -1,
		selectedCategory: -1,
		selectedPrompt:   -1,
	}
}

func (m conducerModel) Init() tea.Cmd {
	return tea.Batch(m.updateBots(), m.updateCategories())
}

func (m conducerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.focusIndex = (m.focusIndex + 1) % 3
		case "shift+tab":
			m.focusIndex = (m.focusIndex - 1 + 3) % 3
		case "up", "k":
			if m.focusIndex == 0 {
				if m.selectedBot > 0 {
					m.selectedBot--
					return m, m.updatePrompts()
				}
			} else if m.focusIndex == 1 {
				if m.selectedCategory > 0 {
					m.selectedCategory--
					return m, m.updatePrompts()
				}
			} else if m.focusIndex == 2 {
				if m.selectedPrompt > 0 {
					m.selectedPrompt--
				}
			}
		case "down", "j":
			if m.focusIndex == 0 {
				if m.selectedBot < len(m.bots)-1 {
					m.selectedBot++
					return m, m.updatePrompts()
				}
			} else if m.focusIndex == 1 {
				if m.selectedCategory < len(m.categories)-1 {
					m.selectedCategory++
					return m, m.updatePrompts()
				}
			} else if m.focusIndex == 2 {
				if m.selectedPrompt < len(m.prompts)-1 {
					m.selectedPrompt++
				}
			}
		}
	case modelUpdateMsg:
		return m, nil
	case error:
		log.Printf("Error: %v", msg)
		return m, nil
	case responseMsg:
		m.output = string(msg)
		m.sending = false
		return m, nil
	}

	var cmd tea.Cmd
	if m.focusIndex == 0 {
		m.botInput, cmd = m.botInput.Update(msg)
	} else if m.focusIndex == 1 {
		m.categoryInput, cmd = m.categoryInput.Update(msg)
	} else if m.focusIndex == 2 {
		m.promptInput, cmd = m.promptInput.Update(msg)
	}
	return m, cmd
}

func (m conducerModel) View() string {
	var sb strings.Builder

	botView := m.botInput.View()
	if m.focusIndex == 0 {
		botView = focusedStyle.Render(botView)
	}
	sb.WriteString(fmt.Sprintf("Bot: %s\n", botView))
	for i, bot := range m.bots {
		if i == m.selectedBot {
			sb.WriteString(fmt.Sprintf("> %s\n", bot.Name))
		} else {
			sb.WriteString(fmt.Sprintf("  %s\n", bot.Name))
		}
	}

	sb.WriteString("\n")

	categoryView := m.categoryInput.View()
	if m.focusIndex == 1 {
		categoryView = focusedStyle.Render(categoryView)
	}
	sb.WriteString(fmt.Sprintf("Category: %s\n", categoryView))
	for i, category := range m.categories {
		if i == m.selectedCategory {
			sb.WriteString(fmt.Sprintf("> %s\n", category.Name))
		} else {
			sb.WriteString(fmt.Sprintf("  %s\n", category.Name))
		}
	}

	sb.WriteString("\n")

	promptView := m.promptInput.View()
	if m.focusIndex == 2 {
		promptView = focusedStyle.Render(promptView)
	}
	sb.WriteString(fmt.Sprintf("Prompt: %s\n", promptView))
	for i, prompt := range m.prompts {
		if i == m.selectedPrompt {
			sb.WriteString(fmt.Sprintf("> %s\n", truncateString(prompt.Content, 50)))
		} else {
			sb.WriteString(fmt.Sprintf("  %s\n", truncateString(prompt.Content, 50)))
		}
	}

	sb.WriteString("\n")

	if m.sending {
		sb.WriteString("Sending request...\n")
	} else {
		sb.WriteString(fmt.Sprintf("Output:\n%s\n", m.output))
	}

	return baseStyle.Render(sb.String())
}

func (m *conducerModel) updateBots() tea.Cmd {
	return func() tea.Msg {
		bots, err := m.fetchBots()
		if err != nil {
			return err
		}
		m.bots = bots
		if len(bots) > 0 {
			m.selectedBot = 0
			return m.updatePrompts()
		} else {
			m.selectedBot = -1
			m.prompts = nil
			m.selectedPrompt = -1
			return nil
		}
	}
}

func (m *conducerModel) fetchBots() ([]Bot, error) {
	rows, err := m.db.Query("SELECT id, name, arch, compatibility_type, quantization, max_context_length, created_at, updated_at FROM bots")
	if err != nil {
		return nil, fmt.Errorf("failed to query bots: %w", err)
	}
	defer rows.Close()

	var bots []Bot
	for rows.Next() {
		var bot Bot
		if err := rows.Scan(&bot.ID, &bot.Name, &bot.Arch, &bot.CompatibilityType, &bot.Quantization, &bot.MaxContextLength, &bot.CreatedAt, &bot.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan bot: %w", err)
		}
		bots = append(bots, bot)
	}
	return bots, nil
}

func (m *conducerModel) updateCategories() tea.Cmd {
	return func() tea.Msg {
		categories, err := m.fetchCategories()
		if err != nil {
			return err
		}
		m.categories = categories
		if len(categories) > 0 {
			m.selectedCategory = 0
			return m.updatePrompts()
		} else {
			m.selectedCategory = -1
			m.prompts = nil
			m.selectedPrompt = -1
			return nil
		}
	}
}

func (m *conducerModel) fetchCategories() ([]Category, error) {
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

func (m *conducerModel) updatePrompts() tea.Cmd {
	return func() tea.Msg {
		prompts, err := m.fetchPrompts()
		if err != nil {
			return err
		}
		m.prompts = prompts
		if len(prompts) > 0 {
			m.selectedPrompt = 0
		} else {
			m.selectedPrompt = -1
		}
		return nil
	}
}

func (m *conducerModel) fetchPrompts() ([]Prompt, error) {
	if m.selectedCategory == -1 {
		return nil, nil
	}
	rows, err := m.db.Query("SELECT id, category_id, content, created_at, updated_at FROM prompts WHERE category_id = ?", m.categories[m.selectedCategory].ID)
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

func (m *conducerModel) sendRequest() tea.Cmd {
	return func() tea.Msg {
		if m.selectedBot == -1 || m.selectedPrompt == -1 {
			return fmt.Errorf("bot or prompt not selected")
		}
		m.sending = true
		bot := m.bots[m.selectedBot]
		prompt := m.prompts[m.selectedPrompt]
		return m.makeRequest(bot, prompt)
	}
}

type responseMsg string

func (m *conducerModel) makeRequest(bot Bot, prompt Prompt) tea.Cmd {
	return func() tea.Msg {
		url := "http://127.0.0.1:1234/v1/chat/completions"
		messages := []map[string]string{
			{"role": "system", "content": "You are an expert translator"},
			{"role": "user", "content": prompt.Content},
		}
		requestBody := map[string]interface{}{
			"model":      bot.Name,
			"stream":     false,
			"max_tokens": -1,
			"messages":   messages,
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}

		if choices, ok := response["choices"].([]interface{}); ok && len(choices) > 0 {
			if choice, ok := choices[0].(map[string]interface{}); ok {
				if message, ok := choice["message"].(map[string]interface{}); ok {
					if content, ok := message["content"].(string); ok {
						m.saveOutput(bot, prompt, content)
						return responseMsg(content)
					}
				}
			}
		}

		return fmt.Errorf("unexpected response format: %v", response)
	}
}

func (m *conducerModel) saveOutput(bot Bot, prompt Prompt, output string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	dirPath := filepath.Join(homeDir, ".llp", "llm_outputs")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}
	filename := fmt.Sprintf("%s-%s.md", bot.Name, time.Now().Format("20060102150405"))
	filePath := filepath.Join(dirPath, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(output)
	if err != nil {
		return fmt.Errorf("failed to write output to file: %w", err)
	}

	_, err = m.db.Exec(
		"UPDATE scores SET output_path = ? WHERE bot_id = ? AND prompt_id = ?",
		filePath, bot.ID, prompt.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update score with output path: %w", err)
	}

	return nil
}

func (m *conducerModel) nextField() tea.Cmd {
	m.focusIndex = (m.focusIndex + 1) % 3
	return nil
}

func (m *conducerModel) prevField() tea.Cmd {
	m.focusIndex = (m.focusIndex - 1 + 3) % 3
	return nil
}
