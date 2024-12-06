# Concurrent Prime Number Generator and Analyzer

## Task Overview

Develop a concurrent system for generating, analyzing, and processing prime numbers that demonstrates advanced Go programming skills, focusing on efficient concurrency patterns, error handling, and clean code design.

## Requirements

### Core Functionality

- Create a package `primeanalyzer` with two main components:
  1. A concurrent prime number generator
  2. A prime number processor with advanced analysis capabilities

### Prime Number Generator (`generator.go`)

- Implement an efficient concurrent prime number generation algorithm
- Support configurable range and concurrency levels
- Generate primes using an optimized sieve method
- Provide channels for prime distribution

### Prime Number Processor (`processor.go`)

- Analyze generated prime numbers
- Perform complex mathematical operations
- Aggregate and categorize prime number characteristics
- Handle potential computational errors

## Detailed Specifications

### Generator Interface

```go
type PrimeGenerator interface {
    Generate(ctx context.Context, start, end int, concurrencyLevel int) <-chan PrimeResult
}

type PrimeResult struct {
    Number     int
    IsPrime    bool
    FindTime   time.Duration
    WorkerID   int
    Error      error
}
```

### Processor Interface

```go
type PrimeProcessor interface {
    Process(results <-chan PrimeResult) ProcessorResult
}

type ProcessorResult struct {
    TotalNumbersProcessed int
    PrimesFound           int
    LargestPrime          int
    SmallestPrime         int
    PrimeDistribution     map[int]int  // Primes grouped by digit count
    PrimePatterns         []PrimePattern
    ComputationTime       time.Duration
}

type PrimePattern struct {
    Type        string  // e.g., "Twin", "Sexy", "Circular"
    PrimeGroup  []int
}
```

## Computational Challenges

- Implement an efficient prime generation algorithm
- Use goroutines to parallelize prime finding
- Detect special prime number patterns
- Minimize computational overhead
- Handle large number ranges efficiently

## Evaluation Criteria

1. Concurrent Implementation (25 points)

   - Efficient use of goroutines
   - Optimal channel management
   - Prevention of goroutine leaks

2. Algorithm Efficiency (20 points)

   - Prime generation speed
   - Memory efficiency
   - Algorithmic complexity

3. Code Quality (25 points)

   - Clean, idiomatic Go code
   - Clear function and variable names
   - Comprehensive documentation
   - Adherence to Go best practices

4. Advanced Features (15 points)

   - Prime pattern detection
   - Sophisticated analysis capabilities
   - Flexible configuration

5. Error Handling (15 points)
   - Robust error management
   - Graceful degradation
   - Comprehensive error logging

## Computational Constraints

- Range: 2 to 1,000,000
- Concurrency Levels: 1 to 16 workers
- Memory Usage: < 500MB
- Computation Time: < 5 seconds

## Bonus Challenges

- Implement advanced prime pattern detection
- Create memory-efficient prime generation
- Develop sophisticated statistical analysis
- Optimize for both small and large number ranges

## Submission Requirements

- Implement `generator.go` and `processor.go`
- Include comprehensive unit tests
- Provide detailed comments explaining algorithmic choices
- Include performance optimization rationale
