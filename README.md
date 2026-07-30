# Web Dev in Go — From Raw TCP to Structured CRUD

A hands-on learning repo that builds a Go web server from scratch, one layer at a time. Each phase adds complexity on top of the previous one — no skipping steps.

## What This Covers

How Go's `net/http` package works under the hood, starting from raw TCP concepts and ending with a structured multi-file JSON API. Every file was handwritten in a notebook first, then coded and pushed.

## Phases

### `tcp-ip.txt` — TCP Foundation (Notes)
How TCP connections work at the OS level: `net.Listen` → `Accept` → goroutine per connection → read/write bytes → close. Understanding this is essential because `net/http` is built directly on top of `net`.

### `phase1.go` — Bare Minimum Server
A single struct implementing the `http.Handler` interface. One `ServeHTTP` method, one `ListenAndServe` call. Responds "Hello, World!" to every request on every path.

**Concepts:** Handler interface, ResponseWriter, ListenAndServe wiring.

### `phase2.go` — Manual Routing (No Mux)
Nested switch statements matching `r.Method` (GET/POST) and `r.URL.Path` (/, /index). Different responses per route, all inside one `ServeHTTP` method.

**Concepts:** HTTP methods, URL path matching, why manual routing breaks down at scale.

### `phase3.go` — ServeMux Routing
Replaced nested switches with `http.NewServeMux()`. Handlers split into named methods on the struct. Clean route registration with `mux.HandleFunc`.

**Concepts:** Mux as a router, separating handler logic, how frameworks like Gin/Chi wrap this same pattern.

### `Phase4/` — Multi-File Structured CRUD API

```
Phase4/
├── main.go    # Server bootstrap and routing
├── Api.go     # Handlers and business logic
└── User.go    # Data model with JSON tags
```

A proper JSON API with:
- `GET /users` — returns all users as JSON
- `POST /users` — accepts JSON body, validates, stores in memory
- Input validation (empty field checks)
- Duplicate detection
- Correct HTTP status codes (200, 201, 400, 500)
- JSON encoding/decoding with `encoding/json`
- Error handling at every layer

**Concepts:** JSON encode/decode, struct tags, Content-Type headers, separation of concerns, service layer pattern.

## How to Run

```bash
# Any phase (standalone files)
go run phase1.go

# Phase 4 (multi-file)
cd Phase4
go run .
```

Test with curl:
```bash
# GET users
curl http://localhost:8084/users

# POST a new user
curl -X POST http://localhost:8084/users \
  -H "Content-Type: application/json" \
  -d '{"first_name": "Hashim", "last_name": "Syed"}'
```

## Learning Approach

Each phase was written by hand in a notebook with comments explaining every line before being typed into an IDE. No copy-pasting from tutorials, no vibe-coding.

## What's Next

- [go-systems-from-scratch](https://github.com/Hashim-777x/go-systems-from-scratch) — raw TCP server, HTTP from scratch, key-value store, Redis clone
- DSA in Go — Neetcode 150

## Stack

Go (standard library only, no frameworks)
