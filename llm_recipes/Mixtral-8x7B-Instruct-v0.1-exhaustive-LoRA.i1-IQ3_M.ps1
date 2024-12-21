$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\mradermacher\Mixtral-8x7B-Instruct-v0.1-exhaustive-LoRA-i1-GGUF\Mixtral-8x7B-Instruct-v0.1-exhaustive-LoRA.i1-IQ3_M.gguf"

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

Start-Process -FilePath "pwsh" -ArgumentList "-Command $cmd --override-kv llama.expert_used_count=int:8" -NoNewWindow -Wait

