# Local LLM Playground

- **LLP**: A lightweight LLM Benchmarking native desktop app to manage the LLMs stats and ingest outputs. (TODO)
  - Golang
  - Fyne GUI Toolkit
  - SQLite
- A repository playground with tools for inferencing with local LLMs.
- Extensive prompt suites to exploring programming and life together with the AIs.

## Why?

Because this is super fun and I like it.

## Tournament Table

- **LLM List**: `llm_list.md`

<details>
    <summary>...more</summary>

qwen2.5-coder-32b-instruct-q5_k_m-00001-of-00003.gguf  
gemma-2-27b-it-Q5_K_M.gguf  
Llama-3.3-70B-Instruct-IQ2_XXS.gguf  
alchemonaut_QuartetAnemoi-70B-b2131-iMat-c32_ch1000-IQ2_XXS.gguf  
Codestral-22B-v0.1-Q5_K_M.gguf  
internlm2_5-20b-chat-q5_k_m.gguf  
DeepSeek-Coder-V2-Lite-Instruct-Q5_K_M.gguf  
Virtuoso-Small-Q5_K_M.gguf  
YiffyEstopianMaid-13B.i1-Q5_K_M.gguf  
Mistral-Nemo-Instruct-2407-Q4_K_M.gguf  
gemma-2-9b-it-Q5_K_M.gguf  
Llama-Doctor-3.2-3B-Instruct.F16.gguf  
Llama-Sentient-3.2-3B-Instruct.F16.gguf  
Nidum-Llama-3.2-3B-Uncensored-F16.gguf  
Yi-1.5-9B-Chat-abliterated.i1-Q5_K_M.gguf  
aya-23-8B-Q5_K_M.gguf  
llava-v1.5-7b-Q5_K_M.gguf  
Yi-Coder-9B-Chat-Q4_K_M.gguf  
Mistral-7B-Instruct-v0.3-Q5_K_M.gguf  
mathstral-7B-v0.1-Q5_K_M.gguf  
Wizard-Vicuna-7B-Uncensored.i1-Q5_K_M.gguf  
Phi-3.1-mini-128k-instruct-Q8_0.gguf  
qwen2.5-coder-3b-instruct-q8_0.gguf  
llama-3.2-3b-instruct-q8_0.gguf  
stable-code-instruct-3b-Q8_0.gguf

(TODO)  
Here’s a prediction of the 5-star rating for each model in the provided list based on their performance across the various evaluation criteria, taking into account your local machine specs:

| Model Name                                                       | Booting (100 pts) | Programming (1200 pts) | General (1145 pts) | AGI Probing (630 pts) | Writing (1425 pts) | Software Engineering (500 pts) | Total (Points) | Predicted Rating |
| ---------------------------------------------------------------- | ----------------- | ---------------------- | ------------------ | --------------------- | ------------------ | ------------------------------ | -------------- | ---------------- |
| qwen2.5-coder-32b-instruct-q5_k_m-00001-of-00003.gguf            | 75                | 1050                   | 950                | 520                   | 1300               | 400                            | 4295           | ★★★★☆            |
| gemma-2-27b-it-Q5_K_M.gguf                                       | 85                | 1100                   | 1000               | 540                   | 1350               | 450                            | 4575           | ★★★★★            |
| Llama-3.3-70B-Instruct-IQ2_XXS.gguf                              | 90                | 1150                   | 1100               | 550                   | 1400               | 475                            | 5165           | ★★★★★            |
| alchemonaut_QuartetAnemoi-70B-b2131-iMat-c32_ch1000-IQ2_XXS.gguf | 80                | 1050                   | 980                | 510                   | 1280               | 460                            | 4360           | ★★★★☆            |
| Codestral-22B-v0.1-Q5_K_M.gguf                                   | 80                | 980                    | 940                | 490                   | 1260               | 470                            | 4220           | ★★★★☆            |
| internlm2_5-20b-chat-q5_k_m.gguf                                 | 85                | 1050                   | 970                | 500                   | 1300               | 450                            | 4355           | ★★★★☆            |
| DeepSeek-Coder-V2-Lite-Instruct-Q5_K_M.gguf                      | 70                | 900                    | 920                | 460                   | 1200               | 440                            | 3990           | ★★★☆☆            |
| Virtuoso-Small-Q5_K_M.gguf                                       | 70                | 900                    | 890                | 450                   | 1180               | 430                            | 3920           | ★★★☆☆            |
| YiffyEstopianMaid-13B.i1-Q5_K_M.gguf                             | 65                | 800                    | 850                | 440                   | 1150               | 420                            | 3825           | ★★★☆☆            |
| Mistral-Nemo-Instruct-2407-Q4_K_M.gguf                           | 90                | 1060                   | 1030               | 510                   | 1320               | 460                            | 4410           | ★★★★★            |
| gemma-2-9b-it-Q5_K_M.gguf                                        | 75                | 980                    | 900                | 470                   | 1250               | 440                            | 4115           | ★★★★☆            |
| Llama-Doctor-3.2-3B-Instruct.F16.gguf                            | 80                | 920                    | 880                | 450                   | 1180               | 430                            | 4060           | ★★★☆☆            |
| Llama-Sentient-3.2-3B-Instruct.F16.gguf                          | 85                | 950                    | 910                | 460                   | 1200               | 440                            | 4345           | ★★★★☆            |
| Nidum-Llama-3.2-3B-Uncensored-F16.gguf                           | 80                | 900                    | 890                | 440                   | 1150               | 420                            | 4080           | ★★★☆☆            |
| Yi-1.5-9B-Chat-abliterated.i1-Q5_K_M.gguf                        | 75                | 850                    | 860                | 430                   | 1120               | 400                            | 4135           | ★★★☆☆            |
| aya-23-8B-Q5_K_M.gguf                                            | 80                | 930                    | 880                | 460                   | 1190               | 440                            | 4090           | ★★★☆☆            |
| llava-v1.5-7b-Q5_K_M.gguf                                        | 70                | 820                    | 840                | 420                   | 1100               | 400                            | 4050           | ★★★☆☆            |
| Yi-Coder-9B-Chat-Q4_K_M.gguf                                     | 70                | 810                    | 830                | 420                   | 1080               | 390                            | 4000           | ★★★☆☆            |
| Mistral-7B-Instruct-v0.3-Q5_K_M.gguf                             | 75                | 940                    | 880                | 470                   | 1200               | 430                            | 4135           | ★★★☆☆            |
| mathstral-7B-v0.1-Q5_K_M.gguf                                    | 70                | 860                    | 830                | 430                   | 1100               | 400                            | 4090           | ★★★☆☆            |
| Wizard-Vicuna-7B-Uncensored.i1-Q5_K_M.gguf                       | 60                | 750                    | 800                | 400                   | 1050               | 380                            | 3840           | ★★★☆☆            |
| Phi-3.1-mini-128k-instruct-Q8_0.gguf                             | 65                | 790                    | 810                | 410                   | 1080               | 390                            | 4000           | ★★★☆☆            |
| qwen2.5-coder-3b-instruct-q8_0.gguf                              | 60                | 720                    | 750                | 380                   | 1020               | 370                            | 3500           | ★★★☆☆            |
| llama-3.2-3b-instruct-q8_0.gguf                                  | 65                | 740                    | 770                | 390                   | 1050               | 380                            | 3700           | ★★★☆☆            |
| stable-code-instruct-3b-Q8_0.gguf                                | 60                | 700                    | 720                | 370                   | 1000               | 360                            | 3550           | ★★★☆☆            |

</details>

## Prerequisites

- Python 3.11 via pyenv (specific version required by Ollama)
- Go 1.21+
- golangci-lint
- Docker/Compose
- LlamaCPP & LM Studio, KoboldCPP, or Ollama & Open Web UI
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
- Run evaluation, and the result should be in `llm_outputs/programming_task/scores/score-<model-name>.json`:

```bash
# Evaluate a all outputed models for swe task
go run evaluate.go
```

## Evaluation Criteria

### Combine Score

- **Max**: 100 + 1200 + 1145 + 630 + 1425 + 500 = 5000 points (5-star equivalent)

### Task 0: Booting

