package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type leaderboardModel struct {
	db            *sql.DB
	table         table.Model
	selectedRow   int
	width         int
	height        int
	ingressorOpen bool
	egressorOpen  bool
	ingressor     ingressorModel
	egressor      egressorModel
	focusIndex    int
	rows          []table.Row
}

func initialLeaderboardModel(db *sql.DB) leaderboardModel {
	columns := []table.Column{
		{Title: "Bot", Width: 20},
		{Title: "Booting", Width: 10},
		{Title: "Programming", Width: 10},
		{Title: "General", Width: 10},
		{Title: "AGI Probing", Width: 10},
		{Title: "Creative Writing", Width: 10},
		{Title: "Software Engineering", Width: 10},
		{Title: "Total", Width: 10},
		{Title: "Avg Speed", Width: 10},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return leaderboardModel{
		db:            db,
		table:         t,
		selectedRow:   -1,
		ingressorOpen: false,
		egressorOpen:  false,
		ingressor:     initialIngressorModel(db),
		egressor:      initialEgressorModel(db),
		focusIndex:    0,
	}
}

func (m leaderboardModel) Init() tea.Cmd {
	return m.updateRows()
}

func (m leaderboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.ingressorOpen {
			updatedModel, cmd := m.ingressor.Update(msg)
			m.ingressor = updatedModel.(ingressorModel)
			return m, cmd
		} else if m.egressorOpen {
			updatedModel, cmd := m.egressor.Update(msg)
			m.egressor = updatedModel.(egressorModel)
			return m, cmd
		}
	case modelUpdateMsg:
		if m.ingressorOpen {
			updatedModel, cmd := m.ingressor.Update(msg.msg)
			m.ingressor = updatedModel.(ingressorModel)
			return m, cmd
		} else if m.egressorOpen {
			updatedModel, cmd := m.egressor.Update(msg.msg)
			m.egressor = updatedModel.(egressorModel)
			return m, cmd
		}
	case error:
		log.Printf("Error: %v", msg)
		return m, nil
	}

	if m.ingressorOpen {
		updatedModel, cmd := m.ingressor.Update(msg)
		m.ingressor = updatedModel.(ingressorModel)
		return m, cmd
	} else if m.egressorOpen {
		updatedModel, cmd := m.egressor.Update(msg)
		m.egressor = updatedModel.(egressorModel)
		return m, cmd
	}

	return m, cmd
}

func (m leaderboardModel) View() string {
	if m.ingressorOpen {
		return m.ingressor.View()
	} else if m.egressorOpen {
		return m.egressor.View()
	}

	return baseStyle.Render(m.table.View())
}

func (m leaderboardModel) nextRow() tea.Cmd {
	if m.selectedRow < len(m.rows)-1 {
		m.selectedRow++
	}
	m.table.SetCursor(m.selectedRow)
	return nil
}

func (m leaderboardModel) prevRow() tea.Cmd {
	if m.selectedRow > 0 {
		m.selectedRow--
	}
	m.table.SetCursor(m.selectedRow)
	return nil
}

func (m leaderboardModel) selectRow() tea.Cmd {
	if m.selectedRow == -1 {
		return nil
	}
	return nil
}

func (m leaderboardModel) openIngressor() tea.Cmd {
	if m.selectedRow == -1 {
		return nil
	}
	m.ingressorOpen = true
	m.ingressor.botID = m.getBotIDFromSelectedRow()
	m.ingressor.focusIndex = 0
	return nil
}

func (m leaderboardModel) openEgressor() tea.Cmd {
	if m.selectedRow == -1 {
		return nil
	}
	m.egressorOpen = true
	m.egressor.botID = m.getBotIDFromSelectedRow()
	return m.egressor.updateBotDetails()
}

func (m leaderboardModel) getBotIDFromSelectedRow() int {
	if m.selectedRow == -1 {
		return 0
	}
	row := m.rows[m.selectedRow]
	botName := row[0]
	var botID int
	err := m.db.QueryRow("SELECT id FROM bots WHERE name = ?", botName).Scan(&botID)
	if err != nil {
		log.Printf("Error getting bot ID: %v", err)
		return 0
	}
	return botID
}

