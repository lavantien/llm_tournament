# Define the model path and parameters
$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\DevQuasar\Mistral-Small-Instruct-2409-GGUF\Mistral-Small-Instruct-2409.Q8_0.gguf"

# Define parameters as an object for readability
$params = @{
    "gpu-layers" = 17
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

