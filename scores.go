package main

import (
	"database/sql"
	"fmt"
	"log/slog"
)

func updateScore(db *sql.DB, botId string, profile string, promptId int, attempt int, elo float64) error {
	slog.Info("Updating score", "botId", botId, "profile", profile, "promptId", promptId, "attempt", attempt, "elo", elo)
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
		slog.Info("Inserted new score", "botId", botId, "profile", profile, "promptId", promptId, "attempt", attempt, "elo", elo)
	} else {
		// Existing score found, so update it
		_, err = tx.Exec("UPDATE scores SET attempt = ?, elo = ? WHERE id = ?", attempt, elo, existingScoreId)
		if err != nil {
			return fmt.Errorf("failed to update score: %v", err)
		}
		slog.Info("Updated score", "botId", botId, "profile", profile, "promptId", promptId, "attempt", attempt, "elo", elo)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Recalculate the leaderboard
	slog.Info("Recalculating leaderboard after score update")
	if err := recalculateLeaderboard(db); err != nil {
		return fmt.Errorf("failed to recalculate leaderboard: %v", err)
	}
	slog.Info("Finished recalculating leaderboard after score update")

	return nil
}
