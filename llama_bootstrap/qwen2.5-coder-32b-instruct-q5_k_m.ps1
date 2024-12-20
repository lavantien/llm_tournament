# Define the model path and parameters
$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\Qwen\Qwen2.5-Coder-32B-Instruct-GGUF\qwen2.5-coder-32b-instruct-q5_k_m-00001-of-00003.gguf"

# Define parameters as an object for readability
$params = @{
    "gpu-layers" = 18
    "ctx-size" = 32768
    "batch-size" = 512
    "threads" = 8
    "keep" = 4096
    "predict" = -1
    "flash-attn" = $true
    "mlock" = $true
    "cache-type-k" = "q8_0"
    "cache-type-v" = "q8_0"
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
Start-Process -FilePath "pwsh" -ArgumentList "-Command $cmd" -NoNewWindow -Wait

