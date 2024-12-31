package main

import (
	"database/sql"
	"fmt"
)

func recalculateLeaderboard(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Get all bots
	bots, err := getBots(db)
	if err != nil {
		return fmt.Errorf("failed to get bots: %v", err)
	}

	// Get all profiles
	profiles, err := getProfiles(db)
	if err != nil {
		return fmt.Errorf("failed to get profiles: %v", err)
	}

	// Clear existing bot scores
	if _, err := tx.Exec("UPDATE bots SET kingOf = NULL"); err != nil {
		return fmt.Errorf("failed to clear existing bot scores: %v", err)
	}

	// Recalculate scores for each bot and profile
	for _, bot := range bots {
		for _, profile := range profiles {
			totalElo, err := calculateBotEloForProfile(db, bot, profile)
			if err != nil {
				return fmt.Errorf("failed to calculate elo for bot %s and profile %s: %v", bot, profile, err)
			}

			// Update the bot's score for the profile
			if _, err := tx.Exec("UPDATE bots SET kingOf = ? WHERE name = ?", profile, bot); err != nil {
				return fmt.Errorf("failed to update bot %s score for profile %s: %v", bot, profile, err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
