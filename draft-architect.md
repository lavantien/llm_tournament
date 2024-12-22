I want to make a spreadsheet-like web app with NextJS and SQLite. To benchmark and conduct tournaments for the LLMs.

- Each prmopt should have:
- Model Manager: CRUD with models. Model schema:

```
"name": "c4ai-command-r-08-2024-GGUF",
"path" = "C:\Users\lavantien\.cache\lm-studio\models\tensorblock\c4ai-command-r-08-2024-GGUF\c4ai-command-r-08-2024-Q5_K_M.gguf"
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
```

- Profile Manager tab: CRUD with sampler profiles. Profile schema:

```
"name" = "Writing Profile"
"System prompt" = "You are a mystical writer adept at blending reality with mythological exposition to captivate readers. Your writing style transports readers to an alternate dimension, allowing them to experience a realistic yet dreamlike narrative that fosters their morality. Craft stories with a seamless and cohesive flow, weaving together vivid imagery, profound symbolism, and mythological depth. Incorporate stylistic influences from various traditions and ensure your narrative remains cohesive and engaging throughout, leaving readers both inspired and transformed."
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
```

- Prompt Manager tab: CRUD with prompts. The listing view is collapsable at category level. Promopt schema:

```
content, solution, profile, category
```

- Leaderboard tab: full scrollable table of all models with all their category total scores and overal score. It sortable on every column.
  - Click on each row will pop up a detail page for entering attempt and scores and output.
