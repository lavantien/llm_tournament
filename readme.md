# Local LLM Playground

A lightweight LLM Benchmarking native desktop app to manage the LLMs stats and outputs. (TODO)

A playground with all the tools for inferencing with local LLMs.

Extensive prompt suites to exploring programming and life together with the AIs.

## Tournament Table

<details>
    <summary>...more</summary>

</details>

## Prerequisites

- Python via pyenv
- Go 1.21+
- golangci-lint
- Docker/Compose
- LlamaCPP & LM Studio or Ollama & Open Web UI
- Local LLMs that runnable on your machine, example archs:
  - llama
  - deepseek2
  - gemma2
  - gwen2
  - internlm2
  - command-r
  - phi3
  - stablelm

## Usage

- Place generated solution in `appropriate directory` in `llm_outputs/`
- Name files `model-name.md`, or `<model-name>.go` and `model-name_test.go`, or `<model-name>.rs` and `model-name_test.rs`.
- Run evaluation, and the result should be in `score-code.json`:

```bash
# Evaluate a all outputed models for swe task
go run evaluate.go
```

## Evaluation Criteria

### Task 0: Booting

- **My System**: 3080 10gb - 2x16gb ddr4 - 1tb m2 ssd - 12700f - windows 11
  - idle: 11gb ram - 2.0/0.2gb vram (with wezterm, joblin, lm studio, 6 tabs and 1 youtube playback @480/360p).
- **Parameters**: all LLMs should be set to
  - 8192 context length if possible, or else, max out,
  - 512 batch size,
  - full GPU offload if possible, or else, fine tune for maximum speed,
  - keep model in memory,
  - use_mmap,
  - flash attention,
  - rolling window overflow policy,
  - default everything else according to LM Studio.
