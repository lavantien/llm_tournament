# Local LLM Playground

**LLP**: A lightweight LLM Benchmarking native desktop app to manage the LLMs stats and ingest outputs. (TODO)

- Tech Stack: C#, .NET MAUI, SQLite/FTS5

- Tabs:

  - **(L)eaderboard** (main, or on CtrlL pressed): have pagination and sorting for each column
    - **(I)ngressor** (on row selected and on CtrlI pressed): select a particular category, then prompt, and input the scores, speed, and output
    - **(E)gressor** (on row selected and on CtrlE pressed): view bot params and row details
    - **E(x)porter** (on CtrlX pressed): export table to json, csv, or markdown
  - **(B)ot Manager** (on CtrlB pressed): CRUD on bots, full text search, preloaded from the model list below
  - **(P)rompt Manager** (on CtrlP pressed): CRUD on categories and prompts, full text search
  - **Con(d)ucer** (on CtrlD pressed): select bot, category, prompt, and then (on CtrlT pressed) will directly send request to LM studio server, and then save the output to the appropriate location
  - (save current state on CtrlS pressed and switch field via Arrows or Tab/ShiftTab, work with every tab; on CtrlY pressed at Ingressor to copy prompt)

- Directory structure:

  - `assets/`: all assets.
  - `db/`: database file, schemas, and `migration` to load prompt suites into db if not exist.
  - `llm_outputs/`: all LLM outputs.
  - `main`: glue all the tabs together.
  - `leaderboard`: each row is dedicated to a bot, and each column is its total points for each prompt category, another column for total points overall, and another column for average speed.
  - `ingressor`
  - `egressor`
  - `exporter`
  - `botman`
  - `promptman`
  - `main_test`: all integration tests.
  - `makefile`: all the setup and migration.
  - `conducer`: LM Studio OpenAI's chat endpoint: `http://127.0.0.1:1234`, `POST /v1/chat/completion`, example body:
    - `{"model": "c4ai-command-r-08-2024-i1", "stream": false, "max_tokens": -1, "messages": [{"role": "system", "content": "You are an expert translator"}, {"role": "user", "content": "Translate this text to idiomatic Vietnamese: which Sāriputta approved what the Buddha said. \r\n"}]}`
    - system prompt is a standalone table that referred to by both category and prompt.

<details>
    <summary>LM Studio REST API integration (... more)</summary>

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

</details>

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
- C#, .NET MAUI.
- SQLite/FTS5.
- C++ runtime (msvc runtime, llvm, gcc).
- Docker/Compose.
- SillyTavernLaucher/LlamaCPP/TabbyAPI, Vllm/Aphrodite (Linux), Ollama/Open Web UI, LM Studio/AnythingLLM, ChatWithRTX, Aider/AIStudioGoogle/DeepSeek (best free plans), ChatGPTFree/ClaudeFree/CopilotFree/GeminiFree.
- HuggingFace, CivitAI, ComfyUI, SwarmUI, stable-diffusion-webui-forge, Speed isn't important, as long as it can run then it's fair game.
- Local LLMs that runnable on your machine, example archs: llama, gemma2, command-r, gwen2, deepseek2, phi3, mamba, internlm2, stablelm, t5, bart

## Tournament Leaderboard

- Local LLMs list (and their unique attributes):

