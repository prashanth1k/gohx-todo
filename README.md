# Simple todo demo with Go, Go/HTTP, and HTMX

There's nothing to look here.

- Only POST and GET handled for now
- No database, just a shared variable
- AlpineJS since checkboxes suck
- Classless CSS since I'm lazy

Couple of bad practices -

1. Sharing of state/ Todos in Go
1. No flexible routing

## Run

Install /refresh.

```bash
go mod tidy
```

Run the server.

```bash
go run .
```
