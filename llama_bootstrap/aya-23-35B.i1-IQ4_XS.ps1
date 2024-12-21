$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\mradermacher\aya-23-35B-i1-GGUF\aya-23-35B.i1-IQ4_XS.gguf"

$params = @{
    "gpu-layers" = 12
    "ctx-size" = 8192
    "batch-size" = 512
    "threads" = 8
    "keep" = 4096
    "predict" = -1
    "flash-attn" = $true
    "mlock" = $true
    "cache-type-k" = "q4_0"
    "cache-type-v" = "q4_0"
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