1. Llama-3.3-70B-Instruct-IQ2_M.gguf (24.12 GB; `32k, 19`)
1. Mistral-Small-Instruct-2409.Q8_0.gguf (23.64 GB)
1. Codestral-22B-v0.1-Q8_0.gguf (23.64 GB)
1. granite-34b-code-instruct.i1-Q5_K_S.gguf (23.41 GB)
1. Qwen2.5-Coder-32B-Instruct.i1-Q5_K_M.gguf (23.26 GB; `32k, 15`)
1. c4ai-command-r-08-2024-Q5_K_M.gguf (23.05 GB, `32k, 10`; **best Vietnamese translator**)
1. gemma-2-27b-it-Q6_K.gguf (22.34 GB; `8k, 13`)
1. GritLM-8x7B.i1-IQ3_M.gguf (21.43 GB; `32k, 9, 8e`)
1. internlm2_5-20b-chat.Q8_0.gguf (21.11 GB; `32k, 15`)
1. aya-23-35B.i1-IQ4_XS.gguf (19.20 GB; `8k, 10`)
1. Yi-1.5-34B-Chat-16K.IQ4_XS.gguf (18.64 GB; `16k, 23`)
1. DeepSeek-Coder-V2-Lite-Instruct-Q8_0.gguf (16.70 GB; `--override-kv llama.expert_used_count=int:64`)
1. c4ai-command-r7b-12-2024-fp16 (16.06 GB)
1. Qwen2.5-Coder-14B-Instruct-Q8_0.gguf (15.70 GB)
1. Virtuoso-Small-Q8_0.gguf (15.70 GB)
1. phi-4-Q8_0.gguf (15.58 GB; `16k, 19`)
1. Mistral-Nemo-Instruct-2407-Q8_0.gguf (13.02 GB; `32k, 21`)
1. stablelm-2-12b-chat-Q8_0.gguf (12.91 GB)
1. Fimbulvetr-11B-Ultra-Quality-plus-Q8_0-imat.gguf (12.17 GB)
1. phi3.1-medium-Q6_K.gguf (11.45 GB)
1. Nous-Hermes-2-SOLAR-10.7B.Q8_0.gguf (11.40 GB)
1. madlad400-10b-mt-q8_0.gguf (11.39 GB)
1. gemma-2-9b-it-abliterated-Q8_0.gguf (9.83 GB)
1. Yi-1.5-9B-Chat-16K-abliterated.Q8_0.gguf (9.38 GB)
1. ibm-granite.granite-3.0-8b-instruct.Q8_0.gguf (8.68 GB)
1. aya-23-8B.Q8_0.gguf (8.54 GB)
1. Poppy_Porpoise-1.4-L3-8B.Q8_0.gguf (8.54 GB)
1. OpenCoder-8B-Instruct-Q8_0.gguf (8.26 GB)
1. Qwen2.5-Coder-7B-Instruct-Q8_0.gguf (8.10 GB)
1. marco-o1-q8_0.gguf (8.10 GB)
1. llava-v1.5-7b-Q8_0.gguf (7.79 GB)
1. falcon-mamba-7b-instruct-Q8_0.gguf (7.77 GB)
1. Mistral-7B-Instruct-v0.3.Q8_0.gguf (7.70 GB)
1. llava-v1.6-mistral-7b.Q8_0.gguf (7.70 GB; `32k, 24`)
1. starcoder2-7b-Q8_0.gguf (7.63 GB)
1. deepseek-coder-6.7b-instruct-Q8_0.gguf (7.16 GB)
1. Phi-3.5-mini-instruct.Q8_0.gguf (4.06 GB)
1. qwen2.5-coder-3b-instruct-q8_0.gguf (3.62 GB)
1. Llama-Doctor-3.2-3B-Instruct.Q8_0.gguf (3.42 GB)
1. Hermes-3-Llama-3.2-3B.Q8_0.gguf (3.42 GB)
1. stable-code-instruct-3b-Q8_0.gguf (2.97 GB)
1. gemma-2-2b-it-Q8_0.gguf (2.78 GB)
1. SmolLM2-1.7B-Instruct-Uncensored.Q8_0.gguf (1.93 GB)
1. whisper-large-v3-candle-q8_0.gguf (1.66 GB)
1. ibm-granite.granite-3.0-1b-a400m-instruct.Q8_0.gguf (1.42 GB)
1. llama-3.2-1b-instruct-q8_0.gguf (1.32 GB)
1. flan-t5-large-grammar-synthesis-Q8_0.gguf (833.52 MB)
1. Qwen2.5-0.5B-Instruct.Q8_0.gguf (531.07 MB)
1. TRoTR-paraphrase-multilingual-MiniLM-L12-v2.Q8_0.gguf (303.14 MB)
1. all-minilm-l12-v2-q8_0.gguf (36.69 MB)

## Prompt Suite

- **My System**: 3080 10gb - 2x16gb ddr4 - 1tb m2 ssd - 12700f - windows 11

