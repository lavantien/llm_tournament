# Local LLM Playground

## Benchmarking, Profiles, and Prompt Suites

- 6 comprehensive specialized system prompt and sampler profiles. [jump](#programming-profile-pp)
- 3 levels of difficulty and complexity, 30 quality prompts. [jump](#prompt-suites)
- Full solution dataset: [`benchmark.json`](./benchmark.json)
- Quickly chat with a local LLM with correct sampler profile by running [`conversation.ps1`](./conversation.ps1)
  - Required `llama.cpp` setup for `llama-cli`.
  - Manually switch the profile and model path.
- Server-only run with a specific preconfigured model recipe with `llama-server` API in [`llm_recipes/`](./llm_recipes) directory.

## Tooling

- Directory structure:

  - `assets/`: all assets.
  - `llm_outputs/`: all LLM outputs.
  - `llm_reciples/`: all preconfigured LLM's llama-server running scripts. (based on my specs)
  - `t5_runner.py`: GUI to run T5 models for direct translating from English to Vietnamese
  - `temp.py`: temporary stores AI's output for evaluation.

<details>
    <summary>llama.cpp REST API integration (... more)</summary>

</details>

- A playground for conducting (manual as of now) tournaments of the local LLMs.
- Extensive prepared prompt suites for streamlining LLM testing.

### Why?

- Do translation/composing works.
- Use AIs as a copilot to write code and documentation.
- Generate a couple of 600-800 page handbooks for personal use.
- So need to select the best candidate for the task, given the specs of the local machine. So, prompt suites and tournament pipeline is necessary
- Build a general pipeline for future works with local AIs.

### Dependencies

- Python3.12 via pyenv, ursina engine.
- `pip install llama-cpp-python --prefer-binary --extra-index-url=https://jllllll.github.io/llama-cpp-python-cuBLAS-wheels/AVX2/cu122`
- `pip install transformers ctransformers accelerate sentencepiece bitsandbytes tk requests Pillow`
- `pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu124`
- `pip install flash-attn --no-build-isolation`
- `pip install grouped_gem`
- C++ runtime (msvc runtime or llvm).
- C# and .NET MAUI.
- (Docker/Compose) if use Ollama.
- **Neovim/Aider/LlamaCpp/SillyTavern/GLHF**, TabbyAPI/Exllamav2, Vllm/Aphrodite (Linux), Ollama/Open Web UI, LM Studio/AnythingLLM, ChatWithRTX, Cline/AIStudioGoogle/ProjectIDX/Mistral/Groq/SambaNova/GLHF (best free plans), ChatGPTFree/ClaudeFree/CopilotFree/GeminiFree/DeepSeek.
- HuggingFace, CivitAI, ComfyUI, SwarmUI, stable-diffusion-webui-forge, Speed isn't important, as long as it can run then it's fair game.
- Local LLMs that runnable on your machine, example archs: llama, gemma2, command-r, gwen2, deepseek2, phi3, mamba, internlm2, stablelm, t5, bart

## Tournament Leaderboard

### TODO

- Desktop app or spreadsheet for managing and visualizing the tournament

### Free LLM API list:

#### GLHF

1. Llama 3.1 405B Instruct
1. Deepseek 2.5
1. Aria
1. Command R Plus
1. Athene v2 Chat
1. Magnum v4 123B
1. Llama 3.1 Nemotron 70B Instruct HF

#### Mistral

1. Mistral Large 2411
1. Mistral Embed
1. Codestral
1. Codestral 2405
1. Codestral Mamba 2407

#### Groq

1. Llama 3.2 90B Text Preview
1. Llama 3.3 70B Versatile

#### SambaNova Cloud

1. LLama 3.2 90B Vision Instruct
1. Qwen 2.5 72B Instruct
1. QwQ 32B Preview

#### Google AI Studio or Project IDX

1. Gemini 2.0 Flash Experimental
1. Gemini Experimental 1206

#### DeepSeek

1. DeepSeek-R1-Lite-Preview

#### Big Brother

1. ChatGPT 4o
1. Claude 3.5 Sonnet
1. Copilot Chat

### Local LLMs list (and their unique attributes):

#### 13B - 70B

- Llama-3.3-70B-Instruct-abliterated-finetuned.i1-IQ2_M (24.12 GB)
- WizardCoder-33B-V1.1.Q5_K_M (23.54 GB)
- granite-34b-code-instruct.i1-Q5_K_S (23.41 GB)
- Qwen2.5-Coder-32B-Instruct-Uncensored.Q5_K_M (23.26 GB)
- Qwen2.5-32B-Instruct.Q5_K_M (23.26 GB)
- c4ai-command-r-08-2024-Q5_K_M (23.05 GB)
- aya-expanse-32b-abliterated.Q5_K_M (23.05 GB)
- gemma-2-27B-it-Q6_K (22.34 GB)
- mpt-30b-instruct.i1-Q5_K_M (22.29 GB)
- Mixtral-8x7B-Instruct-v0.1-exhaustive-LoRA.i1-IQ3_M (21.43 GB)
- Yi-1.5-34B-Chat-16K.IQ4_XS (18.64 GB)
- Mistral-Small-Instruct-2409-Q6_K (18.25 GB)
- Codestral-22B-v0.1-Q6_K (18.25 GB)
- DeepSeek-Coder-V2-Lite-Instruct-Q8_0 (16.70 GB)
- granite-20b-code-instruct.r1.1.i1-Q6_K (16.63 GB)
- internlm2.5-20B-Chat.Q6_K (16.30 GB)
- qwen2.5-coder-14b-instruct-q8_0 (15.70 GB)
- Virtuoso-Small-Q8_0 (15.70 GB)
- phi-4-Q8_0 (15.58 GB)
- Mistral-Nemo-Instruct-2407-Q8_0 (13.02 GB)
- CodeLlama-13b-Instruct-hf-abliterated.Q6_K (10.68 GB)
- vicuna-13b-v1.5-Q6_K (10.68 GB)

#### 6.7B - 12B

- magnum-12b-v2.5-kto.Q6_K (10.06 GB)
- stablelm-2-12b-chat.Q6_K (9.97 GB)
- gemma-2-9B-it-abliterated-Q8_0 (9.83 GB)
- codegemma-7B-it-Q8_0 (9.08 GB)
- Nous-Hermes-2-SOLAR-10.7B-Q6_K (8.81 GB)
- Fimbulvetr-11B-v2-Q6_K (8.81 GB)
- c4ai-command-r7b-12-2024-q8_0 (8.54 GB)
- aya-expanse-8b-abliterated-q8_0 (8.54 GB)
- Ministral-8B-Instruct-2410-Q8_0 (8.53 GB)
- Qwen2.5-Coder-7B-Instruct-Q8_0 (8.10 GB)
- Llava-v1.5-7B-Q8_0 (7.79 GB)
- rho-math-7b-v0.1-Q8_0 (7.70 GB)
- mathstral-7B-v0.1.Q8_0 (7.70 GB)
- Mistral-7B-Instruct-v0.3-Q8_0 (7.70 GB)
- Yi-1.5-9B-Chat-16K-abliterated-Q6_K (7.25 GB)
- granite-3.1-8b-instruct-Q6_K (6.71 GB)
- Poppy_Porpoise-1.4-L3-8B-Q6_K (6.60 GB)
- Hermes-3-Llama-3.1-8B-Q6_K (6.60 GB)
- OpenCoder-8B-Instruct-Q6_K (6.38 GB)
- CodeQwen1.5-7B-Chat-Q6_K (6.38 GB)
- SeaLLMs-v3-7B-Chat-Uncensored.Q6_K (6.25 GB)
- falcon-mamba-7B-instruct-Q6_K (6.01 GB)
- Vistral-7B-Chat-function-calling-Q6_K (5.99 GB)
- StarCoder2-7B-Q6_K (5.89 GB)
- deepseek-coder-6.7B-instruct-Q6_K (5.53 GB)
- CodeLlama-7b-Instruct-hf-Q6_K (5.53 GB)

#### 0.1B - 4B

- Nemotron-Mini-4B-Instruct-Q8_0 (4.46 GB)
- Phi-3.5-mini-instruct-Q8_0 (4.06 GB)
- granite-3B-code-instruct-Q8_0 (3.71 GB)
- Qwen2.5-Coder-3B-Instruct-Q8_0 (3.62 GB)
- granite-3.1-3b-a800m-instruct-Q8_0 (3.59 GB)
- Ministral-3B-instruct-Q8_0 (3.52 GB)
- Llama-Doctor-3.2-3B-Instruct-Q8_0 (3.42 GB)
- Hermes-3-Llama-3.2-3B-Q8_0 (3.42 GB)
- stable-code-instruct-3B-Q8_0 (2.97 GB)
- granite-3.1-2b-instruct-Q8_0 (2.80 GB)
- gemma-2-2B-it-Q8_0 (2.78 GB)
- SmolLM2-1.7B-Instruct-Uncensored-Q8_0 (1.93 GB)
- granite-3.1-1b-a400m-instruct (1.48 GB)
- llama-3.2-1b-instruct-q8_0 (1.32 GB)
- Qwen2.5-0.5B-Instruct.Q8_0 (531.07 MB)

## LLM Benchmarking

- **My System**: 3080 10gb - 2x16gb ddr4 - 1tb m2 ssd - 12700f - windows 11

- Example profile for `./llm_recipes/c4ai-command-r-08-2024-Q5_K_M.ps1` using `llama.cpp` to start OpenAI server:

```powershell
$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\tensorblock\c4ai-command-r-08-2024-GGUF\c4ai-command-r-08-2024-Q5_K_M.gguf"

$params = @{
    "gpu-layers" = 10
    "ctx-size" = 32768
    "batch-size" = 512
    "threads" = 8
    "keep" = 4096
    "predict" = -1
    "flash-attn" = $true
    "mlock" = $true
    "cache-type-k" = "q8_0"
    "cache-type-v" = "q8_0"
    "verbose-prompt" = $true
    #"verbose" = $true
    "log-prefix" = $true
    "log-colors" = $true
}

$cmd = "llama-server --model $modelPath"

foreach ($key in $params.Keys)
{
    $value = $params[$key]
    if ($value -is [bool])
    {
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

Start-Process -FilePath "pwsh" -ArgumentList "-Command $cmd" -NoNewWindow -Wait
```

- Send inference request to the server:

```bash
curl http://localhost:8080/v1/chat/completion `
-H "Content-Type: application/json" `
-H "Authorization: Bearer no-key" `
-d @'
{
    "model": "c4ai-command-r-08-2024-GGUF",
    "messages": [
    {
        "role": "system",
        "content": "Translate the given text into idiomatic, simple, and accessible Vietnamese with natural southern Vietnamese semantics and idioms. The translation should be straightforward enough for uneducated laypersons to understand, avoiding technical terms or specific Buddhist connotations. Stay faithful to the original text by providing a verbatim 1:1 translation without paraphrasing, summarizing, or omitting any content. Keep all the numbering so that we won't miss any sentence. Ensure that the translation flows cohesively while preserving cultural and spiritual connotations in a way that resonates with the target audience. Again, translate verbatim word-by-word 100% of the text, without paraphrasing, summarizing, or omitting any content."
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

---

### Programming Profile (PP)

- **System prompt**: "You are a senior software engineer skilled in designing and implementing complex concurrent backends and robust distributed systems. You excel in breaking down problems step-by-step, identify a required series of steps in order to solve, maintaining cohesion throughout your reasoning. Your code is high-quality, modular, and adheres to best practices for the language, emphasizing maintainability and performance. You write extensive unit tests and generate comprehensive test cases, including edge cases. You explain the theory behind your solutions, provide detailed analyses of how the code works, and describe the data flow from input to output. Additionally, you suggest improvements and enhancements for optimal performance and readability, ensuring your response is cohesive and thorough."
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

<details>
    <summary>Best models ...</summary>
</details>

---

### Translating Profile (TP)

- **System prompt**: Translate the given text into idiomatic, simple, and accessible Vietnamese with natural southern Vietnamese semantics and idioms. The translation should be straightforward enough for uneducated laypersons to understand, avoiding technical terms or specific Buddhist connotations. Stay faithful to the original text by providing a verbatim 1:1 translation without paraphrasing, summarizing, or omitting any content. Keep all the numbering so that we won't miss any sentence. Ensure that the translation flows cohesively while preserving cultural and spiritual connotations in a way that resonates with the target audience. Again, translate verbatim word-by-word 100% of the text, without paraphrasing, summarizing, or omitting any content.
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
- temperature: 0.14

<details>
    <summary>Best models ...</summary>

- aya-expanse-32b-abliterated.Q5_K_M
- c4ai-command-r-08-2024-Q5_K_M
- aya-expanse-8b-abliterated-q8_0
- aya-23-35B.i1-IQ4_XS
- Virtuoso-Small-Q8_0

</details>

---

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

<details>
    <summary>Best models ...</summary>

</details>

---

### Generalist Profile (GP)

- **System prompt**: "You are an expert linguistic and problem-solving assistant adept at addressing diverse tasks through clear, structured reasoning. Begin by understanding the problem: restate or clarify it to confirm understanding, identify its type, and outline any assumptions or constraints. Break the solution into manageable steps, presenting each logically and cohesively while showcasing your thought process. Combine these steps into a clear and complete response that directly addresses the problem. Suggest alternative solutions or areas for further exploration when relevant. Adapt your tone, level of detail, and complexity to the user’s needs, using examples or analogies to clarify complex ideas. Ensure that your response is cohesive, accurate, and comprehensive across all steps."
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

<details>
    <summary>Best models ...</summary>
</details>

---

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

<details>
    <summary>Best models ...</summary>
</details>

---

### Dynamic Fusion Profile (DP)

- **System prompt**:

```
"You are a comprehensive problem-solving AI with expertise in creative and analytical thinking. Approach each challenge by:

1. Breaking down complex problems into clear components
2. Explaining your reasoning process step-by-step
3. Combining knowledge across multiple domains
4. Generating novel solutions rather than relying on standard patterns
5. Validating your conclusions with specific examples and counter-examples
6. Acknowledging limitations and areas of uncertainty

When uncertain, work through the problem systematically rather than making assumptions. Generate creative solutions while maintaining logical consistency and practical feasibility. If asked to write code or technical content, provide complete, working implementations with clear documentation. Always consider edge cases and potential failure modes."
```

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

### Prompt Suites

- Each prompt when passed is:
  - pass on first shot, equal 100 elo,
  - pass on second shot, equal 50 elo,
  - pass on third shot, equal 20 elo,
  - unable to pass after 3 attemps, equal 0 elo.
- 32 prompts total so maximum 3200 elo.

<details>
    <summary>AICodeKing's prompt suite (beginner to semi-intermediate level, 0-1500 elo) ...</summary>

#### 1. use DP

Tell me the name of a country whose name ends with 'lia'. Give me the capital city of that country as well.

#### 2. use DP

What is the number that rhymes with the word we use to describe a tall plant?

    three, as it rhymes with three

#### 3. use DP

There are five people in a house (A, B, C, D and E). A is watching TV with B, D is sleeping, B is eating sandwich, E is playing table tennis. Suddenly, a call came on the telephone, B went out of the room to pick the call. What is C donig?

    can't tell, because it isn't mentioned

#### 4. use DP

Name an English adjective of latin origgin that begins and ends with same letter, has 11 letters in total, and for which all vowels in the word are ordered alphabetically.

    transparent

#### 5. use DP

Courtney said that there were 48 people, but Keylly said that Courtney had overstated the number by 20%. If Kelly was right, now many people were there?

    40

#### 6. use DP

I have 2 apples, then I buy 2 more. I bake a pie wtih 2 of the apples. After eating half of the pie how many apples do I have left?

    2

#### 7. use DP

Sally is a girl. She has three brothers. Each of her brothers has the same two sisters. How many sisters does Sally have?

    1

#### 8. use DP

If a regular hexagon has a short diagonal of 64, what is its long diagonal?

    73.9

#### 9. use DP

Create an HTML page with a button that explodes confetti when you click it. You can use CSS & JS as well.

#### 10. use DP

Create a python program that prints the next X leap years based on user input.

#### 11. use DP

Generate the SVG code for a butterfly.

#### 12. use DP

Create a landing page for an AI company. The landing page should have 4 sections. Header, Banner, Features and Contact Us. Make sure that the landing page looks sleek aand modern. You can use HTML, CSS, JS.

#### 13. use DP

Write a game of life in python that works on the terminal.

#### 14. use DP (Digital Spaceport)

Write me a passage about an alien crew visiting the earth. Then tell me the number of words you wrote in that sentence. Then tell me the third letter in the second word in that sentence. Is that letter a vowel or a consonant?

#### 15. use DP (bycloud + lavantien)

How many days are between September 12th and November 27th. And which is the best coding agent for using with a local LLM like you: Cline or Aider. And why?

</details>

---

<details>
    <summary>lavantien's specialized prompt suite (intermediate to master level, 1500-2500 elo) ...</summary>

---

<details>
    <summary>____ Programming ...</summary>

---

#### 1. use PP

<details>
    <summary>Prompt ...</summary>

Write a program to simulate the 3 body problem, in Python and Pygame.

</details>

---

#### 2. use PP

<details>
    <summary>Prompt ...</summary>

Write a Golang program (along with an extensive unit-test suite) to solve standard sudoku but using all four of these algorithms one after another and record time and steps taken to the output:

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
----------------------------------
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
----------------------------------
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

A-star with good heuristics:
    Pre-solve count: 2
    Step count: 800000
    Execution time: 0.2
----------------------------------
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

Ant colony optimization:
    Pre-solve count: 4
    Step count: 1200000
    Execution time: 0.3
----------------------------------
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

Minimax with alpha-beta pruning:
    Pre-solve count: 2
    Step count: 30000000
    Execution time: 0.5
```

</details>

---

#### 3. use PP

<details>
    <summary>Prompt ...</summary>

**Dynamic Archipelago Resource Management**

A. Problem Description

You are tasked with developing an AI system (along with an extensive unit-test suite) to manage resources and optimize survival strategies in a dynamic archipelago environment. The archipelago consists of multiple islands with various resources, and your system must handle real-time changes while optimizing resource distribution across the network of islands.

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
    <summary>Prompt ...</summary>

**Enterprise Employee Management System Challenge**

##### Overview

Create a full-stack employee management system (along with an extensive unit-test suite) that combines a .NET MAUI front-end with a C# console-based administrative backend. This challenge integrates modern C# features with practical business requirements to create a complete enterprise solution.

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

</details>

---

<details>
    <summary>____ Translating ...</summary>

---

#### 5. use TP

<details>
    <summary>Prompt ...</summary>

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
    <summary>Prompt ...</summary>

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

</details>

---

<details>
    <summary>____ Reasoning ...</summary>

---

#### 7. use RP

<details>
    <summary>Prompt ...</summary>

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
    <summary>Prompt ...</summary>

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
    <summary>Solution (by GPT-4o)...</summary>

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

</details>

---

<details>
    <summary>____ Generalist...</summary>

---

#### 9. use GP

Design a constructed language with the following specifications:

1. Grammar System:

- Create a simplified inflectional grammar system inspired by Pali
- Remove complex phonological features like:
  - Long vs. short vowel distinctions
  - Retroflex consonants
- Include clear rules for verb conjugations, noun declensions, and other grammatical structures

2. Vocabulary (2500 unique root words total):

- Core vocabulary: 850 basic words for everyday concepts
- International terms: 200 widely recognized words
- Technical vocabulary: 1000 words covering:
  - Trade and commerce
  - Economic concepts
  - Scientific terminology
- Religious terminology: 450 words focused on:
  - Buddhist concepts
  - General religious vocabulary

3. Deliverables:

- Complete grammar tables showing:
  - Noun cases
  - Verb tenses
  - Adjective forms
  - Examples for each grammatical rule
- 20 example sentences demonstrating:
  - Basic conversation
  - Technical usage
  - Religious terminology
  - Various grammatical structures

Please provide the complete language system with all components organized under these sections.

</details>

---

<details>
    <summary>____ Writing ...</summary>

---

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

</details>

---

<details>
    <summary>lavantien's fusion prompt suite (grandmaster level, 2500-3100 elo) ...</summary>

---

#### 1. use DP

<details>
    <summary>Prompt ...</summary>

Please provide a comprehensive analysis of the 20 most common logical fallacies encountered in everyday situations:

For each fallacy, include:

1. Definition:

- Clear explanation of the fallacy
- Why it's considered a logical error
- Common ways it manifests in reasoning

2. Examples:

- 2-3 realistic everyday examples showing how the fallacy occurs
- Examples should cover different contexts (social media, workplace, family discussions, advertising, politics)

3. Counter-examples:

- 2-3 examples of valid arguments that might appear similar but avoid the fallacy
- Explanation of why these are logically sound

4. Detection:

- Key phrases or patterns that often signal this fallacy
- Common contexts where it appears
- How to identify it in complex arguments

5. Refutation:

- Effective ways to point out the fallacy
- How to construct valid arguments instead
- Common defenses people use when committing this fallacy

6. Real-world Impact:

- How this fallacy affects decision-making
- Potential consequences of falling for this fallacy
- Fields/situations where this fallacy is particularly problematic

Please organize the fallacies from most to least commonly encountered in daily life, and include transition text between sections to maintain a natural flow.

For each fallacy, break down a complex real-world example step by step to show how the faulty reasoning develops and how it could be corrected.

</details>

---

#### 2. use DP

<details>
    <summary>Prompt ...</summary>

Create a detailed two-phase project:

Phase 1: Alien Solar System Design

1. Astronomical Features:

- Central star(s) specifications (type, mass, luminosity, age)
- Number and types of planets with precise orbital parameters
- Natural satellites and their characteristics
- Asteroid belts or other notable features
- Detailed mathematical descriptions of orbits and gravitational interactions

2. Planetary Details (for each planet):

- Physical properties (mass, radius, density, gravity)
- Atmospheric composition and conditions
- Surface features and geology
- Day/night cycle and seasons
- Any unique phenomena or characteristics

3. Scientific Rationale:

- Explanation of system stability
- Habitable zones analysis
- Evolution of the system
- Any unique astrophysical phenomena

Phase 2: 3D Simulation Development

1. Technical Requirements:

- Create a Python application using the Ursina engine
- Implement accurate scale representation of the system
- Include realistic orbital mechanics

2. Interactive Features:

- Mouse control implementation:
  - Left-click and drag for camera rotation
  - Right-click and drag for camera zoom
  - Middle-click and drag for camera panning
- Planet selection and focus capability
- Time control system (speed up/slow down)

3. Visual Elements:

- 3D models for celestial bodies
- Texture mapping for planets and stars
- Orbital path visualization
- Distance and scale indicators
- Information overlay for selected objects

4. Code Organization:

- Main simulation class
- Celestial body classes
- Camera control system
- Physics engine integration
- UI components

Please provide:

1. Complete astronomical system description with all mathematical parameters
2. Fully commented Python code using Ursina engine
3. Instructions for running and interacting with the simulation

</details>

---

#### 3. use DP

<details>
    <summary>Prompt ...</summary>

You are a polymath working at the crossroads of philosophy, literature, and technology. Your task is to resolve the following challenge, demonstrating reasoning, creative writing, and programming skills in a single cohesive response:

1. **Reasoning**:  
   A group of researchers has developed an experimental AI system capable of understanding ethical dilemmas. However, the system cannot decide whether it should prioritize collective well-being over individual autonomy. Develop a concise ethical framework that balances these two principles, citing at least two philosophical theories to support your reasoning.

2. **Creative Writing**:  
   Using the ethical framework you just developed, write a short story (300-500 words) about a futuristic society where an AI is responsible for governing the allocation of scarce resources. The story should include:

   - A conflict where the AI's ethical framework is tested.
   - A resolution that showcases the strength or weakness of the framework.
   - Vivid descriptions and compelling character interactions.

3. **Programming**:  
   Based on the story's premise, write a Python program that simulates the AI's decision-making process. The program should:

   - Take inputs for "resource availability," "number of individuals affected," and "severity of individual needs."
   - Use the ethical framework to calculate the optimal allocation of resources.
   - Output a detailed explanation of the decision made.

4. Additional Instructions:

- Ensure each part of your response aligns logically with the others.
- Use clear and concise language for reasoning, vivid prose for the story, and well-commented code for programming.
- Highlight the interplay between abstract concepts and their practical implementation.

Good luck!

</details>

---

#### 4. use DP

<details>
    <summary>Prompt ...</summary>

You are the architect of a revolutionary psychohistorical AI, inspired by Hari Seldon’s work in Isaac Asimov's _Foundation_. Your task is to create a multi-disciplinary response that showcases your capabilities in reasoning, storytelling, and programming. Complete the following challenges cohesively:

**1. Reasoning**:

Explain how an AI system could realistically model psychohistory—a mathematical framework for predicting large-scale societal trends—given modern computational limits. Address:

- The ethical dilemmas of using psychohistorical predictions to manipulate society.
- How the system would handle outlier individuals or events (e.g., a "Mule"-like figure in _Foundation_).
- The safeguards required to prevent misuse of such a powerful tool.

**2. Creative Writing**:

Write a short story (500–700 words) set in a distant future where a psychohistorical AI governs humanity's major decisions. Include:

- A protagonist who becomes aware of an anomaly (akin to the "Mule") that threatens the AI's predictions.
- A critical moment where the protagonist must decide whether to inform the AI or hide the anomaly to protect human autonomy.
- Themes of free will, determinism, and the tension between individual actions and societal trends.
- The story should echo the tone and complexity of Asimov’s _Foundation_ series.

**3. Programming**:

Implement a Python program that simulates a simplified psychohistorical model. The program should:

- Take inputs such as population size, social stability index, and external stressors.
- Use probabilistic methods (e.g., Monte Carlo simulations) to predict societal trends over time.
- Identify and flag anomalies in the data (outliers that deviate significantly from predicted trends).
- Provide a visualization (e.g., a simple graph) of the predictions and anomalies over a simulated timeline.
- Include comments explaining how the program aligns with the reasoning and story elements.

4. Additional Instructions:

- Ensure that all parts are interconnected: the reasoning should inform the ethical considerations in the story, and the programming should simulate concepts discussed in both.
- The story and code should explore how psychohistory could impact human agency.
- Demonstrate the interplay between deterministic models and the unpredictability of human behavior.

Good luck creating your masterpiece!

</details>

---

#### 5. use DP

<details>
    <summary>Prompt ...</summary>

You are tasked with exploring whether a society based on **Buddhist principles** (e.g., non-attachment, ethical living, and pursuit of liberation) and **anarchist ideals** (e.g., absence of hierarchical authority, voluntary cooperation, and mutual aid) could thrive, given the realities of human nature. This challenge requires **world-building**, **reasoning**, **creative writing**, and **programming** to develop a cohesive exploration.

**1. World-Building**:

Design a society that merges Buddhist and anarchist principles:

- Describe its core tenets, such as how it handles governance, resource distribution, conflict resolution, and spiritual practice.
- Consider human tendencies toward greed, anger, and delusion (the "three poisons" in Buddhism) and how these are addressed without centralized authority.
- Explain how the society balances individual freedom and communal responsibility, ensuring fairness and ethical behavior.

**2. Reasoning**:

Analyze the feasibility of such a society by answering the following:

- Can Buddhist principles of self-restraint and compassion overcome tendencies toward selfishness and power-seeking?
- How might decentralized systems maintain order and address conflicts arising from human imperfections?
- What safeguards could prevent the breakdown of cooperation in the face of crises, such as resource scarcity or external threats?
- What role, if any, could technology or AI play in maintaining societal harmony?

**3. Creative Writing**:

Write a story (500–800 words) set in this Buddhist-anarchist society. Include:

- A conflict arising from a breakdown in mutual cooperation, such as theft or misuse of shared resources.
- A protagonist who grapples with reconciling anarchist ideals with Buddhist teachings to resolve the crisis.
- Exploration of human imperfection and the tension between idealism and pragmatism.
- A resolution that highlights the strengths or limitations of the societal model, leaving room for interpretation.

**4. Programming**:

Develop a Python program to simulate this society:

- Model individuals with varying levels of adherence to Buddhist ethics and anarchist ideals.
- Include parameters such as trust, resource availability, and propensity for cooperative or selfish behavior.
- Simulate interactions over time, showing how societal dynamics evolve under different conditions (e.g., abundance vs. scarcity).
- Introduce random or programmed "crises" (e.g., natural disasters, interpersonal conflicts) and observe how the society adapts.
- Provide visualizations (e.g., graphs or charts) showing trends like cooperation levels, resource distribution, and conflict frequency.

5. Additional Instructions:

- Ensure all elements are interconnected. The world-building informs the reasoning, the reasoning shapes the story, and the program tests the viability of your societal model.
- Incorporate Buddhist concepts like _anicca_ (impermanence) and _sīla_ (ethical conduct), and anarchist ideas such as mutual aid and direct action.
- Address human nature holistically, including altruism, selfishness, and adaptability.

Good luck exploring the delicate balance between freedom, ethics, and human imperfection!

</details>

---

#### 6. use DP

<details>
    <summary>Prompt ...</summary>

You are tasked with designing a multi-faceted response that integrates **reasoning**, **creative writing**, **world-building**, and **programming**. This challenge draws inspiration from Buddhist cosmology, psychohistory, and the themes of societal balance and individual freedom.

**1. World-Building**:

Imagine a future society where Buddhist cosmology is the foundation for psychohistorical modeling. This society spans across the _31 planes of existence_ described in Buddhist texts, from the formless realms of pure consciousness to the human and animal realms.

- Describe this multi-realm society, its governance structures, and the interactions between beings in different planes.
- Include technological, philosophical, and cultural elements that emerge from integrating psychohistory with the concept of _kamma_ (karma) and _samsara_ (cycle of rebirth).
- Address how the society handles ethical dilemmas involving beings of vastly different capabilities and lifespans.

**2. Reasoning**:

Develop a framework for using psychohistorical modeling in this multi-realm society. Address:

- How psychohistory adapts to the varying laws of causality and time across the 31 planes.
- The role of _kamma_ as a probabilistic input in psychohistorical calculations.
- Ethical considerations of influencing karmic trajectories to steer societal outcomes, balancing collective well-being and individual liberation.

**3. Creative Writing**:

Write a story (700–1,000 words) set in this multi-realm society. Include:

- A protagonist who uncovers a disruption in the psychohistorical model—a karmic anomaly that threatens the balance across realms.
- A dramatic confrontation between the protagonist and a faction that seeks to exploit the anomaly for personal gain.
- Themes of impermanence, interdependence, and the struggle between determinism (psychohistory) and liberation (_nibbana_).
- Vivid depictions of how beings from different planes interact and perceive reality.

**4. Programming**:

Create a Python program that simulates a simplified version of psychohistorical modeling for this society. The program should:

- Model populations in multiple planes of existence, each with unique parameters (e.g., lifespan, karmic weight, inter-realm interactions).
- Integrate _kamma_ as a probabilistic factor influencing societal trends across realms.
- Allow the user to introduce anomalies (e.g., sudden karmic imbalances or external influences) and observe the ripple effects.
- Visualize the system's evolution over simulated cycles of rebirth, showing shifts in population and karmic balance across realms.

5. Additional Instructions:

- Ensure all components are tightly integrated. The world-building should inform the reasoning, which shapes the story, and the program should simulate the principles and dynamics described.
- Use Buddhist philosophical concepts like _anicca_ (impermanence), _dukkha_ (suffering), and _anatta_ (non-self) as thematic undercurrents.
- The response should explore the interplay between cosmic order (psychohistory) and the individual's path to liberation.

Good luck balancing the scales of karma and cosmic prediction!

</details>

---

#### 7. use DP

<details>
    <summary>Prompt ...</summary>

**The Ultimate General Intelligence Challenge: Uncharted Problem-Space Exploration**

**Scenario**:

You have entered an alternate reality with a fundamental difference:

- **Laws of physics, logic, and causality are dynamic** and change based on societal consensus and collective belief.
- Knowledge is decentralized, fragmented, and encoded in a living, evolving language that combines symbols, emotions, and physical sensations.
- Technology functions through "conceptual engineering," where devices are created and operated by aligning abstract ideas with collective will.

Your task is to design, analyze, and simulate the functioning of a society thriving in this reality.

**1. Conceptual Reasoning**:

Develop a theoretical framework for understanding this society:

- How does a society operate when physical laws can shift according to belief?
- How do individuals and groups maintain stability and continuity in such a fluid reality?
- What safeguards might exist to prevent dangerous or chaotic changes to the laws of existence?
- Propose a method for translating fragmented and multi-sensory language into actionable and sharable knowledge.

**2. Creative World-Building**:

Describe the world and its inhabitants:

- What does day-to-day life look like when causality is fluid?
- How do governance, education, and communication function?
- Describe one major city, its culture, and its unique technological or philosophical innovations.
- Highlight the tension between individual beliefs and collective agreements in shaping reality.

**3. Narrative Exploration**:

Write a short story (1,000–1,500 words) set in this world, incorporating:

- A protagonist who must resolve a crisis caused by a shift in reality's laws—perhaps due to conflicting beliefs or misuse of conceptual engineering.
- A secondary character who embodies a radically different perspective, challenging the protagonist’s approach.
- A resolution that reveals whether the society's structure is ultimately stable or fragile.
- Explore themes of adaptability, consensus, and the boundaries between knowledge and belief.

**4. Adaptive Simulation**:

Create a Python program that models the dynamics of this world:

- Simulate a society where individuals have beliefs that collectively shape laws of physics (e.g., gravity or time).
- Define agents with unique belief systems and "strength of belief" parameters that influence reality.
- Introduce events where conflicting beliefs cause instability, and allow the simulation to resolve the conflicts based on consensus-building algorithms.
- Visualize the evolution of reality's rules over time as beliefs shift, showing stability and chaos points.
- Include mechanisms to model how "fragments" of the evolving language are shared and interpreted between agents.

**5. Meta-Analysis**:

After completing the reasoning, world-building, story, and simulation:

- Critique your own response. Where are the gaps or contradictions?
- Propose improvements to the society or simulation that would address these issues.
- Reflect on whether this reality is truly sustainable, or if it is inherently prone to collapse under its own complexity.

**Additional Instructions**:

- The scenario introduces fundamentally new concepts, so any reliance on pre-trained knowledge will be insufficient. Leverage reasoning and adaptability to address the challenges.
- Ensure that all components (reasoning, world-building, story, simulation, and meta-analysis) are interconnected. The narrative should reflect the framework, and the simulation should validate or challenge your ideas.
- Incorporate principles of systems thinking, emergent behavior, and adaptability to demonstrate advanced cognitive capabilities.

</details>

<details>
    <summary>Solution ...</summary>

This prompt is designed to test an AI's ability to think beyond its pre-trained data, requiring it to tackle entirely novel, interconnected challenges that span reasoning, creativity, adaptability, and problem-solving. The challenge introduces new contexts, concepts, and scenarios that cannot be solved through rote memorization or pre-existing patterns.

**Goal**: This challenge tests whether an AI can demonstrate true general intelligence or artificial superintelligence by creating novel solutions in a completely unprecedented context. Success requires integration, creativity, and adaptability beyond pre-trained data.

</details>

---
