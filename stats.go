package main

import (
	"database/sql"
	"fmt"
	"log/slog"
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
	ProfileData map[string][]BotStats   `json:"profileData"`
}

// getStatsData retrieves and compiles statistics data
func getStatsData(db *sql.DB) (StatsData, error) {
	slog.Info("Retrieving statistics data")
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

	slog.Info("Statistics data retrieved successfully")
	return statsData, nil
}

// calculateTopBotsPerProfile calculates the top 10 bots for each profile
func calculateTopBotsPerProfile(db *sql.DB, statsData *StatsData) error {
	slog.Info("Calculating top bots per profile")
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

	slog.Info("Top bots per profile calculated successfully")
	return nil
}

// calculateLordOfLLM determines the bot with the highest total Elo
func calculateLordOfLLM(db *sql.DB, statsData *StatsData) error {
	slog.Info("Calculating Lord of LLM")
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

	slog.Info("Lord of LLM calculated successfully")
	return nil
}

// prepareProfileGraphData prepares data for the per-profile line graphs
func prepareProfileGraphData(db *sql.DB, statsData *StatsData) error {
	slog.Info("Preparing profile graph data")
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

	slog.Info("Profile graph data prepared successfully")
	return nil
}

// prepareTotalEloGraphData prepares data for the total Elo line graph
func prepareTotalEloGraphData(db *sql.DB, statsData *StatsData) error {
	slog.Info("Preparing total Elo graph data")
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

	slog.Info("Total Elo graph data prepared successfully")
	return nil
}

// concludeStats updates the database with the concluded stats
func concludeStats(db *sql.DB) error {
	slog.Info("Concluding stats")
	statsData, err := getStatsData(db)
	if err != nil {
		return fmt.Errorf("failed to get stats: %v", err)
	}

	slog.Debug("Stats data", "statsData", statsData)

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Reset kingOf for all bots
	slog.Info("Resetting kingOf for all bots")
	_, err = tx.Exec(`UPDATE bots SET kingOf = ''`)
	if err != nil {
		return fmt.Errorf("failed to reset kingOf for all bots: %v", err)
	}

	// Update bestBots for each profile
	for profileName, profileStats := range statsData.Profiles {
		slog.Info("Updating bestBots for profile", "profile", profileName)
		for _, botStats := range profileStats.TopBots {
			slog.Debug("Adding bot to bestBots", "bot", botStats.Name, "profile", profileName)
			_, err = tx.Exec(`
                INSERT INTO profile_bot (profile_name, bot_name)
                VALUES (?, ?)
                ON CONFLICT (profile_name, bot_name) DO NOTHING
            `, profileName, botStats.Name)
			if err != nil {
				return fmt.Errorf("failed to update bestBots for profile %s: %v", profileName, err)
			}
		}
	}

	// Update kingOf for each bot based on their best-performing profile
	botKingOf := make(map[string]string)
	for profileName, profileStats := range statsData.Profiles {
		slog.Info("Checking kingOf for profile", "profile", profileName)
		if len(profileStats.TopBots) > 0 {
			kingBot := profileStats.TopBots[0]
			slog.Debug("Bot is king of profile", "bot", kingBot.Name, "profile", profileName)
			botKingOf[kingBot.Name] = profileName
		}
	}

	// Find the bot with the highest total Elo
	var lordOfLLM string
	var maxTotalElo float64
	for botName, totalElo := range statsData.TotalElos {
		if totalElo > maxTotalElo {
			maxTotalElo = totalElo
			lordOfLLM = botName
		}
	}

	for botName, profileName := range botKingOf {
		slog.Info("Updating kingOf for bot", "bot", botName, "profile", profileName)
		_, err = tx.Exec(`
            UPDATE bots SET kingOf = ? WHERE name = ?
        `, profileName, botName)
		if err != nil {
			return fmt.Errorf("failed to update kingOf for bot %s: %v", botName, err)
		}
	}

	// Update the bot with the highest total Elo to have "Lord of LLM"
	if lordOfLLM != "" {
		slog.Info("Setting Lord of LLM", "bot", lordOfLLM)
		_, err = tx.Exec(`
            UPDATE bots SET kingOf = 'Lord of LLM' WHERE name = ?
        `, lordOfLLM)
		if err != nil {
			return fmt.Errorf("failed to update Lord of LLM: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	slog.Info("Stats concluded successfully")
	return nil
}
