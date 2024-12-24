$baseProfile = @{
    "gpu-layers" = 9
    "ctx-size" = 32768
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

$modelPath = $rootPath + "bartowski/aya-expanse-32b-GGUF/aya-expanse-32b-Q5_K_L.gguf"
$profileName = "Translating"
Run-LlamaCli -modelPath $modelPath -profileName $profileName -baseProfile $baseProfile
