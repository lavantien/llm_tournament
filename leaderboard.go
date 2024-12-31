package main

import (
	"database/sql"
	"fmt"
)

func getLeaderboardData(db *sql.DB) (LeaderboardData, error) {
	profiles, err := getProfiles(db)
	if err != nil {
		return LeaderboardData{}, fmt.Errorf("failed to get profiles: %v", err)
	}

	bots, err := getBots(db)
	if err != nil {
		return LeaderboardData{}, fmt.Errorf("failed to get bots: %v", err)
	}

	leaderboardData := LeaderboardData{
		Profiles: profiles,
		Bots:     make([]struct {
			Name     string             `json:"name"`
			Elos     map[string]float64 `json:"elos"`
			TotalElo float64            `json:"totalElo"`
		}, len(bots)),
	}

	for i, bot := range bots {
		leaderboardData.Bots[i].Name = bot
		leaderboardData.Bots[i].Elos = make(map[string]float64)

		for _, profile := range profiles {
			elo, err := calculateBotEloForProfile(db, bot, profile)
			if err != nil {
				return LeaderboardData{}, fmt.Errorf("failed to calculate elo for bot %s and profile %s: %v", bot, profile, err)
			}
			leaderboardData.Bots[i].Elos[profile] = elo
		}

		// Calculate totalElo for the bot
		totalElo := 0.0
		for _, elo := range leaderboardData.Bots[i].Elos {
			totalElo += elo
		}
		leaderboardData.Bots[i].TotalElo = totalElo
	}

	return leaderboardData, nil
}

type LeaderboardBot struct {
	Name     string             `json:"name"`
	Elos     map[string]float64 `json:"elos"`
	TotalElo float64            `json:"totalElo"`
}

func getProfiles(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM profiles")
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

	return profiles, nil
}

func getBots(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM bots")
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

	return bots, nil
}

func calculateBotEloForProfile(db *sql.DB, botName, profileName string) (float64, error) {
	rows, err := db.Query(`
        SELECT s.elo
        FROM scores s
        JOIN prompts p ON s.promptId = p.number
        WHERE s.botId = ? AND p.profile = ?
    `, botName, profileName)
	if err != nil {
		return 0, fmt.Errorf("failed to query scores for bot %s and profile %s: %v", botName, profileName, err)
	}
	defer rows.Close()

	var totalElo float64
	for rows.Next() {
		var elo float64
		if err := rows.Scan(&elo); err != nil {
			return 0, fmt.Errorf("failed to scan elo for bot %s and profile %s: %v", botName, profileName, err)
		}
		totalElo += elo
	}

	if err = rows.Err(); err != nil {
		return 0, fmt.Errorf("error iterating over scores for bot %s and profile %s: %v", botName, profileName, err)
	}

	return totalElo, nil
}
