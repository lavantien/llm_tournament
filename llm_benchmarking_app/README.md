# LLM Benchmarking App

This application is designed to benchmark Large Language Models (LLMs) using a variety of prompts and profiles. It provides a user-friendly interface for managing prompts, entering scores, viewing leaderboards, and visualizing results.

## Features

- Prompt Management: Load, edit, and save prompts.
- Score Entry: Enter scores for LLM attempts.
- Leaderboard: View the leaderboard of LLM scores.
- Visualization: Visualize LLM scores using charts.

## Installation

1. Clone the repository:
   ```sh
   git clone <repository_url>
   ```

2. Navigate to the project directory:
   ```sh
   cd llm_benchmarking_app
   ```

3. Install the required dependencies:
   ```sh
   pip install -r requirements.txt
   ```

## Usage

1. Run the application:
   ```sh
   python src/main.py
   ```

2. Use the GUI to manage prompts, enter scores, view the leaderboard, and visualize results.

## Directory Structure

- `data/`: Contains JSON files for prompts and profiles.
- `db/`: Contains the SQLite database file.
- `src/`: Contains the source code for the application.
- `tests/`: Contains test cases for the application.
- `requirements.txt`: List of required dependencies.
- `README.md`: Documentation for the application.
