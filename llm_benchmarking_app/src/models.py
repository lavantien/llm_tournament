from dataclasses import dataclass
from typing import List

@dataclass
class Prompt:
    id: int
    content: str
    category: str

@dataclass
class Profile:
    id: int
    name: str
    system_prompt: str
    parameters: dict

@dataclass
class LLM:
    id: int
    name: str
    version: str
    description: str

@dataclass
class Attempt:
    id: int
    llm_id: int
    prompt_id: int
    score: float
    output: str
    timestamp: str
