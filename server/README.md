# Freelands — Server Implementation Roadmap

**Version:** 1.0  
**Status:** Initial implementation  
**Technology:** Go  
**Scope:** First working multiplayer server

---

## 1. Objective

Build the first functional version of the Freelands game server.

The server must be able to:

- Start and listen for incoming connections.
- Accept multiple simultaneous players.
- Create and maintain player sessions.
- Place connected players into a shared game world.
- Receive and validate player movement input.
- Update player positions authoritatively.
- Broadcast world updates to connected clients.
- Handle player disconnections cleanly.

The goal is to establish a working multiplayer foundation that the SvelteKit + Three.js client can connect to.

This is not yet a complete MMORPG server. It is the foundation upon which gameplay systems will be built.

## 2. Initial Technical Decisions

Keep the initial implementation simple. Avoid introducing infrastructure before it is necessary.

| Area | Initial approach |
|---|---|
| Language | Go |
| HTTP server | Go standard library |
| Real-time communication | WebSockets |
| Message format | JSON |
| World simulation | Server-authoritative |
| Game loop | Fixed update interval |
| Player identification | Server-generated ID |
| Player persistence | In-memory only |
| Database | Not required initially |
| Authentication | Not required initially |
| Deployment | Local development |

Use a WebSocket library such as `github.com/coder/websocket` rather than implementing the WebSocket protocol yourself.

### Server authority

The client sends intentions, not authoritative game state.

For example:

- Client: "I want to move east."
- Server: validates the input and updates the player's position.
- Server: broadcasts the resulting position to other players.

The client must never be trusted to decide its own authoritative position, movement speed or other gameplay outcomes.

## 3. Implementation Roadmap

### Milestone 1 — Basic HTTP Server

**Goal:** Get the Go server running and expose basic HTTP endpoints.

Tasks:

1. Create the server entry point.
2. Configure the HTTP server.
3. Add a health-check endpoint.
4. Add basic startup and shutdown handling.
5. Configure the listening address and port.

Suggested structure:

```text
server/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── httpapi/
│       └── health.go
├── go.mod
└── README.md
```

**Acceptance criteria:**

- The server starts successfully.
- The server listens on a configurable port.
- `GET /health` returns HTTP 200.
- The endpoint returns a simple JSON response.
- The server shuts down cleanly when interrupted.

Example response:

```json
{
  "status": "ok"
}
```

### Milestone 2 — Player Connections

**Goal:** Allow multiple clients to connect to the server.

Tasks:

1. Add a WebSocket endpoint, for example `/ws`.
2. Accept incoming WebSocket connections.
3. Generate a unique server-side player ID for each connection.
4. Register connected players in a central player manager.
5. Handle incoming and outgoing messages.
6. Detect disconnected clients.
7. Remove disconnected players from the active world.

Suggested additions:

```text
internal/
├── httpapi/
│   ├── health.go
│   └── websocket.go
└── game/
    ├── player.go
    └── world.go
```

**Acceptance criteria:**

- A client can establish a WebSocket connection.
- Every connection receives a unique player ID.
- Multiple clients can connect simultaneously.
- The server tracks connected players.
- Disconnected players are removed without crashing the server.
- One client's disconnection does not affect other players.

### Milestone 3 — Shared World and Player Spawning

**Goal:** Introduce a minimal shared world in which connected players exist.

Tasks:

1. Create a world manager.
2. Define the initial player state.
3. Assign a spawn position when a player connects.
4. Register the player in the world.
5. Notify the connecting client of its initial state.
6. Make existing player information available to newly connected players.

Initial player state:

```go
type Player struct {
    ID       string
    X        float64
    Y        float64
    Z        float64
    Rotation float64
}
```

The coordinate system and world dimensions should be consistent with the client. The initial implementation does not need actual terrain, collision or character models.

**Acceptance criteria:**

- Every connected player receives a spawn position.
- The server maintains a shared list of active players.
- A newly connected player receives the current player list.
- Existing players can be identified by their server-generated IDs.
- Player state is not stored independently by each connection.

### Milestone 4 — Authoritative Movement

**Goal:** Allow players to move around the shared world.

Tasks:

1. Define a movement input message.
2. Receive movement input from connected clients.
3. Validate the input.
4. Update player positions on the server.
5. Run world updates at a fixed interval.
6. Broadcast updated player states to connected clients.

Start with simple movement. Do not implement acceleration curves, terrain navigation, combat movement or advanced physics yet.

Example client message:

```json
{
  "type": "move",
  "payload": {
    "x": 1,
    "z": 0
  }
}
```

Example server message:

```json
{
  "type": "world_state",
  "payload": {
    "players": [
      {
        "id": "player-123",
        "x": 10.5,
        "y": 0,
        "z": 20.0,
        "rotation": 1.57
      }
    ]
  }
}
```