- **My System**: 3080 10gb - 2x16gb ddr4 - 1tb m2 ssd - 12700f - windows 11
  - idle: 10gb ram - 1.5/0.1gb vram (with wezterm, joblin, lm studio, 6 tabs and 1 youtube playback @480/360p).
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
- **Evaluation**: read through and evaluate them manually
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

- **Total**: 1145 points

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

- **Total**: 630 points

### Task 4: Writing

- **System Prompt**: `empty`
- **Prompt**: `prompt-writing.md`
- **Output**: `llm_outputs/writing_task/<model-name>.md`, points and speed go into [Tournament Table](#tournament-table).

1. **Prompt 1: Near Future**: (40 points)
1. **Prompt 2: The Resonance Cascade’s Echo**: (100 points)
1. **Prompt 3: The Last Transit Nexus**: (105 points)
1. **Prompt 4: The Archivist’s Dilemma**: (110 points)
1. **Prompt 5: Romantic and Celibate**: (225 points)
1. **Prompt 6: Journey Through Hell**: (240 points)
1. **Prompt 7: The Quantum Consciousness Convergence**: (260 points)
1. **Prompt 8: The Celestial Pilgrimage**: (345 points)

- **Total**: 1425 points

### Task 5: Software Engineering

- **System Prompt**: `system-prompt.md`
- **Prompt**: `prompt-swe.md`
- **Evaluation Script**: `evaluate.go`
- **Output**: `llm_outputs/swe_task/<model-name>.go,<model-name>_test.go`, points and speed go into [Tournament Table](#tournament-table).
- **Added Information**: on top of the code file, add the following information:
  - Line 1: Speed: `// 2.28 tok/sec • 476 tokens • 69.21s to first token • Stop: eosFound`
  - Line 2: Manual checks: `// adjustment: 0; explanation: 50; suggestions: 15`

1. **Instruction Following** (300 points)
   - Runnable on first try (automated, if `adjustment == 0`): 50 points.
   - Or runnable after minor adjustment (manual check, then set `adjustment = 1`, and rerun evaluation): 10 points
   - Or gabbage code (manual check, then set `adjustment = -1`, and disqualify the candidate): 0 point.
   - Correct Code (automated): 250 points
2. **Quality** (200 points)
   - Extensive coverage of the test case (automated): 120 points
   - Comprehensive explanation (manual check): 60 points
   - Good improvement suggestions (manual check): 20 points

- **Total**: 500 points

- **Debug Prompts**:

<details>
    <summary>...more</summary>

**Generate Prompt**: upload `llm_list.md`.

I'm doing a Local LLM Tournament to determine which AIs (the list is in `llm_list.md`) will be the best suite for my machine and my use case.

Currently I need a prompt for a software engineering task which will be evaluated based on the provided criteria (should be in Golang, the AIs should also generate unit tests along side with the code, a comprehensive explanation of how to code works, and improvement suggestions). The test should be able to evaluate the coding skill of the AIs and their ability to handle concurrency, but should not rely on any third party libraries or tools or interacting with the internet beside Golang for a streamline evaluation.

I will run the AIs on LM Studio and manually copy the output to `llm_outputs/programming_task/`, e.g. `llm_outputs/programming_task/Qwen2.5-Coder-7B-Instruct-Q6_K.go` and `llm_outputs/programming_task/Qwen2.5-Coder-7B-Instruct-Q6_K_test.go`, alongside with the speed information recorded on the UI as a comment on top of the solution code

And I also need a script to automatically evaluate the output and tests of a certain AI and output the result to the file `/llm_outputs/programming_task/scores/score-<model-name>.json`. The script should cover all the evaluation criteria that can be evaluated automatically, the three other criteria (runnable after adjustments, explanation clarity, and improvement suggestions) should also be retrieved via another comment on top of the file at 2nd line.

`<paste all the above>`

.

**Debug Prompt**: upload `staging/evaluate.go`, `staging/prompt-code.md`.

This `evaluate.go` is to evaluate the outputs of local LLMs after they've generated the `SWE Task` according to the `prompt-code.md`.

It's now missing scoring logic. Please fix the code and ensure it correctness.

</details>

## Output

- Detailed console output
- JSON results file for each model
- Performance profiles available

## Troubleshooting

- Ensure Go is installed
- Check `golangci-lint` is available
- Verify generated code compiles
