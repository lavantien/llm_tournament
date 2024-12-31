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
			elo, err := getBotEloForProfile(db, bot, profile)
			if err != nil {
				return LeaderboardData{}, fmt.Errorf("failed to get elo for bot %s and profile %s: %v", bot, profile, err)
			}
			leaderboardData.Bots[i].Elos[profile] = elo
			leaderboardData.Bots[i].TotalElo += elo
		}
	}

	return leaderboardData, nil
}

func getProfiles(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM profiles")
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

func getBots(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM bots")
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

func getBotEloForProfile(db *sql.DB, botName, profileName string) (float64, error) {
	var totalElo float64
	rows, err := db.Query(`
        SELECT COALESCE(SUM(s.elo), 0)
        FROM scores s
        JOIN prompts p ON s.promptId = p.number
        WHERE s.botId = ? AND p.profile = ?
    `, botName, profileName)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&totalElo); err != nil {
			return 0, err
		}
	}

	return totalElo, nil
}
