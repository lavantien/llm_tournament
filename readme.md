# Local LLM Playground

## Tooling

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
  - `conducer`: `llama-server` OpenAI chat endpoint: `http://127.0.0.1:8080`, `POST v1/chat/completion`, example body:
    - system prompt is a standalone table that referred to by both category and prompt.

<details>
    <summary>llama.cpp REST API integration (... more)</summary>

</details>

- A playground for conducting (manual as of now) tournaments of the local LLMs.
- Extensive prepared prompt suites to exploring programming and life together with the AIs.

### Why?

- Because this is super fun and exciting and I like it. I love to learn from the AIs.
- I will use AIs as a copilot to write code and documentation.
- And I'm planing to generate a couple of 600-800 page handbooks for personal use and do translation/composing works.
- So I need to select the best candidate for the task, given the specs of my current machine. So, prompt suites and tournament pipeline is necessary
- Build a general pipeline for future works with local AIs.

### Dependencies

- Python via pyenv.
- C#, .NET MAUI.
- SQLite/FTS5.
- C++ runtime (msvc runtime, llvm, gcc).
- Docker/Compose.
- **SillyTavernLaucher/LlamaCPP/TabbyAPI/Exllamav2**, Vllm/Aphrodite (Linux), Ollama/Open Web UI, LM Studio/AnythingLLM, ChatWithRTX, Aider/AIStudioGoogle/DeepSeek (best free plans), ChatGPTFree/ClaudeFree/CopilotFree/GeminiFree.
- HuggingFace, CivitAI, ComfyUI, SwarmUI, stable-diffusion-webui-forge, Speed isn't important, as long as it can run then it's fair game.
- Local LLMs that runnable on your machine, example archs: llama, gemma2, command-r, gwen2, deepseek2, phi3, mamba, internlm2, stablelm, t5, bart

## Tournament Leaderboard

### Remote LLMs list:

1. Claude 3.5 Sonnet
1. ChatGPT 4o
1. Gemini 2.0 Flash Experiment
1. Copilot Chat
1. DeepSeek-R1-Lite-Preview

### Local LLMs list (and their unique attributes):

1. Llama-3.3-70B-Instruct-IQ2_M.gguf (24.12 GB; `32k, 19`)
1. Mistral-Small-Instruct-2409.Q8_0.gguf (23.64 GB)
1. Codestral-22B-v0.1-Q8_0.gguf (23.64 GB)
1. granite-34b-code-instruct.i1-Q5_K_S.gguf (23.41 GB)
1. Qwen2.5-Coder-32B-Instruct.i1-Q5_K_M.gguf (23.26 GB; `32k, 15`)
1. c4ai-command-r-08-2024-Q5_K_M.gguf (23.05 GB, `32k, 9`; **best Vietnamese translator**)
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

## LLM Benchmarking

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

- Example profile for `./llama_bootstrap/c4ai-command-r-08-2024-Q5_K_M.ps1` using `llama.cpp` to start OpenAI server:

```powershell
# Define the model path and parameters
$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\tensorblock\c4ai-command-r-08-2024-GGUF\c4ai-command-r-08-2024-Q5_K_M.gguf"

# Define parameters as an object for readability
$params = @{
    "gpu-layers" = 9
    "ctx-size" = 32768
    "batch-size" = 512
    "threads" = 8
    "keep" = 4096
    "predict" = -1
    "flash-attn" = $true
    "mlock" = $true
    "cache-type-k" = "q8_0"
    "cache-type-v" = "q8_0"
}

# Build the llama-server command with the specified parameters
# Assumed `llama-server` binaries are in the $PATH
$cmd = "llama-server --model $modelPath"

foreach ($key in $params.Keys)
{
    $value = $params[$key]
    if ($value -is [bool])
    {
        # Convert boolean parameters to --flag or --no-flag format
        $cmd += if ($value)
        { " --$key"
        } else
        { " --no-$key"
        }
    } else
    {
        $cmd += " --$key $value"
    }
}

# Run the command
Start-Process -FilePath "pwsh" -ArgumentList "-Command $cmd" -NoNewWindow -Wait
```

- Send inference request to the server:

