# Detailed Design Document

![System UML](./llp-uml.png)

<details>
    <summary>...expand whole UML script</summary>

```plantuml
@startuml
class MainWindow {
  +openApplication()
  +displayLLMTable()
  +viewStatistics()
  +exportData()
}

MainWindow --> LLMList : displays list of models
MainWindow --> LLMDetail : opens selected model's details
MainWindow --> DB : retrieves benchmark results
MainWindow --> DB : retrieves export data



class LLMList {
  +selectLLM()
}

LLMList --> LLMDetail : opens the LLM details when selected



class LLMDetail {
  +pasteOutput(output: String)
  +viewModelDetails()
}

LLMDetail --> DB : stores output data
LLMDetail --> DB : stores prompt data
LLMDetail --> BR : stores benchmark results



class Database {
  +storeLLMData()
  +storeBenchmarkResults()
  +retrieveBenchmarkResults()
  +storePromptData()
  +retrieveExportData()
}

Database --> BenchmarkResults : stores and retrieves benchmark results
Database --> Prompts : stores and retrieves prompt data



class BenchmarkResults {
  +saveBenchmarkResult(modelID: String, result: String)
  +calculateScore(modelID: String): Score
}

BenchmarkResults --> Database : stores results



class Prompts {
  +storePromptText(prompt: String)
  +retrievePromptData(modelID: String): String
}

Prompts --> Database : stores and retrieves prompt data
@enduml
```

</details>

![System Sequence Diagram](./llp-sequence-diagram.png)

<details>
    <summary>...expand whole Sequence Diagram script</summary>

```plantuml
@startuml
actor User
participant MainWindow
participant LLMList

User -> MainWindow : openApplication()
MainWindow -> LLMList : displayLLMTable()
LLMList -> MainWindow : return LLM List
MainWindow -> User : show list of available LLMs



actor User
participant LLMList
participant LLMDetail

User -> LLMList : selectLLM()
LLMList -> LLMDetail : openLLMDetails()
LLMDetail -> User : show LLM Details



actor User
participant LLMDetail
participant Database
participant BenchmarkResults

User -> LLMDetail : pasteOutput(output)
LLMDetail -> Database : storeLLMData(modelDetails)
LLMDetail -> BenchmarkResults : saveBenchmarkResult(result)
BenchmarkResults -> Database : storeBenchmarkResults(result)



actor User
participant MainWindow
participant Database
participant BenchmarkResults

User -> MainWindow : viewStatistics()
MainWindow -> Database : retrieveBenchmarkResults()
Database -> BenchmarkResults : getBenchmarkResults()
BenchmarkResults -> MainWindow : return results
MainWindow -> User : show statistics



actor User
participant MainWindow
participant Database

User -> MainWindow : exportData()
MainWindow -> Database : retrieveExportData()
Database -> MainWindow : return export data (CSV/Markdown)
MainWindow -> User : show export data



participant LLMDetail
participant Database
participant BenchmarkResults
participant Prompts

LLMDetail -> Database : storeLLMData(modelDetails)
Database -> LLMDetail : acknowledge storage

LLMDetail -> Database : storeBenchmarkResults(results)
Database -> BenchmarkResults : store results

LLMDetail -> Database : storePromptData(promptText)
Database -> Prompts : store prompt

Database -> LLMDetail : retrieveBenchmarkResults()
Database -> Prompts : retrievePromptData(modelID)
@enduml
```

</details>

To create a detailed UML and sequence diagram for each component of the Local LLM Playground (LLP) system, we’ll break it down into several parts: the Main Window, LLM List, LLM Detail View, Database, Benchmark Results, and Prompts.

### 1. **Main Window Component UML**

```plantuml
@startuml
class MainWindow {
  +openApplication()
  +displayLLMTable()
  +viewStatistics()
  +exportData()
}

MainWindow --> LLMList : displays list of models
MainWindow --> LLMDetail : opens selected model's details
MainWindow --> DB : retrieves benchmark results
MainWindow --> DB : retrieves export data
@enduml
```

### 2. **LLM List Component UML**

```plantuml
@startuml
class LLMList {
  +selectLLM()
}

LLMList --> LLMDetail : opens the LLM details when selected
@enduml
```

### 3. **LLM Detail View Component UML**

```plantuml
@startuml
class LLMDetail {
  +pasteOutput(output: String)
  +viewModelDetails()
}

LLMDetail --> DB : stores output data
LLMDetail --> DB : stores prompt data
LLMDetail --> BR : stores benchmark results
@enduml
```

