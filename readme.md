# Local LLM Tournament

## Tournament Table

| LLM                                         | Size-Context-Batch-Offload | Starting Score          | FSS    | Task 1 Speed (tok/s) | Task 2 Speed (tok/s) | Task 1 Quality | Task 2 Quality | Task 1 Score | Task 2 Score | Overall Score |
| ------------------------------------------- | -------------------------- | ----------------------- | ------ | -------------------- | -------------------- | -------------- | -------------- | ------------ | ------------ | ------------- |
| qwen2.5-coder-32b-instruct-q5_k_m.gguf      | 23.2 - 32k - 8k - 20/64    | 2.25 - 22.46 - 24.71    | 215.35 |                      |                      |                |                |              |              |               |
| gemma-2-27b-it-Q5_K_M.gguf                  | 19.4 - 8k - 8k - 24/46     | 3.03 - 8.16 - 11.19     | 69.53  |                      |                      |                |                |              |              |               |
| Codestral-22B-v0.1-Q5_K_M.gguf              | 15.7 - 32k - 8k - 36/56    | 4.23 - 4.72 - 8.95      | 39.65  |                      |                      |                |                |              |              |               |
| internlm2_5-20b-chat-q5_k_m.gguf            | 14.0 - 32k - 8k - 36/48    | 5.3 - 7 - 12.3          | 47.6   |                      |                      |                |                |              |              |               |
| DeepSeek-Coder-V2-Lite-Instruct-Q5_K_M.gguf | 11.8 - 32k - 8k - 13/27    | 19.64 - 8.02 - 27.66    | 66.66  |                      |                      |                |                |              |              |               |
| qwen2.5-coder-14b-instruct-q5_k_m.gguf      | 10.5 - 65k - 8k - 44/48    | 9.23 - 3.84 - 13.07     | 33.7   |                      |                      |                |                |              |              |               |
| Qwen2.5-Coder-7B-Instruct-Q8_0.gguf         | 8.1 - 16k - 8k - 28/28     | 63.67 - 4.7 - 68.37     | 109.79 |                      |                      |                |                |              |              |               |
| Mistral-Nemo-Instruct-2407-Q4_K_M.gguf      | 7.5 - 8k - 8k - 40/40      | 61.7 - 8.5 - 70.2       | 112.73 |                      |                      |                |                |              |              |               |
| Qwen2.5-Coder-7B-Instruct-Q6_K.gguf         | 6.3 - 32k - 8k - 28/28     | 72.38 - 7.42 - 79.8     | 128.14 |                      |                      |                |                |              |              |               |
| gemma-2-9b-it-Q4_K_M.gguf                   | 5.8 - 8k - 8k - 42/42      | 49.18 - 7.04 - 56.22    | 103.36 |                      |                      |                |                |              |              |               |
| llava-v1.5-7b-Q5_K_M.gguf                   | 5.4 - 4k - 4k - 32/32      | 87.65 - 3.58 - 91.23    | 146.5  |                      |                      |                |                |              |              |               |
| Yi-Coder-9B-Chat-Q4_K_M.gguf                | 5.3 - 32k - 8k - 48/48     | 67.86 - 5.54 - 73.4     | 134.94 |                      |                      |                |                |              |              |               |
| starcoder2-7b-Q5_K_M.gguf                   | 5.1 - 16k - 8k - 32/32     | 75.26 - 6.72 - 81.98    | 131.64 |                      |                      |                |                |              |              |               |
| Qwen2.5-Coder-7B-Instruct-Q4_K_M.gguf       | 4.7 - 32k - 8k - 28/28     | 82.83 - 6.36 - 89.19    | 143.22 |                      |                      |                |                |              |              |               |
| llama-3.2-3b-instruct-q8_0.gguf             | 3.4 - 32k - 8k - 28/28     | 108.73 - 9.5 - 118.23   | 144.84 |                      |                      |                |                |              |              |               |
| stable-code-instruct-3b-Q8_0.gguf           | 3.0 - 16k - 8k - 32/32     | 103.52 - 11.72 - 115.24 | 141.17 |                      |                      |                |                |              |              |               |
| Phi-3.1-mini-128k-instruct-Q5_K_M.gguf      | 2.8 - 16k - 8k - 32/32     | 65.31 - 7.66 - 72.97    | 89.39  |                      |                      |                |                |              |              |               |