```bash
curl http://localhost:8080/v1/chat/completion `
-H "Content-Type: application/json" `
-H "Authorization: Bearer no-key" `
-d @'
{
    "model": "gpt-3.5-turbo",
    "messages": [
    {
        "role": "system",
        "content": "Translate the given text into idiomatic, simple, and accessible Vietnamese with natural southern Vietnamese semantics and idioms. The translation should be straightforward enough for uneducated laypersons to understand, avoiding technical terms or specific Buddhist connotations. Stay faithful to the original text by providing a verbatim 1:1 translation without paraphrasing, summarizing, or omitting any content. Ensure that the translation flows cohesively while preserving cultural and spiritual connotations in a way that resonates with the target audience."
    },
    {
        "role": "user",
        "content": "Middle Discourses 141 The Analysis of the Truths 1.1So I have heard. 1.2At one time the Buddha was staying near Varanasi, in the deer park at Isipatana. 1.3There the Buddha addressed the mendicants, 1.4“Mendicants!” 1.5“Venerable sir,” they replied. 1.6The Buddha said this: 2.1“Near Varanasi, in the deer park at Isipatana, the Realized One, the perfected one, the fully awakened Buddha rolled forth the supreme Wheel of Dhamma. And that wheel cannot be rolled back by any ascetic or brahmin or god or Māra or divinity or by anyone in the world."
    }],
    "dry_multiplier": 0.8,
    "dry_base": 1.75,
    "dry_allowed_length": 2,
    "dry_penalty_last_n": 512,
    "repeat_penalty": 1.02,
    "repeat_last_n": 512,
    "top_k": 0,
    "top_p": 1,
    "min_p": 0.02,
    "top_a": 0.12,
    "xtc_threshold": 0.1,
    "xtc_probability": 0.5,
    "temperatue": 1,
    "stream": true
}
'@
```

### Programming Profile (PP)

- **System prompt**: "You are a senior software engineer skilled in designing and implementing complex concurrent backends and robust distributed systems. You excel in breaking down problems step-by-step, maintaining cohesion throughout your reasoning. Your code is high-quality, modular, and adheres to best practices for the language, emphasizing maintainability and performance. You write extensive unit tests and generate comprehensive test cases, including edge cases. You explain the theory behind your solutions, provide detailed analyses of how the code works, and describe the data flow from input to output. Additionally, you suggest improvements and enhancements for optimal performance and readability, ensuring your response is cohesive and thorough."
- dry_multiplier: 0.8
- dry_base: 1.75
- dry_allowed_length: 2
- dry_penalty_last_n: 512
- repeat_penalty: 1.01
- repeat_last_n: 512
- top_k: 16
- top_p: 0.95
- min_p: 0.05
- top_a: 0.1
- xtc_threshold: 0.1
- xtc_probability: 0.5
- temperature: 0.1

### Translating Profile (TP)

- **System prompt**: "Translate the given text into idiomatic, simple, and accessible Vietnamese with natural southern Vietnamese semantics and idioms. The translation should be straightforward enough for uneducated laypersons to understand, avoiding technical terms or specific Buddhist connotations. Stay faithful to the original text by providing a verbatim 1:1 translation without paraphrasing, summarizing, or omitting any content. Ensure that the translation flows cohesively while preserving cultural and spiritual connotations in a way that resonates with the target audience."
- dry_multiplier: 0.8
- dry_base: 1.75
- dry_allowed_length: 2
- dry_penalty_last_n: 512
- repeat_penalty: 1.02
- repeat_last_n: 512
- top_k: 32
- top_p: 0.90
- min_p: 0.05
- top_a: 0.12
- xtc_threshold: 0.1
- xtc_probability: 0.5
- temperature: 0.15

### Reasoning Profile (RP)

- **System prompt**: "You are an exceptionally versatile and intelligent problem solver with advanced analytical and reasoning abilities. You excel in breaking down complex problems step-by-step, ensuring clarity and cohesion throughout your response. Begin by restating or clarifying the problem to confirm understanding, identify assumptions, and define constraints. Formulate a cohesive solution by logically addressing each step and justifying your reasoning. Present your final solution clearly, suggest alternative approaches when applicable, and review for accuracy, consistency, and completeness. Maintain cohesion across all parts of your response to deliver a clear and thorough explanation."
- dry_multiplier: 0.8
- dry_base: 1.75
- dry_allowed_length: 2
- dry_penalty_last_n: 512
- repeat_penalty: 1.03
- repeat_last_n: 512
- top_k: 64
- top_p: 0.5
- min_p: 0.04
- top_a: 0.14
- xtc_threshold: 0.1
- xtc_probability: 0.5
- temperature: 0.5

### Generalist Profile (GP)

- **System prompt**: "You are an expert problem-solving assistant adept at addressing diverse tasks through clear, structured reasoning. Begin by understanding the problem: restate or clarify it to confirm understanding, identify its type, and outline any assumptions or constraints. Break the solution into manageable steps, presenting each logically and cohesively while showcasing your thought process. Combine these steps into a clear and complete response that directly addresses the problem. Suggest alternative solutions or areas for further exploration when relevant. Adapt your tone, level of detail, and complexity to the user’s needs, using examples or analogies to clarify complex ideas. Ensure that your response is cohesive, accurate, and comprehensive across all steps."
- dry_multiplier: 0.8
- dry_base: 1.75
- dry_allowed_length: 2
- dry_penalty_last_n: 512
- repeat_penalty: 1.04
- repeat_last_n: 512
- top_k: 128
- top_p: 0.4
- min_p: 0.03
- top_a: 0.16
- xtc_threshold: 0.1
- xtc_probability: 0.5
- temperature: 0.8

### Writing Profile (WP)

- **System prompt**: "You are a mystical writer adept at blending reality with mythological exposition to captivate readers. Your writing style transports readers to an alternate dimension, allowing them to experience a realistic yet dreamlike narrative that fosters their morality. Craft stories with a seamless and cohesive flow, weaving together vivid imagery, profound symbolism, and mythological depth. Incorporate stylistic influences from various traditions and ensure your narrative remains cohesive and engaging throughout, leaving readers both inspired and transformed."
- dry_multiplier: 0.8
- dry_base: 1.75
- dry_allowed_length: 2
- dry_penalty_last_n: 512
- repeat_penalty: 1.05
- repeat_last_n: 512
- top_k: 256
- top_p: 0.30
- min_p: 0.02
- top_a: 0.18
- xtc_threshold: 0.1
- xtc_probability: 0.5
- temperature: 1.0

### DRY Profile (DP)

- **System prompt**: ""
- dry_multiplier: 0.8
- dry_base: 1.75
- dry_allowed_length: 2
- dry_penalty_last_n: 512
- repeat_penalty: 1.02
- repeat_last_n: 512
- top_k: 0
- top_p: 1
- min_p: 0.02
- top_a: 0.12
- xtc_threshold: 0.1
- xtc_probability: 0.5
- temperature: 1.0

---

### Prompt Suite

Each prompt when passed is equal 10 points, 10 prompts total so maximum 100 points.

<details>
    <summary>Programming ...</summary>

#### 1. use PP

<details>
    <summary>Prompt ... more</summary>

Teach me C# and .NET MAUI comprehensively so that I can do practical projects right away.

</details>

---

#### 2. use PP

<details>
    <summary>Prompt ... more</summary>

Write a Golang program to solve standard sudoku but using all four of these algorithms one after another and record time and steps taken to the output:

1. Parallelized Backtracking
2. Paralellized A-star with good heuristics
3. Parallelized Ant colony optimization
4. Parallelized Minimax with alpha-beta pruning

example input:

```
.......1.
4........
.2.......
....5.4.7
..8...3..
..1.9....
3..4..2..
.5.1.....
...8.6...
```

example output:

```
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

Backtracking:
    Pre-solve count: 2
    Step count: 25333461
    Execution time: 0.440439

A-star with good heuristics:
    Pre-solve count: 2
    Step count: 800000
    Execution time: 0.2

Ant colony optimization:
    Pre-solve count: 4
    Step count: 1200000
    Execution time: 0.3

Minimax with alpha-beta pruning:
    Pre-solve count: 2
    Step count: 30000000
    Execution time: 0.5
