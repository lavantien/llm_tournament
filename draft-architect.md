I am building a native desktop LLM benchmarking and visualization app using Python and PyQt6 for Windows 11. The app will manage and display benchmarking results entered manually. Here are the requirements:

### Features:

1. **Prompt Suite Management:**

   - Load JSON prompt suites and profiles from local files. I've prepared the data in `profiles_and_prompt_suites.md`.
   - Display loaded prompts in list format.
   - Allow users to view, edit, categorize prompts.

2. **Score and Output Entry:**

   - Enable manual entry of LLM scores and outputs for each prompt.
   - Include fields for individual scores and a calculated total score.
   - Validate and save data to a SQLite database.

3. **Database Integration:**

   - Use SQLite for persistent storage.
   - Store prompt data, LLM metadata, attempts, scores, and outputs.
   - Provide import/export options for JSON formats.

4. **Leaderboard:**

   - Display all LLMs with their scores and total scores in a ranked format.
   - Allow sorting and filtering based on different metrics.

5. **Visualization Tab:**

   - Include charts for comparing scores across LLMs and prompt suites.
   - Provide bar charts, line graphs, or pie charts for key metrics.
   - Enable filtering and customization of displayed data.

6. **User Experience:**
   - Design a clean and intuitive PyQt6 interface with tabs for prompt management, leaderboard, and visualization.
   - Include responsive elements suitable for desktop use.
   - Provide tooltips and help dialogs for guidance.

### Constraints:

- Python and PyQt6 should be used exclusively for the GUI and logic.
- SQLite will be the database backend.
- The app must be fully offline and optimized for Windows 11.

### Questions:

- What is the recommended structure for this app to ensure maintainability?
- How should I design the SQLite schema to efficiently handle prompt data, scores, and outputs?
- What libraries or tools can simplify chart generation and data visualization in PyQt6?

Provide a detailed implementation plan, including recommendations for file handling, SQLite integration, and PyQt6 GUI design best practices.