### 4. **Database Component UML**

```plantuml
@startuml
class Database {
  +storeLLMData()
  +storeBenchmarkResults()
  +retrieveBenchmarkResults()
  +storePromptData()
  +retrieveExportData()
}

Database --> BenchmarkResults : stores and retrieves benchmark results
Database --> Prompts : stores and retrieves prompt data
@enduml
```

### 5. **Benchmark Results Component UML**

```plantuml
@startuml
class BenchmarkResults {
  +saveBenchmarkResult(modelID: String, result: String)
  +calculateScore(modelID: String): Score
}

BenchmarkResults --> Database : stores results
@enduml
```

### 6. **Prompts Component UML**

```plantuml
@startuml
class Prompts {
  +storePromptText(prompt: String)
  +retrievePromptData(modelID: String): String
}

Prompts --> Database : stores and retrieves prompt data
@enduml
```

---

## Sequence Diagrams

### 1. **Main Window - Opening the Application and Viewing LLM List**

```plantuml
@startuml
actor User
participant MainWindow
participant LLMList

User -> MainWindow : openApplication()
MainWindow -> LLMList : displayLLMTable()
LLMList -> MainWindow : return LLM List
MainWindow -> User : show list of available LLMs
@enduml
```

### 2. **LLM List - Selecting an LLM and Viewing Details**

```plantuml
@startuml
actor User
participant LLMList
participant LLMDetail

User -> LLMList : selectLLM()
LLMList -> LLMDetail : openLLMDetails()
LLMDetail -> User : show LLM Details
@enduml
```

### 3. **LLM Detail View - Pasting Output and Storing Benchmark Results**

```plantuml
@startuml
actor User
participant LLMDetail
participant Database
participant BenchmarkResults

User -> LLMDetail : pasteOutput(output)
LLMDetail -> Database : storeLLMData(modelDetails)
LLMDetail -> BenchmarkResults : saveBenchmarkResult(result)
BenchmarkResults -> Database : storeBenchmarkResults(result)
@enduml
```

### 4. **Main Window - Viewing Statistics**

```plantuml
@startuml
actor User
participant MainWindow
participant Database
participant BenchmarkResults

User -> MainWindow : viewStatistics()
MainWindow -> Database : retrieveBenchmarkResults()
Database -> BenchmarkResults : getBenchmarkResults()
BenchmarkResults -> MainWindow : return results
MainWindow -> User : show statistics
@enduml
```

### 5. **Main Window - Exporting Data**

```plantuml
@startuml
actor User
participant MainWindow
participant Database

User -> MainWindow : exportData()
MainWindow -> Database : retrieveExportData()
Database -> MainWindow : return export data (CSV/Markdown)
MainWindow -> User : show export data
@enduml
```

### 6. **Database - Storing and Retrieving Data**

```plantuml
@startuml
participant LLMDetail
participant Database
participant BenchmarkResults
participant Prompts

LLMDetail -> Database : storeLLMData(modelDetails)
Database -> LLMDetail : acknowledge storage

LLMDetail -> Database : storeBenchmarkResults(results)
Database -> BenchmarkResults : store results

LLMDetail -> Database : storePromptData(promptText)
Database -> Prompts : store prompt

Database -> LLMDetail : retrieveBenchmarkResults()
Database -> Prompts : retrievePromptData(modelID)
@enduml
```

---

### Breakdown of Components in UML and Sequence:

1. **Main Window** handles the user interface. It interacts with the LLM List to display available models and with the LLM Detail View to display the selected model’s details. It also retrieves benchmark statistics and export data from the database.

2. **LLM List** displays all the LLMs. When a user selects a model, it triggers the LLM Detail View.

3. **LLM Detail View** shows the detailed view of a selected model. The user can paste the model's output here, which is then stored in the database. Benchmark results are also saved and calculated in this component.

4. **Database** is responsible for storing and retrieving all necessary data, including LLM details, benchmark results, and prompts.

5. **Benchmark Results** stores benchmark results and calculates scores. It interfaces with the Database for persistent storage.

6. **Prompts** stores the prompt data used during LLM evaluations. It also retrieves prompt data for display in the LLM Detail View.

This architecture uses simple interactions between components to ensure scalability, maintainability, and ease of understanding, adhering to the limitations of using Golang, SQLite, and Fyne as the core technologies for the system.