## Prerequisites

- Go 1.21+
- golangci-lint
- Make (optional but recommended)
- LM Studio or Ollama or equivalent
- Local LLMs that runnable on your machine

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

### Task 0: Starting

**My System**: 3080 10gb - 2x16gb ddr4 - 1tb m2 ssd - 12700f ; idle: 15.3gb ram - 2.6/0.2gb vram  
**System Prompt**: `system-prompt.md`  
**Statring Prompt**: What are your capabilities regarding programming and writing technical documentation?  
**Outputs**: `/llm_outputs/starting_task/`

1. **Coverage Range** (50 points)

   - 1 point per 50 tokens

2. Speed (100+ points)

   - 1 point per 1 tok/s
   - Unable to boot even after tuning: 0 point

**Final Starting Score (FSS)** = (TPS + tokens/50) _ (1.07^B); example: `qwen2.5-coder-32b` has FSS = 24.71 _ (1.07^32) = 215.35

### Task 1: Documentation

**Input tokens**: 1,622  
**Expected output tokens**: 38,675  
**Detailed breakdown**:

<details>
    <summary>...more</summary>

Here’s the detailed token breakdown for each section of the **Senior Software Engineer Handbook**:

#### Token Estimates by Section:

1. **Section 1: Concurrency in Go** – **4,875 tokens**
2. **Section 2: Data Structures and Algorithms** – **15,600 tokens**
3. **Section 3: Functional Programming in Go** – **2,600 tokens**
4. **Section 4: Design Patterns** – **2,600 tokens**
5. **Section 5: Testing in Go** – **2,600 tokens**
6. **Section 6: Systems Design Foundations** – **2,600 tokens**
7. **Section 7: Advanced Software Design** – **3,250 tokens**
8. **Section 8: Observability and Performance** – **1,300 tokens**
9. **Section 9: Real-World Projects** – **1,950 tokens**
10. **Section 10: Career Development and Professional Growth** – **1,300 tokens**

#### Total Tokens for Full Document:

~**38,675 tokens**

#### Observations:

- **Section 2: Data Structures and Algorithms** is the largest, accounting for **40% of the total tokens**.
- Other sizable sections like **Concurrency in Go** and **Advanced Software Design** will also demand significant processing power.

</details>

**Prompt**: `prompt_doc.md`

1. **Instruction Following** (125 points)

   - Based on 125 distinct details across all sections
   - Each one covered is equal to one point

2. **Coverage Quality** (250 points)

   - Have proper code (if applied) and explanation
   - Extensive coverage of the subject

3. **Speed** (150+ points)

   - 75+ tok/s: 150 points
   - 0.5 TPS per 1 point

Total: 3x125 + 150 = 525 points

### Task 2: Programming

**Prompt**: `prompt_code.md`

1. **Instruction Following** (70 points)

   - Runnable on first try: 15 points
   - Runnable after minor adjustment: 5 points
   - Correct Code: 50 points

2. **Coverage Quality** (80 points)

   - Extensive coverage of the test case: 25 points
   - Comprehensive explanation: 40 points
   - Good improvement suggestions: 15 points

3. **Speed** (100+ points)

   - 1 point per 1 tok/s

Total: 250 points

### Combine Score

Max: 275 + 525 + 250 = 1000 points

## Output

- Detailed console output
- JSON results file for each model
- Performance profiles available

## Troubleshooting

- Ensure Go is installed
- Check `golangci-lint` is available
- Verify generated code compiles
