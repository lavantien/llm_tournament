package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type egressorModel struct {
	db      *sql.DB
	botID   int
	bot     Bot
	width   int
	height  int
	details string
}

func initialEgressorModel(db *sql.DB) egressorModel {
	return egressorModel{
		db: db,
	}
}

func (m egressorModel) Init() tea.Cmd {
	return nil
}

func (m egressorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" {
			return m, func() tea.Msg {
				return modelUpdateMsg{msg: tabChangeMsg(0)}
			}
		}
	case modelUpdateMsg:
		return m, nil
	case error:
		log.Printf("Error: %v", msg)
		return m, nil
	}
	return m, nil
}

func (m egressorModel) View() string {
	var sb strings.Builder
	sb.WriteString("Bot Details:\n\n")
	sb.WriteString(fmt.Sprintf("ID: %d\n", m.bot.ID))
	sb.WriteString(fmt.Sprintf("Name: %s\n", m.bot.Name))
	sb.WriteString(fmt.Sprintf("Arch: %s\n", m.bot.Arch))
	sb.WriteString(fmt.Sprintf("Compatibility Type: %s\n", m.bot.CompatibilityType))
	sb.WriteString(fmt.Sprintf("Quantization: %s\n", m.bot.Quantization))
	sb.WriteString(fmt.Sprintf("Max Context Length: %d\n", m.bot.MaxContextLength))
	sb.WriteString(fmt.Sprintf("Created At: %s\n", m.bot.CreatedAt.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("Updated At: %s\n", m.bot.UpdatedAt.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("\n%s\n", m.details))
	sb.WriteString("\nPress ESC to return to leaderboard\n")
	return baseStyle.Render(sb.String())
}

func (m *egressorModel) updateBotDetails() tea.Cmd {
	return func() tea.Msg {
		bot, err := m.fetchBotDetails()
		if err != nil {
			return err
		}
		m.bot = bot
		m.details, err = m.fetchScoreDetails()
		if err != nil {
			return err
		}
		return nil
	}
}

func (m *egressorModel) fetchBotDetails() (Bot, error) {
	var bot Bot
	err := m.db.QueryRow("SELECT id, name, arch, compatibility_type, quantization, max_context_length, created_at, updated_at FROM bots WHERE id = ?", m.botID).Scan(
		&bot.ID, &bot.Name, &bot.Arch, &bot.CompatibilityType, &bot.Quantization, &bot.MaxContextLength, &bot.CreatedAt, &bot.UpdatedAt,
	)
	if err != nil {
		return Bot{}, fmt.Errorf("failed to query bot details: %w", err)
	}
	return bot, nil
}

func (m *egressorModel) fetchScoreDetails() (string, error) {
	rows, err := m.db.Query(`
		SELECT 
			p.content,
			s.score,
			s.speed,
			s.output_path
		FROM scores s
		JOIN prompts p ON s.prompt_id = p.id
		WHERE s.bot_id = ?
	`, m.botID)
	if err != nil {
		return "", fmt.Errorf("failed to query score details: %w", err)
	}
	defer rows.Close()

	var sb strings.Builder
	for rows.Next() {
		var content string
		var score int
		var speed float64
		var outputPath string
		if err := rows.Scan(&content, &score, &speed, &outputPath); err != nil {
			return "", fmt.Errorf("failed to scan score details: %w", err)
		}
		sb.WriteString(fmt.Sprintf("Prompt: %s\n", truncateString(content, 50)))
		sb.WriteString(fmt.Sprintf("Score: %d, Speed: %.2f, Output Path: %s\n\n", score, speed, outputPath))
	}
	return sb.String(), nil
}
