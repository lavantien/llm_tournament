package main

import (
	"database/sql"
	"fmt"
	"log/slog"
)

func recalculateLeaderboard(db *sql.DB) error {
	slog.Info("Recalculating leaderboard")
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
		slog.Info("Recalculating total Elo for bot", "bot", bot)
		totalElo := 0.0
		for _, profile := range profiles {
			elo, err := calculateBotEloForProfileReCalc(tx, bot, profile)
			if err != nil {
				return fmt.Errorf("failed to calculate elo for bot %s and profile %s: %v", bot, profile, err)
			}
			totalElo += elo
			slog.Info("Recalculating Elo for bot and profile", "bot", bot, "profile", profile, "elo", elo, "totalElo", totalElo)
		}

		// Update the bot's total Elo score
		_, err = tx.Exec("UPDATE bots SET totalElo = ? WHERE name = ?", totalElo, bot)
		if err != nil {
			return fmt.Errorf("failed to update total Elo for bot %s: %v", bot, err)
		}
		slog.Info("Updated total Elo for bot", "bot", bot, "totalElo", totalElo)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	slog.Info("Leaderboard recalculated successfully")
	return nil
}

func getBotsReCalc(tx *sql.Tx) ([]string, error) {
	slog.Info("Getting bots for recalculation")
	rows, err := tx.Query("SELECT name FROM bots")
	if err != nil {
		return nil, fmt.Errorf("failed to query bots: %v", err)
	}
	defer rows.Close()

	var bots []string
	for rows.Next() {
		var bot string
		if err := rows.Scan(&bot); err != nil {
			return nil, fmt.Errorf("failed to scan bot: %v", err)
		}
		bots = append(bots, bot)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over bots: %v", err)
	}

	slog.Info("Bots retrieved for recalculation", "bots", bots)
	return bots, nil
}

func getProfilesReCalc(tx *sql.Tx) ([]string, error) {
	slog.Info("Getting profiles for recalculation")
	rows, err := tx.Query("SELECT name FROM profiles")
	if err != nil {
		return nil, fmt.Errorf("failed to query profiles: %v", err)
	}
	defer rows.Close()

	var profiles []string
	for rows.Next() {
		var profile string
		if err := rows.Scan(&profile); err != nil {
			return nil, fmt.Errorf("failed to scan profile: %v", err)
		}
		profiles = append(profiles, profile)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over profiles: %v", err)
	}

	slog.Info("Profiles retrieved for recalculation", "profiles", profiles)
	return profiles, nil
}

func calculateBotEloForProfileReCalc(tx *sql.Tx, botName, profileName string) (float64, error) {
	slog.Info("Calculating Elo for bot and profile", "bot", botName, "profile", profileName)
	rows, err := tx.Query(`
        SELECT COALESCE(SUM(s.elo), 0)
        FROM scores s
        JOIN prompts p ON s.promptId = p.number
        WHERE s.botId = ? AND p.profile = ?
    `, botName, profileName)
	if err != nil {
		return 0, fmt.Errorf("failed to query scores for bot %s and profile %s: %v", botName, profileName, err)
	}
	defer rows.Close()

	var totalElo float64
	if rows.Next() {
		if err := rows.Scan(&totalElo); err != nil {
			return 0, fmt.Errorf("failed to scan elo for bot %s and profile %s: %v", botName, profileName, err)
		}
	}
	if err = rows.Err(); err != nil {
		return 0, fmt.Errorf("error iterating over scores for bot %s and profile %s: %v", botName, profileName, err)
	}
	slog.Info("Elo calculated for bot and profile", "bot", botName, "profile", profileName, "elo", totalElo)

	return totalElo, nil
}
