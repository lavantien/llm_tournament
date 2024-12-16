package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ModelEvaluation struct {
	ModelName             string  `json:"model_name"`
	TokensPerSecond       float64 `json:"tokens_per_second"`
	TotalTokens           int     `json:"total_tokens"`
	TimeToFirstToken      float64 `json:"time_to_first_token"`
	ConcurrencyScore      int     `json:"concurrency_score"`
	AlgorithmEfficiency   int     `json:"algorithm_efficiency_score"`
	CodeQualityScore      int     `json:"code_quality_score"`
	ErrorHandlingScore    int     `json:"error_handling_score"`
	AdvancedFeaturesScore int     `json:"advanced_features_score"`
	TotalScore            int     `json:"total_score"`
}

func evaluateModelCode(modelName, codePath string) ModelEvaluation {
	eval := ModelEvaluation{ModelName: modelName}

	// Read file contents
	codeContent, err := os.ReadFile(codePath)
	if err != nil {
		log.Printf("Error reading code file: %v", err)
		return eval
	}

	// Extract performance metrics from comment
	performanceMetrics := extractPerformanceMetrics(codeContent)
	eval.TokensPerSecond = performanceMetrics.TokensPerSecond
	eval.TotalTokens = performanceMetrics.TotalTokens
	eval.TimeToFirstToken = performanceMetrics.TimeToFirstToken

	// Assess different aspects of the implementation
	eval.ConcurrencyScore = evaluateConcurrencyImplementation(codeContent)
	eval.AlgorithmEfficiency = evaluateAlgorithmEfficiency(codeContent)
	eval.CodeQualityScore = evaluateCodeQuality(codePath)
	eval.ErrorHandlingScore = evaluateErrorHandling(codeContent)
	eval.AdvancedFeaturesScore = evaluateAdvancedFeatures(codeContent)

	// Calculate total score
	eval.TotalScore =
		eval.ConcurrencyScore*25/100 +
			eval.AlgorithmEfficiency*20/100 +
			eval.CodeQualityScore*25/100 +
			eval.ErrorHandlingScore*15/100 +
			eval.AdvancedFeaturesScore*15/100

	return eval
}

func extractPerformanceMetrics(content []byte) struct {
	TokensPerSecond  float64
	TotalTokens      int
	TimeToFirstToken float64
} {
	re := regexp.MustCompile(`// (\d+\.\d+) tok/sec • (\d+) tokens • (\d+\.\d+)s to first token`)
	matches := re.FindSubmatch(content)

	if len(matches) < 4 {
		return struct {
			TokensPerSecond  float64
			TotalTokens      int
			TimeToFirstToken float64
		}{}
	}

	tokPerSec, _ := strconv.ParseFloat(string(matches[1]), 64)
	totalTokens, _ := strconv.Atoi(string(matches[2]))
	timeToFirstToken, _ := strconv.ParseFloat(string(matches[3]), 64)

	return struct {
		TokensPerSecond  float64
		TotalTokens      int
		TimeToFirstToken float64
	}{
		TokensPerSecond:  tokPerSec,
		TotalTokens:      totalTokens,
		TimeToFirstToken: timeToFirstToken,
	}
}

func evaluateConcurrencyImplementation(content []byte) int {
	score := 0

	concurrencyChecks := []struct {
		requirement string
		points      int
		check       func([]byte) bool
	}{
		{"Goroutine Usage", 10, func(code []byte) bool {
			return bytes.Contains(code, []byte("go ")) &&
				bytes.Contains(code, []byte("func("))
		}},
		{"Channel Communication", 10, func(code []byte) bool {
			return bytes.Contains(code, []byte("chan ")) &&
				bytes.Contains(code, []byte("<-"))
		}},
		{"Context Usage", 5, func(code []byte) bool {
			return bytes.Contains(code, []byte("context.Context"))
		}},
	}

	for _, check := range concurrencyChecks {
		if check.check(content) {
			score += check.points
		}
	}

	return score
}

