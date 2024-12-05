# LLM Task Processing System Evaluation

## Prerequisites

- Go 1.21+
- golangci-lint
- Make (optional but recommended)

## Setup

1. Install dependencies:

```bash
make setup
```

## Usage

- Place generated solution in current directory
- Name files `<model-name>.go` and `model-name_test.go`
- Run evaluation:

```bash
# Evaluate a specific model
go run evaluation.go "Model Name"

# Or use make to evaluate all models
make evaluate
```

## Evaluation Criteria

- Code Quality (40 points)
- Concurrency Performance (15 points)
- System Robustness (15 points)
- Testing Coverage (10 points)

## Output

- Detailed console output
- JSON results file for each model
- Performance profiles available

## Troubleshooting

- Ensure Go is installed
- Check `golangci-lint` is available
- Verify generated code compiles
