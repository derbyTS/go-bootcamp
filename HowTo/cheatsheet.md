# Go Command Cheat Sheet

A practical reference for Go commands, especially **testing, debugging, race detection, synchronization bugs, memory/goroutine leaks, profiling, and performance analysis**.

---

## 1. Basic Go Commands

```bash
go version
```

Show installed Go version.

```bash
go env
```

Show Go environment variables.

```bash
go env GOPATH GOMOD GOROOT
```

Show specific environment values.

```bash
go mod init example.com/myapp
```

Create a new Go module.

```bash
go mod tidy
```

Add missing dependencies and remove unused ones.

```bash
go get package/path
```

Add or update a dependency.

```bash
go list -m all
```

List module dependencies.

```bash
go list ./...
```

List all packages in the module.

---

# 2. Run a Go Program

```bash
go run .
```

Run the current package.

```bash
go run main.go
```

Run a specific file.

```bash
go run ./cmd/server
```

Run a package.

Pass arguments:

```bash
go run . --port 8080
```

---

# 3. Build

```bash
go build
```

Compile the current package.

```bash
go build ./...
```

Build every package.

```bash
go build -o app .
```

Create an executable named `app`.

```bash
go install ./...
```

Build and install packages/binaries.

### Build for another OS/architecture

```bash
GOOS=linux GOARCH=amd64 go build -o app
```

Examples:

```bash
GOOS=linux GOARCH=arm64 go build
GOOS=darwin GOARCH=arm64 go build
GOOS=windows GOARCH=amd64 go build
```

---

# 4. Tests

Run all tests:

```bash
go test ./...
```

Run tests in the current package:

```bash
go test
```

Run a specific test:

```bash
go test -run TestMyFunction
```

Run tests matching a pattern:

```bash
go test -run 'TestUser.*'
```

Verbose output:

```bash
go test -v ./...
```

Run a test multiple times:

```bash
go test -count=10 ./...
```

Disable test caching:

```bash
go test -count=1 ./...
```

Set a timeout:

```bash
go test -timeout=30s ./...
```

---

# 5. Race Detection

## Detect Data Races

The most important command for concurrent Go code:

```bash
go test -race ./...
```

This instruments the program and attempts to detect **data races**.

Example:

```go
var counter int

go func() {
    counter++
}()

go func() {
    counter++
}()
```

Run:

```bash
go test -race
```

You may see:

```text
WARNING: DATA RACE
Read at ...
Write at ...
```

### Run a program with race detection

```bash
go run -race .
```

### Build with race detection

```bash
go build -race -o app .
```

Then:

```bash
./app
```

### Race detection notes

`-race` is relatively expensive.

Expect:

- higher CPU usage
- higher memory usage
- slower execution

Therefore:

```bash
go test -race ./...
```

is primarily a **testing/debugging command**, not normally something you deploy with.

### Useful workflow

```bash
go test ./...
go test -race ./...
```

Run both.

Normal tests tell you whether behavior is correct.

Race tests tell you whether concurrent access is unsafe.

---

# 6. `go vet`

`go vet` detects suspicious constructs that compile successfully but are likely bugs.

Run:

```bash
go vet ./...
```

Current package:

```bash
go vet
```

---

## Mutex Misuse & Invalid Synchronization

`go vet` can detect several synchronization mistakes.

For example, copying a mutex:

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

