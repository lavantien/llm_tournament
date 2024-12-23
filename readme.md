# Local LLM Playground

## Benchmarking, Profiles, and Prompt Suites

- Full solution dataset and evaluation methodology: [`profiles_and_prompt_suites.md`](./profiles_and_prompt_suites.md).
  - 6 comprehensive specialized system prompt and sampler profiles.
  - 3 levels of difficulty and complexity, 32 quality prompts.
- Quickly chat with a local LLM with correct sampler profile by running:
  - [`converse.ps1`](./converse.ps1)
  - or translate a text to Vietnamese with [`translate.ps1`](./translate.ps1)
  - or do quick programming task with [`code.ps1`](./code.ps1)
  - Required `llama.cpp` setup for `llama-cli`, and `Qwen2.5 32B Instruct`, `Aya Expanse 32B Abliterated`, `Qwen2.5 Coder 32B Instruct`
  - Manually switch the profile and model path.
- Server-only run with a specific preconfigured model recipe with `llama-server` API in [`llm_recipes/`](./llm_recipes) directory.

- Powered by:

  - **Local Dev**: Neovim, Aider, MistralAPI Free _(500,000 tpm - 1,000,000,000 tpM)_, GeminiAPI Free _(10 rpm, 1500 rpd)_, Crawl4AI
  - **Local Gen**: LlamaCpp + {a set of best local models}, SillyTavern/Agnai, StableDiffusionWebUIForge, FluxDev + ControlNets/LoRAs
  - **External Chat**: AIStudio, ChatGPT, ClaudeAI, CopilotChat, Cohere, Mistral, DeepSeek, SambaNova, Groq, GLHF.
  - **Extras**: ProjectIDX, Cline, OpenCanvas, MetaGPT, DeepInfra (non-free)

- **LLM Benchmarking** [`preliminary_design.md`](./preliminary_design.md)
- `npm install && npm run dev` to start `LLM Benchmarking` app on localhost:
  - TypeScript, Next.js
  - Sahdcn/ui, D3.js
  - Zustand, React-Query
  - DrizzleORM, SQLite
  - Vitest, Playright

![ER Diagram](./assets/design/erdiagram.png)

## Tooling

- Directory structure:

  - `assets/`: all assets.
  - `llm_outputs/`: all LLM outputs.
  - `llm_reciples/`: all preconfigured LLM's llama-server running scripts. (based on my specs)
  - `conversation.ps1`
  - `profiles_and_prompt_suites`
  - `t5_runner.py`: GUI to run T5 models for direct translating from English to Vietnamese
  - `temp.py`: temporary stores AI's output for evaluation.

- A playground for conducting (manual as of now) tournaments of the local LLMs.
- Extensive prepared prompt suites for streamlining LLM testing.

### Why?

- Do translation/composing works.
- Use AIs as a copilot to write code and documentation.
- Generate a couple of 600-1000 page books for personal use.
- So need to select the best candidate for the task, given the specs of the local machine.
- So, prompt suites and tournament pipeline is necessary
- Build a general pipeline for future works with local AIs.

### Dependencies

- Python3.12 via pyenv
- `pip install llama-cpp-python --prefer-binary --extra-index-url=https://jllllll.github.io/llama-cpp-python-cuBLAS-wheels/AVX2/cu122`
- `pip install transformers ctransformers accelerate sentencepiece bitsandbytes tk requests Pillow`
- `pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu124`
- `pip install flash-attn --no-build-isolation`
- `pip install grouped_gem`
- Node.js
- C++ runtime (msvc runtime or llvm).
- C# and .NET MAUI.
- (Docker/Compose) if use Ollama.
- **Neovim/Aider/Mistral/Crawl4AI**, **LlamaCpp/SillyTavern**, TabbyAPI/Exllamav2, Vllm/Aphrodite (Linux), Ollama/Open Web UI, LM Studio/AnythingLLM, etc.
- HuggingFace, CivitAI, stable-diffusion-webui-forge, ComfyUI, SwarmUI, Speed isn't important, as long as it can run then it's fair game.
- Local LLMs with highest quant that runnable on your machine, example archs: llama, gemma2, command-r, gwen2, deepseek2, phi3, mamba, internlm2, stablelm, t5, bart

## Tournament Leaderboard

### TODO

- Desktop app or spreadsheet for managing and visualizing the tournament

### Free LLM API list:

<details>
    <summary>...more</summary>

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
1. Mistral 8x22B
1. Codestral

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
1. Gemini 2.0 Flash Thinking Experimental

#### DeepSeek

1. DeepSeek-R1-Lite-Preview

#### Big Brother

1. ChatGPT 4o
1. Claude 3.5 Sonnet
1. Copilot Chat

</details>

### Local LLMs list (and their unique attributes):

<details>
    <summary>...more</summary>

**Disk Usage**: 678 GB

#### 13B - 70B

