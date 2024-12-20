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