func (m *leaderboardModel) updateRows() tea.Cmd {
	return func() tea.Msg {
		rows, err := m.fetchLeaderboardData()
		if err != nil {
			return err
		}
		m.rows = rows
		m.table.SetRows(rows)
		return nil
	}
}

func (m *leaderboardModel) fetchLeaderboardData() ([]table.Row, error) {
	bots, err := m.fetchBots()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bots: %w", err)
	}

	categories, err := m.fetchCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}

	rows := make([]table.Row, len(bots))
	for i, bot := range bots {
		row := make(table.Row, len(categories)+3)
		row[0] = bot.Name
		totalScore := 0
		totalSpeed := 0.0
		validScores := 0

		for j, category := range categories {
			score, speed, err := m.fetchScoreForBotAndCategory(bot.ID, category.ID)
			if err != nil {
				log.Printf("Error fetching score for bot %d and category %d: %v", bot.ID, category.ID, err)
				row[j+1] = "N/A"
				continue
			}
			row[j+1] = strconv.Itoa(score)
			totalScore += score
			totalSpeed += speed
			validScores++
		}

		row[len(categories)+1] = strconv.Itoa(totalScore)
		if validScores > 0 {
			row[len(categories)+2] = fmt.Sprintf("%.2f", totalSpeed/float64(validScores))
		} else {
			row[len(categories)+2] = "N/A"
		}
		rows[i] = row
	}
	return rows, nil
}

func (m *leaderboardModel) fetchBots() ([]Bot, error) {
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

func (m *leaderboardModel) fetchCategories() ([]Category, error) {
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

func (m *leaderboardModel) fetchScoreForBotAndCategory(botID, categoryID int) (int, float64, error) {
	var totalScore int
	var totalSpeed float64
	var count int

	query := `
		SELECT 
			SUM(s.score), 
			SUM(s.speed),
			COUNT(s.id)
		FROM scores s
		JOIN prompts p ON s.prompt_id = p.id
		WHERE s.bot_id = ? AND p.category_id = ?
	`

	err := m.db.QueryRow(query, botID, categoryID).Scan(&totalScore, &totalSpeed, &count)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, fmt.Errorf("failed to query score: %w", err)
	}

	if count == 0 {
		return 0, 0, nil
	}

	return totalScore, totalSpeed, nil
}

func (m *leaderboardModel) exportData() tea.Cmd {
	return func() tea.Msg {
		data, err := m.fetchLeaderboardDataForExport()
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("leaderboard-%s.json", time.Now().Format("20060102150405"))
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		filePath := filepath.Join(homeDir, ".llp", filename)

		err = exportToJSON(data, filePath)
		if err != nil {
			log.Printf("Failed to export data: %v", err)
			return err
		}

		return nil
	}
}

func (m *leaderboardModel) fetchLeaderboardDataForExport() ([]map[string]interface{}, error) {
	bots, err := m.fetchBots()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bots: %w", err)
	}

	categories, err := m.fetchCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}

	var results []map[string]interface{}
	for _, bot := range bots {
		botData := map[string]interface{}{
			"bot_name": bot.Name,
			"scores":   make(map[string]interface{}),
		}

		totalScore := 0
		totalSpeed := 0.0
		validScores := 0

		for _, category := range categories {
			score, speed, err := m.fetchScoreForBotAndCategory(bot.ID, category.ID)
			if err != nil {
				log.Printf("Error fetching score for bot %d and category %d: %v", bot.ID, category.ID, err)
				continue
			}
			botData["scores"].(map[string]interface{})[category.Name] = score
			totalScore += score
			totalSpeed += speed
			validScores++
		}

		botData["total_score"] = totalScore
		if validScores > 0 {
			botData["average_speed"] = totalSpeed / float64(validScores)
		} else {
			botData["average_speed"] = "N/A"
		}

		results = append(results, botData)
	}
	return results, nil
}

func (m *leaderboardModel) nextField() tea.Cmd {
	m.focusIndex++
	return nil
}

func (m *leaderboardModel) prevField() tea.Cmd {
	m.focusIndex--
	return nil
}
