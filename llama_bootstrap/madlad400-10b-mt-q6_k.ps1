$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\sander3\madlad400-10b-mt-Q6_K-GGUF\madlad400-10b-mt-q6_k.gguf"

$params = @{
    "gpu-layers" = 32
    "ctx-size" = 512
    "batch-size" = 512
    "threads" = 8
    "keep" = 512
    "predict" = -1
    #"flash-attn" = $true
    "mlock" = $true
    #"cache-type-k" = "q8_0"
    #"cache-type-v" = "q8_0"
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

