# Web Dev in Go — From Raw TCP to Structured CRUD

A hands-on learning repo that builds a Go web server from scratch, one layer at a time. Each phase adds complexity on top of the previous one — no skipping steps.

## What This Covers

How Go's `net/http` package works under the hood, starting from raw TCP concepts and ending with a concurrency-safe, structured multi-file JSON API. Every file was handwritten in a notebook with comments first, then coded and pushed.

---

## Phases

### `tcp-ip.txt` — TCP Foundation (Notes)

How TCP connections work at the OS level: `net.Listen` → `Accept` (blocks) → client `net.Dial` triggers the SYN/SYN-ACK/ACK handshake → `Accept` unblocks → **a goroutine per connection** → read/write bytes → `conn.Close()` on disconnect.

Understanding this is essential, because `net/http` is built directly on top of `net` — and the goroutine-per-connection detail turns out to matter a lot later (see Phase 4).

### `phase1.go` — Bare Minimum Server

A single struct implementing the `http.Handler` interface. One `ServeHTTP` method, one `ListenAndServe` call. Responds "Hello, World!" to every request on every path.

**Concepts:** Handler interface, ResponseWriter, ListenAndServe wiring.
**Limitation:** no way to distinguish routes or HTTP methods.

### `phase2.go` — Manual Routing (No Mux)

Nested switch statements matching `r.Method` (GET/POST) and `r.URL.Path` (`/`, `/index`). Different responses per route, all inside one `ServeHTTP` method.

**Concepts:** HTTP methods, URL path matching.
**Limitation:** nested switches become unreadable fast — this is the problem a router solves.

### `phase3.go` — ServeMux Routing

Replaced nested switches with `http.NewServeMux()`. Handlers split into named methods on the struct. Routes registered with `mux.HandleFunc("GET /users", ...)`.

**Concepts:** mux as a router, separating handler logic per route, `http.Server{}` config.
**Insight:** this is the same pattern frameworks like Gin and Chi wrap — they're ergonomics on top of this, not magic.

### `Phase4/` — Multi-File Structured CRUD API

```
Phase4/
├── main.go    # Server bootstrap and routing
├── api.go     # Handlers, storage, and concurrency control
└── user.go    # Data model with JSON tags
```

**Endpoints:**

| Method | Path     | Behaviour                                      |
|--------|----------|------------------------------------------------|
| GET    | `/users` | Returns all users as JSON                      |
| POST   | `/users` | Accepts a JSON body, validates, stores in memory |

**Features:**
- JSON encoding/decoding with `encoding/json` and struct tags
- Input validation (empty field checks)
- Duplicate detection
- Correct status codes — 201 Created, 400 Bad Request, 500 Internal Server Error
- In-memory slice as a stand-in for a real database
- **Concurrency-safe access to shared state** (see below)

---

## The Concurrency Problem (and why this repo bothers with it)

Most Go CRUD tutorials store data in a slice and stop there. That code has a data race, and it's rarely mentioned.

**Why:** `net/http` spawns a **goroutine per request**. Handlers do not run one at a time — two requests arriving close together execute simultaneously on different goroutines. So this:

```go
for _, user := range a.users { ... }   // read
a.users = append(a.users, u)           // write
```

...can be running in two goroutines at once. Both read the slice, both decide the user doesn't exist, both append. That's a lost update, or worse — memory corruption while the slice reallocates.

**The fix** — a `sync.Mutex` on the struct that owns the data:

```go
type Api struct {
    addr  string
    users []User
    mu    sync.Mutex
}
```

Locked around the read-modify-write, in the function that owns the data (not the caller — so every future caller is safe by default):

```go
a.mu.Lock()
defer a.mu.Unlock()
```

`defer` goes immediately after `Lock()` so the mutex is released on **every** exit path, including early returns on validation errors. Placing the `defer` after an early `return` leaves the mutex locked forever and deadlocks every subsequent request — a bug this repo hit and fixed.

**Verify it yourself:**

```bash
cd Phase4
go run -race .
```

Then fire concurrent requests:

```bash
for i in {1..50}; do
  curl -s -X POST http://localhost:8084/users \
    -H "Content-Type: application/json" \
    -d "{\"first_name\":\"user$i\",\"last_name\":\"test\"}" &
done
```

Remove the mutex and the race detector reports `WARNING: DATA RACE` on the users slice. Put it back and the output is clean.

---

## How to Run

Phases 1–3 are standalone single-file servers:

```bash
go run phase1.go
go run phase2.go
go run phase3.go
```

Phase 4 is a multi-file package:

```bash
cd Phase4
go run .
```

**Test with curl:**

```bash
# GET users
curl http://localhost:8084/users

# POST a new user
curl -X POST http://localhost:8084/users \
  -H "Content-Type: application/json" \
  -d '{"first_name": "Hashim", "last_name": "Syed"}'
```

---

## Learning Approach

Each phase was written by hand in a notebook, with comments explaining every line, before being typed into an IDE. No copy-pasting from tutorials, no vibe-coding. If a line is in this repo, it's because I understood why it needed to be there.

---

## Known Limitations

Deliberate — this is a learning progression, not a production service:

- In-memory storage only (data is lost on restart)
- No PUT or DELETE yet
- No middleware (logging, auth, recovery)
- No tests
- The GET handler holds the mutex while encoding to the network; copying the slice under lock and encoding outside it would be the production pattern

## What's Next

- PUT and DELETE to complete CRUD
- Middleware for logging and panic recovery
- A real database (PostgreSQL) behind a repository layer
- Tests, including a concurrent test run with `go test -race`

## Related Repos

- [go-systems-from-scratch](https://github.com/Hashim-777x/go-systems-from-scratch) — raw TCP server, HTTP from scratch, key-value store, Redis clone
- [DSA-in-GO](https://github.com/Hashim-777x/DSA-in-GO) — Neetcode 150 in Go

## Stack

Go — standard library only, no frameworks.