func foo(c Counter) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.n++
}
```

Passing `Counter` by value copies the mutex.

Prefer:

```go
func foo(c *Counter) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.n++
}
```

### `go vet` is particularly useful for:

- copying locks
- incorrect synchronization patterns
- unreachable code
- invalid printf arguments
- suspicious struct tags
- incorrect method signatures
- common API misuse

Run:

```bash
go vet ./...
```

---

# 7. Static Analysis

`go vet` is only one analyzer.

Run:

```bash
go tool vet
```

For a more extensive static-analysis setup, use tools such as:

```bash
staticcheck ./...
```

If Staticcheck is installed.

Typical workflow:

```bash
go vet ./...
staticcheck ./...
```

---

# 8. Format Code

Format one package/file:

```bash
gofmt -w .
```

Format a file:

```bash
gofmt -w main.go
```

Check formatting without modifying files:

```bash
gofmt -l .
```

Modern alternative:

```bash
go fmt ./...
```

---

# 9. Find Bugs + Test Everything

A very useful basic CI sequence:

```bash
go test ./...
go test -race ./...
go vet ./...
gofmt -l .
```

A stronger setup:

```bash
gofmt -l .
go vet ./...
staticcheck ./...
go test ./...
go test -race ./...
```

---

# 10. Goroutine Leaks

Go does not have a simple:

```bash
go test -leaks
```

flag.

A **goroutine leak** happens when a goroutine remains alive when it should have exited.

For example:

```go
func worker(ch <-chan int) {
    for {
        value := <-ch
        fmt.Println(value)
    }
}
```

If the channel is never closed and the goroutine is no longer needed, it can remain blocked forever.

---

## Check Number of Goroutines

Inside a program:

```go
runtime.NumGoroutine()
```

Example:

```go
fmt.Println("goroutines:", runtime.NumGoroutine())
```

If this continuously increases during repeated operations, investigate.

---

## Goroutine Leak Test

A useful pattern:

```go
func TestNoGoroutineLeak(t *testing.T) {
    before := runtime.NumGoroutine()

    // Run operation that creates goroutines.
    doSomething()

    time.Sleep(100 * time.Millisecond)

    after := runtime.NumGoroutine()

    if after > before {
        t.Fatalf(
            "possible goroutine leak: before=%d after=%d",
            before,
            after,
        )
    }
}
```

However, this can be flaky because the Go runtime itself creates goroutines.

For serious testing, use a dedicated leak-testing library such as:

```bash
go get go.uber.org/goleak
```

Then:

```go
func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}
```

Or:

```go
func TestSomething(t *testing.T) {
    defer goleak.VerifyNone(t)

    // test code
}
```

This is much better than simply comparing `runtime.NumGoroutine()`.

---

# 11. Check for Memory Leaks

Go has garbage collection, but you can still have **memory leaks**.

For example:

```go
var cache = make(map[string][]byte)
```

If entries are continuously added and never removed, GC cannot help because the map still references the objects.

---

## Heap Profiling

Run your application with the `pprof` HTTP server:

```go
import (
    _ "net/http/pprof"
    "net/http"
)

func main() {
    go http.ListenAndServe("localhost:6060", nil)

    // application...
}
```

Then inspect the heap:

```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

Interactive commands include:

```text
top
list functionName
web
```

---

## Heap Profile

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap
```

Open:

```text
http://localhost:8080
```

This gives you a visual profile.

---

# 12. Memory Leak Investigation

A very useful technique is comparing heap profiles.

Take a baseline:

```bash
curl http://localhost:6060/debug/pprof/heap > heap1
```

Exercise the application.

Take another:

```bash
curl http://localhost:6060/debug/pprof/heap > heap2
```

Then inspect:

```bash
go tool pprof heap2
```

You can also compare profiles:

```bash
go tool pprof -base heap1 heap2
```

If memory continuously grows after repeated operations and does not return after GC, investigate retained references.

---

# 13. Force Garbage Collection

For debugging:

```go
runtime.GC()
```

Then inspect memory.

This can help distinguish:

**memory that is temporarily allocated**

from

**memory that is still reachable and therefore cannot be collected.**

Do not normally call `runtime.GC()` manually in production code.

---

# 14. Escape Analysis

Find out what values escape to the heap:

```bash
go build -gcflags="-m" .
```

More detailed:

```bash
go build -gcflags="-m -m" .
```

Example:

```go
func foo() *int {
    x := 42
    return &x
}
```

The compiler may report:

```text
moved to heap: x
```

This is useful when investigating allocations.

---

# 15. Allocation Statistics

Benchmark with memory statistics:

```bash
go test -bench=. -benchmem
```

Example output:

```text
BenchmarkFoo-8    1000000    1200 ns/op    128 B/op    3 allocs/op
```

Meaning:

```text
1200 ns/op     time per operation
128 B/op       bytes allocated per operation
3 allocs/op    allocations per operation
```

---

# 16. Benchmarks

Run benchmarks:

```bash
go test -bench=.
```

Run all benchmarks:

```bash
go test -bench=. ./...
```

Run one:

```bash
go test -bench=BenchmarkFoo
```

Run for a fixed amount of time:

```bash
go test -bench=. -benchtime=10s
```

Run multiple times:

```bash
go test -bench=. -count=5
```

Include allocations:

```bash
go test -bench=. -benchmem
```

---

# 17. CPU Profiling

Benchmark CPU profile:

```bash
go test -bench=. -cpuprofile=cpu.out
```

Inspect:

```bash
go tool pprof cpu.out
```

Web UI:

```bash
go tool pprof -http=:8080 cpu.out
```

Useful interactive commands:

```text
top
top -cum
list FunctionName
web
```

---

# 18. Memory / Allocation Profiling

```bash
go test -bench=. -memprofile=mem.out
```

Inspect:

```bash
go tool pprof mem.out
```

Web UI:

```bash
go tool pprof -http=:8080 mem.out
```

---

# 19. Goroutine Profiling

With `net/http/pprof` enabled:

```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

