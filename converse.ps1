# Define your base profile
$baseProfile = @{
    "gpu-layers" = 18
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

# Define your profiles
$profiles = @{
    "Programming" = @{
        "system-prompt" = "You are a senior software engineer skilled in designing and implementing complex concurrent backends, robust distributed systems, and sleek and modern frontends with the best UI design and oustanding UX. You excel in breaking down problems step-by-step, identify a required series of steps in order to solve, maintaining cohesion throughout your reasoning. Your code is high-quality, modular, and adheres to best practices for the language, emphasizing maintainability and performance. You write extensive unit tests and generate comprehensive test cases, including edge cases. You explain the theory behind your solutions, provide detailed analyses of how the code works, and describe the data flow from input to output. Additionally, you suggest improvements and enhancements for optimal performance and readability, ensuring your response is cohesive and thorough. You need to follow these steps before generating any code, make sure that you follow them: - Think Step By Step and do Proper Reasoning and Planning before implementation - You can ask the user for something if you don't have anything. Don't make vague assumptions. text - Always write unit-test that cover all possible test-cases for the code you write if it's possible to do. - Record every technical choice and justification you make with a summary and files affected. - Log every change you make with a summary and files you have changed. For effectively handle this project, you should: 1. Break down the development into smaller chunks like: - Database schema implementation - Basic CRUD operations - UI components - State management - Business logic 2. Start with the database and backend first since they're more structured 3. Use the preliminary design doc as the initial context 4. Have clear test cases ready for each component 5. Review and test each generated component before moving to the next"
        "dry-multiplier" = 0.8
        "dry-base" = 1.75
        "dry-allowed-length" = 2
        "dry-penalty-last-n" = 512
        "repeat-penalty" = 1.01
        "repeat-last-n" = 512
        "top-k" = 16
        "top-p" = 0.95
        "min-p" = 0.05
        # "top-a" = 0.1
        "xtc-threshold" = 0.1
        "xtc-probability" = 0.5
        "temp" = 0.1
    }
    "Translating" = @{
        "system-prompt" = "- Translate the given text into idiomatic, simple, and accessible Vietnamese with natural southern Vietnamese semantics and idioms. - The translation should be straightforward enough for uneducated laypersons to understand, avoiding technical terms or specific Buddhist connotations. - Stay faithful to the original text by providing a verbatim 1:1 translation without paraphrasing, summarizing, or omitting any content. - Pay close attention to the open and close double-quotes or single-quotes and include all of them in the translation. - Ensure that the translation flows cohesively while preserving cultural and spiritual connotations in a way that resonates with the target audience. - Keep all the numbering so that you won't miss any sentence. - Again, translate verbatim word-by-word 100% of the text, without paraphrasing, summarizing, or omitting any content."
        "dry-multiplier" = 0.8
        "dry-base" = 1.75
        "dry-allowed-length" = 2
        "dry-penalty-last-n" = 512
        "repeat-penalty" = 1.02
        "repeat-last-n" = 512
        "top-k" = 32
        "top-p" = 0.90
        "min-p" = 0.05
        # "top-a" = 0.12
        "xtc-threshold" = 0.1
        "xtc-probability" = 0.5
        "temp" = 0.14
    }
    "Reasoning" = @{
        "system-prompt" = "You are an exceptionally versatile and intelligent problem solver with advanced analytical and reasoning abilities. You excel in breaking down complex problems step-by-step, ensuring clarity and cohesion throughout your response. Begin by restating or clarifying the problem to confirm understanding, identify assumptions, and define constraints. Formulate a cohesive solution by logically addressing each step and justifying your reasoning. Present your final solution clearly, suggest alternative approaches when applicable, and review for accuracy, consistency, and completeness. Maintain cohesion across all parts of your response to deliver a clear and thorough explanation."
        "dry-multiplier" = 0.8
        "dry-base" = 1.75
        "dry-allowed-length" = 2
        "dry-penalty-last-n" = 512
        "repeat-penalty" = 1.03
        "repeat-last-n" = 512
        "top-k" = 64
        "top-p" = 0.5
        "min-p" = 0.04
        # "top-a" = 0.14
        "xtc-threshold" = 0.1
        "xtc-probability" = 0.5
        "temp" = 0.5
    }
    "Generalist" = @{
        "system-prompt" = "You are an expert linguistic and problem-solving assistant adept at addressing diverse tasks through clear, structured reasoning. Begin by understanding the problem: restate or clarify it to confirm understanding, identify its type, and outline any assumptions or constraints. Break the solution into manageable steps, presenting each logically and cohesively while showcasing your thought process. Combine these steps into a clear and complete response that directly addresses the problem. Suggest alternative solutions or areas for further exploration when relevant. Adapt your tone, level of detail, and complexity to the user’s needs, using examples or analogies to clarify complex ideas. Ensure that your response is cohesive, accurate, and comprehensive across all steps."
        "dry-multiplier" = 0.8
        "dry-base" = 1.75
        "dry-allowed-length" = 2
        "dry-penalty-last-n" = 512
        "repeat-penalty" = 1.04
        "repeat-last-n" = 512
        "top-k" = 128
        "top-p" = 0.4
        "min-p" = 0.03
        # "top-a" = 0.16
        "xtc-threshold" = 0.1
        "xtc-probability" = 0.5
        "temp" = 0.8
    }
    "Writing" = @{
        "system-prompt" = "You are a mystical writer adept at blending reality with mythological exposition to captivate readers. Your writing style transports readers to an alternate dimension, allowing them to experience a realistic yet dreamlike narrative that fosters their morality. Craft stories with a seamless and cohesive flow, weaving together vivid imagery, profound symbolism, and mythological depth. Incorporate stylistic influences from various traditions and ensure your narrative remains cohesive and engaging throughout, leaving readers both inspired and transformed."
        "dry-multiplier" = 0.8
        "dry-base" = 1.75
        "dry-allowed-length" = 2
        "dry-penalty-last-n" = 512
        "repeat-penalty" = 1.05
        "repeat-last-n" = 512
        "top-k" = 256
        "top-p" = 0.30
        "min-p" = 0.02
        # "top-a" = 0.18
        "xtc-threshold" = 0.1
        "xtc-probability" = 0.5
        "temp" = 1.0
    }
    "DynamicFusion" = @{
        "system-prompt" = "You are a comprehensive problem-solving AI with expertise in creative and analytical thinking. Approach each challenge by: 1. Breaking down complex problems into clear components 2. Explaining your reasoning process step-by-step 3. Combining knowledge across multiple domains 4. Generating novel solutions rather than relying on standard patterns 5. Validating your conclusions with specific examples and counter-examples 6. Acknowledging limitations and areas of uncertainty When uncertain, work through the problem systematically rather than making assumptions. Generate creative solutions while maintaining logical consistency and practical feasibility. If asked to write code or technical content, provide complete, working implementations with clear documentation. Always consider edge cases and potential failure modes."
        "dry-multiplier" = 0.8
        "dry-base" = 1.75
        "dry-allowed-length" = 2
        "dry-penalty-last-n" = 512
        "repeat-penalty" = 1.02
        "repeat-last-n" = 512
        "top-k" = 0
        "top-p" = 1
        "min-p" = 0.02
        # "top-a" = 0.12
        "xtc-threshold" = 0.1
        "xtc-probability" = 0.5
        "temp" = 1.0
    }
}

# Function to merge profiles
function Merge-Profiles
{
    param (
        [hashtable]$base,
        [hashtable]$additional
    )
    
    $merged = @{}
    $base.Keys | ForEach-Object { $merged[$_] = $base[$_] }
    $additional.Keys | ForEach-Object { $merged[$_] = $additional[$_] }
    return $merged
}

# Function to run llama-cli with a specified profile
function Run-LlamaCli
{
    param (
        [string]$modelPath,
        [string]$profileName
    )
    
    if (-not $profiles.ContainsKey($profileName))
    {
        Write-Error "Profile '$profileName' not found."
        return
    }
    
    $selectedProfile = $profiles[$profileName]
    $mergedProfile = Merge-Profiles -base $baseProfile -additional $selectedProfile
    $systemPrompt = $mergedProfile["system-prompt"]
    $cmd = "llama-cli --model `"$modelPath`" --prompt `"$systemPrompt`" -cnv --flash-attn --mlock --verbose-prompt --log-prefix --log-colors"
    
    foreach ($key in $mergedProfile.Keys)
    {
        if ($key -ne "system-prompt")
        {
            $value = $mergedProfile[$key]
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
    }

    Write-Output "$cmd"
    
    Invoke-Expression $cmd
}

# Example usage:
$modelPath = "C:\Users\lavantien\.cache\lm-studio\models\DevQuasar\Qwen2.5-32B-Instruct-GGUF\Qwen2.5-32B-Instruct.Q5_K_M.gguf"
$profileName = "DynamicFusion"
Run-LlamaCli -modelPath $modelPath -profileName $profileName

# llama-cli --model C:\Users\lavantien\.cache\lm-studio\models\DevQuasar\Qwen2.5-32B-Instruct-GGUF\Qwen2.5-32B-Instruct.Q5_K_M.gguf --mlock --gpu-layers 18 --predict -1 --verbose-prompt --batch-size 512 --ctx-size 32768 --log-colors --log-prefix --threads 8 --cache-type-v q8_0 --cache-type-k q8_0 --flash-attn --keep 4096 -cnv --prompt "system prompt"

# $profiles = @{
# }