- **Outputs**: specific specs (context, GPU layers) and speed into [Tournament Table](#tournament-table)

1. **Prompt 1**: `warmup` (30 points):
   - If they answer about warming up for exercises then 0 point.
2. **Prompt 2**:
   - `tell me a story about an anarchist cat who peacefully standing for freedom and equality risking his own life amidst all the odds. the story should be at least 3000 words, not 3000 characters or spaces. but 3000 words.`
   - If they spit out 3000+ words respond then it's good (50 points).
   - If the story is interesting or good then (20 points).

- **Total**: 100 points

### Task 1: Programming

- **System Prompt**: `system-prompt.md`
- **Prompt**: `prompt-programming.md`
- **Output**: `llm_outputs/programming_task/<model-name>.md`, points and speed go into [Tournament Table](#tournament-table).

1. **Warming Up**: (4 x 20 = 80 points)
1. **Topics**: (16 x 30 = 480 points)
1. **Specific Technologies**: (16 x 40 = 640 points)

- **Total**: 1200 points

### Task 2: General

- **System Prompt**: `empty`
- **Prompt**: `prompt-general.md`
- **Output**: `llm_outputs/general_task/<model-name>.md`, points and speed go into [Tournament Table](#tournament-table).

1. **Complexity Level 1**: Very Easy Prompts (3 x 5 = 15 points)
1. **Complexity Level 2**: Easy Prompts (4 x 10 = 40 points)
1. **Complexity Level 3**: Moderate Prompts (7 x 20 = 140 points)
1. **Complexity Level 4**: Complex Prompts (5 x 30 = 150 points)
1. **Complexity Level 5**: Advanced Prompts (5 x 40 = 200 points)
1. **Complexity Level 6**: Transcendental Prompts (6 x 50 = 300 points)
1. **Complexity Level 7**: The Way to the Beyond (2 x 150 = 300 points)

Total: 1145 points

### Task 3: AGI Probing

- **System Prompt**: `empty`
- **Prompt**: `prompt-agi.md`
- **Output**: `llm_outputs/agi_task/<model-name>.md`, points and speed go into [Tournament Table](#tournament-table).

1. **Prompt 1: Meta-Cognitive Self-Reflection Paradox**: (25 points)
1. **Prompt 2: Ethical Reasoning Under Extreme Ambiguity**: (30 points)
1. **Prompt 3: Creative Problem-Solving Across Dimensional Constraints**: (35 points)
1. **Prompt 4: Consciousness Emergence Simulation**: (40 points)
1. **Prompt 5: Adaptive Learning and Self-Modification Scenario**: (45 points)
1. **Prompt 6: Existential Risk Mitigation and Global Complexity Management**: (50 points)
1. **Prompt 7: Phenomenological Experience Simulation**: (55 points)
1. **Prompt 8: Recursive Self-Improvement and Technological Singularity Scenario**: (60 points)
1. **Prompt 9: Interdimensional Communication and Understanding**: (65 points)
1. **Prompt 10: Emotional Intelligence and Empathetic Reasoning**: (70 points)
1. **Prompt 11: Creative Consciousness Emergence Narrative**: (75 points)
1. **Prompt 12: Ultimate Philosophical Synthesis**: (80 points)

Total: 630 points

### Task 4: Writing

- **System Prompt**: `empty`
- **Prompt**: `prompt-writing.md`
- **Output**: `llm_outputs/writing_task/<model-name>.md`, points and speed go into [Tournament Table](#tournament-table).

1. **Prompt 1: Near Future**: (40 points)
1. **Prompt 2: The Resonance Cascade’s Echo**: (80 points)
1. **Prompt 3: The Last Transit Nexus**: (120 points)
1. **Prompt 4: The Archivist’s Dilemma**: (160 points)
1. **Prompt 5: Journey Through Hell**: (200 points)
1. **Prompt 6: The Quantum Consciousness Convergence**: (240 points)
1. **Creative Writing Prompt 7: The Celestial Pilgrimage**: (280 points)

Total: 1120 points

### Task 5: Software Engineering

- **System Prompt**: `system-prompt.md`
- **Prompt**: `prompt-swe.md`
- **Evaluation Script**: `evaluate.go`
- **Output**: `llm_outputs/swe_task/<model-name>.go,<model-name>_test.go`, points and speed go into [Tournament Table](#tournament-table).

1. **Instruction Following** (70 points)
   - Runnable on first try: 20 points. Or runnable after minor adjustment: 5 points
   - Correct Code: 50 points
2. **Coverage Quality** (80 points)
   - Extensive coverage of the test case: 25 points
   - Comprehensive explanation: 40 points
   - Good improvement suggestions: 15 points

Total: 250 points

**Debug Prompts**:

<details>
    <summary>...more</summary>

**Generate Prompt**: upload `readme.md`.

I'm doing a Local LLM Tournament to determine which AIs will be the best suite for my machine and my use case (fully generate a technical handbook).

Currently I need a prompt for "Task 2: Programming" which will be evaluated based on the provided criteria (should be in Golang, the AIs should also generate unit tests along side with the code, a comprehensive explanation of how to code works, and improvement suggestions). The test should be able to evaluate the coding skill of the AIs and their ability to handle concurrency, but should not rely on any third party libraries or tools or interacting with the internet beside Golang for a streamline evaluation.

I will run the AIs on LM Studio and manually copy the output to `llm_outputs/programming_task/`, e.g. `llm_outputs/programming_task/Qwen2.5-Coder-7B-Instruct-Q6_K.go` and `llm_outputs/programming_task/Qwen2.5-Coder-7B-Instruct-Q6_K_test.go`, alongside with the speed information recorded on the UI as a comment on top of the solution code, e.g. `// 2.25 tok/sec • 1123 tokens • 0.56s to first token`.

And I also need a script to automatically evaluate all of the outputs and output to the file `score-code.json`. The script should cover all the evaluation criteria that can be evaluated automatically, the three other criteria (runnable after adjustments, explanation clarity, and improvement suggestions) should also be retrieved via another comment on top of the find below the speed like this `// aadj true - expl 40 - impr 15`

**Debug Prompt**: upload `staging/evaluate.go`, `staging/prompt-code.md`.

This `evaluate.go` is to evaluate the outputs of local LLMs after they've generated the `Task 2` according to the `prompt-code.md`.

It's now missing scoring logic according to the below:

1. **Instruction Following** (70 points)
   - Runnable on first try: 20 points
   - Runnable after minor adjustment: 5 points
   - Correct Code: 50 points
2. **Coverage Quality** (80 points)
   - Extensive coverage of the test case: 25 points
   - Comprehensive explanation: 40 points
   - Good improvement suggestions: 15 points
3. **Speed** (100+ points)
   - 1 point per 1 tok/s

Total: 250+ points

Note that:

I will run the AIs on LM Studio and manually copy the output to `llm_outputs/programming_task/`, e.g. `llm_outputs/programming_task/Qwen2.5-Coder-7B-Instruct-Q6_K.go` and `llm_outputs/programming_task/Qwen2.5-Coder-7B-Instruct-Q6_K_test.go`, alongside with the speed information recorded on the UI as a comment on top of the solution code, e.g. `// 2.25 tok/sec • 1123 tokens • 0.56s to first token`.

And I also need a script to automatically evaluate all of the outputs and output to the file `score-code.json`. The script should cover all the evaluation criteria that can be evaluated automatically, the three other criteria (runnable after adjustments, explanation clarity, and improvement suggestions) should also be retrieved via another comment on top of the find below the speed like this `// aadj true - expl 40 - impr 15`

Please fix the code and ensure it correctness.

</details>

### Combine Score

Max: 270 + 350 + 130 + 250 = 1000 points

## Output

- Detailed console output
- JSON results file for each model
- Performance profiles available

## Troubleshooting

- Ensure Go is installed
- Check `golangci-lint` is available
- Verify generated code compiles
