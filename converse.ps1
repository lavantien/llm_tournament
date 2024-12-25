$baseProfile = @{
    "gpu-layers" = 20
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

$modelPath = $rootPath + "arcee-ai/Virtuoso-Small-GGUF/Virtuoso-Small-Q8_0.gguf"
$profileName = "Generalist"
Run-LlamaCli -modelPath $modelPath -profileName $profileName -baseProfile $baseProfile
