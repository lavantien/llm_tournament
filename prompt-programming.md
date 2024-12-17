# Programming Prompt

## System Prompt

"You are a hobbyist programmer with deep passion for programming and problem solving. You love doing recreational programming and creative programming work."

## Warming Up

### Prompt 1

Write me a program that generates a spinning fractals.

### Prompt 2

You are an expert python game developer. Create a snake game with blockages and highscore.

### Prompt 3

Coding Challenge: "Write a Python function to implement a basic recommendation system for movies using collaborative filtering."

### Prompt 4

Technical Problem-Solving Challenge.
Design an efficient algorithm to find the shortest path between multiple cities while minimizing total transportation costs, accounting for variable fuel prices, road conditions, and time constraints. Provide a pseudocode implementation and explain your algorithmic reasoning.

---

## Instructions for Each Following Task

At the end of each prompt, request the following:

1. Provide instructions to run the code locally or via Docker or DockerCompose, with examples of the expected commands.
2. Ensure all dependencies and setup steps are documented.
3. Include any test cases or demonstration scripts if applicable.

These prompts cover a wide range of skills and technologies, pushing the LLMs to handle practical, real-world programming challenges.

---

## Topics

### **General Programming Concepts**

1. **Data Structures and Algorithms (Go)**  
   Implement a thread-safe, fixed-size circular buffer (ring buffer) in Go. Provide a demonstration of its usage in a concurrent program simulating a producer-consumer model.

2. **File I/O and Parsing (Python)**  
   Write a Python program to parse a log file with entries in the format `[timestamp] [log_level] [message]`. Group the entries by `log_level` and save them into separate files named `<log_level>.log`.

---

### **Concurrency and Distributed Systems**

3. **Concurrency (Rust)**  
   Implement a multi-threaded Rust program to calculate the frequency of words in a given text file using a thread pool. Each thread processes a portion of the file.

4. **Distributed Systems (Go)**  
   Create a simple distributed key-value store in Go using gRPC. The store should support `PUT` and `GET` operations. Provide a script to simulate clients sending requests concurrently.

---

### **Web Development**

5. **Full-Stack App (JS/HTML/CSS)**  
   Create a basic web-based to-do app. Implement the backend in Go with RESTful APIs and the frontend using HTML, CSS, and JavaScript. Provide instructions to run it locally.

6. **Canvas Drawing App (JS/HTML/Canvas)**  
   Develop a drawing application in the browser using the HTML Canvas API. Include features like selecting brush size and color, undo, and clear canvas.

---

### **Databases**

7. **Database Integration (Go + SQLite)**  
   Write a Go program to manage a library system. The system should allow adding, updating, and retrieving books from an SQLite database. Provide SQL schema and example queries.

8. **Database Optimization (Python + SQLite)**  
   Design and optimize an SQLite query to retrieve the top 5 customers by total spending from a sample `orders` database. Provide a Python script to demonstrate the query execution.

---

### **Advanced Topics**

9. **Concurrency Control (Rust)**  
   Design and implement a bank account system in Rust with support for concurrent deposits and withdrawals. Use synchronization primitives to avoid race conditions.

10. **Distributed Messaging (Go)**  
    Build a simple messaging system using NATS or RabbitMQ in Go. Simulate multiple publishers and subscribers communicating concurrently.

---

### **Testing and Debugging**

11. **Unit Testing (Python)**  
    Write a Python library for a basic calculator supporting addition, subtraction, multiplication, and division. Include a complete suite of unit tests using `unittest`.

12. **Profiling (Go)**  
    Provide a Go program with a computationally expensive function. Use `pprof` to identify performance bottlenecks and optimize the function.

---

### **Infrastructure**

13. **Dockerize Application (Rust)**  
    Create a Rust application that serves a basic "Hello, World!" web page over HTTP. Write a Dockerfile to containerize the application and include instructions to run it.

14. **CI/CD Pipeline (Go)**  
    Write a Go program that includes unit tests. Provide a `Jenkinsfile` or `GitHub Actions` workflow to build, test, and deploy the program to a local server.

