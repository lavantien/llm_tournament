# Local LLM Playground

**LLP**: A lightweight LLM Benchmarking native desktop app to manage the LLMs stats and ingest outputs. (TODO)

- Tech Stack: Go, BubbleTea/Bubbles, SQLite/FTS5, Mermaid

- Tabs:

  - **(L)eaderboard** (main, or on CtrlL pressed)
    - **(I)ngressor** (on row selected and on CtrlI pressed): select a particular category, then prompt, and input the scores, speed, and output
    - **(E)gressor** (on row selected and on CtrlE pressed): view bot params and row details
    - **E(x)porter** (on CtrlX pressed): export table to json, csv, or markdown
  - **(B)ot Manager** (on CtrlB pressed): CRUD on bots, full text search, preloaded from LM Studio model list
  - **(P)rompt Manager** (on CtrlP pressed): CRUD on categories and prompts, full text search
  - **Con(d)ucer** (on CtrlD pressed): select bot, category, prompt, and then (on CtrlT pressed) will directly send request to LM studio server, and then save the output to the appropriate location
  - (save current state on CtrlS pressed and switch field via Arrows or Tab/ShiftTab, work with every tab; on CtrlY pressed at Ingressor to copy prompt)

- Directory structure:

  - `assets/`: all assets.
  - `db/`: database file, schemas, and `migration.go` to load prompt suites into db if not exist.
  - `llm_outputs/`: all LLM outputs.
  - `main.go`: glue all the tabs together.
  - `leaderboard.go`: each row is dedicated to a bot, and each column is its total points for each prompt category, another column for total points overall, and another column for average speed.
  - `ingressor.go`
  - `egressor.go`
  - `exporter.go`
  - `botman.go`
  - `promptman.go`
  - `main_test.go`: all integration tests.
  - `makefile`: all the setup and migration.
  - `conducer.go`: LM Studio OpenAI's chat endpoint: `http://127.0.0.1:1234`, `POST /v1/chat/completion`, example body:
    - `{"model": "c4ai-command-r-08-2024-i1", "stream": true, "max_tokens": -1, "messages": [{"role": "system", "content": "You are an expert translator"}, {"role": "user", "content": "Translate this text to idiomatic Vietnamese: which Sāriputta approved what the Buddha said. \r\n"}]}`
    - Stream the respond to output, then upon complete, will offer the option to save the result.

- `GET /api/v0/models` - List available models, example response:

```json
{
  "object": "list",
  "data": [
    {
      "id": "qwen2-vl-7b-instruct",
      "object": "model",
      "type": "vlm",
      "publisher": "mlx-community",
      "arch": "qwen2_vl",
      "compatibility_type": "mlx",
      "quantization": "4bit",
      "state": "not-loaded",
      "max_context_length": 32768
    },
    {
      "id": "meta-llama-3.1-8b-instruct",
      "object": "model",
      "type": "llm",
      "publisher": "lmstudio-community",
      "arch": "llama",
      "compatibility_type": "gguf",
      "quantization": "Q4_K_M",
      "state": "not-loaded",
      "max_context_length": 131072
    },
    {
      "id": "text-embedding-nomic-embed-text-v1.5",
      "object": "model",
      "type": "embeddings",
      "publisher": "nomic-ai",
      "arch": "nomic-bert",
      "compatibility_type": "gguf",
      "quantization": "Q4_0",
      "state": "not-loaded",
      "max_context_length": 2048
    }
  ]
}
```

- `GET /api/v0/models/{model}` - Get info about a specific model, example response:

```json
curl http://localhost:1234/api/v0/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "granite-3.0-2b-instruct",
    "messages": [
      { "role": "system", "content": "Always answer in rhymes." },
      { "role": "user", "content": "Introduce yourself." }
    ],
    "temperature": 0.7,
    "max_tokens": -1,
    "stream": false
  }'
```

- `POST /api/v0/chat/completions` - Chat Completions (messages -> assistant response), example request and response:

```json
curl http://localhost:1234/api/v0/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "granite-3.0-2b-instruct",
    "messages": [
      { "role": "system", "content": "Always answer in rhymes." },
      { "role": "user", "content": "Introduce yourself." }
    ],
    "temperature": 0.7,
    "max_tokens": -1,
    "stream": false
  }'
```

```json
{
  "id": "chatcmpl-i3gkjwthhw96whukek9tz",
  "object": "chat.completion",
  "created": 1731990317,
  "model": "granite-3.0-2b-instruct",
  "choices": [
    {
      "index": 0,
      "logprobs": null,
      "finish_reason": "stop",
      "message": {
        "role": "assistant",
        "content": "Greetings, I'm a helpful AI, here to assist,\nIn providing answers, with no distress.\nI'll keep it short and sweet, in rhyme you'll find,\nA friendly companion, all day long you'll bind."
      }
    }
  ],
  "usage": {
    "prompt_tokens": 24,
    "completion_tokens": 53,
    "total_tokens": 77
  },
  "stats": {
    "tokens_per_second": 51.43709529007664,
    "time_to_first_token": 0.111,
    "generation_time": 0.954,
    "stop_reason": "eosFound"
  },
  "model_info": {
    "arch": "granite",
    "quant": "Q4_K_M",
    "format": "gguf",
    "context_length": 4096
  },
  "runtime": {
    "name": "llama.cpp-mac-arm64-apple-metal-advsimd",
    "version": "1.3.0",
    "supported_formats": ["gguf"]
  }
}
```

- A playground for conducting (manual as of now) tournaments of the local LLMs.
- Extensive prepared prompt suites to exploring programming and life together with the AIs.

## Why?

- Because this is super fun and exciting and I like it. I love to learn from the AIs.
- I will use AIs as a copilot to write code and documentation.
- And I'm planing to generate a couple of 600-800 page handbooks for personal use and do translation/composing works.
- So I need to select the best candidate for the task, given the specs of my current machine. So, prompt suites and tournament pipeline is necessary
- Build a general pipeline for future works with local AIs.

## Dependencies

- Python via pyenv.
- C++ runtime (msvc runtime, llvm, gcc).
- Go & golangci-lint & BubbleTea/Bubbles.
- SQLite/FTS5.
- Docker/Compose.
- LM Studio/AnythingLLM, SillyTavernLaucher/TabbyAPI/LlamaCPP, Vllm/Aphrodite (Linux), Ollama/Open Web UI, ChatWithRTX, Aider/AIStudioGoogle/DeepSeek (best free plans), LMSYS/ChatGPTFree/ClaudeFree.
- HuggingFace, CivitAI, ComfyUI, SwarmUI, stable-diffusion-webui-forge, Speed isn't important, as long as it can run then it's fair game.
- Local LLMs that runnable on your machine, example archs: llama, gemma2, command-r, gwen2, deepseek2, phi3, internlm2, stablelm, t5, bart

## Tournament Leaderboard

- **My System**: 3080 10gb - 2x16gb ddr4 - 1tb m2 ssd - 12700f - windows 11

  - idle: 7.3gb ram - 0.9/0.1gb vram (with wezterm, lm studio).

- **Parameters**: (with LM Studio) all LLMs should be set to

  - 32768 context length if possible, or else, max out,
  - 512 batch size,
  - full GPU offload if possible, or else (\> 6.4gb), fine tune for fitting entirely in 9gb dedicated VRAM,
  - 16 threads,
  - keep model in memory,
  - use_mmap,
  - flash attention,
  - rolling window overflow policy,
  - 0.1 temp for programming and translating, 0.5 for general and agi probing, 0.9 for creative tasks.
  - default everything else.

```json
{
  "n_gpu_layers": -1,
  "n_ctx": 32768,
  "n_batch": 512,
  "n_threads": 16,
  "max_tokens": -1,
  "n_predict": -1
  "flash_attn": true,
  "use_mmap": true,
  "use_mlock": true,
  "temperature": 0.1,
  "top_k": 16,
  "top_p": 0.95,
  "min_p": 0.05,
  "repeat_penalty": 1.1,
  "frequency_penalty": 0,
  "presence_penalty": 0,
}
```

- Local LLMs list (and their unique attributes):