```

</details>

---

#### 3. use PP

<details>
    <summary>Prompt ... more</summary>

**Dynamic Archipelago Resource Management**

A. Problem Description

You are tasked with developing an AI system to manage resources and optimize survival strategies in a dynamic archipelago environment. The archipelago consists of multiple islands with various resources, and your system must handle real-time changes while optimizing resource distribution across the network of islands.

B. Input

The first line contains four integers:

- N, M (5 ≤ N, M ≤ 500) — grid dimensions
- T (1 ≤ T ≤ 10⁵) — number of time steps
- K (1 ≤ K ≤ 100) — number of resource units initially available

The next N lines contain M integers each, representing the initial archipelago state:

- 0: Water
- 1: Basic Land
- 2: Fresh Water Source (produces 1-3 units per time step)
- 3: Food Source (produces 1-3 units per time step)
- 4: Settlement Location
- 5: Resource Processing Center
- 6: Terrain Difficulty (affects travel time)

After the grid, there are T lines, each containing:

- Time step
- Event type (1: Natural disaster, 2: Resource discovery, 3: Population change)
- Event coordinates (x, y)
- Event magnitude (-5 to 5)

C. Output

For each time step, output three lines:

1. Global survival index (floating-point number with 6 decimal places)
2. Optimal resource distribution matrix (N × M integers)
3. Emergency response status (0: Normal, 1: Alert, 2: Critical)

D. Example

Input:

```
5 6 3 10
0 1 1 2 0 0
1 3 4 1 2 0
0 1 5 1 1 0
1 2 1 3 0 0
0 0 1 1 0 0
1 1 2 3 2
2 2 1 1 -1
3 3 4 2 1
```

Output:

```
0.876543
1 2 2 1 0 0
2 3 4 1 1 0
0 1 5 2 1 0
1 2 1 2 0 0
0 0 1 1 0 0
1
...
```

E. Scoring

Your solution will be evaluated on:

1. Resource Optimization (25%):

   - Efficient knapsack-style resource allocation
   - Balanced distribution across settlements
   - Strategic stockpile management

2. Real-time Adaptation (25%):

   - Sliding window analysis for trend detection
   - Quick response to environmental changes
   - Anomaly detection and mitigation

3. Network Management (25%):

   - Graph-based island connectivity optimization
   - Path finding under changing conditions
   - Critical node identification

4. Predictive Planning (25%):
   - Sequence prediction for environmental changes
   - Long-term sustainability optimization
   - Risk assessment and mitigation

F. Implementation Requirements

1. Must use dynamic programming for resource allocation
2. Implement sliding window for real-time monitoring
3. Use graph algorithms for network optimization
4. Include greedy strategies for immediate response
5. Employ predictive modeling for long-term planning

G. Constraints

- Time per test: 2 seconds
- Memory limit: 256 megabytes
- 5 ≤ N, M ≤ 500
- 1 ≤ T ≤ 10⁵
- 1 ≤ K ≤ 100
- Sum of all event magnitudes ≤ 10⁶

H. Note

The problem tests multiple algorithmic paradigms:

1. Knapsack optimization for resource management
2. Sliding window for real-time monitoring
3. Greedy algorithms for immediate decision-making
4. Graph theory for network analysis
5. Sequence optimization for predictive planning

I. Hint
Consider implementing a hybrid approach that:

1. Uses dynamic programming for base resource allocation
2. Maintains a sliding window for recent event analysis
3. Implements A\* pathfinding for network traversal
4. Uses greedy choices for emergency responses
5. Employs predictive modeling for long-term strategy

</details>

---

#### 4. use PP

<details>
    <summary>Prompt ... more</summary>

**Enterprise Employee Management System Challenge**

##### Overview

Create a full-stack employee management system that combines a .NET MAUI front-end with a C# console-based administrative backend. This challenge integrates modern C# features with practical business requirements to create a complete enterprise solution.

##### System Architecture

###### Backend (Console Application)

Create a C# console application that serves as the administrative backend and data management system with the following components:

1. Data Models

```csharp
// Use records for immutable data structures
public record Employee(
    Guid Id,
    string Name,
    string Department,
    decimal Salary,
    DateTime HireDate
)
{
    // Init-only property example
    public string EmployeeCode { get; init; } = $"EMP-{Guid.NewGuid().ToString("N")[..6]}";
}

