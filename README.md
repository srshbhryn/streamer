# streamer

A lightweight WebSocket pub/sub message broker written in Go. Clients connect over
WebSocket, subscribe to topics, and receive every message published to those topics.
Messages are read from a pluggable *source* and fanned out to all subscribed clients.

## How it works

```
                 +------------------+        +------------------+
   source ─────► |   streamer.Run   | ─────► |  WebSocket       | ─────► subscribed
 (a Reader)      |  (fan-out loop)  |        |  clients         |        clients
                 +------------------+        +------------------+
                          ▲                           │
                          └───── subscribe / ─────────┘
                                 unsubscribe
```

- A **source** is anything implementing `Read() (string, error)`. The bundled
  [`pullers.MockPuller`](lib/pullers/mock.go) emits a message every 300ms, alternating
  between topics `a` and `b`.
- Each connected client is a `ReaderWriter` (read commands, write messages). The
  WebSocket implementation lives in [`lib/websocket/websocket.go`](lib/websocket/websocket.go).
- The [streamer](lib/streamer/streamer.go) keeps a registry of clients and the topics
  they subscribe to, then forwards each incoming source message to matching clients.

### Message format

Both source messages and client commands are comma-separated strings.

- **Source / outbound messages:** `topic,payload` — e.g. `a,42`. The first field is the
  topic used for routing; the full string is delivered to subscribers.
- **Client commands:** sent by a client over its WebSocket connection:
  - `subscribe,<topic>` — start receiving messages for `<topic>`
  - `unsubscribe,<topic>` — stop receiving messages for `<topic>`

Only configured topics are accepted. Valid topics are currently hard-coded in
[`lib/streamer/config.go`](lib/streamer/config.go) as `a`, `b`, and `c`.

## HTTP / WebSocket API

The web server is built on [Gin](https://github.com/gin-gonic/gin) and listens on
`:8080` ([`lib/webserver/webserver.go`](lib/webserver/webserver.go)).

| Method | Path         | Description                                              |
|--------|--------------|----------------------------------------------------------|
| GET    | `/ws`        | Upgrade to a WebSocket connection and register a client. |
| GET    | `/getTopics` | Return the list of available topics as JSON.             |
| GET    | `/`          | Demo HTML page (debug mode only).                        |
| GET    | `/static/*`  | Static assets (debug mode only).                         |

## Getting started

### Prerequisites

- Go 1.18 or newer

### Run

```sh
go run ./bin/streamer
```

This starts the mock source and the web server on `http://localhost:8080`.

### Build

```sh
go build -o bin/streamer/streamer ./bin/streamer
./bin/streamer/streamer
```

### Try it

Connect with any WebSocket client (e.g. [`websocat`](https://github.com/vi/websocat)):

```sh
websocat ws://localhost:8080/ws
# then type:
subscribe,a
```

You should start receiving `a,<counter>` messages roughly every 600ms.

## Project layout

```
bin/streamer/        Entry point (main)
lib/streamer/        Client registry, topic routing, fan-out loop
lib/websocket/       WebSocket connection handler (ReaderWriter implementation)
lib/webserver/       Gin HTTP server and routes
lib/pullers/         Message sources; MockPuller is the bundled example
```

## Extending

To plug in a real message source, implement the `Reader` interface and pass it to
`streamer.Init`:

```go
type Reader interface {
    Read() (string, error)
}
```

Swap `pullers.CreateMock()` in [`bin/streamer/streamer.go`](bin/streamer/streamer.go)
for your own source.

## Status

Early-stage / v0. Configuration (topics, port, debug mode) is currently hard-coded;
several values are intended to move to environment variables (see the commented-out
`TOPICS` handling in [`lib/streamer/config.go`](lib/streamer/config.go)).
</content>
</invoke>
