# I D:\a\llama.cpp\llama.cpp\src\llama.cpp:20117: llama_new_context_with_model: n_seq_max     = 1
# GGML_ASSERT(hparams.n_embd_head_k % ggml_blck_size(type_k) == 0) failed

$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\mradermacher\mpt-30b-instruct-i1-GGUF\mpt-30b-instruct.i1-Q5_K_M.gguf"

$params = @{
    "gpu-layers" = 14
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

Start-Process -FilePath "pwsh" -ArgumentList "-Command $cmd" -NoNewWindow -Wait