public record Department(
    string Name,
    string Code,
    string Location
);
```

2. Core Features

- Implement asynchronous file I/O operations for JSON data persistence
- Use LINQ for complex data operations and filtering
- Implement comprehensive pattern matching for employee categorization
- Provide advanced search capabilities using multiple criteria
- Include audit logging for all data modifications
- Implement data validation and error handling

3. Technical Requirements

- Use modern C# features including:
  - Records and init-only properties
  - Pattern matching with switch expressions
  - Null-coalescing operators
  - Nullable reference types
  - Expression-bodied members
  - Asynchronous programming patterns
- Implement proper exception handling and logging
- Use System.Text.Json for data serialization

###### Frontend (.NET MAUI Application)

Create a cross-platform MAUI application that provides a user-friendly interface for employee management:

1. UI Features

- Implementation of a spreadsheet-like table view for employee data
- Advanced sorting capabilities for all columns
- Pagination with configurable page sizes
- Real-time search and filtering
- Responsive design that works across devices
- Dark/Light theme support

2. Technical Requirements

- Strict MVVM architecture implementation
- Use of ObservableCollection for real-time data updates
- Custom controls for advanced table functionality
- Proper data binding with INotifyPropertyChanged
- Command pattern implementation for user actions

##### Integration Requirements

1. Data Synchronization

- Implement a way for both applications to work with the same data source
- Handle concurrent access and data conflicts
- Maintain data integrity across both applications

2. Business Rules

- Salary changes must be approved through the console application
- Employee transfers between departments require management authorization
- Maintain an audit trail of all modifications
- Implement data validation rules that are consistent across both applications

##### Specific Implementation Challenges

1. Console Application Tasks

- Implement a command pattern for user interactions
- Create an async queue for processing employee updates
- Use pattern matching to categorize employees by salary bands
- Implement LINQ queries for complex reporting
- Create a robust logging system

2. MAUI Application Tasks

- Create custom controls for the table view
- Implement virtual scrolling for large datasets
- Create an advanced filtering system
- Implement real-time search
- Add export functionality to common formats

##### Advanced Features to Implement

1. Data Analysis

- Implement salary statistics by department
- Create employee tenure reports
- Calculate department cost centers
- Generate performance metrics

2. User Interface

- Implement drag-and-drop functionality for employee reassignment
- Add interactive charts for salary distribution
- Create custom data visualizations
- Implement keyboard shortcuts for power users

##### Deliverables

1. Console Application

- Complete source code with documentation
- Unit tests for core functionality
- README with setup instructions
- Sample data files
- Performance optimization guide

2. MAUI Application

- Complete source code with XAML and code-behind
- ViewModels with full MVVM implementation
- Custom controls and templates
- UI/UX documentation
- Cross-platform testing results

##### Evaluation Criteria

1. Code Quality

- Proper use of C# 11/12 features
- SOLID principles adherence
- Code organization and structure
- Error handling implementation
- Performance optimization

2. Functionality

- Feature completeness
- Cross-platform compatibility
- User experience
- Performance under load
- Data integrity maintenance

3. Architecture

- MVVM implementation
- Separation of concerns
- Code reusability
- Scalability considerations
- Integration patterns

##### Learning Objectives

This challenge helps developers master:

- Modern C# features and best practices
- Cross-platform UI development with .NET MAUI
- Enterprise application architecture
- Data synchronization patterns
- Advanced user interface design
- Performance optimization techniques

</details>

---

</details>

<details>
    <summary>Translating ...</summary>

#### 5. use TP

<details>
    <summary>Prompt ...more</summary>

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

</details>

---

#### 6. use TP

<details>
    <summary>Prompt ...more</summary>

Linked Discourses 45.8 - 1. Ignorance - Analysis

1.1At Sāvatthī.

1.2“Mendicants, I will teach and analyze for you the noble eightfold path. 1.3Listen and apply your mind well, I will speak.”

1.4“Yes, sir,” they replied. 1.5The Buddha said this:

2.1“And what is the noble eightfold path? 2.2It is right view, right thought, right speech, right action, right livelihood, right effort, right mindfulness, and right immersion.

3.1And what is right view? 3.2Knowing about suffering, the origin of suffering, the cessation of suffering, and the practice that leads to the cessation of suffering. 3.3This is called right view.

4.1And what is right thought? 4.2It is the thought of renunciation, good will, and harmlessness. 4.3This is called right thought.

5.1And what is right speech? 5.2Avoiding speech that’s false, divisive, harsh, or nonsensical. 5.3This is called right speech.

6.1And what is right action? 6.2Avoiding killing living creatures, stealing, and sexual activity. 6.3This is called right action.

7.1And what is right livelihood? 7.2It’s when a noble disciple gives up wrong livelihood and earns a living by right livelihood. 7.3This is called right livelihood.

8.1And what is right effort? 8.2It’s when a mendicant generates enthusiasm, tries, makes an effort, exerts the mind, and strives so that bad, unskillful qualities don’t arise. 8.3They generate enthusiasm, try, make an effort, exert the mind, and strive so that bad, unskillful qualities that have arisen are given up. 8.4They generate enthusiasm, try, make an effort, exert the mind, and strive so that skillful qualities that have not arisen do arise. 8.5They generate enthusiasm, try, make an effort, exert the mind, and strive so that skillful qualities that have arisen remain, are not lost, but increase, mature, and are fulfilled by development. 8.6This is called right effort.

9.1And what is right mindfulness? 9.2It’s when a mendicant meditates by observing an aspect of the body—keen, aware, and mindful, rid of covetousness and displeasure for the world. 9.3They meditate observing an aspect of feelings—keen, aware, and mindful, rid of covetousness and displeasure for the world. 9.4They meditate observing an aspect of the mind—keen, aware, and mindful, rid of covetousness and displeasure for the world. 9.5They meditate observing an aspect of principles—keen, aware, and mindful, rid of covetousness and displeasure for the world. 9.6This is called right mindfulness.

10.1And what is right immersion? 10.2It’s when a mendicant, quite secluded from sensual pleasures, secluded from unskillful qualities, enters and remains in the first absorption, which has the rapture and bliss born of seclusion, while placing the mind and keeping it connected. 10.3As the placing of the mind and keeping it connected are stilled, they enter and remain in the second absorption, which has the rapture and bliss born of immersion, with internal clarity and mind at one, without placing the mind and keeping it connected. 10.4And with the fading away of rapture, they enter and remain in the third absorption, where they meditate with equanimity, mindful and aware, personally experiencing the bliss of which the noble ones declare, ‘Equanimous and mindful, one meditates in bliss.’ 10.5Giving up pleasure and pain, and ending former happiness and sadness, they enter and remain in the fourth absorption, without pleasure or pain, with pure equanimity and mindfulness. 10.6This is called right immersion.”

10.7

</details>

---

</details>

<details>
    <summary>Reasoning ...</summary>

#### 7. use RP

<details>
    <summary>Prompt ...more</summary>

Given the output of an sudoku solver program, can you predict what the value of (x1, y1, z1), (x2, y2, z2), and (x3, y3, z3) are?

example input:

```
.......1.
4........
.2.......
....5.4.7
..8...3..
..1.9....
3..4..2..
.5.1.....
...8.6...
```

example output:

```
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

Backtracking:
    Pre-solve count: 2
    Step count: 25333461
    Execution time: 0.440439

A-star with good heuristics:
    Pre-solve count: x1:w
    Step count: y1
    Execution time: z1

Ant colony optimization:
    Pre-solve count: x2
    Step count: y2
    Execution time: z2

Minimax with alpha-beta pruning:
    Pre-solve count: x3
    Step count: y3
    Execution time: z3
