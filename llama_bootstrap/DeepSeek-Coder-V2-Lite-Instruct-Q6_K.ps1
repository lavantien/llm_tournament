# Define the model path and parameters
$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\lmstudio-community\DeepSeek-Coder-V2-Lite-Instruct-GGUF\DeepSeek-Coder-V2-Lite-Instruct-Q6_K.gguf"

# Define parameters as an object for readability
$params = @{
    "gpu-layers" = 8
    "ctx-size" = 32768
    "batch-size" = 512
    "threads" = 8
    "keep" = 4096
    "predict" = -1
    "mlock" = $true
}

# Build the llama-server command with the specified parameters
# Assumed `llama-server` binaries are in the $PATH
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

# Run the command
Start-Process -FilePath "pwsh" -ArgumentList "-Command $cmd --no-context-shift" -NoNewWindow -Wait

