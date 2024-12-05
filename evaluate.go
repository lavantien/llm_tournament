package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// EvaluationResult represents the comprehensive evaluation of a solution
type EvaluationResult struct {
	ModelName        string    `json:"model_name"`
	Timestamp        time.Time `json:"timestamp"`
	CodeQualityScore int       `json:"code_quality_score"`
	PerformanceScore int       `json:"performance_score"`
	RobustnessScore  int       `json:"robustness_score"`
	TestingScore     int       `json:"testing_score"`
	TotalScore       int       `json:"total_score"`
	Logs             []string  `json:"logs"`
	Warnings         []string  `json:"warnings"`
	Errors           []string  `json:"errors"`
}

// Evaluator manages the comprehensive evaluation process
type Evaluator struct {
	modelName string
	sourceDir string
}

// NewEvaluator creates a new evaluation instance
func NewEvaluator(modelName string) *Evaluator {
	return &Evaluator{
		modelName: modelName,
		sourceDir: ".",
	}
}

// RunCodeQualityChecks performs static analysis
func (e *Evaluator) RunCodeQualityChecks() (int, []string) {
	warnings := []string{}
	score := 40 // Maximum score

	// Run go vet
	vetCmd := exec.Command("go", "vet", "./...")
	vetOutput, err := vetCmd.CombinedOutput()
	if err != nil {
		warnings = append(warnings, string(vetOutput))
		score -= 10
	}

	// Run golangci-lint
	lintCmd := exec.Command("golangci-lint", "run")
	lintOutput, err := lintCmd.CombinedOutput()
	if err != nil {
		warnings = append(warnings, string(lintOutput))
		score -= 15
	}

	return score, warnings
}

// RunConcurrencyTests checks race conditions and concurrency safety
func (e *Evaluator) RunConcurrencyTests() (int, []string) {
	cmd := exec.Command("go", "test", "-race", "-v", "./...")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return 0, []string{string(output)}
	}

	return 15, nil
}

// RunPerformanceBenchmarks measures system performance
func (e *Evaluator) RunPerformanceBenchmarks() (int, []string) {
	cmd := exec.Command("go", "test", "-bench=.", "-benchmem")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return 0, []string{string(output)}
	}

	// Parse benchmark results (simplified)
	benchResults := string(output)
	score := 15 // Base performance score

	// Adjust score based on benchmark characteristics
	if strings.Contains(benchResults, "ns/op") {
		score += 5
	}
	if strings.Contains(benchResults, "B/op") {
		score += 5
	}

	return score, nil
}

// Evaluate runs the full evaluation suite
func (e *Evaluator) Evaluate() *EvaluationResult {
	result := &EvaluationResult{
		ModelName: e.modelName,
		Timestamp: time.Now(),
		Logs:      []string{},
		Warnings:  []string{},
		Errors:    []string{},
	}

	// Prepare environment
	prepCmd := exec.Command("go", "mod", "tidy")
	if output, err := prepCmd.CombinedOutput(); err != nil {
		result.Errors = append(result.Errors, string(output))
		return result
	}

	// Code Quality Check
	qualityScore, qualityWarnings := e.RunCodeQualityChecks()
	result.CodeQualityScore = qualityScore
	result.Warnings = append(result.Warnings, qualityWarnings...)

	// Concurrency Tests
	concurrencyScore, concurrencyWarnings := e.RunConcurrencyTests()
	result.RobustnessScore = concurrencyScore
	result.Warnings = append(result.Warnings, concurrencyWarnings...)

	// Performance Benchmarks
	performanceScore, performanceWarnings := e.RunPerformanceBenchmarks()
	result.PerformanceScore = performanceScore
	result.Warnings = append(result.Warnings, performanceWarnings...)

	// Testing Score (integration/unit tests)
	result.TestingScore = 10 // Assuming passing tests

	result.TotalScore = result.CodeQualityScore +
		result.RobustnessScore +
		result.PerformanceScore +
		result.TestingScore

	return result
}

// SaveResults writes evaluation results to a JSON file
func (e *Evaluator) SaveResults(result *EvaluationResult) error {
	filename := fmt.Sprintf("evaluation_results_%s.json",
		strings.ReplaceAll(e.modelName, " ", "_"))

	file, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, file, 0644)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a model name")
	}

	modelName := os.Args[1]
	evaluator := NewEvaluator(modelName)

	result := evaluator.Evaluate()

	// Print results to console
	fmt.Printf("Evaluation Results for %s:\n", modelName)
	fmt.Printf("Total Score: %d/100\n", result.TotalScore)
	fmt.Printf("Code Quality: %d/40\n", result.CodeQualityScore)
	fmt.Printf("Performance: %d/15\n", result.PerformanceScore)
	fmt.Printf("Robustness: %d/15\n", result.RobustnessScore)
	fmt.Printf("Testing: %d/10\n", result.TestingScore)

	if err := evaluator.SaveResults(result); err != nil {
		log.Printf("Error saving results: %v", err)
	}

	// Exit with non-zero status if total score is too low
	if result.TotalScore < 60 {
		os.Exit(1)
	}
}