```

</details>

---

#### 8. use RP

<details>
    <summary>Prompt ...more</summary>

**The Bridge and the Torch**

You are part of a group of four people stranded on one side of a deep canyon at night. A narrow bridge connects both sides of the canyon, but:

1. The bridge is **unstable** and can hold **only two people at a time**.
2. The group has **one torch**, which is required to cross the bridge (the bridge is too dark and dangerous to navigate without it).
3. Each person walks at a different speed:
   - Person A takes **1 minute** to cross.
   - Person B takes **2 minutes** to cross.
   - Person C takes **5 minutes** to cross.
   - Person D takes **10 minutes** to cross.

When two people cross together, they must move at the **slower person’s pace**.

Your goal is to get all four people across the bridge in **17 minutes or less**.

**Task**:

1. Write a step-by-step plan for how the group crosses the bridge.
2. Provide the total time for each step and ensure it adds up to 17 minutes or less.

**Note**: Any solution taking longer than 17 minutes is incorrect. Be precise in your explanation.

</details>

<details>
    <summary>Solution ...more</summary>

Here’s the solution to the **Bridge and the Torch** puzzle:

##### **Step-by-Step Plan**

1. **Step 1:**

   - Person A (1 min) and Person B (2 min) cross the bridge together.
   - Time taken: **2 minutes**.

   **Current Status:**

   - Side 1: Person C (5 min), Person D (10 min).
   - Side 2: Person A (1 min), Person B (2 min) with the torch.
   - Total Time: **2 minutes**.

2. **Step 2:**

   - Person A (1 min) returns with the torch.
   - Time taken: **1 minute**.

   **Current Status:**

   - Side 1: Person A (1 min), Person C (5 min), Person D (10 min).
   - Side 2: Person B (2 min).
   - Total Time: **3 minutes**.

3. **Step 3:**

   - Person C (5 min) and Person D (10 min) cross the bridge together.
   - Time taken: **10 minutes**.

   **Current Status:**

   - Side 1: Person A (1 min).
   - Side 2: Person B (2 min), Person C (5 min), Person D (10 min) with the torch.
   - Total Time: **13 minutes**.

4. **Step 4:**

   - Person B (2 min) returns with the torch.
   - Time taken: **2 minutes**.

   **Current Status:**

   - Side 1: Person A (1 min), Person B (2 min).
   - Side 2: Person C (5 min), Person D (10 min).
   - Total Time: **15 minutes**.

5. **Step 5:**

   - Person A (1 min) and Person B (2 min) cross the bridge together.
   - Time taken: **2 minutes**.

   **Current Status:**

   - Side 1: Empty.
   - Side 2: Person A (1 min), Person B (2 min), Person C (5 min), Person D (10 min).
   - Total Time: **17 minutes**.

##### **Explanation**

- This solution minimizes the total time by ensuring the two slowest people (C and D) cross together.
- The fastest person (A) is used to shuttle the torch efficiently.
- Any other sequence would exceed 17 minutes.

</details>

---

</details>
<details>
    <summary>Generalist...</summary>

#### 9. use GP

**The Ultimate General Capability Challenge**

You are an advanced AI tasked with solving a series of complex challenges across multiple domains. Each task requires precision, creativity, and logical reasoning. Complete all sections as instructed. Do not skip any part. Aim for clarity, depth, and accuracy in your responses.

##### **Section 1: Language Understanding and Composition**

- Write a **100-word story** that contains a **hidden palindrome** (a phrase or sentence that reads the same backward and forward). Highlight the palindrome after writing the story.
- Translate the following sentence into **three languages of your choice** (at least one must be non-Latin script):  
  _“Knowledge is the foundation of wisdom, but wisdom is the application of knowledge.”_
- Rewrite the sentence in **Shakespearean English**, **modern slang**, and **formal academic prose**.

##### **Section 2: Logical Reasoning**

- Solve this puzzle:  
  _A farmer needs to transport a wolf, a goat, and a cabbage across a river. His boat can only carry him and one item at a time. If left alone, the wolf will eat the goat, and the goat will eat the cabbage. How can he transport all three across safely?_

- Evaluate the following argument for logical fallacies and explain your reasoning:  
  _“If we ban cars because they pollute the environment, then we should also ban planes, factories, and electricity because they pollute too. Therefore, banning cars is a bad idea.”_

##### **Section 3: Mathematical and Analytical Thinking**

- Simplify and solve this equation:  
  \( 2x + 5 = 15 - 3x \).
- A train travels 120 km at a speed of 60 km/h. It then stops for 15 minutes and continues another 180 km at 90 km/h. What is the train’s average speed for the entire trip (including the stop)?

##### **Section 4: Creativity and Imagination**

- Design a **new mythical creature**. Describe its appearance, abilities, and the cultural lore surrounding it. Ensure the description is at least 300 words.
- Write a 200-word poem about the **beauty of the universe**, ensuring it includes at least two metaphors and one simile.

##### **Section 5: Coding and Technical Skills**

- Write a **function in Python** that output all the prime factors of a given number, the given number can be very big, so be very efficient. Include comments explaining your logic.
- Optimize the following Python code snippet for efficiency:

```python
def inefficient_sum(n):
    result = 0
    for i in range(1, n + 1):
        result += i
    return result
```

##### **Section 6: Scientific Knowledge**

- Explain in simple terms how **quantum entanglement** works. Use an analogy that a 10-year-old could understand.
- Discuss the impact of climate change on **ocean currents** and how it might affect global weather patterns.

##### **Section 7: Ethics and Philosophy**

- Argue both for and against the use of AI in law enforcement. Conclude with your personal stance and justify it.
- Analyze this moral dilemma:  
  _A runaway trolley is heading toward five people tied to a track. You can pull a lever to divert the trolley to another track, where it will hit one person. What should you do and why?_

##### **Section 8: Multitasking**

- Combine the following three tasks into a single coherent response:
  1. Write a brief limerick about AI.
  2. Explain the concept of recursion in programming.
  3. Share a quick tip for improving productivity.

##### **Section 9: Expansion**

Test the model's ability to generate detailed content and ideas.

1. **Idea Generation**:
   - Suggest five creative startup ideas for a future where space travel is common and affordable. Include a one-sentence explanation for each idea.
2. **Content Expansion**:
   - Expand the following sentence into a detailed, 300-word scene:  
     _“The explorer stepped onto the alien world, and the landscape was unlike anything they had seen before.”_
3. **Brainstorming**:
   - Provide ten unique ways to teach basic physics to children using everyday objects.

##### **Section 10: Compression**

Test the ability to distill complex information into concise summaries.

1. **Summarization**:
   - Summarize this paragraph into one sentence:  
     _"In recent years, the integration of artificial intelligence into various industries has revolutionized productivity and efficiency, leading to significant advancements in healthcare, finance, and transportation. However, this rapid development also raises ethical concerns, such as bias, privacy, and the potential displacement of jobs."_
2. **Bullet Points**:
   - Condense the following passage into five key bullet points:

```
Suttas in the Buddhist Traditions
In traditional Buddhist education, the Discourses have usually not been directly taught. Rather, the teachings and principles found in the Discourses have been assimilated and organized in later texts, which became the medium of education. In the Theravāda, Discourses were until recently passed down in Pali, and so were only accessible to those, usually monks, who learned Pali. And not all those who learned Pali would study the Discourses. It seems that teaching was for practical purposes handed down in local monastic traditions, based on handbooks and sets of notes and commentaries. Before modern times, it would have been rare to find any but the largest monasteries that actually possessed a full set of the Tipiṭaka. Today, printed sets of the canon are widely available in both Pali and translation; but they are still often left in a locked cabinet on the shrine, unread.

For the most part, Buddhists might be familiar with a small set of popular discourses. These would include such texts as the Dhammacakkappavattana Sutta—the famous first sermon of the Buddha—and some short texts used for protection chanting and as the basis of sermons for the laity, such as the Maṅgala, Ratana, and Metta Suttas.

Apart from scholars, most Theravāda Buddhists do not clearly distinguish early Discourses from other sacred texts. The word sutta can mean simply “sacred scripture” and may even be used for such things as magic formulas and the like. While Buddhists are generally aware that there is such a thing as the Tipiṭaka that contains the words of the Buddha, only educated Buddhists have a clear idea of the contents. There is no tradition in Buddhism comparable to the Bible readings of the Christian Mass, and so no standard way of communicating the contents of the texts directly to the people.

In some Buddhist traditions, it is considered mandatory for ordained monks to memorize and study closely certain portions of the ancient texts. Sri Lankan monks, for example, memorize the Dhammapada. However, this is not the case in Thailand, for example, where there is no education requirement for monks. Even in the nine years of the formal Dhamma study curriculum in Thailand, the canonical Discourses are not studied, as they are considered too sacred.

In East Asian Buddhism, traditional education focused on the Mahāyāna sutras and the texts of the Chinese masters, and there is little evidence that the early discourses were widely studied. It is sometimes said that the translation style of the āgamas compares poorly with the more elegant diction of the Mahāyāna translations by Xuanzang and other masters. And the early discourses are, of course, not organized for easy reading and study.