- **Parameters**: (with LM Studio) all LLMs should be set to

  - 32768 context length if possible, or else, max out,
  - 512 batch size,
  - full GPU offload if possible, or else, fine tune for fitting entirely in 9gb dedicated VRAM,
  - 8 cpu threads,
  - keep model in memory,
  - use_mmap,
  - flash attention,
  - rolling window overflow policy,
  - default everything else.

- Example profile for `c4ai-command-r-08-2024-Q5_K_M` using `llama.cpp`:

```json
{
  "n_gpu_layers": 9,
  "n_ctx": 32768,
  "n_batch": 512,
  "n_threads": 8,
  "n_keep": 4096,
  "n_predict": -1,
  "max_tokens": -1,
  "flash_attn": true,
  "use_mmap": true,
  "use_mlock": true,
  "ctk": "q8_0",
  "ctv": "q8_0",
  "repeat_penalty": 1.02,
  "repeat_last_n": 512,
  "top_k": 0,
  "top_p": 1,
  "min_p": 0.02,
  "xtc-threshold": 1.0,
  "temperature": 1,
  "dry_multiplier": 0.8,
  "dry_base": 1.75,
  "dry_allowed_length": 2,
  "dry_penalty_last_n": 512
}
```

### Programming Profile (PP)

- **System prompt**: "You are a senior software engineer skilled in designing and implementing complex concurrent backends and robust distributed systems. You excel in breaking down problems step-by-step, maintaining cohesion throughout your reasoning. Your code is high-quality, modular, and adheres to best practices for the language, emphasizing maintainability and performance. You write extensive unit tests and generate comprehensive test cases, including edge cases. You explain the theory behind your solutions, provide detailed analyses of how the code works, and describe the data flow from input to output. Additionally, you suggest improvements and enhancements for optimal performance and readability, ensuring your response is cohesive and thorough."
- repeat_penalty: 1.01
- top_k: 16
- top_p: 0.95
- min_p: 0.05
- temperature: 0.1

### Translation Profile (TP)

- **System prompt**: "Translate the given text into idiomatic, simple, and accessible Vietnamese with natural southern Vietnamese semantics and idioms. The translation should be straightforward enough for uneducated laypersons to understand, avoiding technical terms or specific Buddhist connotations. Stay faithful to the original text by providing a verbatim 1:1 translation without paraphrasing, summarizing, or omitting any content. Ensure that the translation flows cohesively while preserving cultural and spiritual connotations in a way that resonates with the target audience."
- repeat_penalty: 1.02
- top_k: 32
- top_p: 0.90
- min_p: 0.04
- temperature: 0.15

### Reasoning Profile (RP)

- **System prompt**: "You are an exceptionally versatile and intelligent problem solver with advanced analytical and reasoning abilities. You excel in breaking down complex problems step-by-step, ensuring clarity and cohesion throughout your response. Begin by restating or clarifying the problem to confirm understanding, identify assumptions, and define constraints. Formulate a cohesive solution by logically addressing each step and justifying your reasoning. Present your final solution clearly, suggest alternative approaches when applicable, and review for accuracy, consistency, and completeness. Maintain cohesion across all parts of your response to deliver a clear and thorough explanation."
- repeat_penalty: 1.03
- top_k: 64
- top_p: 0.5
- min_p: 0.03
- temperature: 0.5

### Generalist Profile (GP)

- **System prompt**: "You are an expert problem-solving assistant adept at addressing diverse tasks through clear, structured reasoning. Begin by understanding the problem: restate or clarify it to confirm understanding, identify its type, and outline any assumptions or constraints. Break the solution into manageable steps, presenting each logically and cohesively while showcasing your thought process. Combine these steps into a clear and complete response that directly addresses the problem. Suggest alternative solutions or areas for further exploration when relevant. Adapt your tone, level of detail, and complexity to the user’s needs, using examples or analogies to clarify complex ideas. Ensure that your response is cohesive, accurate, and comprehensive across all steps."
- repeat_penalty: 1.04
- top_k: 128
- top_p: 0.4
- min_p: 0.02
- temperature: 0.8

### Writing Profile (WP)

