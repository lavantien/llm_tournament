# Advanced Go Programming Challenge: Distributed Task Processing System Evaluation

## Objective

Design a production-ready, distributed task processing system that demonstrates:

- Advanced Go concurrency patterns
- Robust system design
- Performance optimization
- Comprehensive testing and observability

## Evaluation Criteria

The generated solution will be assessed across multiple dimensions:

1. **Code Quality** (40 points)

   - Idiomatic Go code
   - Clean, readable design
   - Proper error handling
   - Adherence to Go best practices

2. **Concurrency & Performance** (30 points)

   - Efficient use of goroutines
   - Minimal resource contention
   - Low-overhead synchronization
   - Benchmarkable performance characteristics

3. **System Robustness** (20 points)

   - Handling of edge cases
   - Graceful error and failure management
   - Configurable and flexible design

4. **Testing & Observability** (10 points)
   - Comprehensive unit and integration tests
   - Meaningful logging and tracing
   - Performance profiling readiness

## Detailed Requirements

### System Architecture

Implement a distributed task processing system with the following core components:

#### 1. Work Queue Management

- Dynamically scalable worker pool
- Priority-based task queuing
- Configurable concurrency limits
- Support for task cancellation and timeouts

#### 2. Task Execution Framework

- Rate limiting mechanism
- Circuit breaker pattern implementation
- Centralized error handling
- Distributed tracing support

#### 3. Advanced Concurrency Patterns

- Use `sync.Once` for singleton initialization
- Semaphore-like worker pool control
- Channel-based communication and synchronization
- Context-driven cancellation

### Constraints & Guidelines

- **Language**: Pure Go (1.21+)
- **External Dependencies**: Limited to:
  - `prometheus/client_golang` for metrics
  - `opentelemetry` for distributed tracing
  - `testify` for testing
- **Code Structure**:
  - Single package with clear separation of concerns
  - Comprehensive godoc comments
  - No external library usage beyond specified
  - Production-ready, zero-allocation critical paths

### Deliverables

1. Complete Go package with:
   - Implementation files
   - Comprehensive tests
   - Performance benchmarks
2. Detailed `README.md` with:
   - System architecture overview
   - Usage examples
   - Performance considerations
   - Profiling recommendations

### Bonus Challenges

- Implement lock-free algorithms
- Minimize memory allocations in critical paths
- Add advanced rate limiting with adaptive strategies
- Demonstrate zero-downtime worker scaling

## Evaluation Methodology

Submissions will be rigorously tested via:

- `go vet` for code quality
- `go test -race` for concurrency safety
- `go test -bench=.` for performance benchmarks
- Manual code review against evaluation criteria

## Sample Validation Process

Evaluators will run the following commands to validate the implementation:

```bash
# Initial setup
go mod init taskprocessor
go mod tidy

# Run tests with race detector
go test -v -race ./...

# Run benchmarks
go test -bench=. -benchmem

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Scoring Notes

- Partial credit for incomplete but well-structured solutions
- Emphasis on clean, maintainable, and performant code
- Creative solutions that demonstrate deep understanding of Go's concurrency model

### Final Output Expectations

- Fully runnable package
- Passing all tests
- Comprehensive documentation
- Clear, idiomatic Go implementation
