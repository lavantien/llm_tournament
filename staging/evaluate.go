package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// ModelResult stores the evaluation results for a single LLM model.
type ModelResult struct {
	ModelName        string  `json:"model_name"`
	TestPassRate     float64 `json:"test_pass_rate"`
	CodeQualityScore float64 `json:"code_quality_score"`
	Speed            string  `json:"speed"`
	Runnable         bool    `json:"runnable"`
	Explanation      int     `json:"explanation"`
	Improvement      int     `json:"improvement"`
}

func main() {
	// Directory containing the generated code and test files.
	codeDir := "llm_outputs/programming_task"

	// Load all Go files from the directory.
	files, err := os.ReadDir(codeDir)
	if err != nil {
		panic(err)
	}

	// Create a slice to store the evaluation results for each model.
	results := make([]ModelResult, 0)

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			modelName := strings.TrimSuffix(file.Name(), ".go")
			filePath := filepath.Join(codeDir, file.Name())

			// Evaluate the code and generate a ModelResult struct.
			result := evaluateCode(modelName, filePath)
			results = append(results, result)
		}
	}

	// Marshal the results into JSON format.
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write the JSON data to a file named "score-code.json".
	err = os.WriteFile("score-code.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Evaluation results saved to score-code.json")
}

func evaluateCode(modelName, filePath string) ModelResult {
	// Extract speed information from the comment at the top of the file.
	speedInfo := extractSpeedInfo(filePath)

	// Run the Golang tests and capture the output.
	cmd := exec.Command("go", "test", "-v", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running tests for %s: %v\n", modelName, err)
		return ModelResult{
			ModelName:        modelName,
			TestPassRate:     0,
			CodeQualityScore: 0,
			Speed:            speedInfo,
			Runnable:         false,
		}
	}

	// Analyze the test results.
	testResults := parseTestOutput(string(output))

	// Run golangci-lint to analyze code quality.
	cmd = exec.Command("golangci-lint", "run", filePath)
	codeQualityOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running golangci-lint for %s: %v\n", modelName, err)
		return ModelResult{
			ModelName:        modelName,
			TestPassRate:     testResults.PassRate,
			CodeQualityScore: 0,
			Speed:            speedInfo,
			Runnable:         true,
		}
	}

	codeQualityScore := parseCodeQualityOutput(string(codeQualityOutput))

	// Extract manual evaluation scores from the comment.
	runnable, explanation, improvement := extractManualScores(filePath)

	return ModelResult{
		ModelName:        modelName,
		TestPassRate:     testResults.PassRate,
		CodeQualityScore: codeQualityScore,
		Speed:            speedInfo,
		Runnable:         runnable,
		Explanation:      explanation,
		Improvement:      improvement,
	}
}

// extractSpeedInfo extracts the speed information from the comment at the top of the file.
func extractSpeedInfo(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`//.*tok/sec.*`)
	match := re.FindString(string(data))
	return match
}

// parseTestOutput parses the output of the Go test command and returns a struct containing the pass rate.
func parseTestOutput(output string) struct {
	PassRate float64
} {
	passCount := strings.Count(output, "PASS")
	failCount := strings.Count(output, "FAIL")
	totalTests := passCount + failCount

	if totalTests == 0 {
		return struct{ PassRate float64 }{PassRate: 0}
	}

	passRate := float64(passCount) / float64(totalTests)
	return struct{ PassRate float64 }{PassRate: passRate}
}

// parseCodeQualityOutput parses the output of golangci-lint and returns a code quality score.
func parseCodeQualityOutput(output string) float64 {
	// TODO: Implement logic to extract a meaningful code quality score from the golangci-lint output.
	return 0 // Placeholder for now
}

// extractManualScores extracts the manual evaluation scores (runnable, explanation clarity, improvement suggestions) from the comment at the top of the file.
func extractManualScores(filePath string) (bool, int, int) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`//.*aadj\s+(true|false)\s+- expl\s+(\d+)\s+- impr\s+(\d+)`)
	match := re.FindStringSubmatch(string(data))
	if len(match) != 4 {
		return false, 0, 0 // Default values if no match found
	}

	runnable, _ := strconv.ParseBool(match[1])
	explanation, _ := strconv.Atoi(match[2])
	improvement, _ := strconv.Atoi(match[3])
	return runnable, explanation, improvement
}