1. Llama-3.3-70B-Instruct-IQ2_M.gguf (24.12 GB; `32k, 19`)
1. Codestral-22B-v0.1-Q8_0.gguf (23.64 GB)
1. Qwen2.5-Coder-32B-Instruct.i1-Q5_K_M.gguf (23.26 GB; `32k, 15`)
1. c4ai-command-r-08-2024-Q5_K_M.gguf (23.05 GB, `32k, 9`)
1. gemma-2-27b-it-Q6_K.gguf (22.34 GB; `8k, 13`)
1. GritLM-8x7B.i1-IQ3_M.gguf (21.43 GB; `32k, 9, 8e`)
1. internlm2_5-20b-chat.Q8_0.gguf (21.11 GB; `32k, 15`)
1. aya-23-35B.i1-IQ4_XS.gguf (19.20 GB; `8k, 10`)
1. Yi-1.5-34B-Chat-16K.IQ4_XS.gguf (18.64 GB; `16k, 23`)
1. WizardCoder-33B-V1.1.i1-IQ4_XS.gguf (17.86 GB)
1. DeepSeek-Coder-V2-Lite-Instruct-Q8_0.gguf (16.70 GB)
1. Qwen2.5-Coder-14B-Instruct-Q8_0.gguf (15.70 GB)
1. Virtuoso-Small-Q8_0.gguf (15.70 GB)
1. TowerInstruct-13B-v0.1.Q8_0.gguf (13.83 GB)
1. vicuna-13b-v1.5-16k-Q8_0.gguf (13.83 GB)
1. Mistral-Nemo-Instruct-2407-Q8_0.gguf (13.02 GB; `32k, 21`)
1. stablelm-2-12b-chat-Q8_0.gguf (12.91 GB)
1. Fimbulvetr-11B-Ultra-Quality-plus-Q8_0-imat.gguf (12.17 GB)
1. phi3.1-medium-Q6_K.gguf (11.45 GB)
1. madlad400-10b-mt-q8_0.gguf (11.39 GB)
1. gemma-2-9b-it-abliterated-Q8_0.gguf (9.83 GB)
1. Yi-1.5-9B-Chat-16K-abliterated.Q8_0.gguf (9.38 GB)
1. aya-23-8B.Q8_0.gguf (8.54 GB)
1. Poppy_Porpoise-v0.7-L3-8B.Q8_0.gguf (8.54 GB)
1. OpenCoder-8B-Instruct-Q8_0.gguf (8.26 GB)
1. Qwen2.5-7B-Instruct_Q8_0.gguf (8.10 GB)
1. marco-o1-q8_0.gguf (8.10 GB)
1. falcon-mamba-7b-instruct-Q8_0.gguf (7.77 GB)
1. Mistral-7B-Instruct-v0.3.Q8_0.gguf (7.70 GB)
1. llava-v1.6-mistral-7b.Q8_0.gguf (7.70 GB)
1. Phi-3.5-3.8B-vision-instruct-Q8_0.gguf (4.71 GB)
1. Phi-3.5-mini-instruct.Q8_0.gguf (4.06 GB)
1. Qwen2.5-3B-Instruct-abliterated-Evol-CoT_SLERP.Q8_0.gguf (3.62 GB)
1. qwen2.5-coder-3b-instruct-q8_0.gguf (3.62 GB)
1. Llama-Doctor-3.2-3B-Instruct.Q8_0.gguf (3.42 GB)
1. Hermes-3-Llama-3.2-3B.Q8_0.gguf (3.42 GB)
1. stable-code-instruct-3b-Q8_0.gguf (2.97 GB)
1. whisper-large-v3-candle-q8_0.gguf (1.66 GB)
1. gemma-2-2b-it-IQ4_XS.gguf (1.57 GB)
1. llama-3.2-1b-instruct-q8_0.gguf (1.32 GB)
1. flan-t5-large-grammar-synthesis-Q8_0.gguf (833.52 MB)
1. TRoTR-paraphrase-multilingual-MiniLM-L12-v2.IQ4_XS.gguf (211.89 MB)
1. all-minilm-l12-v2-q8_0.gguf (36.69 MB)

