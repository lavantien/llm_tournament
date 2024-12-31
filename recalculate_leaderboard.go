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
	bots, err := getBotsReCalc(tx)
	if err != nil {
		return fmt.Errorf("failed to get bots: %v", err)
	}

	// Get all profiles
	profiles, err := getProfilesReCalc(tx)
	if err != nil {
		return fmt.Errorf("failed to get profiles: %v", err)
	}

	// Recalculate scores for each bot and profile
	for _, bot := range bots {
		var totalElo float64
		for _, profile := range profiles {
			elo, err := calculateBotEloForProfileReCalc(tx, bot, profile)
			if err != nil {
				return fmt.Errorf("failed to calculate elo for bot %s and profile %s: %v", bot, profile, err)
			}
			totalElo += elo
		}

		// Update the bot's total Elo score
		_, err = tx.Exec("UPDATE bots SET totalElo = ? WHERE name = ?", totalElo, bot)
		if err != nil {
			return fmt.Errorf("failed to update total Elo for bot %s: %v", bot, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func getBotsReCalc(tx *sql.Tx) ([]string, error) {
	rows, err := tx.Query("SELECT name FROM bots")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bots []string
	for rows.Next() {
		var bot string
		if err := rows.Scan(&bot); err != nil {
			return nil, err
		}
		bots = append(bots, bot)
	}

	return bots, nil
}

func getProfilesReCalc(tx *sql.Tx) ([]string, error) {
	rows, err := tx.Query("SELECT name FROM profiles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []string
	for rows.Next() {
		var profile string
		if err := rows.Scan(&profile); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func calculateBotEloForProfileReCalc(tx *sql.Tx, botName, profileName string) (float64, error) {
	rows, err := tx.Query(`
        SELECT s.elo
        FROM scores s
        JOIN prompts p ON s.promptId = p.number
        WHERE s.botId = ? AND p.profile = ?
    `, botName, profileName)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var totalElo float64
	for rows.Next() {
		var elo float64
		if err := rows.Scan(&elo); err != nil {
			return 0, err
		}
		totalElo += elo
	}

	return totalElo, nil
}
