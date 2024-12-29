$baseProfile = @{
    "gpu-layers" = 29
    "ctx-size" = 8192
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

$modelPath = $rootPath + "TheDrummer/Tiger-Gemma-9B-v3-GGUF/Tiger-Gemma-9B-v3-Q8_0.gguf"
$profileName = "Default"
Run-LlamaCli -modelPath $modelPath -profileName $profileName -baseProfile $baseProfile