Or:

```bash
go tool pprof -http=:8080 \
    http://localhost:6060/debug/pprof/goroutine
```

This can help identify:

- stuck goroutines
- blocked goroutines
- goroutine leaks
- unexpected goroutine creation

---

# 20. Block Profiling

Block profiling helps find goroutines waiting on synchronization.

Enable it:

```go
runtime.SetBlockProfileRate(1)
```

Then:

```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

Useful for investigating:

- channel blocking
- mutex contention
- synchronization delays

---

# 21. Mutex Profiling

Enable mutex profiling:

```go
runtime.SetMutexProfileFraction(1)
```

Then:

```bash
go tool pprof http://localhost:6060/debug/pprof/mutex
```

Useful when many goroutines are fighting over the same mutex.

---

# 22. Deadlocks

Go can detect certain deadlocks automatically.

Example:

```go
func main() {
    ch := make(chan int)

    <-ch
}
```

The runtime reports something similar to:

```text
fatal error: all goroutines are asleep - deadlock!
```

However, **not every deadlock is automatically detected**.

For example, two goroutines waiting on each other may require profiling/debugging to diagnose.

Useful tools:

```bash
go test -race ./...
```

and:

```bash
go tool pprof
```

---

# 23. Goroutine Dump

If the application exposes pprof:

```bash
curl http://localhost:6060/debug/pprof/goroutine?debug=2
```

This produces a detailed goroutine dump.

Look for states such as:

```text
chan receive
chan send
sync.Mutex.Lock
select
IO wait
```

A large number of goroutines stuck in the same state is often a clue.

---

# 24. Runtime Trace

Generate a trace during tests:

```bash
go test -trace=trace.out
```

View it:

```bash
go tool trace trace.out
```

For a benchmark:

```bash
go test -bench=. -trace=trace.out
```

Runtime tracing is useful for investigating:

- goroutine scheduling
- blocking
- network activity
- GC
- synchronization
- CPU utilization
- latency

---

# 25. Test Coverage

Basic coverage:

```bash
go test -cover ./...
```

Generate coverage profile:

```bash
go test -coverprofile=coverage.out ./...
```

View it:

```bash
go tool cover -func=coverage.out
```

HTML visualization:

```bash
go tool cover -html=coverage.out
```

---

# 26. Race + Coverage

You can combine useful checks:

```bash
go test -race -cover ./...
```

Or:

```bash
go test -race -coverprofile=coverage.out ./...
```

---

# 27. Detect Tests That Hang

Set a timeout:

```bash
go test -timeout=30s ./...
```

For a specific package:

```bash
go test -timeout=10s
```

This is especially useful for detecting:

- goroutines waiting forever
- deadlocks
- channels never being closed
- tests waiting for an event that never occurs

---

# 28. Run a Test Repeatedly

Useful for finding flaky concurrency bugs:

```bash
go test -run TestSomething -count=100
```

With race detection:

```bash
go test -race -run TestSomething -count=100
```

This is an excellent way to expose timing-dependent races.

---

# 29. Randomized / Stress Testing

Run tests repeatedly:

```bash
go test -count=100 ./...
```

Combine with race detection:

```bash
go test -race -count=100 ./...
```

For concurrency-heavy code, this can reveal bugs that a single run misses.

---

# 30. Package Test Output

Show package names:

```bash
go test -v ./...
```

Stop after first failure:

```bash
go test -failfast ./...
```

Run tests in parallel where applicable:

```bash
go test -parallel 8 ./...
```

---

# 31. Inspect Dependencies

```bash
go list -m all
```

Dependency graph:

```bash
go mod graph
```

Why is a module needed?

```bash
go mod why
```

Find package dependencies:

```bash
go list -deps ./...
```

---

# 32. Find Build Information

```bash
go version -m ./app
```

Shows build/module information embedded in a Go binary.

---

# 33. Compiler Diagnostics

Show compiler optimization decisions:

```bash
go build -gcflags="-m" .
```

More verbose:

```bash
go build -gcflags="-m -m" .
```

Useful for investigating:

- escape analysis
- inlining
- heap allocations
- compiler optimizations

---

# 34. Disable Optimizations for Debugging

Sometimes useful when debugging with Delve:

```bash
go build -gcflags="all=-N -l" .
```

Where:

```text
-N = disable optimizations
-l = disable inlining
```

Typical Delve workflow:

```bash
dlv debug --build-flags='-gcflags=all=-N -l'
```

---

# 35. Delve Debugger

Install:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Debug current package:

```bash
dlv debug
```

Debug a test:

```bash
dlv test
```

Attach to a process:

```bash
dlv attach <PID>
```

Debug a binary:

```bash
dlv exec ./app
```

Useful commands inside Delve:

```text
break main.main
continue
next
step
stepout
print variable
locals
args
goroutines
stack
bt
```

---

# 36. Check Goroutines in Delve

Inside Delve:

```text
goroutines
```

Inspect a goroutine:

```text
goroutine <id>
```

Stack:

```text
stack
```

This is useful when investigating deadlocks or goroutine leaks.

---

# 37. Useful Environment Variables

Show Go environment:

```bash
go env
```

Compiler environment:

```bash
go env GOOS GOARCH CGO_ENABLED
```

Use a specific compiler:

```bash
CC=clang go build
```

---

# 38. CGO

Check whether CGO is enabled:

```bash
go env CGO_ENABLED
```

Build without CGO:

```bash
CGO_ENABLED=0 go build
```

Build with CGO:

```bash
CGO_ENABLED=1 go build
```

---

# 39. Clean Build Cache

```bash
go clean
```

Clean build cache:

```bash
go clean -cache
```

Clean test cache:

```bash
go clean -testcache
```

Clean module cache:

```bash
go clean -modcache
```

Be careful with:

```bash
go clean -modcache
```

It removes downloaded modules and they will need to be downloaded again.

---

# 40. Useful "Everything" Check

For a normal Go project:

```bash
gofmt -l .
go vet ./...
go test ./...
go test -race ./...
```

If using Staticcheck:

```bash
staticcheck ./...
```

---

# 41. Recommended CI Check

A good basic CI pipeline:

```bash
gofmt -l .
go vet ./...
go test ./...
go test -race ./...
```

If formatting should fail the CI job:

```bash
test -z "$(gofmt -l .)"
go vet ./...
go test ./...
go test -race ./...
```

With Staticcheck:

```bash
test -z "$(gofmt -l .)"
go vet ./...
staticcheck ./...
go test ./...
go test -race ./...
```

---

# 42. Quick Diagnostic Table

| Problem               | Command / Tool                     |
| --------------------- | ---------------------------------- |
| Does it compile?      | `go build ./...`                   |
| Run program           | `go run .`                         |
| Run tests             | `go test ./...`                    |
| Verbose tests         | `go test -v ./...`                 |
| Flaky test            | `go test -count=100`               |
| Data race             | `go test -race ./...`              |
| Race in program       | `go run -race .`                   |
| Mutex misuse          | `go vet ./...`                     |
| Static analysis       | `staticcheck ./...`                |
| Test hangs            | `go test -timeout=30s`             |
| Goroutine leak        | `goleak` / pprof                   |
| Goroutine count       | `runtime.NumGoroutine()`           |
| Goroutine dump        | `/debug/pprof/goroutine?debug=2`   |
| Memory growth         | pprof heap                         |
| Heap profile          | `go tool pprof .../heap`           |
| CPU bottleneck        | CPU pprof                          |
| Allocation bottleneck | `-memprofile` / `-benchmem`        |
| Mutex contention      | mutex pprof                        |
| Blocking              | block pprof                        |
| Deadlock              | timeout + goroutine dump + Delve   |
| Scheduler behavior    | `go tool trace`                    |
| Heap allocations      | `go build -gcflags="-m"`           |
| Coverage              | `go test -cover`                   |
| Coverage HTML         | `go tool cover -html=coverage.out` |
| Debugger              | `dlv debug`                        |
| Debug tests           | `dlv test`                         |
| Dependencies          | `go mod graph`                     |
| Module cleanup        | `go mod tidy`                      |

---

# 43. The Important Distinction: Race vs Leak

These are different problems.

### Data race

Two or more goroutines access the same memory concurrently and at least one access is a write, without proper synchronization.

Use:

```bash
go test -race ./...
```

### Goroutine leak

A goroutine remains alive because it is blocked or otherwise never exits.

Use:

```text
goleak
pprof goroutine
Delve
runtime.NumGoroutine()
```

### Memory leak

Memory remains reachable and therefore cannot be garbage-collected.

Use:

```text
pprof heap
heap profile comparison
runtime.GC()
```

### Mutex misuse

Synchronization itself is incorrectly implemented.

Use:

```bash
go vet ./...
```

and:

```bash
go test -race ./...
```

### Mutex contention

The synchronization is correct, but goroutines spend too much time waiting for locks.

Use:

```text
mutex pprof
block pprof
runtime trace
```

---

# 44. My Go Debugging Checklist

When something is suspicious, run these first:

```bash
# 1. Formatting
gofmt -l .

