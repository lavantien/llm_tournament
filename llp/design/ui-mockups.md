# UI Mockups

Here’s how your UI mockups can transition to Bubble Tea, leveraging its TUI capabilities. I'll adapt each section of your GUI mockup to a Bubble Tea-friendly terminal interface.

---

### **Main Menu**

```
+-------------------------------------+
|          LLM Benchmark App          |
+-------------------------------------+
|  [1] Model Manager                  |
|  [2] Prompt Manager                 |
|  [3] Benchmark Results              |
|  [4] Statistics                     |
|  [5] Export Data                    |
|  [Q] Quit                           |
+-------------------------------------+
```

**Navigation:**

- Users press the corresponding number/letter key to navigate.

---

### **Model Manager**

```
+---------------------------------------------+
|                Model Manager                |
+---------------------------------------------+
|  # | Model Name      | Type       | Status  |
|----|-----------------|------------|---------|
|  1 | LLaMA          | Language   | Ready   |
|  2 | Qwen           | Coder      | Running |
|  3 | DeepSeek       | Lite       | Idle    |
+---------------------------------------------+
|  [A] Add New Model  [D] Delete Model       |
|  [V] View Details   [B] Back to Main Menu  |
+---------------------------------------------+
```

**Interaction:**

- Users press `A` to add a model, `D` to delete a selected model, or `V` to view model details.

---

### **Model Details**

```
+---------------------------------------------+
|                Model Details                |
+---------------------------------------------+
| Model Name:       LLaMA                    |
| Type:             Language Model           |
| Status:           Ready                    |
| Max Context:      8192                     |
| Speed:            2.3 tok/sec              |
+---------------------------------------------+
|  [E] Edit Model  [B] Back to Model Manager |
+---------------------------------------------+
```

---

### **Prompt Manager**

```
+---------------------------------------------+
|                Prompt Manager               |
+---------------------------------------------+
|  # | Task Name               | Points       |
|----|-------------------------|--------------|
|  1 | Warmup                 | 10           |
|  2 | Story Writing          | 45           |
|  3 | Translation            | 55           |
|  4 | Complex Problems       | 75           |
+---------------------------------------------+
|  [A] Add Prompt  [E] Edit Prompt           |
|  [V] View Prompt [B] Back to Main Menu     |
+---------------------------------------------+
```

---

### **Prompt Input Form**

```
+---------------------------------------------+
|                Add/Edit Prompt              |
+---------------------------------------------+
| Task Name: _______________________________  |
| Points:     _______________________________  |
| Prompt:                                     |
| -----------------------------------------   |
| |                                       |   |
| |                                       |   |
| |                                       |   |
| -----------------------------------------   |
+---------------------------------------------+
|  [S] Save  [C] Cancel                      |
+---------------------------------------------+
```

---

### **Benchmark Results**

```
+---------------------------------------------+
|             Benchmark Results               |
+---------------------------------------------+
|  Model Name      | Task Name     | Points   |
|------------------|---------------|----------|
| LLaMA            | Warmup        | 10       |
| Qwen             | Translation   | 55       |
| DeepSeek         | Story Writing | 45       |
+---------------------------------------------+
|  [V] View Details  [B] Back to Main Menu   |
+---------------------------------------------+
```

---

### **Statistics**

```
+---------------------------------------------+
|                 Statistics                  |
+---------------------------------------------+
|  Total Points:           5000              |
|  Models Evaluated:       15                |
|  Tasks Completed:        45                |
+---------------------------------------------+
|  [B] Back to Main Menu                      |
+---------------------------------------------+
```

---

### **Export Data**

```
+---------------------------------------------+
|                Export Data                  |
+---------------------------------------------+
|  File Name: _______________________________ |
|  Format:                                    |
|   [1] CSV                                   |
|   [2] JSON                                  |
|   [3] XML                                   |
+---------------------------------------------+
|  [S] Save  [C] Cancel                      |
+---------------------------------------------+
```

---

### **Next Steps**

1. Implement navigation logic using **Bubble Tea's Model-Update-View pattern**.
2. Use Bubble Tea's **textinput** and **tables** packages to handle user inputs and display tabular data.
3. Assign data persistence to SQLite, connecting models, prompts, and results.