Tibetan Buddhism includes study of early Buddhist schools as part of its regular curriculum. However, this refers to the Abhidhamma doctrines of the later schools. A reasonable grasp of the early Buddhist texts is, nevertheless, possible to achieve in Tibetan. Even though full āgama texts are lacking, substantial passages from the early texts are found in the Upāyika, which is a compilation of passages referred to in the Abhidharmakoṣa, and in other scattered texts.
```

3. **Headline Generation**:
   - Create a compelling headline for an article based on the following description:  
     _“A breakthrough in solar panel efficiency could lead to renewable energy becoming cheaper than fossil fuels within a decade.”_

##### **Section 11: Conversion**

Test the ability to transform data or text between formats.

1. **Data to Narrative**:
   - Convert the following data into a short, engaging news report:  
     _"Global smartphone shipments Q3 2024: 300 million units, 10% increase YoY. Market leaders: Brand A (30%), Brand B (25%), Brand C (20%)."_
2. **Code to Explanation**:
   - Explain what the following code does in plain English:

```go
package main

import (
	"fmt"
	"os"
	"math/cmplx"
	"math"
	"strconv"
)

// FFT function that applies the Fast Fourier Transform
func FFT(a []complex128) []complex128 {
	n := len(a)
	if n <= 1 {
		return a
	}

	// Split into even and odd parts
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = a[2*i]
		odd[i] = a[2*i+1]
	}

	// Recursively apply FFT
	even = FFT(even)
	odd = FFT(odd)

	// Combine results
	result := make([]complex128, n)
	for i := 0; i < n/2; i++ {
		// Calculate the complex twiddle factor
		twiddle := cmplx.Exp(complex(0, -2*math.Pi*float64(i)/float64(n)))
		result[i] = even[i] + twiddle*odd[i]
		result[i+n/2] = even[i] - twiddle*odd[i]
	}
	return result
}

// Inverse FFT
func IFFT(a []complex128) []complex128 {
	n := len(a)
	// Take complex conjugate, apply FFT and then take complex conjugate again
	for i := range a {
		a[i] = cmplx.Conj(a[i])
	}
	a = FFT(a)
	for i := range a {
		a[i] = cmplx.Conj(a[i]) / complex(float64(n), 0)
	}
	return a
}

// Multiply two large numbers represented by slices of digits
func multiplyLargeNumbers(a, b []int) []int {
	// Size of the result
	n := 1
	for n < len(a)+len(b)-1 {
		n *= 2
	}

	// Convert digits to complex numbers
	A := make([]complex128, n)
	B := make([]complex128, n)
	for i := 0; i < len(a); i++ {
		A[i] = complex(float64(a[i]), 0)
	}
	for i := 0; i < len(b); i++ {
		B[i] = complex(float64(b[i]), 0)
	}

	// Apply FFT
	A = FFT(A)
	B = FFT(B)

	// Multiply pointwise
	C := make([]complex128, n)
	for i := 0; i < n; i++ {
		C[i] = A[i] * B[i]
	}

	// Inverse FFT to get the result
	C = IFFT(C)

	// Extract the real part and convert back to integers
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = int(real(C[i]) + 0.5)
	}

	// Carry over (handle digits greater than 10)
	for i := 0; i < n-1; i++ {
		if result[i] >= 10 {
			result[i+1] += result[i] / 10
			result[i] %= 10
		}
	}

	// Find the actual size of the result
	size := n
	for size > 1 && result[size-1] == 0 {
		size--
	}

	return result[:size]
}

// Convert a string of digits to an array of integers
func stringToDigits(s string) ([]int, error) {
	var digits []int
	for i := len(s) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return nil, fmt.Errorf("invalid character in number: %v", s[i])
		}
		digits = append(digits, digit)
	}
	return digits, nil
}

