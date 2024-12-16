package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type botmanModel struct {
	db          *sql.DB
	bots        []Bot
	selectedRow int
	focusIndex  int
	nameInput   textinput.Model
	archInput   textinput.Model
	compatInput textinput.Model
	quantInput  textinput.Model
	ctxInput    textinput.Model
	width       int
	height      int
	viewport    viewport.Model
	creating    bool
	editing     bool
	deleting    bool
	searchQuery textinput.Model
}

func initialBotmanModel(db *sql.DB) botmanModel {
	nameInput := textinput.New()
	nameInput.Placeholder = "Bot Name"
	nameInput.CharLimit = 50
	nameInput.Width = 30

	archInput := textinput.New()
	archInput.Placeholder = "Architecture"
	archInput.CharLimit = 50
	archInput.Width = 30

	compatInput := textinput.New()
	compatInput.Placeholder = "Compatibility Type"
	compatInput.CharLimit = 50
	compatInput.Width = 30

	quantInput := textinput.New()
	quantInput.Placeholder = "Quantization"
	quantInput.CharLimit = 50
	quantInput.Width = 30

	ctxInput := textinput.New()
	ctxInput.Placeholder = "Max Context Length"
	ctxInput.CharLimit = 10
	ctxInput.Width = 30

	searchQuery := textinput.New()
	searchQuery.Placeholder = "Search Bots"
	searchQuery.Width = 30

	return botmanModel{
		db:          db,
		selectedRow: -1,
		focusIndex:  0,
		nameInput:   nameInput,
		archInput:   archInput,
		compatInput: compatInput,
		quantInput:  quantInput,
		ctxInput:    ctxInput,
		searchQuery: searchQuery,
		viewport:    viewport.New(80, 20),
	}
}

func (m botmanModel) Init() tea.Cmd {
	return m.updateBots()
}

func (m botmanModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.creating || m.editing {
				m.focusIndex = (m.focusIndex + 1) % 5
			}
		case "shift+tab":
			if m.creating || m.editing {
				m.focusIndex = (m.focusIndex - 1 + 5) % 5
			}
		case "enter":
			if m.creating {
				if m.focusIndex == 4 {
					return m, m.createBot()
				}
			} else if m.editing {
				if m.focusIndex == 4 {
					return m, m.updateBot()
				}
			}
		case "ctrl+n":
			m.creating = true
			m.editing = false
			m.focusIndex = 0
			m.nameInput.SetValue("")
			m.archInput.SetValue("")
			m.compatInput.SetValue("")
			m.quantInput.SetValue("")
			m.ctxInput.SetValue("")
		case "ctrl+e":
			if m.selectedRow != -1 {
				m.editing = true
				m.creating = false
				m.focusIndex = 0
				m.loadSelectedBot()
			}
		case "ctrl+d":
			if m.selectedRow != -1 {
				m.deleting = true
				return m, nil
			}
		case "y":
			if m.deleting {
				return m, m.deleteBot()
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
			if m.selectedRow < len(m.bots)-1 {
				m.selectedRow++
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
			m.nameInput, cmd = m.nameInput.Update(msg)
		case 1:
			m.archInput, cmd = m.archInput.Update(msg)
		case 2:
			m.compatInput, cmd = m.compatInput.Update(msg)
		case 3:
			m.quantInput, cmd = m.quantInput.Update(msg)
		case 4:
			m.ctxInput, cmd = m.ctxInput.Update(msg)
		}
	} else {
		m.searchQuery, cmd = m.searchQuery.Update(msg)
		if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "enter" {
			return m, m.updateBots()
		}
	}
	return m, cmd
}

func (m botmanModel) View() string {
	var sb strings.Builder

	if m.creating || m.editing {
		sb.WriteString("Creating/Editing Bot:\n\n")

		nameView := m.nameInput.View()
		if m.focusIndex == 0 {
			nameView = focusedStyle.Render(nameView)
		}
		sb.WriteString(fmt.Sprintf("Name: %s\n", nameView))

		archView := m.archInput.View()
		if m.focusIndex == 1 {
			archView = focusedStyle.Render(archView)
		}
		sb.WriteString(fmt.Sprintf("Arch: %s\n", archView))

		compatView := m.compatInput.View()
		if m.focusIndex == 2 {
			compatView = focusedStyle.Render(compatView)
		}
		sb.WriteString(fmt.Sprintf("Compatibility: %s\n", compatView))

		quantView := m.quantInput.View()
		if m.focusIndex == 3 {
			quantView = focusedStyle.Render(quantView)
		}
		sb.WriteString(fmt.Sprintf("Quantization: %s\n", quantView))

		ctxView := m.ctxInput.View()
		if m.focusIndex == 4 {
			ctxView = focusedStyle.Render(ctxView)
		}
		sb.WriteString(fmt.Sprintf("Max Context Length: %s\n", ctxView))

		submitButton := "Submit"
		if m.focusIndex == 4 {
			submitButton = focusedStyle.Render(submitButton)
		}
		sb.WriteString(fmt.Sprintf("\n%s\n", submitButton))
	} else if m.deleting {
		sb.WriteString("Are you sure you want to delete this bot? (y/n)\n")
	} else {
		sb.WriteString("Bot Manager:\n\n")
		sb.WriteString(fmt.Sprintf("Search: %s\n\n", m.searchQuery.View()))
		for i, bot := range m.bots {
			if i == m.selectedRow {
				sb.WriteString(fmt.Sprintf("> %s\n", bot.Name))
			} else {
				sb.WriteString(fmt.Sprintf("  %s\n", bot.Name))
			}
		}
		m.viewport.SetContent(sb.String())
		return baseStyle.Render(m.viewport.View())
	}

	return baseStyle.Render(sb.String())
}

