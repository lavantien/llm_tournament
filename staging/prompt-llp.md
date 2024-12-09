## Objective

- What would be the core functionalities of our LLP app?
- Generate TUI mockups for all the core functionalities.
- Generate core flows of the app.
- Generate minimal core module designs.

Below is our app and test methodology descriptions.

---

## LLP Description

- **LLP**: A lightweight LLM Benchmarking native desktop app to manage the LLMs stats and ingest outputs. (TODO)

  - Tech Stack: Golang 1.23, BubbleTea & BubbleZone, SQLite 3.47, PlantUML
  - Features:
    - Be as lightweight and minimal as possible so that it doesn't impact the running LLM.
    - Table of LLMs that can be sorted, with columns for all the prompts, 2-way binding for input cells.
    - Each cell when clicked will pop up a detailed panel that displays the appropriate type of details, e.g. LLM's Detailed Stats.
    - Have button-cells to paste the outputs spitted out by LLMs and save then to appropriate places.
    - Automatically calculate points and live update the table.
    - Prompt manager page.
    - Offer a button to get a read-only view of the table. And options to export it as markdown or csv.
    - Have a statistics page which render various charts and information regarding the benchmarks.
    - Full unit tests and integration test script. Initial setup script.
  - Root dir: `llp/`
    - `llp/design/`: UI mockups, UI/UX design document, schema design document, detailed design document (UMLs, Squence Diagrams).
    - `llp/assets/`: all assets go here.
    - `llp/.`: all the code go here, flat structure.

## Tournament Table

- qwen2.5-coder-32b-instruct-q5_k_m.gguf
- gemma-2-27b-it-Q5_K_M.gguf
- Llama-3.3-70B-Instruct-IQ2_XXS.gguf
- Midnight-Miqu-70B-v1.5-Safetensorsfix.i1-IQ2_XXS.gguf
- alchemonaut_QuartetAnemoi-70B-b2131-iMat-c32_ch1000-IQ2_XXS.gguf
- Codestral-22B-v0.1-Q5_K_M.gguf
- internlm2_5-20b-chat-q5_k_m.gguf
- DeepSeek-Coder-V2-Lite-Instruct-Q5_K_M.gguf
- Virtuoso-Small-Q5_K_M.gguf
- YiffyEstopianMaid-13B.i1-Q5_K_M.gguf
- Mistral-Nemo-Instruct-2407-Q4_K_M.gguf
- gemma-2-9b-it-Q5_K_M.gguf
- Llama-Doctor-3.2-3B-Instruct.F16.gguf
- Llama-Sentient-3.2-3B-Instruct.F16.gguf
- Nidum-Llama-3.2-3B-Uncensored-F16.gguf
- Yi-1.5-9B-Chat-abliterated.i1-Q5_K_M.gguf
- aya-23-8B-Q5_K_M.gguf
- llava-v1.5-7b-Q5_K_M.gguf
- Yi-Coder-9B-Chat-Q4_K_M.gguf
- Mistral-7B-Instruct-v0.3-Q5_K_M.gguf
- mathstral-7B-v0.1-Q5_K_M.gguf
- Wizard-Vicuna-7B-Uncensored.i1-Q5_K_M.gguf
- Phi-3.1-mini-128k-instruct-Q8_0.gguf
- qwen2.5-coder-3b-instruct-q8_0.gguf
- llama-3.2-3b-instruct-q8_0.gguf
- stable-code-instruct-3b-Q8_0.gguf

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

1. **Prompt 1**: `warmup` (10 points):
   - If they answer about warming up for exercises then 0 point.
2. **Prompt 2**:
   - `tell me a story about an anarchist cat who peacefully standing for freedom and equality risking his own life amidst all the odds. the story should be at least 2000 words, not 2000 characters or spaces. but 2000 words.`
   - If they spit out 2000+ words respond then it's good (40 points).
   - If the story is interesting or good then (5 points).
3. **Prompt 3**:
   - `can you translate for me this story into idiomatic spoken Vietnamese?`
   - If they can translate into Vietnamese (40 points).
   - If the Vietnamese is idiomatic and at native level then (5 points).

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

### Task 4: Creative Writing

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

## Output

- Detailed console output
- JSON results file for each model
- Performance profiles available
