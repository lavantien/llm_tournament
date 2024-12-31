package main

import (
	"database/sql"
	"fmt"
	"sort"
)

// BotStats represents the statistics for a bot
type BotStats struct {
	Name     string             `json:"name"`
	Profile  string             `json:"profile"`
	Elo      float64            `json:"elo"`
	TotalElo float64            `json:"totalElo"`
	Scores   map[string]float64 `json:"scores"`
}

// ProfileStats represents the statistics for a profile
type ProfileStats struct {
	Name      string             `json:"name"`
	TopBots   []BotStats         `json:"topBots"`
	BotScores map[string]float64 `json:"botScores"`
}

// StatsData represents the overall statistics data
type StatsData struct {
	LordOfLLM   string                  `json:"lordOfLLM"`
	Profiles    map[string]ProfileStats `json:"profiles"`
	TotalElos   map[string]float64      `json:"totalElos"`
	ProfileData map[string][]BotStats  `json:"profileData"`
}

// getStatsData retrieves and compiles statistics data
func getStatsData(db *sql.DB) (StatsData, error) {
	statsData := StatsData{
		Profiles:    make(map[string]ProfileStats),
		TotalElos:   make(map[string]float64),
		ProfileData: make(map[string][]BotStats),
	}

	if err := calculateTopBotsPerProfile(db, &statsData); err != nil {
		return statsData, fmt.Errorf("failed to calculate top bots per profile: %v", err)
	}

	if err := calculateLordOfLLM(db, &statsData); err != nil {
		return statsData, fmt.Errorf("failed to calculate Lord of LLM: %v", err)
	}

	if err := prepareProfileGraphData(db, &statsData); err != nil {
		return statsData, fmt.Errorf("failed to prepare profile graph  %v", err)
	}

	if err := prepareTotalEloGraphData(db, &statsData); err != nil {
		return statsData, fmt.Errorf("failed to prepare total elo graph  %v", err)
	}

	return statsData, nil
}

// calculateTopBotsPerProfile calculates the top 10 bots for each profile
func calculateTopBotsPerProfile(db *sql.DB, statsData *StatsData) error {
	profiles, err := getProfiles(db)
	if err != nil {
		return err
	}

	for _, profile := range profiles {
		rows, err := db.Query(`
            SELECT b.name, COALESCE(SUM(s.elo), 0) as elo
            FROM bots b
            LEFT JOIN scores s ON b.name = s.botId
            LEFT JOIN prompts p ON s.promptId = p.number
            WHERE p.profile = ?
            GROUP BY b.name
            ORDER BY elo DESC
            LIMIT 10
        `, profile)
		if err != nil {
			return err
		}
		defer rows.Close()

		profileStats := ProfileStats{
			Name:      profile,
			TopBots:   []BotStats{},
			BotScores: make(map[string]float64),
		}

		for rows.Next() {
			var botStats BotStats
			if err := rows.Scan(&botStats.Name, &botStats.Elo); err != nil {
				return err
			}
			profileStats.TopBots = append(profileStats.TopBots, botStats)
			profileStats.BotScores[botStats.Name] = botStats.Elo
		}

		statsData.Profiles[profile] = profileStats
	}

	return nil
}

// calculateLordOfLLM determines the bot with the highest total Elo
func calculateLordOfLLM(db *sql.DB, statsData *StatsData) error {
	rows, err := db.Query(`
        SELECT b.name, COALESCE(SUM(s.elo), 0) as totalElo
        FROM bots b
        LEFT JOIN scores s ON b.name = s.botId
        GROUP BY b.name
        ORDER BY totalElo DESC
        LIMIT 1
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		var botName string
		var totalElo float64
		if err := rows.Scan(&botName, &totalElo); err != nil {
			return err
		}
		statsData.LordOfLLM = botName
	}

	return nil
}

// prepareProfileGraphData prepares data for the per-profile line graphs
func prepareProfileGraphData(db *sql.DB, statsData *StatsData) error {
	profiles, err := getProfiles(db)
	if err != nil {
		return err
	}

	for _, profile := range profiles {
		rows, err := db.Query(`
            SELECT b.name, COALESCE(SUM(s.elo), 0) as elo
            FROM bots b
            LEFT JOIN scores s ON b.name = s.botId
            LEFT JOIN prompts p ON s.promptId = p.number
            WHERE p.profile = ?
            GROUP BY b.name
            ORDER BY b.name
        `, profile)
		if err != nil {
			return err
		}
		defer rows.Close()

		var botStatsList []BotStats
		for rows.Next() {
			var botStats BotStats
			if err := rows.Scan(&botStats.Name, &botStats.Elo); err != nil {
				return err
			}
			botStatsList = append(botStatsList, botStats)
		}

		statsData.ProfileData[profile] = botStatsList
	}

	return nil
}

// prepareTotalEloGraphData prepares data for the total Elo line graph
func prepareTotalEloGraphData(db *sql.DB, statsData *StatsData) error {
	rows, err := db.Query(`
        SELECT b.name, COALESCE(SUM(s.elo), 0) as totalElo
        FROM bots b
        LEFT JOIN scores s ON b.name = s.botId
        GROUP BY b.name
        ORDER BY b.name
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var botName string
		var totalElo float64
		if err := rows.Scan(&botName, &totalElo); err != nil {
			return err
		}
		statsData.TotalElos[botName] = totalElo
	}

	return nil
}

// concludeStats updates the database with the concluded stats
func concludeStats(db *sql.DB) error {
	statsData, err := getStatsData(db)
	if err != nil {
		return fmt.Errorf("failed to get stats  %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Update bestBots for each profile
	for profileName, profileStats := range statsData.Profiles {
		for _, botStats := range profileStats.TopBots {
			_, err := tx.Exec(`
                INSERT INTO profile_bot (profile_name, bot_name)
                VALUES (?, ?)
                ON CONFLICT (profile_name, bot_name) DO NOTHING
            `, profileName, botStats.Name)
			if err != nil {
				return fmt.Errorf("failed to update bestBots for profile %s: %v", profileName, err)
			}
		}
	}

	// Update kingOf for each bot based on the profile they performed best in
	for profileName, profileStats := range statsData.Profiles {
		if len(profileStats.TopBots) > 0 {
			topBot := profileStats.TopBots[0].Name
			_, err := tx.Exec(`
                UPDATE bots SET kingOf = ? WHERE name = ?
            `, profileName, topBot)
			if err != nil {
				return fmt.Errorf("failed to update kingOf for bot %s: %v", topBot, err)
			}
		}
	}

	// Update the bot with the highest total Elo to have "Lord of LLM"
	if statsData.LordOfLLM != "" {
		_, err := tx.Exec(`
            UPDATE bots SET kingOf = 'Lord of LLM' WHERE name = ?
        `, statsData.LordOfLLM)
		if err != nil {
			return fmt.Errorf("failed to update Lord of LLM: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