- Llama-3.3-70B-Instruct.i1-IQ2_M (24.12 GB)
- Qwen2.5-Coder-32B-Instruct-Q5_K_L (23.74 GB)
- Qwen2.5-32B-Instruct-Q5_K_L (23.74 GB)
- Codestral-22B-v0.1-Q8_0 (23.64 GB)
- aya-expanse-32b-Q5_K_L (23.56 GB)
- c4ai-command-r-08-2024-Q5_K_L (23.56 GB)
- WizardCoder-33B-V1.1.Q5_K_M (23.54 GB)
- gemma-2-27B-it-Q6_K (22.34 GB)
- Mixtral-8x7B-Instruct-v0.1-exhaustive-LoRA.i1-IQ3_M (21.43 GB)
- internlm2_5-20b-chat-Q8_0 (21.11 GB)
- Yi-1.5-34B-Chat-16K.IQ4_XS (18.64 GB)
- Mistral-Small-Instruct-2409-Q6_K (18.25 GB)
- starcoder2-15b-instruct-v0.1-Q8_0 (16.97 GB)
- DeepSeek-Coder-V2-Lite-Instruct-Q8_0 (16.70 GB)
- SuperNova-Medius-Q8_0 (15.70 GB)
- qwen2.5-coder-14b-instruct-q8_0 (15.70 GB)
- Virtuoso-Small-Q8_0 (15.70 GB)
- phi-4-Q8_0 (15.58 GB)
- CodeLlama-13B-Instruct-fp16-Q8_0 (13.83 GB)
- vicuna-13b-v1.5.Q8_0 (13.83 GB)

#### 6.7B - 12B

- Mistral-Nemo-Instruct-2407-Q8_0 (13.02 GB)
- magnum-12b-v2.5-kto-Q8_0 (13.02 GB)
- stablelm-2-12b-chat-Q8_0 (12.91 GB)
- Fimbulvetr-11B-v2-Q8_0 (11.40 GB)
- Nous-Hermes-2-SOLAR-10.7B.Q8_0 (11.40 GB)
- gemma-2-9b-it-Q8_0_L (10.69 GB)
- codegeex4-all-9b-Q8_0 (9.99 GB)
- codegemma-7B-it-Q8_0 (9.08 GB)
- c4ai-command-r7b-12-2024-q8_0 (8.54 GB)
- aya-expanse-8b-Q8_0 (8.54 GB)
- Poppy_Porpoise-1.4-L3-8B.Q8_0 (8.54 GB)
- Hermes-3-Llama-3.1-8B-Q8_0 (8.54 GB)
- Ministral-8B-Instruct-2410-Q8_0 (8.53 GB)
- OpenCoder-8B-Instruct-Q8_0 (8.26 GB)
- Qwen2.5-Coder-7B-Instruct-Q8_0 (8.10 GB)
- SeaLLMs-v3-7B-Chat-Q8_0 (8.10 GB)
- Llava-v1.5-7B-Q8_0 (7.79 GB)
- falcon-mamba-7b-instruct-Q8_0 (7.77 GB)
- rho-math-7b-v0.1-Q8_0 (7.70 GB)
- mathstral-7B-v0.1.Q8_0 (7.70 GB)
- Mistral-7B-Instruct-v0.3-Q8_0 (7.70 GB)
- starcoder2-7b-instruct.Q8_0 (7.63 GB)
- OpenCoder-8B-Instruct-Q8_0 (8.26 GB)
- falcon-mamba-7b-instruct-Q8_0 (7.77 GB)
- codeqwen-1_5-7b-chat-q8_0 (7.71 GB)
- deepseek-coder-6.7B-instruct-Q6_K (5.53 GB)
- CodeLlama-7b-Instruct-hf-Q6_K (5.53 GB)

#### 0.1B - 4B

- gemma-2-2b-it-f16 (5.24 GB)
- Nemotron-Mini-4B-Instruct-Q8_0 (4.46 GB)
- Phi-3.5-mini-instruct-Q8_0 (4.06 GB)
- Qwen2.5-Coder-3B-Instruct-Q8_0 (3.62 GB)
- Ministral-3B-instruct-Q8_0 (3.52 GB)
- Llama-Doctor-3.2-3B-Instruct-Q8_0 (3.42 GB)
- SmolLM2-1.7B-Instruct-f16 (3.42 GB)
- Hermes-3-Llama-3.2-3B-Q8_0 (3.42 GB)
- stable-code-instruct-3B-Q8_0 (2.97 GB)
- Llama-3.2-1B-Instruct-f16 (2.48 GB)
- Qwen2.5-Coder-0.5B-Instruct-f16 (994.16 MB)
- Qwen2.5-0.5B-Instruct-f16 (994.16 MB)
- SmolLM2-360M-Instruct-f16 (725.55 MB)
- SmolLM2-135M-Instruct-f16 (270.89 MB)

</details>
