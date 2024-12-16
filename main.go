package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"llm_tournament/db"
)

type Bot struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Arch              string    `json:"arch"`
	CompatibilityType string    `json:"compatibility_type"`
	Quantization      string    `json:"quantization"`
	MaxContextLength  int       `json:"max_context_length"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Prompt struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Score struct {
	ID         int       `json:"id"`
	BotID      int       `json:"bot_id"`
	PromptID   int       `json:"prompt_id"`
	Score      int       `json:"score"`
	Speed      float64   `json:"speed"`
	OutputPath string    `json:"output_path"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type model struct {
	db          *sql.DB
	activeTab   int
	leaderboard leaderboardModel
	ingressor   ingressorModel
	egressor    egressorModel
	botman      botmanModel
	promptman   promptmanModel
	conducer    conducerModel
	quitting    bool
}

func initialModel(db *sql.DB) model {
	return model{
		db:          db,
		activeTab:   0,
		leaderboard: initialLeaderboardModel(db),
		ingressor:   initialIngressorModel(db),
		egressor:    initialEgressorModel(db),
		botman:      initialBotmanModel(db),
		promptman:   initialPromptmanModel(db),
		conducer:    initialConducerModel(db),
		quitting:    false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "ctrl+l":
			m.activeTab = 0
			return m, nil
		case "ctrl+b":
			m.activeTab = 1
			return m, nil
		case "ctrl+p":
			m.activeTab = 2
			return m, nil
		case "ctrl+d":
			m.activeTab = 3
			return m, nil
		case "ctrl+x":
			if m.activeTab == 0 {
				return m, m.leaderboard.exportData()
			}
			return m, nil
		case "tab":
			if m.activeTab == 0 {
				return m, m.leaderboard.nextField()
			} else if m.activeTab == 1 {
				return m, m.botman.nextField()
			} else if m.activeTab == 2 {
				return m, m.promptman.nextField()
			} else if m.activeTab == 3 {
				return m, m.conducer.nextField()
			}
			return m, nil
		case "shift+tab":
			if m.activeTab == 0 {
				return m, m.leaderboard.prevField()
			} else if m.activeTab == 1 {
				return m, m.botman.prevField()
			} else if m.activeTab == 2 {
				return m, m.promptman.prevField()
			} else if m.activeTab == 3 {
				return m, m.conducer.prevField()
			}
			return m, nil
		case "up", "k":
			if m.activeTab == 0 {
				return m, m.leaderboard.prevRow()
			} else if m.activeTab == 1 {
				return m, m.botman.prevRow()
			} else if m.activeTab == 2 {
				return m, m.promptman.prevRow()
			}
			return m, nil
		case "down", "j":
			if m.activeTab == 0 {
				return m, m.leaderboard.nextRow()
			} else if m.activeTab == 1 {
				return m, m.botman.nextRow()
			} else if m.activeTab == 2 {
				return m, m.promptman.nextRow()
			}
			return m, nil
		case "enter":
			if m.activeTab == 0 {
				return m, m.leaderboard.selectRow()
			} else if m.activeTab == 1 {
				return m, m.botman.selectRow()
			} else if m.activeTab == 2 {
				return m, m.promptman.selectRow()
			}
			return m, nil
		case "ctrl+i":
			if m.activeTab == 0 {
				return m, m.leaderboard.openIngressor()
			}
			return m, nil
		case "ctrl+e":
			if m.activeTab == 0 {
				return m, m.leaderboard.openEgressor()
			}
			return m, nil
		case "ctrl+t":
			if m.activeTab == 3 {
				return m, m.conducer.sendRequest()
			}
			return m, nil
		case "ctrl+y":
			if m.activeTab == 0 && m.leaderboard.ingressorOpen {
				return m, m.leaderboard.ingressor.copyPrompt()
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.leaderboard.width = msg.Width
		m.leaderboard.height = msg.Height
		m.ingressor.width = msg.Width
		m.ingressor.height = msg.Height
		m.egressor.width = msg.Width
		m.egressor.height = msg.Height
		m.botman.width = msg.Width
		m.botman.height = msg.Height
		m.promptman.width = msg.Width
		m.promptman.height = msg.Height
		m.conducer.width = msg.Width
		m.conducer.height = msg.Height
	case tabChangeMsg:
		m.activeTab = int(msg)
		return m, nil
	case modelUpdateMsg:
		if m.activeTab == 0 {
			updatedModel, cmd := m.leaderboard.Update(msg.msg)
			m.leaderboard = updatedModel.(leaderboardModel)
			return m, cmd
		} else if m.activeTab == 1 {
			updatedModel, cmd := m.botman.Update(msg.msg)
			m.botman = updatedModel.(botmanModel)
			return m, cmd
		} else if m.activeTab == 2 {
			updatedModel, cmd := m.promptman.Update(msg.msg)
			m.promptman = updatedModel.(promptmanModel)
			return m, cmd
		} else if m.activeTab == 3 {
			updatedModel, cmd := m.conducer.Update(msg.msg)
			m.conducer = updatedModel.(conducerModel)
			return m, cmd
		}
	case error:
		log.Printf("Error: %v", msg)
		return m, nil
	}

	if m.activeTab == 0 {
		updatedModel, cmd := m.leaderboard.Update(msg)
		m.leaderboard = updatedModel.(leaderboardModel)
		return m, cmd
	} else if m.activeTab == 1 {
		updatedModel, cmd := m.botman.Update(msg)
		m.botman = updatedModel.(botmanModel)
		return m, cmd
	} else if m.activeTab == 2 {
		updatedModel, cmd := m.promptman.Update(msg)
		m.promptman = updatedModel.(promptmanModel)
		return m, cmd
	} else if m.activeTab == 3 {
		updatedModel, cmd := m.conducer.Update(msg)
		m.conducer = updatedModel.(conducerModel)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Exiting...\n"
	}

	var view string
	switch m.activeTab {
	case 0:
		view = m.leaderboard.View()
	case 1:
		view = m.botman.View()
	case 2:
		view = m.promptman.View()
	case 3:
		view = m.conducer.View()
	}

	return view
}

type tabChangeMsg int
type modelUpdateMsg struct {
	msg tea.Msg
}

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(true)
)

func main() {
	db, err := db.Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	defer db.Close()

	p := tea.NewProgram(initialModel(db), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
