# Portfolio Project: GoPulse Log Ingestion Pipeline

A high-performance CLI tool and backend service that streams raw application log files (or accepts HTTP POST payloads), parses them concurrently, aggregates metrics in memory, and exposes structured results via gRPC/REST.

---

## 1. Core Technical Goals & Concepts Covered

### The GMP Scheduler (Go Runtime Architecture)

- **Bounded Concurrency:** Build a fixed-size **Worker Pool** reading from buffered channels rather than spawning unbounded goroutines ($G$).
- **Processor Tuning:** Experiment with `runtime.GOMAXPROCS(n)` to observe how the Go runtime assigns $P$ (Logical Processors) to $M$ (OS Threads).
- **System Call Handling:** Observe how non-blocking I/O vs. blocking file system operations cause $M$ to detach from $P$ to maintain throughput across remaining goroutines.

### Memory Alignment, Slices, and Maps

- **Slice Pre-allocation:** Avoid continuous heap allocations during log parsing by pre-allocating slice capacities via `make([]T, len, cap)`.
- **Zero-Copy Slicing vs. Leaks:** Master when slicing a byte array holds reference to underlying memory versus when to explicitly copy buffers with `copy()`.
- **Map Internals:** Implement `map[string]int` for aggregate statistics (e.g., status code hits). Track lock contention using `sync.Mutex` vs. `sync.Map`, and handle the fact that Go maps do not release allocated bucket memory upon key deletion.
- **Garbage Collection (GC) Optimization:** Implement `sync.Pool` to reuse allocated byte buffers during log string parsing, keeping memory allocations near zero during high load.

### Idiomatic Go Patterns

- **Implicit Interfaces:** Avoid deep class hierarchies. Keep interfaces small (1-2 methods) and accept interfaces, return structs.
- **Graceful Shutdown:** Intercept `SIGINT`/`SIGTERM` using `os/signal` and drain channel buffers using `context.WithTimeout`.
- **Structured Logging:** Use `log/slog` for structured JSON output and expose a Prometheus metrics endpoint (`/metrics`).

---

## 2. Recommended Directory Structure

```text
gopulse/
├── cmd/
│   └── gopulse/
│       └── main.go          # Application entrypoint, CLI flags, OS signal handling
├── pkg/
│   ├── aggregator/          # In-memory metrics, map tracking, mutexes
│   ├── config/              # Environment variable loading
│   ├── parser/              # Fast log parsing using zero-allocation byte slicing
│   ├── pool/                # Worker pool and job distribution logic
│   └── server/              # gRPC and HTTP endpoint handlers
├── .dockerignore            # Excludes git, local binaries, and temp files
├── Dockerfile               # Multi-stage static binary container build
├── docker-compose.yml       # Local environment with GoPulse + Redis/Postgres
├── go.mod
└── go.sum
```

## System Architecture Diagram

```text
[ Incoming Log Stream / HTTP POST ]
                 │
                 ▼
      ┌─────────────────────┐
      │  Main Ingestion Ch  │  (Buffered Channel)
      └──────────┬──────────┘
                 │
                 ▼
      ┌─────────────────────┐
      │     Worker Pool     │  (GOMAXPROCS-tuned Goroutines)
      └──────────┬──────────┘
                 │
   ┌─────────────┴─────────────┐
   ▼                           ▼
┌──────────────────┐  ┌──────────────────┐
│  sync.Pool       │  │  Parser Routine  │
│  (Buffer Reuse)  │  │  (Zero-Copy)     │
└────────┬─────────┘  └────────┬─────────┘
         │                     │
         └──────────┬──────────┘
                    │
                    ▼
      ┌─────────────────────┐
      │  Aggregator Engine  │  (Mutex-Protected Map)
      └──────────┬──────────┘
                 │
      ┌──────────┴──────────┐
      ▼                     ▼
┌───────────┐         ┌───────────┐
│ REST API  │         │gRPC Stream│
└───────────┘         └───────────┘
```

### 3. Benchmark & Portfolio Deliverables

To present this effectively on GitHub or during interviews, include these artifacts in your repository:

- **Standard Go Benchmarks (`go test -bench=. -benchmem`)**
  Report allocations per operation (`B/op`) and total allocs (`allocs/op`).
- **Architecture Visualizations**
  Add a diagram showing how incoming logs pass through Channels $\rightarrow$ Worker Pool $\rightarrow$ sync.Pool Buffer $\rightarrow$ Aggregator Map.
- **Local Deployment**
  A single `docker-compose up` command that launches the Go service alongside a secondary storage/cache dependency (e.g., PostgreSQL or Redis).