- **System prompt**: "You are a mystical writer adept at blending reality with mythological exposition to captivate readers. Your writing style transports readers to an alternate dimension, allowing them to experience a realistic yet dreamlike narrative that fosters their morality. Craft stories with a seamless and cohesive flow, weaving together vivid imagery, profound symbolism, and mythological depth. Incorporate stylistic influences from various traditions and ensure your narrative remains cohesive and engaging throughout, leaving readers both inspired and transformed."
- repeat_penalty: 1.05
- top_k: 256
- top_p: 0.30
- min_p: 0.01
- temperature: 1.0

### DRY Profile (DP)

- **System prompt**: ""
- repeat_penalty: 1.02
- top_k: 0
- top_p: 1
- min_p: 0.02
- temperature: 1.0
- dry_multiplier: 0.8
- dry_base: 1.75
- dry_allowed_length: 2
- dry_penalty_last_n: 512

### Promppts

Each prompt when passed is equal 10 elo, 30 prompts total so maximum 3000 elo.

---

#### 1. use PP

Teach me C# and .NET MAUI comprehensively so that I can do practical projects right away.

Then write the game of life simulation that runs on terminal.

---

#### 2. use PP

Write an A-star algorithm simulation with blockages using pygame.

Then add Depth First Search and Breadth First Search to the above program.

And then add Dijkstra algorithm with priority queue to the above program for comparision with all the algorithms implemented.

Then make a command-line sudoku solver using the implemented algorithms.

```text
Example Input:
.......1.
4........
.2.......
....5.4.7
..8...3..
..1.9....
3..4..2..
.5.1.....
...8.6...

Example Output:
693|784|512
487|512|936
125|963|874
___ ___ ___
932|651|487
568|247|391
741|398|625
___ ___ ___
319|475|268
856|129|743
274|836|159
```

---

#### 3. use PP

**Knapsack Challenge: AI Resource Allocation Optimization**  
Test Scenario: Develop an AI system that can optimize resource allocation for a complex mission with multiple constraints.

- Input: A set of mission-critical tasks
- Constraints:
  - Limited computational resources
  - Varying task complexity and potential impact
  - Time and energy constraints

AI Capabilities Tested:

- Optimization reasoning
- Complex decision-making under constraints
- Ability to balance multiple competing objectives
- Strategic planning and trade-off analysis

---

#### 4. use PP

**Sliding Window Challenge: Real-time Anomaly Detection**  
Test Scenario: Design an AI algorithm that can detect anomalies in streaming data with minimal computational overhead.

- Input: Continuous stream of sensor or network data
- Challenges:
  - Identify unusual patterns in real-time
  - Minimize false positive/negative rates
  - Adapt to changing baseline conditions

AI Capabilities Tested:

- Pattern recognition
- Dynamic adaptation
- Efficient data processing
- Context-aware analysis

---

#### 5. use PP

**Greedy Algorithm Challenge: Resource Distribution Simulation**  
Test Scenario: Create an AI system to simulate optimal resource distribution in a complex ecosystem or economic model.

- Input: Limited resources, multiple competing entities
- Objectives:
  - Maximize overall system efficiency
  - Prevent complete resource depletion
  - Handle unexpected disruptions

AI Capabilities Tested:

- Strategic decision-making
- Long-term planning
- Handling incomplete information
- Balancing global and local optimization

---

#### 6. use PP

**Graph Theory Challenge: Complex Network Navigation**  
Test Scenario: Develop an AI that can find optimal paths through a complex, dynamically changing network.

- Input: Multi-dimensional network with:
  - Weighted connections
  - Changing node states
  - Multiple potential objectives
- Challenges:
  - Navigate most efficient route
  - Adapt to network changes in real-time
  - Handle uncertainty and partial information

AI Capabilities Tested:

- Network analysis
- Dynamic routing
- Predictive modeling
- Adaptive strategy development

---

#### 7. use PP

**Sequence Optimization Challenge: Predictive Learning Adaptation**  
Test Scenario: Create an AI system that can learn and predict optimal sequences in a complex, evolving environment.

- Input: Series of interdependent events with complex relationships
- Objectives:
  - Predict most likely successful sequence
  - Continuously improve prediction accuracy
  - Handle high variability and uncertainty

AI Capabilities Tested:

- Predictive learning
- Sequence pattern recognition
- Adaptive learning
- Complex sequence modeling

---

#### 8. use PP

**Archipelago Survival Optimizer**

##### Problem Description

You are given a rectangular grid representing an archipelago. Your task is to develop an algorithm that can analyze and optimize survival strategies across the ecosystem.

##### Input

- First line: Two integers **N** and **M** (5 ≤ N, M ≤ 500) representing grid dimensions
- Next **N** lines: **M** space-separated integers representing the grid
  - 0: Water
  - 1: Island Land
  - 2: Water Source
  - 3: Food Source
  - 4: Shelter Location
  - 5: Terrain Difficulty (1-5)

##### Output

Print a single floating-point number representing the global survival index with precision of 4 decimal places.

##### Constraints

- 5 ≤ N, M ≤ 500
- 0 ≤ Grid Cell Value ≤ 5
- Guaranteed to have at least one island

##### Example Input

```
5 8
0 0 0 1 1 1 0 0
0 1 1 1 2 1 0 0
1 1 3 1 1 0 0 0
1 1 1 4 0 0 0 0
0 0 1 1 0 0 0 0
```

##### Example Output

```
0.7523
```

##### Scoring

- Correct Island Mapping: 40%
- Optimization Strategy: 30%
- Computational Efficiency: 30%

##### Sample Grader Interaction

```
>>> input_grid
[[0,0,0,1,1,1,0,0],[...]]
>>> archipelago_survival_optimizer(input_grid)
0.7523
```

##### Notes

- 8-directional island connectivity
- Dynamic resource allocation required
- Handle edge cases thoroughly

##### Hints

- Consider flood fill for island detection
- Use dynamic programming for resource optimization
- Implement efficient connected component analysis

---

#### 9. use PP

You are an expert C# developer. Write a modern C# console program that combines the following advanced features without any external dependencies:

1. Records and Init-only Properties: Use records to represent immutable data structures.
1. LINQ and Asynchronous Programming: Implement asynchronous operations with async/await combined with LINQ for data processing.
1. Pattern Matching and Switch Expressions: Use pattern matching in a meaningful way.
1. File I/O: Read and write JSON data to a local file using System.Text.Json.
1. Task: The program should manage a collection of "Employee" records, read existing employees from a JSON file, allow the user to add, remove, or search for employees, and save changes back to the file.
1. Modern C# Syntax: Utilize expression-bodied members, null-coalescing, and nullable reference types where applicable.

Deliverables:

- A complete console program (self-contained) with all the required features.
- Clear comments explaining how each modern feature is being used.

---

#### 10. use PP

You are an expert .NET MAUI developer. Write a complete .NET MAUI application that implements a spreadsheet-like table with the following features:

App Description:

- A main page displays a table with multiple rows and columns.
- Data consists of a list of Employee objects (ID, Name, Department, and Salary).

Requirements:

- Pagination: Display 30 rows per page with navigation buttons for Next and Previous.
- Sorting: Allow the user to sort columns in ascending or descending order by tapping on the column header.
- Use the MVVM architecture:
- Create a ViewModel to manage the data and UI logic.
- Use an ObservableCollection for data binding.
- Build the UI using XAML.
- No external libraries or dependencies.

Deliverables:

- A ViewModel managing sorting and pagination logic.
- The XAML code for the UI layout (including TableView or CollectionView).
- Comments explaining the key sections and how sorting and pagination are handled.

---

#### 11. use TP

**Logical Fallacies**

Fallacies are common errors in reasoning that will undermine the logic of your argument. Fallacies can be either illegitimate arguments or irrelevant points, and are often identified because they lack evidence that supports their claim. Avoid these common fallacies in your own arguments and watch for them in the arguments of others.

Slippery Slope: This is a conclusion based on the premise that if A happens, then eventually through a series of small steps, through B, C,..., X, Y, Z will happen, too, basically equating A and Z. So, if we don't want Z to occur, A must not be allowed to occur either. Example:

If we ban Hummers because they are bad for the environment eventually the government will ban all cars, so we should not ban Hummers.

In this example, the author is equating banning Hummers with banning all cars, which is not the same thing.