## Usage

- Place generated solution in `appropriate directory` in `llm_outputs/`
- Name files `model-name.md`, or `<model-name>.go` and `model-name_test.go`, or `<model-name>.rs` and `model-name_test.rs`.
- Run evaluation, and the result should be in `llm_outputs/programming/scores/score-<model-name>.json`:

## Evaluation Criteria

### Combine Score

- **Max**: 385 + 1200 + 1145 + 630 + 1140 + 500 = 5000 points (5-star equivalent)

### Task 0: Booting

- **Outputs**: `llm_outputs/booting/<model-name>.md`, specific specs (context, GPU layers) and speed into LLP
- **System Prompt** (from Prompt 4 onward): inside prompt doc.
- **Prompts**: [`prommpt-booting.md`](./prompt-booting.md)

1. **Prompt 1**: Teaching (40 points):
2. **Prompt 2**: Programming and Explanation (45 points)
3. **Prompt 3**: Simple creative writing (40 points).
4. **Prompt 4**: Simple Translating (45 points).
5. **Prompt 5**: Creative poetry in foreign language (45 points).
6. **Prompt 6**: Real-world Translating Work (170 points).

- **Total**: 385 points

### Task 1: Programming

- **System Prompt**: [`system-prompt.md`](./system-prompt.md)
- **Prompt**: [`prompt-programming.md`](./prompt-programming.md)
- **Evaluation**: read through and evaluate them manually
- **Output**: `llm_outputs/programming/<model-name>.md`, points and speed go into LLP.

1. **Warming Up**: (4 x 20 = 80 points)
1. **Topics**: (16 x 30 = 480 points)
1. **Specific Technologies**: (16 x 40 = 640 points)

- **Total**: 1200 points

### Task 2: General

- **System Prompt**: inside prompt doc.
- **Prompt**: [`prompt-general.md`](./prompt-general.md)
- **Output**: `llm_outputs/general/<model-name>.md`, points and speed go into LLP.

1. **Complexity Level 1**: Very Easy Prompts (3 x 5 = 15 points)
1. **Complexity Level 2**: Easy Prompts (4 x 10 = 40 points)
1. **Complexity Level 3**: Moderate Prompts (7 x 20 = 140 points)
1. **Complexity Level 4**: Complex Prompts (5 x 30 = 150 points)
1. **Complexity Level 5**: Advanced Prompts (5 x 40 = 200 points)
1. **Complexity Level 6**: Transcendental Prompts (6 x 50 = 300 points)
1. **Complexity Level 7**: The Way to the Beyond (2 x 150 = 300 points)

- **Total**: 1145 points

### Task 3: AGI Probing

- **System Prompt**: inside prompt doc.
- **Prompt**: [`prompt-agi.md`](./prompt-agi.md)
- **Output**: `llm_outputs/agi/<model-name>.md`, points and speed go into LLP.

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

- **System Prompt**: inside prompt doc.
- **Prompt**: [`prompt-writing.md`](./prompt-writing.md)
- **Output**: `llm_outputs/writing/<model-name>.md`, points and speed go into LLP.

1. **Prompt 1: Near Future**: (40 points)
1. **Prompt 2: The Resonance Cascade’s Echo**: (80 points)
1. **Prompt 3: The Last Transit Nexus**: (90 points)
1. **Prompt 4: The Archivist’s Dilemma**: (100 points)
1. **Prompt 5: Romantic and Celibate**: (150 points)
1. **Prompt 6: Journey Through Hell**: (180 points)
1. **Prompt 7: The Quantum Consciousness Convergence**: (220 points)
1. **Prompt 8: The Celestial Pilgrimage**: (280 points)

- **Total**: 1140 points

### Task 5: Software Engineering

- **System Prompt**: [`system-prompt.md`](./system-prompt.md)
- **Prompt**: [`prompt-swe.md`](./prompt-swe.md)
- **Evaluation Script**: [`evaluate.go`](./evaluate.go)
- **Output**: `llm_outputs/swe/<model-name>.go,<model-name>_test.go`, points and speed go into LLP.
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

---
