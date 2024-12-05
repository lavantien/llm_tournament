# Makefile for LLM Task Processing System Evaluation

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet

# Evaluation parameters
MODELS := \
	"qwen2.5-coder-32b-instruct-q5km" \
	"gemma-2-27b-it-q5km" \
	"Codestral-22B-v0.1-q5km" \
	"internlm2_5-20b-chat-q5km" \
	"DeepSeek-Coder-V2-Lite-Instruct-q5km" \
	"qwen2.5-coder-14b-instruct-q5km" \
	"Qwen2.5-Coder-7B-Instruct-q80" \
	"Mistral-Nemo-Instruct-2407-q4km" \
	"gemma-2-9b-it-q4km" \
	"llava-v1.5-7b-q5km" \
	"Yi-Coder-9B-Chat-q4km" \
	"starcoder2-7b-q5km" \
	"Qwen2.5-Coder-7B-Instruct-q4km" \
	"llama-3.2-3b-instruct-q80" \
	"stable-code-instruct-3b-q80" \
	"Phi-3.1-mini-128k-instruct-q5km"

# Default target
all: setup test evaluate

# Setup project dependencies
setup:
	$(GOCMD) mod tidy
	go install github.com/golangci/golangci-lint@latest

# Run tests
test:
	$(GOTEST) -v ./...
	$(GOTEST) -race ./...

# Lint code
lint:
	golangci-lint run

# Evaluate all models
evaluate:
	@for model in $(MODELS); do \
		echo "Evaluating $$model"; \
		go run evaluation.go "$$model"; \
	done

# Generate performance profile
profile:
	$(GOTEST) -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof
	go tool pprof cpu.prof
	go tool pprof mem.prof

# Clean up generated files
clean:
	rm -f *.prof
	rm -f evaluation_results_*.json

# Help target
help:
	@echo "Available targets:"
	@echo "  all       - Setup, test, and evaluate all models"
	@echo "  setup     - Install dependencies"
	@echo "  test      - Run unit and race tests"
	@echo "  lint      - Run golangci-lint"
	@echo "  evaluate  - Run evaluation for all models"
	@echo "  profile   - Generate performance profiles"
	@echo "  clean     - Remove generated files"

.PHONY: all setup test lint evaluate profile clean help