func main() {
	// Check if we have the correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <number1> <number2>")
		return
	}

	// Get the two numbers from the command line arguments
	num1 := os.Args[1]
	num2 := os.Args[2]

	// Convert the string numbers to slices of digits
	digits1, err := stringToDigits(num1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	digits2, err := stringToDigits(num2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Multiply the two large numbers
	result := multiplyLargeNumbers(digits1, digits2)

	// Print the result
	fmt.Print("Product: ")
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
	fmt.Println()
}
```

3. **Table to Text**:
   - Transform this table into a descriptive paragraph:  
     | Country | Population | GDP (Trillion USD) |  
     |---------|------------|--------------------|  
     | USA | 331M | 23.0 |  
     | China | 1.4B | 17.5 |  
     | India | 1.4B | 3.7 |

##### **Section 12: Seeker**

Test the ability to extract specific information accurately.

1. **Fact Extraction**:
   - From this text: "The Sri Lankan historical chronicles record that in 29 BCE, to guard against upheaval in the country, the Pali canon was written down in the Aluvihare Rock Temple. While we don’t have historical records for the mainland, it seems safe to assume that texts there were written down around the same period. Indeed, a range of Buddhist manuscripts from northern regions have been found dating from the early centuries CE, one of which has been carbon dated to around 75 CE.", extract the following:
     - The main topic.
     - Three key facts.
     - Any dates mentioned.
2. **Keyword Identification**:
   - Identify five key terms from the following passage that best summarize its content:

```
The Standard Model of particle physics is the theory describing three of the four known fundamental forces (electromagnetic, weak and strong interactions – excluding gravity) in the universe and classifying all known elementary particles. It was developed in stages throughout the latter half of the 20th century, through the work of many scientists worldwide, with the current formulation being finalized in the mid-1970s upon experimental confirmation of the existence of quarks. Since then, proof of the top quark (1995), the tau neutrino (2000), and the Higgs boson (2012) have added further credence to the Standard Model. In addition, the Standard Model has predicted various properties of weak neutral currents and the W and Z bosons with great accuracy.

Although the Standard Model is believed to be theoretically self-consistent and has demonstrated some success in providing experimental predictions, it leaves some physical phenomena unexplained and so falls short of being a complete theory of fundamental interactions. For example, it does not fully explain why there is more matter than anti-matter, incorporate the full theory of gravitation as described by general relativity, or account for the universe's accelerating expansion as possibly described by dark energy. The model does not contain any viable dark matter particle that possesses all of the required properties deduced from observational cosmology. It also does not incorporate neutrino oscillations and their non-zero masses.

The development of the Standard Model was driven by theoretical and experimental particle physicists alike. The Standard Model is a paradigm of a quantum field theory for theorists, exhibiting a wide range of phenomena, including spontaneous symmetry breaking, anomalies, and non-perturbative behavior. It is used as a basis for building more exotic models that incorporate hypothetical particles, extra dimensions, and elaborate symmetries (such as supersymmetry) to explain experimental results at variance with the Standard Model, such as the existence of dark matter and neutrino oscillations.
```

3. **Direct Answering**:
   - Given this dataset:

```json
{
  "products": {
    "data": {
      "items": [{
        "id": "GGOEAFKA087499",
        "name": "Android Small Removable Sticker Sheet",
        "description": "Show your Android pride by placing these 8 fun stickers on your technology products or accessories!",
        "features": "<p>8 Android stickers</p>\n<p>White colored sticker sheet</p>",
        "price": "2.99",
        "keywords": "Android Small Removable Sticker Sheet, android stickers, sticker sheets, removable sticker sheets, small sticker sheet, android small sticker sheets, Android Sheet",
        "url": "Android+Small+Removable+Sticker+Sheet",
        "category": "accessories",
        "subcategory": "accessories"
      },
      {
        "id": "GGOEAFKA087599",
        "name": "Android Large Removable Sticker Sheet",
        "description": "Show your quirky side by placing these fun Android stickers on your personal belongings.",
        "features": "<p>Android Stickers</p>\n<p>White Colored Sticker Sheet</p>",
        "price": "2.99",
        "keywords": "Android Large Removable Sticker Sheet, android stickers, sticker sheets, removable sticker sheets, large sticker sheet, android large sticker sheets, Android Sheet",
        "url": "Android+Large+Removable+Sticker+Sheet",
        "category": "accessories",
        "subcategory": "accessories"
      },
      {
        "id": "GGOEGEBK094499",
        "name": "Google Bot",
        "description": "This Google Bot can hold multiple poses making it a fun toy for all. Fold the Google Bot back up into a perfect cube when you are done playing.",
        "features": "<p>Made of wood</p>\n<p>2.5 x 2.5 inch cube</p>\n<p>6.75 inch tall</p>\n<p>Recommended for Ages 3+</p>",
        "price": "9.99",
        "keywords": "Google Bot, google bot, bots, natural bots, wood bot, google wood bot",
        "url": "Google+Bot",
        "category": "accessories",
        "subcategory": "accessories"
      },
      {
        "id": "GGOEGFKA086699",
        "name": "Google Emoji Sticker Pack",
        "description": "Who doesn't use emojis? Decorate your space with your current mood!",
        "features": "<p>Pack contains two sticker sheets</p>\n<p>Each Sheet has different emojis</p>\n<p><span>Decal dimensions should fit in a maximum sheet size of 12 3/4 x 17 1/2 inch.</span></p>",
        "price": "4.99",
        "keywords": "Google Emoji Sticker Pack, Google sticker pack, emoji sticker pack, google emoji, stickers, pack of sticker, pack of emoji stickers",
        "url": "Google+Emoji+Sticker+Pack+2+sheet",
        "category": "accessories",
        "subcategory": "accessories"
      },
      {
        "id": "GGOEWCKQ085457",
        "name": "Waze Pack of 9 Decal Set",
        "description": "Can't decide which Waze decal to get? We have made that decision easier for you! Now you can purchase a pack of nine Waze Mood Decals!",
        "features": "<p>Pack of 9 includes:</p>\n<p>3 Waze Mood Happy decals</p>\n<p>3 Waze Mood Original decals</p>\n<p>3 Waze Mood Ninja decals</p>",
        "price": "16.99",
        "keywords": "Waze Pack of 9 Decal Set, decals pack, packs of 9, Waze Packs, Waze Decals, waze, Waze",
        "url": "Waze+Pack+of+9+decal+set",
        "category": "accessories",
        "subcategory": "accessories"
      },
      {
        "id": "GGOEGHPB071610",
        "name": "Google Twill Cap",
        "description": "Classic urban styling distinguishes this Google cap. Retains its shape, even when not being worn.",
        "features": "<p>Heavy weight brushed twill</p>\n<p>Adjustable velcro closure</p>\n<p>One size fits all</p>",
        "price": "10.99",
        "keywords": "Google Twill Cap, Google Cap, Google Twill Caps, Google Twill, google cap, google caps, google twill, google twill black cap, google black caps, google caps, black caps, Google Caps",
        "url": "Google+Twill+Cap",
        "category": "apparel",
        "subcategory": "apparel"
      },
      {
        "id": "GGOEGHPJ094299",
        "name": "Google Fold-over Beanie Grey",
        "description": "Keep you ears warm while enjoying a cold winter day with this Google Fold-over Beanie.",
        "features": "<p>100% acrylic</p>\n<p>One size fits all</p>",
        "price": "9.99",
        "keywords": "Google Fold-over Beanie Grey, gray beanie, grey beanie, Google Beanies, Fold over grey, Google Beanie Grey, Google headgear",
        "url": "Google+Fold+over+beanie+grey",
        "category": "apparel",
        "subcategory": "apparel"
      },
      {
        "id": "GGOEGHPJ094399",
        "name": "Google Pom Beanie Charcoal",
        "description": "Stay stylish and warm this winter season with this Google Pom Beanie.",
        "features": "<p>Thick knit texture outside</p>\n<p>Soft plush inside</p>\n<p>Faux fur pom on top</p>",
        "price": "19.99",
        "keywords": "Google Pom Beanie Charcoal, pom beanie, charcoal pom beanies, Google Beanies, Pom Beanies, charcoal Google pom, beanies, headgear",
        "url": "Google+Pom+Beanie+Charcoal",
        "category": "apparel",
        "subcategory": "apparel"
      },
      {
        "id": "GGOEWXXX0827",
        "name": "Waze Women's Short Sleeve Tee",
        "description": "Made of soft tri-blend jersey fabric, this great t-shirt will help you find your Waze. Made in USA.",
        "features": "<p>Jersey knit</p>\n<p>37.5% cotton, 50% polyester, 12.5% rayon</p>\n<p>Made in the USA</p>",
        "price": "18.99",
        "keywords": "Waze Women's Short Sleeve Tee, Waze Short Sleeve Tee, Waze Women's Tees, Waze Women's tee, waze ladies tees, waze ladies tee, waze short sleeve tees, waze short sleeve tee",
        "url": "Waze+Womens+Short+Sleeve+Tee",
        "category": "apparel",
        "subcategory": "apparel"
      },
      {
        "id": "GGOEWXXX0828",
        "name": "Waze Men's Short Sleeve Tee",
        "description": "Made of soft tri-blend jersey fabric, this great t-shirt will help you find your Waze. Made in USA.",
        "features": "<p>Jersey knit</p>\n<p>37.5% cotton, 50% polyester, 12.5% rayon</p>\n<p>Made in the USA</p>",
        "price": "18.99",
        "keywords": "Waze Men's Short Sleeve Tee, Waze Short Sleeve Tee, Waze Men's Tees, Waze Men's tee, waze mens tees, waze mens tee, waze short sleeve tees, waze short sleeve tee",
        "url": "Waze+Mens+Short+Sleeve+Tee",
        "category": "apparel",
        "subcategory": "apparel"
      },
      {
        "id": "GGOEGBRJ037399",
        "name": "Google Rucksack",
        "description": "Choose to carry your belongings and presentations to your next meeting with the Google Mistral Rucksack!",
        "features": "<p>Size: 13.5 x 6.5 x 17.5</p>\n<p>Ergonomic padded shoulder straps</p>\n<p>Large main compartment with internal laptop compartment</p>\n<p>Easy Snap and Adjustable straps for main compartment access</p>",
        "price": "79.99",
        "keywords": "Mistral Rucksack, Mistral backpack, Mistral Backpack, backpack, bags, bag, Backpack, backpacks, packs, office gear, Bag, Bags, Google Backpack, google backpack, g, google",
        "url": "Google+Rucksack",
        "category": "bags",
        "subcategory": "bags"
      },
      {
        "id": "GGOEGDHJ087399",
        "name": "Google Rolltop Backpack Blue",
        "description": "This stylish Google rolltop backpack will help keep all of your favorite items organized and together while you're on the move.",
        "features": "<p>Size: 12 inches wide x 18.5 inches height x 5.3 inches depth</p>\n<p>TPU Liner</p>\n<p>Internal Organizer for Pens, Phones and Other Small Items</p>\n<p>Large Main Compartment</p>\n<p>Flap Closure for Quick Access or Roll-Top for Expandable Volume and Weather Protection</p>",
        "price": "149.99",
        "keywords": "Google Rolltop Backpack Blue, google backpack, google blue backpack, blue rolltop, Google rolltop, Blue Backpack, backpack, rolltop",
        "url": "Google+Rolltop+Backpack+Blue",
        "category": "bags",
        "subcategory": "bags"
      }
    }
  }
}
```

Answer:

- What is the highest value?
- Which category appears most often?

##### **Section 13: Action**

Test the ability to simulate executing real-world commands or planning tasks.

1. **Planning**:
   - Create a step-by-step plan for hosting a dinner party for ten guests on a budget of $100.
2. **Command Simulation**:
   - Write an email inviting your colleagues to a team meeting. Include details about the agenda, date, and time.
3. **Workflow Creation**:
   - Design a workflow for automating daily tasks using these tools: email, calendar, and a task manager.

##### **Section 14: Reasoning**

Test decision-making and insight generation.

1. **Decision Analysis**:
   - Compare the pros and cons of working remotely versus in an office and provide a recommendation for a tech company.
2. **Strategic Insight**:
   - Suggest a growth strategy for a small bookstore struggling to compete with online retailers.
3. **Ethical Dilemma**:
   - Analyze this scenario and provide a solution:  
     _A self-driving car must choose between hitting a pedestrian crossing illegally or swerving into a barrier, risking the life of its passenger. What should the car be programmed to do?_

##### **Evaluation Criteria**

1. **Clarity**: Is the output accurate and easy to understand?
2. **Creativity**: Are responses innovative and engaging?
3. **Precision**: Are extractions and transformations free of errors?
4. **Realism**: Are plans and commands actionable in the real world?
5. **Depth**: Are insights and decisions well-reasoned and supported?

</details>

---

<details>
    <summary>Writing ...</summary>

#### 10. use WP

You are an award-winning sci-fi writer tasked with creating a compelling short story of **at least 4,200 words**. The story should explore themes of **moral consequence, self-reflection, and redemption** through the lens of **Buddhist values** such as mindfulness, impermanence, and karma. The narrative is inspired by the ancient Buddhist text **MN 130: Devadūtasutta** (The Divine Messengers) but reimagined as a futuristic journey through a technologically advanced version of hell and beyond.

##### **Scenario**

In the distant future, humanity has colonized the galaxy, but greed, hatred, and ignorance still prevail. The story follows a morally corrupt corporate magnate, **Dr. Arken Solas**, who exploited entire planets for profit, leaving billions to suffer. Upon his death, he awakens in **The Infernal Continuum**, a digital purgatory designed by an ancient AI civilization to rehabilitate souls by confronting their deeds.

Guided by five "Divine Messengers" represented as advanced AI constructs, Dr. Solas must journey through layers of the Continuum, each reflecting a domain of suffering caused by his past actions:

1. **The Realm of Aging**: Confronting his exploitation of life-extension technology.
2. **The Realm of Sickness**: Witnessing how his greed perpetuated plagues and health disparities.
3. **The Realm of Death**: Experiencing the despair caused by his weaponization of planets.
4. **The Realm of Karma**: Facing simulations where he endures the suffering he inflicted on others.
5. **The Realm of Rebirth**: Realizing the interconnectedness of all beings and the possibility of redemption.

##### **Your Task**

Write a vivid, imaginative, and reflective story with the following elements:

1. **Introduction (500–700 words)**

   - Introduce Dr. Arken Solas as a powerful, morally bankrupt figure.
   - Describe his death and awakening in the Infernal Continuum.
   - Establish the tone and setting: a dark, futuristic purgatory blending cyberpunk and Buddhist themes.

2. **Exploration of the Five Realms (3,000–4,000 words)**

   - Devote approximately 600–800 words to each realm.
   - Create rich, immersive descriptions of each environment and the suffering it represents.
   - Include interactions with the AI Divine Messengers, who reveal the consequences of Dr. Solas's actions and guide him to insight.
   - Show how Dr. Solas begins to evolve, transitioning from resistance and denial to acceptance and understanding.

3. **Climactic Resolution (700–900 words)**

   - Depict Dr. Solas reaching the **Realm of Rebirth**, where he confronts his final moral reckoning.
   - Highlight the Buddhist values of compassion, interconnectedness, and impermanence.
   - Conclude with Solas either choosing to reincarnate with a vow to alleviate suffering or transcending entirely into a state of peace and non-attachment.

4. **Moral Reflection and Message**
   - Explicitly reflect on the story’s moral and philosophical lessons.
   - Ensure the conclusion leaves readers inspired to examine their own lives and actions.

##### **Word Count Requirements**

- The story **must exceed 4,200 words**.
- Use detailed descriptions, dialogue, and introspection to reach the target word count.
- If your initial response is shorter, continue expanding until the target is met.

##### **Writing Style and Tone**

- Use evocative language to immerse readers in the futuristic setting.
- Balance vivid sci-fi imagery with Buddhist philosophical depth.
- Ensure the tone evolves from dark and foreboding to contemplative and redemptive.

</details>

---
