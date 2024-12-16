package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ingressorModel struct {
	db    *sql.DB
	botID int
	// promptID       int
	scoreInput     textinput.Model
	speedInput     textinput.Model
	outputInput    textinput.Model
	focusIndex     int
	width          int
	height         int
	prompts        []Prompt
	selectedPrompt int
}

func initialIngressorModel(db *sql.DB) ingressorModel {
	scoreInput := textinput.New()
	scoreInput.Placeholder = "Score"
	scoreInput.CharLimit = 3
	scoreInput.Width = 20

	speedInput := textinput.New()
	speedInput.Placeholder = "Speed"
	speedInput.CharLimit = 10
	speedInput.Width = 20

	outputInput := textinput.New()
	outputInput.Placeholder = "Output Path"
	outputInput.Width = 40

	return ingressorModel{
		db:             db,
		scoreInput:     scoreInput,
		speedInput:     speedInput,
		outputInput:    outputInput,
		focusIndex:     0,
		selectedPrompt: -1,
	}
}

func (m ingressorModel) Init() tea.Cmd {
	return m.updatePrompts()
}

func (m ingressorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.focusIndex = (m.focusIndex + 1) % 4
		case "shift+tab":
			m.focusIndex = (m.focusIndex - 1 + 4) % 4
		case "enter":
			if m.focusIndex == 3 {
				return m, m.saveScore()
			}
		case "up", "k":
			if m.selectedPrompt > 0 {
				m.selectedPrompt--
			}
		case "down", "j":
			if m.selectedPrompt < len(m.prompts)-1 {
				m.selectedPrompt++
			}
		}
	case modelUpdateMsg:
		return m, nil
	case error:
		log.Printf("Error: %v", msg)
		return m, nil
	}

	var cmd tea.Cmd
	if m.focusIndex == 0 {
		m.scoreInput, cmd = m.scoreInput.Update(msg)
	} else if m.focusIndex == 1 {
		m.speedInput, cmd = m.speedInput.Update(msg)
	} else if m.focusIndex == 2 {
		m.outputInput, cmd = m.outputInput.Update(msg)
	}
	return m, cmd
}

func (m ingressorModel) View() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Bot ID: %d\n", m.botID))
	sb.WriteString("Prompts:\n")
	for i, prompt := range m.prompts {
		if i == m.selectedPrompt {
			sb.WriteString(fmt.Sprintf("> %s\n", truncateString(prompt.Content, 50)))
		} else {
			sb.WriteString(fmt.Sprintf("  %s\n", truncateString(prompt.Content, 50)))
		}
	}

	sb.WriteString("\n")

	scoreInputView := m.scoreInput.View()
	if m.focusIndex == 0 {
		scoreInputView = focusedStyle.Render(scoreInputView)
	}
	sb.WriteString(fmt.Sprintf("Score: %s\n", scoreInputView))

	speedInputView := m.speedInput.View()
	if m.focusIndex == 1 {
		speedInputView = focusedStyle.Render(speedInputView)
	}
	sb.WriteString(fmt.Sprintf("Speed: %s\n", speedInputView))

	outputInputView := m.outputInput.View()
	if m.focusIndex == 2 {
		outputInputView = focusedStyle.Render(outputInputView)
	}
	sb.WriteString(fmt.Sprintf("Output Path: %s\n", outputInputView))

	submitButton := "Submit"
	if m.focusIndex == 3 {
		submitButton = focusedStyle.Render(submitButton)
	}
	sb.WriteString(fmt.Sprintf("\n%s\n", submitButton))

	return baseStyle.Render(sb.String())
}

func (m *ingressorModel) updatePrompts() tea.Cmd {
	prompts, err := m.fetchPrompts()
	if err != nil {
		return func() tea.Msg {
			return err
		}
	}
	m.prompts = prompts
	if len(prompts) > 0 {
		m.selectedPrompt = 0
	}
	return nil
}

func (m *ingressorModel) fetchPrompts() ([]Prompt, error) {
	rows, err := m.db.Query("SELECT id, category_id, content, created_at, updated_at FROM prompts")
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

func (m *ingressorModel) saveScore() tea.Cmd {
	return func() tea.Msg {
		if m.selectedPrompt == -1 {
			return fmt.Errorf("no prompt selected")
		}
		score, err := strconv.Atoi(m.scoreInput.Value())
		if err != nil {
			return fmt.Errorf("invalid score: %w", err)
		}
		speed, err := strconv.ParseFloat(m.speedInput.Value(), 64)
		if err != nil {
			return fmt.Errorf("invalid speed: %w", err)
		}
		outputPath := m.outputInput.Value()

		_, err = m.db.Exec(
			"INSERT INTO scores (bot_id, prompt_id, score, speed, output_path) VALUES (?, ?, ?, ?, ?)",
			m.botID, m.prompts[m.selectedPrompt].ID, score, speed, outputPath,
		)
		if err != nil {
			return fmt.Errorf("failed to insert score: %w", err)
		}

		m.scoreInput.SetValue("")
		m.speedInput.SetValue("")
		m.outputInput.SetValue("")
		return modelUpdateMsg{msg: tabChangeMsg(0)}
	}
}

func (m *ingressorModel) copyPrompt() tea.Cmd {
	return func() tea.Msg {
		if m.selectedPrompt == -1 {
			return fmt.Errorf("no prompt selected")
		}
		m.outputInput.SetValue(m.prompts[m.selectedPrompt].Content)
		return nil
	}
}