These examples describe the initial protocol, not a finalized network specification.

**Important implementation rules:**

- The server controls movement speed.
- Movement input must be validated and normalized.
- Clients cannot submit arbitrary authoritative positions.
- Invalid or excessively large inputs must be rejected or safely handled.
- One player's movement must not block the processing of other players.
- The server must not allow one slow client to stall world updates.

**Acceptance criteria:**

- A client can request movement.
- The server calculates the resulting position.
- Movement is constrained by server-defined speed.
- All connected clients receive updated player positions.
- Two clients can move simultaneously and see each other's movement.
- Disconnecting a player removes that player from subsequent world updates.

### Milestone 5 — Connect the Three.js Client

**Goal:** Verify the server through the actual Freelands client.

Tasks:

1. Add a WebSocket connection manager to the SvelteKit client.
2. Connect to the Go server.
3. Receive the initial player state.
4. Render the local player.
5. Render other connected players.
6. Send movement input from the client.
7. Update rendered positions using server messages.
8. Handle connection loss and reconnection attempts.

The client can initially use simple colored shapes instead of finished character models.

**Acceptance criteria:**

- The SvelteKit client connects to the Go server.
- The local player appears in the world.
- A second browser session can connect as another player.
- Both players can see each other.
- Moving one player updates its position for the other player.
- Disconnecting a client removes its representation from the world.

### Milestone 6 — Basic Stability and Validation

**Goal:** Make the first multiplayer prototype reliable enough for continued development.

Tasks:

1. Add basic structured logging.
2. Handle malformed messages safely.
3. Limit incoming message sizes.
4. Validate message types and payloads.
5. Prevent duplicate player registration.
6. Ensure concurrent access to shared world state is safe.
7. Add unit tests for movement and player management.
8. Add an integration test for connecting multiple clients.
9. Test server shutdown and disconnected clients.

**Acceptance criteria:**

- Malformed messages do not crash the server.
- Invalid movement inputs cannot produce arbitrary movement.
- Concurrent player updates do not cause data races.
- Players are cleaned up correctly after disconnection.
- Basic server tests pass.
- The server can run with multiple connected clients without blocking the world update loop.

---

## 4. Suggested Final Project Structure

The structure below is a starting point. Add files when they become necessary rather than creating empty packages for systems that do not exist yet.

```text
server/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── httpapi/
│   │   ├── health.go
│   │   └── websocket.go
│   └── game/
│       ├── player.go
│       ├── world.go
│       └── movement.go
├── go.mod
└── README.md
```

Responsibilities:

- `cmd/server`: Application entry point and startup.
- `config`: Server configuration.
- `httpapi`: HTTP endpoints and WebSocket connection handling.
- `game`: Player state, shared world and authoritative gameplay updates.

Keep the architecture straightforward. A small number of clear components is preferable to a large framework of managers, interfaces and abstractions.

## 5. Initial Network Protocol

Use a simple message envelope:

```json
{
  "type": "message_type",
  "payload": {}
}
```

Initial message types:

| Direction | Message | Purpose |
|---|---|---|
| Client → Server | `move` | Request player movement |
| Server → Client | `welcome` | Confirm connection and provide player ID |
| Server → Client | `world_state` | Provide current player states |
| Server → Client | `player_joined` | Notify clients about a new player |
| Server → Client | `player_left` | Notify clients about a disconnected player |
| Server → Client | `error` | Report an invalid request |

Keep the protocol small. Additional messages can be introduced as new gameplay systems are implemented.

## 6. Deliberately Out of Scope

Do not implement these systems during the first server version:

- Account registration and authentication.
- PostgreSQL or Redis.
- Character creation and character persistence.
- Inventory and equipment.
- Combat and damage calculations.
- Skills, attributes and progression.
- Quests and NPCs.
- Crafting and gathering.
- Trading and economy.
- Housing.
- Guilds and parties.
- World persistence.
- Advanced anti-cheat systems.
- Multiple maps, instances or server regions.

These systems will eventually depend on the multiplayer foundation. Implementing them now would increase complexity without helping validate the first playable server.

## 7. Definition of Done

The first working server version is complete when:

- [ ] The Go server starts successfully.
- [ ] The health endpoint responds correctly.
- [ ] Multiple clients can connect simultaneously.
- [ ] Each client receives a unique player ID.
- [ ] Players spawn in a shared world.
- [ ] Clients can request movement.
- [ ] The server validates and processes movement.
- [ ] Updated player positions are broadcast to all connected clients.
- [ ] Players can see each other moving in the Three.js client.
- [ ] Disconnecting players are removed from the world.
- [ ] Basic tests pass and the server handles malformed input safely.

**The final result:** Open two browser sessions, connect both to the Go server, see both characters in the same world, and move them around while each session sees the other's movement.

That is the first meaningful multiplayer milestone for Freelands. Everything else can grow from there.