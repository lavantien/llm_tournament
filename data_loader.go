package main

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func loadData(db *sql.DB) error {
	if err := loadBots(db, "data/models.md"); err != nil {
		return fmt.Errorf("failed to load bots: %v", err)
	}
	if err := loadProfiles(db, "data/profiles.md"); err != nil {
		return fmt.Errorf("failed to load profiles: %v", err)
	}
	if err := loadPrompts(db, "data/prompts.md"); err != nil {
		return fmt.Errorf("failed to load prompts: %v", err)
	}
	return nil
}

func loadBots(db *sql.DB, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	botRegex := regexp.MustCompile(`(?m)^\s*-\s*Path:\s*"([^"]+)"\s*\n\s*-\s*Size:\s*([\d.]+)\s*\n\s*-\s*Param:\s*([\d.]+)\s*\n\s*-\s*Quant:\s*"([^"]+)"\s*\n\s*-\s*GPU\s*Layers:\s*(\d+)\s*\n\s*-\s*GPU\s*Layers\s*Used:\s*(\d+)\s*\n\s*-\s*Ctx:\s*(\d+)\s*\n\s*-\s*Ctx\s*Used:\s*(\d+)\s*$`)
	matches := botRegex.FindAllStringSubmatch(string(content), -1)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO bots(name, path, size, param, quant, gpuLayers, gpuLayersUsed, ctx, ctxUsed) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, match := range matches {
		path := match[1]
		name := extractNameFromPath(path)
		size, _ := strconv.ParseFloat(match[2], 64)
		param, _ := strconv.ParseFloat(match[3], 64)
		quant := match[4]
		gpuLayers, _ := strconv.Atoi(match[5])
		gpuLayersUsed, _ := strconv.Atoi(match[6])
		ctx, _ := strconv.Atoi(match[7])
		ctxUsed, _ := strconv.Atoi(match[8])

		_, err = stmt.Exec(name, path, size, param, quant, gpuLayers, gpuLayersUsed, ctx, ctxUsed)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func loadProfiles(db *sql.DB, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	profileRegex := regexp.MustCompile(`(?m)^##\s*(.*)\s*\n\n###\s*System\s*prompt\n\n([\s\S]*?)(?=(?:###|$))`)
	paramRegex := regexp.MustCompile(`(?m)^\s*-\s*([a-z_]+):\s*([\d.]+)\s*$`)

	matches := profileRegex.FindAllStringSubmatch(string(content), -1)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO profiles(name, systemPrompt, repeatPenalty, topK, topP, minP, topA) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, match := range matches {
		name := match[1]
		systemPrompt := ""
		repeatPenalty := 1.0
		topK := 0
		topP := 0.0
		minP := 0.0
		topA := 0.0

		params := paramRegex.FindAllStringSubmatch(match[2], -1)
		for _, param := range params {
			switch param[1] {
			case "System prompt":
				systemPrompt = strings.TrimSpace(param[2])
			case "repeat_penalty":
				repeatPenalty, _ = strconv.ParseFloat(param[2], 64)
			case "top_k":
				topK, _ = strconv.Atoi(param[2])
			case "top_p":
				topP, _ = strconv.ParseFloat(param[2], 64)
			case "min_p":
				minP, _ = strconv.ParseFloat(param[2], 64)
			case "top_a":
				topA, _ = strconv.ParseFloat(param[2], 64)
			}
		}

		_, err = stmt.Exec(name, systemPrompt, repeatPenalty, topK, topP, minP, topA)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func loadPrompts(db *sql.DB, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	promptRegex := regexp.MustCompile(`(?m)^###\s*(\d+)\n\n####\s*Content\n\n([\s\S]*?)(?=(?:####\s*Solution|$))####\s*Solution\n\n([\s\S]*?)(?=(?:---|$))`)
	profileTypeRegex := regexp.MustCompile(`(?m)^##\s*(.*)\s*Profile\s*$`)

	profileTypeMatches := profileTypeRegex.FindAllStringSubmatch(string(content), -1)
	currentProfileType := ""
	if len(profileTypeMatches) > 0 {
		currentProfileType = profileTypeMatches[0][1]
	}

	matches := promptRegex.FindAllStringSubmatch(string(content), -1)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO prompts(number, content, solution, profile) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, match := range matches {
		number, _ := strconv.Atoi(match[1])
		content := strings.TrimSpace(match[2])
		solution := strings.TrimSpace(match[3])

		_, err = stmt.Exec(number, content, solution, currentProfileType)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func extractNameFromPath(path string) string {
	parts := strings.Split(path, "/")
	filename := parts[len(parts)-1]
	nameParts := strings.Split(filename, "-")
	return strings.Join(nameParts[:len(nameParts)-1], "-")
}