func (m *botmanModel) updateBots() tea.Cmd {
	return func() tea.Msg {
		bots, err := m.fetchBots()
		if err != nil {
			return err
		}
		m.bots = bots
		if len(bots) > 0 {
			m.selectedRow = 0
		} else {
			m.selectedRow = -1
		}
		return nil
	}
}

func (m *botmanModel) fetchBots() ([]Bot, error) {
	var query string
	var rows *sql.Rows
	var err error

	search := m.searchQuery.Value()
	if search != "" {
		query = `
			SELECT b.id, b.name, b.arch, b.compatibility_type, b.quantization, b.max_context_length, b.created_at, b.updated_at
			FROM bots b
			JOIN bots_fts f ON b.id = f.rowid
			WHERE f.name MATCH ? OR f.arch MATCH ? OR f.compatibility_type MATCH ? OR f.quantization MATCH ?
		`
		rows, err = m.db.Query(query, search, search, search, search)
	} else {
		query = "SELECT id, name, arch, compatibility_type, quantization, max_context_length, created_at, updated_at FROM bots"
		rows, err = m.db.Query(query)
	}

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

func (m *botmanModel) createBot() tea.Cmd {
	return func() tea.Msg {
		name := m.nameInput.Value()
		arch := m.archInput.Value()
		compat := m.compatInput.Value()
		quant := m.quantInput.Value()
		ctx, err := strconv.Atoi(m.ctxInput.Value())
		if err != nil {
			return fmt.Errorf("invalid max context length: %w", err)
		}

		_, err = m.db.Exec(
			"INSERT INTO bots (name, arch, compatibility_type, quantization, max_context_length) VALUES (?, ?, ?, ?, ?)",
			name, arch, compat, quant, ctx,
		)
		if err != nil {
			return fmt.Errorf("failed to insert bot: %w", err)
		}

		m.creating = false
		m.focusIndex = 0
		m.nameInput.SetValue("")
		m.archInput.SetValue("")
		m.compatInput.SetValue("")
		m.quantInput.SetValue("")
		m.ctxInput.SetValue("")
		return m.updateBots()
	}
}

func (m *botmanModel) updateBot() tea.Cmd {
	return func() tea.Msg {
		if m.selectedRow == -1 {
			return fmt.Errorf("no bot selected")
		}
		bot := m.bots[m.selectedRow]
		name := m.nameInput.Value()
		arch := m.archInput.Value()
		compat := m.compatInput.Value()
		quant := m.quantInput.Value()
		ctx, err := strconv.Atoi(m.ctxInput.Value())
		if err != nil {
			return fmt.Errorf("invalid max context length: %w", err)
		}

		_, err = m.db.Exec(
			"UPDATE bots SET name = ?, arch = ?, compatibility_type = ?, quantization = ?, max_context_length = ? WHERE id = ?",
			name, arch, compat, quant, ctx, bot.ID,
		)
		if err != nil {
			return fmt.Errorf("failed to update bot: %w", err)
		}

		m.editing = false
		m.focusIndex = 0
		m.nameInput.SetValue("")
		m.archInput.SetValue("")
		m.compatInput.SetValue("")
		m.quantInput.SetValue("")
		m.ctxInput.SetValue("")
		return m.updateBots()
	}
}

func (m *botmanModel) deleteBot() tea.Cmd {
	return func() tea.Msg {
		if m.selectedRow == -1 {
			return fmt.Errorf("no bot selected")
		}
		bot := m.bots[m.selectedRow]
		_, err := m.db.Exec("DELETE FROM bots WHERE id = ?", bot.ID)
		if err != nil {
			return fmt.Errorf("failed to delete bot: %w", err)
		}
		m.deleting = false
		return m.updateBots()
	}
}

func (m *botmanModel) loadSelectedBot() {
	if m.selectedRow == -1 {
		return
	}
	bot := m.bots[m.selectedRow]
	m.nameInput.SetValue(bot.Name)
	m.archInput.SetValue(bot.Arch)
	m.compatInput.SetValue(bot.CompatibilityType)
	m.quantInput.SetValue(bot.Quantization)
	m.ctxInput.SetValue(strconv.Itoa(bot.MaxContextLength))
}

func (m *botmanModel) nextField() tea.Cmd {
	m.focusIndex = (m.focusIndex + 1) % 5
	return nil
}

func (m *botmanModel) prevField() tea.Cmd {
	m.focusIndex = (m.focusIndex - 1 + 5) % 5
	return nil
}

func (m *botmanModel) nextRow() tea.Cmd {
	if m.selectedRow < len(m.bots)-1 {
		m.selectedRow++
	}
	return nil
}

func (m *botmanModel) prevRow() tea.Cmd {
	if m.selectedRow > 0 {
		m.selectedRow--
	}
	return nil
}

func (m *botmanModel) selectRow() tea.Cmd {
	if m.selectedRow == -1 {
		return nil
	}
	return nil
}