func evaluateAlgorithmEfficiency(content []byte) int {
	score := 0

	algorithmChecks := []struct {
		requirement string
		points      int
		check       func([]byte) bool
	}{
		{"Optimized Prime Generation", 10, func(code []byte) bool {
			return bytes.Contains(code, []byte("Sieve")) ||
				bytes.Contains(code, []byte("sqrt(")) ||
				bytes.Contains(code, []byte("isPrime"))
		}},
		{"Concurrent Computation", 10, func(code []byte) bool {
			return bytes.Contains(code, []byte("worker")) &&
				bytes.Contains(code, []byte("range"))
		}},
	}

	for _, check := range algorithmChecks {
		if check.check(content) {
			score += check.points
		}
	}

	return score
}

func evaluateCodeQuality(codePath string) int {
	score := 0
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, codePath, nil, parser.ParseComments)
	if err != nil {
		return 0
	}

	// Check for comments
	var commentCount int
	for range node.Comments {
		commentCount++
	}
	if commentCount > 5 {
		score += 10
	}

	// Check naming conventions
	score += checkNamingConventions(node)

	// Check complexity
	complexity := calculateCyclomaticComplexity(node)
	if complexity <= 10 {
		score += 15
	} else if complexity <= 20 {
		score += 10
	}

	return score
}

func evaluateErrorHandling(content []byte) int {
	score := 0

	errorHandlingChecks := []struct {
		requirement string
		points      int
		check       func([]byte) bool
	}{
		{"Error Checking", 10, func(code []byte) bool {
			return bytes.Contains(code, []byte("if err !=")) &&
				bytes.Contains(code, []byte("return"))
		}},
		{"Comprehensive Error Types", 5, func(code []byte) bool {
			return bytes.Contains(code, []byte("type Error")) ||
				bytes.Contains(code, []byte("NewError"))
		}},
	}

	for _, check := range errorHandlingChecks {
		if check.check(content) {
			score += check.points
		}
	}

	return score
}

func evaluateAdvancedFeatures(content []byte) int {
	score := 0

	advancedFeatureChecks := []struct {
		requirement string
		points      int
		check       func([]byte) bool
	}{
		{"Prime Pattern Detection", 5, func(code []byte) bool {
			return bytes.Contains(code, []byte("TwinPrime")) ||
				bytes.Contains(code, []byte("SexyPrime"))
		}},
		{"Statistical Analysis", 5, func(code []byte) bool {
			return bytes.Contains(code, []byte("Distribution")) ||
				bytes.Contains(code, []byte("Aggregate"))
		}},
		{"Flexible Configuration", 5, func(code []byte) bool {
			return bytes.Contains(code, []byte("WithConcurrency")) ||
				bytes.Contains(code, []byte("SetRange"))
		}},
	}

	for _, check := range advancedFeatureChecks {
		if check.check(content) {
			score += check.points
		}
	}

	return score
}

func calculateCyclomaticComplexity(node *ast.File) int {
	complexity := 0
	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.SelectStmt, *ast.SwitchStmt:
			complexity++
		case *ast.CaseClause:
			complexity++
		}
		return true
	})
	return complexity
}

func checkNamingConventions(node *ast.File) int {
	nameScore := 0
	ast.Inspect(node, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.FuncDecl:
			if matchesGoNamingConvention(v.Name.Name) {
				nameScore += 5
			}
		case *ast.TypeSpec:
			if matchesGoNamingConvention(v.Name.Name) {
				nameScore += 3
			}
		}
		return true
	})
	return nameScore
}

func matchesGoNamingConvention(name string) bool {
	return regexp.MustCompile(`^[A-Z][a-zA-Z0-9]*$`).MatchString(name) ||
		regexp.MustCompile(`^[a-z][a-zA-Z0-9]*$`).MatchString(name)
}

func Eval() {
	modelOutputDir := "llm_outputs/programming_task"
	results := []ModelEvaluation{}

	err := filepath.Walk(modelOutputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" && !strings.Contains(path, "_test.go") {
			modelName := strings.TrimSuffix(filepath.Base(path), ".go")
			// testPath := filepath.Join(filepath.Dir(path), modelName+"_test.go")

			// if _, err := os.Stat(testPath); !os.IsNotExist(err) {
			modelEval := evaluateModelCode(modelName, path)
			results = append(results, modelEval)
			// }
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Sort results by total score
	sort.Slice(results, func(i, j int) bool {
		return results[i].TotalScore > results[j].TotalScore
	})

	// Write results to JSON
	outputFile, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("score-code.json", outputFile, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Evaluation complete. Results written to score-code.json")
}
