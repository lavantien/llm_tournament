$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\tensorblock\DeepSeek-Coder-V2-Lite-Instruct-GGUF\DeepSeek-Coder-V2-Lite-Instruct-Q8_0.gguf"

$params = @{
    "gpu-layers" = 7
    "ctx-size" = 32768
    "batch-size" = 512
    "threads" = 8
    "keep" = 4096
    "predict" = -1
    "mlock" = $true
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

Start-Process -FilePath "pwsh" -ArgumentList "-Command $cmd --no-context-shift --override-kv llama.expert_used_count=int:64" -NoNewWindow -Wait

