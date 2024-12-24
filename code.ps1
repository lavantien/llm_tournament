$baseProfile = @{
    "gpu-layers" = 20
    "ctx-size" = 16384
    "batch-size" = 512
    "threads" = 8
    "keep" = 4096
    "predict" = -1
    "cache-type-k" = "q8_0"
    "cache-type-v" = "q8_0"
    # "flash-attn" = $true
    # "mlock" = $true
    # "verbose-prompt" = $true
    # "verbose" = $true
    # "log-prefix" = $true
    # "log-colors" = $true
}

. ./shared.ps1

$modelPath = $rootPath + "Qwen/Qwen2.5-Coder-14B-Instruct-GGUF/qwen2.5-coder-14b-instruct-q8_0-00001-of-00002.gguf"
$profileName = "Programming"
Run-LlamaCli -modelPath $modelPath -profileName $profileName -baseProfile $baseProfile