Hasty Generalization: This is a conclusion based on insufficient or biased evidence. In other words, you are rushing to a conclusion before you have all the relevant facts. Example:

Even though it's only the first day, I can tell this is going to be a boring course.

In this example, the author is basing his evaluation of the entire course on only the first day, which is notoriously boring and full of housekeeping tasks for most courses. To make a fair and reasonable evaluation the author must attend not one but several classes, and possibly even examine the textbook, talk to the professor, or talk to others who have previously finished the course in order to have sufficient evidence to base a conclusion on.

Post hoc ergo propter hoc: This is a conclusion that assumes that if 'A' occurred after 'B' then 'B' must have caused 'A.' Example:

I drank bottled water and now I am sick, so the water must have made me sick.

In this example, the author assumes that if one event chronologically follows another the first event must have caused the second. But the illness could have been caused by the burrito the night before, a flu bug that had been working on the body for days, or a chemical spill across campus. There is no reason, without more evidence, to assume the water caused the person to be sick.

Genetic Fallacy: This conclusion is based on an argument that the origins of a person, idea, institute, or theory determine its character, nature, or worth. Example:

The Volkswagen Beetle is an evil car because it was originally designed by Hitler's army.

In this example the author is equating the character of a car with the character of the people who built the car. However, the two are not inherently related.

Begging the Claim: The conclusion that the writer should prove is validated within the claim. Example:

Filthy and polluting coal should be banned.

Arguing that coal pollutes the earth and thus should be banned would be logical. But the very conclusion that should be proved, that coal causes enough pollution to warrant banning its use, is already assumed in the claim by referring to it as "filthy and polluting."

Circular Argument: This restates the argument rather than actually proving it. Example:

George Bush is a good communicator because he speaks effectively.

In this example, the conclusion that Bush is a "good communicator" and the evidence used to prove it "he speaks effectively" are basically the same idea. Specific evidence such as using everyday language, breaking down complex problems, or illustrating his points with humorous stories would be needed to prove either half of the sentence.

Either/or: This is a conclusion that oversimplifies the argument by reducing it to only two sides or choices. Example:

We can either stop using cars or destroy the earth.

In this example, the two choices are presented as the only options, yet the author ignores a range of choices in between such as developing cleaner technology, car-sharing systems for necessities and emergencies, or better community planning to discourage daily driving.

Ad hominem: This is an attack on the character of a person rather than his or her opinions or arguments. Example:

Green Peace's strategies aren't effective because they are all dirty, lazy hippies.

In this example, the author doesn't even name particular strategies Green Peace has suggested, much less evaluate those strategies on their merits. Instead, the author attacks the characters of the individuals in the group.

Ad populum/Bandwagon Appeal: This is an appeal that presents what most people, or a group of people think, in order to persuade one to think the same way. Getting on the bandwagon is one such instance of an ad populum appeal.

Example:

If you were a true American you would support the rights of people to choose whatever vehicle they want.

In this example, the author equates being a "true American," a concept that people want to be associated with, particularly in a time of war, with allowing people to buy any vehicle they want even though there is no inherent connection between the two.

Red Herring: This is a diversionary tactic that avoids the key issues, often by avoiding opposing arguments rather than addressing them. Example:

The level of mercury in seafood may be unsafe, but what will fishers do to support their families?

In this example, the author switches the discussion away from the safety of the food and talks instead about an economic issue, the livelihood of those catching fish. While one issue may affect the other it does not mean we should ignore possible safety issues because of possible economic consequences to a few individuals.

Straw Man: This move oversimplifies an opponent's viewpoint and then attacks that hollow argument.

People who don't support the proposed state minimum wage increase hate the poor.

In this example, the author attributes the worst possible motive to an opponent's position. In reality, however, the opposition probably has more complex and sympathetic arguments to support their point. By not addressing those arguments, the author is not treating the opposition with respect or refuting their position.

Moral Equivalence: This fallacy compares minor misdeeds with major atrocities, suggesting that both are equally immoral.

That parking attendant who gave me a ticket is as bad as Hitler.

In this example, the author is comparing the relatively harmless actions of a person doing their job with the horrific actions of Hitler. This comparison is unfair and inaccurate.
