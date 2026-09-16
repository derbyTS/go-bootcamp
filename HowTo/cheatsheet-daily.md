# Go Daily-Use Debugging Cheat Sheet

The ~30 Go commands worth memorizing for **development, testing, races, leaks, performance, and debugging**.

---

## 🏃 Run & Build

### 1. Run current project

```bash
go run .
```

### 2. Build everything

```bash
go build ./...
```

### 3. Build an executable

```bash
go build -o app .
```

### 4. Install a Go tool

```bash
go install <package>@latest
```

Example:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

---

# 🧪 Testing

### 5. Run all tests

```bash
go test ./...
```

### 6. Verbose tests

```bash
go test -v ./...
```

### 7. Run one test

```bash
go test -run TestName
```

### 8. Run tests repeatedly

Great for flaky/concurrency bugs:

```bash
go test -run TestName -count=100
```

### 9. Disable test cache

```bash
go test -count=1 ./...
```

### 10. Put a limit on hanging tests

```bash
go test -timeout=30s ./...
```

---

# 🏎️ Race Detection

### 11. Detect data races

**Memorize this one.**

```bash
go test -race ./...
```

### 12. Run program with race detector

```bash
go run -race .
```

### 13. Build with race detector

```bash
go build -race -o app .
```

### Race workflow

```bash
go test ./...
go test -race ./...
```

If concurrency is suspicious:

```bash
go test -race -count=100 ./...
```

---

# 🔍 Static Analysis

### 14. Check suspicious Go code

```bash
go vet ./...
```

Useful for things such as **mutex copying/misuse**, incorrect printf calls, suspicious constructs, etc.

### 15. Deeper static analysis

If Staticcheck is installed:

```bash
staticcheck ./...
```

---

# 🧹 Formatting

### 16. Format code

```bash
gofmt -w .
```

### 17. Check which files need formatting

```bash
gofmt -l .
```

---

# 🧵 Goroutines & Leaks

### 18. Check current goroutine count

Inside Go:

```go
runtime.NumGoroutine()
```

Useful when you suspect goroutines are accumulating.

### 19. Goroutine profile

With `net/http/pprof` running:

```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

### 20. Detailed goroutine dump

```bash
curl http://localhost:6060/debug/pprof/goroutine?debug=2
```

Look for many goroutines stuck in:

```text
chan receive
chan send
sync.Mutex.Lock
select
IO wait
```

### 21. Detect goroutine leaks in tests

Install:

```bash
go get go.uber.org/goleak
```

Then use `goleak.VerifyNone(t)` or `goleak.VerifyTestMain(...)`.

**Important:** `go test -race` does **not** detect goroutine leaks.

---

# 🧠 Memory / Allocation

### 22. Heap profile

With pprof running:

```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

### 23. Open heap profile in browser

```bash
go tool pprof -http=:8080 \
  http://localhost:6060/debug/pprof/heap
```

### 24. Force garbage collection while debugging

Inside Go:

```go
runtime.GC()
```

Useful for determining whether memory is actually collectible.

### 25. Check escape analysis

```bash
go build -gcflags="-m" .
```

More detail:

```bash
go build -gcflags="-m -m" .
```

Look for:

```text
moved to heap
```

---

# ⚡ Performance

### 26. Run benchmarks

```bash
go test -bench=.
```

### 27. Benchmark with allocation information

```bash
go test -bench=. -benchmem
```

Look at:

```text
ns/op
B/op
allocs/op
```

### 28. CPU profile

```bash
go test -bench=. -cpuprofile=cpu.out
```

Then:

```bash
go tool pprof -http=:8080 cpu.out
```

### 29. Memory profile

```bash
go test -bench=. -memprofile=mem.out
```

Then:

```bash
go tool pprof -http=:8080 mem.out
```

---

# 🐛 Debugging

### 30. Start Delve

```bash
dlv debug
```

Debug tests:

```bash
dlv test
```

Attach to a running process:

```bash
dlv attach <PID>
```

Useful Delve commands:

```text
break main.main
continue
next
step
stepout
print variable
locals
goroutines
stack
```

---

# 🔒 Mutex / Blocking Problems

These aren't commands you need every day, but memorize the **pprof endpoints**.

### Mutex contention

Enable:

```go
runtime.SetMutexProfileFraction(1)
```

Then:

```bash
go tool pprof http://localhost:6060/debug/pprof/mutex
```

### Blocking

Enable:

```go
runtime.SetBlockProfileRate(1)
```

Then:

```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

Useful for finding:

- mutex contention
- channel blocking
- synchronization bottlenecks

---

# 🕵️ Runtime Trace

### Trace a test

```bash
go test -trace=trace.out
```

### Open trace

```bash
go tool trace trace.out
```

Use this when you need to understand:

- goroutine scheduling
- blocking
- GC
- synchronization
- latency

---

# 📊 Coverage

### Coverage percentage

```bash
go test -cover ./...
```

### Generate coverage profile

```bash
go test -coverprofile=coverage.out ./...
```

### Open coverage in browser

```bash
go tool cover -html=coverage.out
```

---

# 🧰 Modules

### Update/clean dependencies

```bash
go mod tidy
```

### See dependencies

```bash
go list -m all
```

### Why is a dependency needed?

```bash
go mod why
```

---

# 🚨 The Commands to Memorize

If you remember nothing else, remember these:

```bash
# Run
go run .

# Build
go build ./...

# Test
go test ./...

# Race detector
go test -race ./...

# Repeat concurrency tests
go test -race -count=100 ./...

# Static analysis
go vet ./...
staticcheck ./...

# Format
gofmt -w .

# Benchmark
go test -bench=. -benchmem

# CPU profile
go test -bench=. -cpuprofile=cpu.out

# Memory profile
go test -bench=. -memprofile=mem.out

# Heap
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap

# Goroutines
curl http://localhost:6060/debug/pprof/goroutine?debug=2

# Trace
go test -trace=trace.out
go tool trace trace.out

# Debugger
dlv debug
```

---

# 🧠 Which Tool Should I Use?

| Problem                    | First thing to try                   |
| -------------------------- | ------------------------------------ |
| Code doesn't compile       | `go build ./...`                     |
| Test failing               | `go test -v ./...`                   |
| Flaky test                 | `go test -count=100`                 |
| **Data race**              | **`go test -race ./...`**            |
| Race happens rarely        | **`go test -race -count=100 ./...`** |
| Mutex copied/misused       | **`go vet ./...`**                   |
| General suspicious code    | `go vet` / `staticcheck`             |
| Goroutines keep increasing | `goleak` / goroutine pprof           |
| Memory keeps increasing    | **heap pprof**                       |
| Too many allocations       | `go test -bench=. -benchmem`         |
| CPU too slow               | CPU pprof                            |
| Mutex contention           | mutex pprof                          |
| Channel blocking           | block pprof                          |
| Deadlock                   | goroutine dump + Delve               |
| Weird scheduler behavior   | `go tool trace`                      |
| Unexpected heap allocation | `go build -gcflags="-m"`             |
| Need code coverage         | `go test -cover`                     |
| Need interactive debugging | `dlv debug`                          |

---

# 🔥 My Practical Debugging Sequence

When a Go project starts behaving strangely:

```bash
gofmt -l .
go vet ./...
go test ./...
go test -race ./...
```

If it looks like a **concurrency bug**:

```bash
go test -race -count=100 ./...
```

If it looks like a **goroutine leak**:

```text
goleak
    ↓
goroutine pprof
    ↓
goroutine dump
    ↓
Delve
```

If it looks like a **memory leak**:

```text
heap pprof
    ↓
compare heap usage
    ↓
runtime.GC()
    ↓
escape analysis
```

If it looks like a **performance problem**:

```text
benchmark
    ↓
-benchmem
    ↓
CPU profile
    ↓
memory profile
    ↓
trace
```

---

## ⭐ The 5 Most Important

If you want the absolute minimum:

```bash
go test ./...
go test -race ./...
go vet ./...
go test -bench=. -benchmem
go tool pprof -http=:8080 <profile>
```

And for actual debugging:

```bash
dlv debug
```

**`-race` = data races**
**`go vet` = suspicious/invalid Go code, including synchronization mistakes**
**`goleak` = goroutine leaks**
**`pprof heap` = memory/retained allocations**
**`pprof cpu` = CPU bottlenecks**
**`trace` = runtime/scheduler behavior**
**`dlv` = step through the actual program**
