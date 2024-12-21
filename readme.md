# Local LLM Playground

## Benchmarking, Profiles, and Prompt Suites

- 6 comprehensive specialized system prompt and sampler profiles. [jump](#programming-profile-pp)
- 3 levels of difficulty and complexity, 30 quality prompts. [jump](#prompt-suites)
- Full solution dataset: [`profiles_and_prompt_suites.md`](./profiles_and_prompt_suites.md), or [`./llm_benchmarking_app/data/prompts/prompt_suite.json`](./llm_benchmarking_app/data/prompts/prompt_suite.json)
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
- **Neovim/Aider/Mistral**, **LlamaCpp/SillyTavern**, TabbyAPI/Exllamav2, Vllm/Aphrodite (Linux), Ollama/Open Web UI, LM Studio/AnythingLLM, ChatWithRTX, Cline/AIStudioGoogle/ProjectIDX/Mistral/Groq/SambaNova/GLHF (best free plans), ChatGPTFree/ClaudeFree/CopilotFree/GeminiFree/DeepSeek.
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