# 2. Static correctness
go vet ./...

# 3. Tests
go test ./...

# 4. Race detector
go test -race ./...

# 5. Repeat concurrency tests
go test -race -count=100 ./...

# 6. Staticcheck
staticcheck ./...
```

If the problem is **memory**:

```text
pprof heap
go tool pprof
```

If the problem is **goroutines**:

```text
pprof goroutine
goleak
runtime.NumGoroutine()
```

If the problem is **deadlock/blocking**:

```text
pprof block
pprof mutex
go tool trace
dlv
```

If the problem is **performance**:

```bash
go test -bench=. -benchmem
```

then:

```bash
go test -bench=. -cpuprofile=cpu.out
go tool pprof -http=:8080 cpu.out
```

If the problem is **unexpected allocations**:

```bash
go build -gcflags="-m" .
```

and:

```bash
go test -bench=. -benchmem
```

---

# 45. One-Liner "Health Check"

For a project where you just want to quickly check for common problems:

```bash
gofmt -l . && \
go vet ./... && \
go test ./... && \
go test -race ./...
```

With Staticcheck:

```bash
test -z "$(gofmt -l .)" && \
go vet ./... && \
staticcheck ./... && \
go test ./... && \
go test -race ./...
```

**Recommended order:**

```text
format
  ↓
go vet
  ↓
staticcheck
  ↓
normal tests
  ↓
race detector
  ↓
profiling / leak investigation
```

The key point is that **`go test -race` does not detect memory leaks or goroutine leaks**. For leaks, use the appropriate profiler/tool: **heap pprof for retained memory, goroutine pprof or goleak for goroutine leaks, and mutex/block profiles for synchronization problems**.
