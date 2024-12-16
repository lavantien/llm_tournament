package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func exportToJSON(data interface{}, filename string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	filePath := filepath.Join(homeDir, ".llp", filename)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode to json: %w", err)
	}
	return nil
}

func truncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen] + "..."
	}
	return s
}

// func exportToCSV(data [][]string, filename string) error {
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		return fmt.Errorf("failed to get user home directory: %w", err)
// 	}
// 	filePath := filepath.Join(homeDir, ".llp", filename)
//
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return fmt.Errorf("failed to create file: %w", err)
// 	}
// 	defer file.Close()
//
// 	writer := csv.NewWriter(file)
// 	if err := writer.WriteAll(data); err != nil {
// 		return fmt.Errorf("failed to write to csv: %w", err)
// 	}
// 	return nil
// }

// func exportToMarkdown(data [][]string, filename string) error {
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		return fmt.Errorf("failed to get user home directory: %w", err)
// 	}
// 	filePath := filepath.Join(homeDir, ".llp", filename)
//
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return fmt.Errorf("failed to create file: %w", err)
// 	}
// 	defer file.Close()
//
// 	var sb strings.Builder
// 	for _, row := range data {
// 		sb.WriteString("| ")
// 		sb.WriteString(strings.Join(row, " | "))
// 		sb.WriteString(" |\n")
// 	}
//
// 	_, err = file.WriteString(sb.String())
// 	if err != nil {
// 		return fmt.Errorf("failed to write to markdown: %w", err)
// 	}
// 	return nil
// }
