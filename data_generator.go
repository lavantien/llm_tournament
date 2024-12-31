package main

import (
	"database/sql"
	"fmt"
	"math/rand"
)

func generateMockScores(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT name FROM bots")
	if err != nil {
		return fmt.Errorf("failed to query bots: %v", err)
	}
	defer rows.Close()

	var bots []string
	for rows.Next() {
		var botName string
		if err := rows.Scan(&botName); err != nil {
			return fmt.Errorf("failed to scan bot name: %v", err)
		}
		bots = append(bots, botName)
	}

	rows, err = tx.Query("SELECT number FROM prompts")
	if err != nil {
		return fmt.Errorf("failed to query prompts: %v", err)
	}
	defer rows.Close()

	var prompts []int
	for rows.Next() {
		var promptNumber int
		if err := rows.Scan(&promptNumber); err != nil {
			return fmt.Errorf("failed to scan prompt number: %v", err)
		}
		prompts = append(prompts, promptNumber)
	}

	stmt, err := tx.Prepare("INSERT INTO scores(attempt, elo, botId, promptId) VALUES(?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	for _, bot := range bots {
		for _, prompt := range prompts {
			attempt := rand.Intn(4) + 1
			var elo float64
			switch attempt {
			case 1:
				elo = 100
			case 2:
				elo = 50
			case 3:
				elo = 20
			default:
				elo = 0
			}

			if _, err := stmt.Exec(attempt, elo, bot, prompt); err != nil {
				return fmt.Errorf("failed to insert score: %v", err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
