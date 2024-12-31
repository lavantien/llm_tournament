package main

import (
	"database/sql"
	"fmt"
)

func updateScore(db *sql.DB, botId string, profile string, promptId int, attempt int, elo float64) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Check if a score already exists for the given bot, profile and prompt
	var existingScoreId int
	err = tx.QueryRow("SELECT id FROM scores WHERE botId = ? AND promptId = ? AND profile = ?", botId, promptId, profile).Scan(&existingScoreId)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check for existing score: %v", err)
	}

	if err == sql.ErrNoRows {
		// No existing score, so insert a new one
		_, err = tx.Exec("INSERT INTO scores(attempt, elo, botId, promptId, profile) VALUES(?, ?, ?, ?, ?)", attempt, elo, botId, promptId, profile)
		if err != nil {
			return fmt.Errorf("failed to insert new score: %v", err)
		}
	} else {
		// Existing score found, so update it
		_, err = tx.Exec("UPDATE scores SET attempt = ?, elo = ? WHERE id = ?", attempt, elo, existingScoreId)
		if err != nil {
			return fmt.Errorf("failed to update score: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Recalculate the leaderboard
	if err := recalculateLeaderboard(db); err != nil {
		return fmt.Errorf("failed to recalculate leaderboard: %v", err)
	}

	return nil
}
