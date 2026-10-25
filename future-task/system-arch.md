# Go Systems Architecture & Performance Engineering: A Study Guide

This guide covers high-performance backend patterns, memory management, and database optimization strategies tailored for **Go (Golang)**.

---

## 1. The Core Disciplines

When studying how to optimize backend systems and database communication, you are working across three main domains:

- **Systems Engineering & Performance Engineering:** Minimizing CPU cycles, optimizing memory footprints, and managing Garbage Collection (GC) pressure.
- **Distributed Systems & Architecture:** Managing network I/O, latency vs. throughput, connection pooling, and payload efficiency.
- **Database Internals & Data Engineering:** Optimizing query execution, reducing network round-trips via batching, and handling raw data streams.

---

## 2. Zero-Copy & Raw Byte Streaming (`json.RawMessage`)

### The Concept

Traditional web servers scan database rows into heavy Go structs, unmarshal them, and then re-marshal them into JSON bytes to send over HTTP. This wastes CPU cycles and spikes memory allocations.

### How it Works in Go

If your database stores pre-formatted data (like PostgreSQL `JSONB`), you can scan it directly into a byte slice or `json.RawMessage` and pipe it straight to the HTTP response writer without parsing.

```go
package main

import (
    "database/sql"
    "encoding/json"
    "net/http"
)

func handleUser(db *sql.DB, w http.ResponseWriter) {
    var rawJSON json.RawMessage

    // Scan directly from DB into raw bytes, bypassing struct unmarshaling
    err := db.QueryRow("SELECT data FROM users WHERE id = 1").Scan(&rawJSON)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(rawJSON) // Stream straight to client
}
```

### Trade-offs: When NOT to do it

- **No Transformation:** You cannot easily redact sensitive fields, modify fields, or change snake_case to camelCase.
- **No Validation:** Type safety is bypassed; corrupted data in the DB will break the client directly.
- **DX Loss:** Code readability and maintainability drop because you lose strongly typed structs (`user.Email`).

---

## 3. High-Throughput Batching & Bulk Processing

### The Concept

Doing operations one-by-one introduces massive latency overhead due to network round-trips, database lock contention, and frequent memory allocations.

### Why Batching Wins

1. **Eliminates Network Latency:**
   - 1,000 individual queries $\times$ 5ms = **5,000ms** total latency.
   - 1 bulk query pulling 1,000 rows = **5ms** total latency.
2. **Reduces GC Pressure in Go:**
   Allocating thousands of tiny objects in a loop triggers the Go Garbage Collector frequently. Pre-allocating slices with exact capacities (`make([]T, 0, cap)`) allows memory reuse and keeps latency smooth.

### Example: Pre-allocating Slices for Performance

```go
// Pre-allocate slice capacity to avoid underlying array re-allocations during growth
users := make([]User, 0, 1000)

rows, _ := db.Query("SELECT id, name FROM users LIMIT 1000")
defer rows.Close()

for rows.Next() {
    var u User
    rows.Scan(&u.ID, &u.Name)
    users = append(users, u)
}
```

---

## 4. Key Go Optimization Keywords to Research

- **Escape Analysis:** How the Go compiler decides whether a variable lives on the stack or the heap.
- **`sync.Pool`:** Reusing objects to minimize GC overhead in high-throughput loops.
- **Buffer Pools (`bytes.Buffer` / `bufio`):** Reducing memory allocations when writing heavy network payloads.
- **Database Connection Pooling (`database/sql` settings):** Tuning `MaxOpenConns`, `MaxIdleConns`, and `ConnMaxLifetime`.

---

_Keep this guide handy when designing high-concurrency microservices or tuning API bottlenecks in Go!_