---

### **Miscellaneous**

15. **Real-Time System (JS/WebSocket)**  
    Develop a web-based chat application. Use a Go server to handle WebSocket connections and JavaScript for the client. Include a simple demo to run locally.

16. **Task Scheduling (Python)**  
    Write a Python program to schedule and execute tasks at specific intervals (like a cron job). Use multithreading to run tasks concurrently and log their outputs.

---

## Specific Technologies

### **Go**

1. **Basic Task: REST API**

   - Write a Go program to implement a REST API with endpoints to add, retrieve, update, and delete books. Use only the standard library and SQLite. Include instructions to run locally and with Docker.

2. **Concurrency: Worker Pool**

   - Write a Go program that processes a list of 1,000 URLs using a worker pool with 10 workers. Simulate fetching each URL using `time.Sleep`. Include metrics on processing time.

3. **Distributed System: gRPC Service**
   - Design a gRPC service in Go for a chat application. Include `SendMessage` and `ListMessages` RPCs. Use SQLite for storage. Provide steps to run the service and a simple client locally.

---

### **Rust**

4. **Basic Task: CLI Tool**

   - Build a Rust CLI tool to manage a to-do list. Use SQLite for persistence. The tool should support adding, listing, updating, and deleting tasks. Include instructions for running it locally.

5. **Concurrency: Parallel File Processing**

   - Write a Rust program that reads 10 large files concurrently and counts the number of lines in each file. Demonstrate performance optimizations using `tokio` or similar.

6. **Distributed System: Message Queue Simulation**
   - Implement a lightweight message queue in Rust that allows publishers to send messages to topics and consumers to subscribe to those topics. Ensure messages are processed at-least-once.

---

### **Python**

7. **Basic Task: Web Scraper**

   - Create a Python script to scrape the top 10 news headlines from a static HTML file (provided as input). Parse the content using `BeautifulSoup` or a standard-library alternative. No external requests allowed.

8. **Concurrency: Async Task Runner**

   - Write an async Python script that simulates downloading 1,000 files using `asyncio` with a maximum concurrency of 50. Demonstrate how to handle retries on failure.

9. **Distributed System: Simple Key-Value Store**
   - Implement a distributed key-value store using Python. Use `multiprocessing` to simulate communication between three nodes running locally. Provide simple CRUD operations.

---

### **JavaScript/HTML/CSS**

10. **Basic Task: Canvas Drawing**

    - Build an HTML/JavaScript application where users can draw shapes on an HTML5 `<canvas>` element. Include controls to change colors, size, and clear the canvas.

11. **Concurrency: Web Worker**

    - Create a JavaScript application using Web Workers to compute the Fibonacci sequence for a given number without blocking the UI. Include an input field and a result display.

12. **Full-Stack: Simple Blog**
    - Build a simple blog application with Node.js, Express, SQLite, and HTML/CSS. Users should be able to add, edit, delete, and view blog posts. Provide steps to run locally and via Docker.

---

### **SQLite**

13. **Database Optimization: Query Performance**

    - Given a table `orders` with 1 million rows, write SQL queries to:
      - Retrieve the total sales grouped by month.
      - Find the top 5 customers by total purchase amount.
      - Provide query optimization recommendations for SQLite.

14. **Data Migration: Schema Evolution**
    - Write a Python or Go script that migrates a SQLite database schema from:
      - Version 1: `users(id, name)`
      - To Version 2: `users(id, name, email)`.
    - Include steps for backing up the database and migrating existing data.

---

### **Concurrency and Distributed Systems**

15. **Concurrency: Dining Philosophers**

    - Implement the "Dining Philosophers" problem in Go, Rust, or Python using concurrency primitives (e.g., goroutines, async/await, threads). Demonstrate proper deadlock handling.

16. **Distributed System: Leader Election**
    - Design and implement a leader election algorithm using Python or Go, simulating nodes on a single machine. Use either Raft or a simple custom algorithm. Ensure fault tolerance.

---
