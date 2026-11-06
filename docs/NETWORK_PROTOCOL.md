# Freelands — Network Protocol Specification

**Document:** Network Protocol  
**Version:** 0.1  
**Status:** Initial specification  
**Transport:** WebSocket  
**Serialization:** JSON  
**Server:** Go  
**Client:** SvelteKit + TypeScript

---

## 1. Overview

This document defines the network communication protocol used between Freelands clients and the authoritative game server.

The initial protocol supports:

- Establishing a multiplayer connection.
- Identifying the connected player.
- Joining a shared world.
- Synchronizing player state.
- Sending movement input.
- Receiving player updates.
- Notifying clients about player connections and disconnections.
- Communicating protocol errors.

The protocol is designed to be extended as additional gameplay systems are implemented.

### 1.1 Protocol principles

1. The server is authoritative.
2. Clients send player intentions, not authoritative gameplay results.
3. All connected players in the same world instance share authoritative world state.
4. Messages use explicitly defined structures.
5. Invalid messages must be handled safely.
6. Protocol changes must be reflected in both client and server implementations.

## 2. Transport

### 2.1 Connection

The client connects using WebSocket.

Development endpoint:

`ws://localhost:8080/ws`

Production endpoint:

`wss://<server-domain>/ws`

The endpoint must be configurable by the client.

### 2.2 Message encoding

All application messages are JSON objects encoded as UTF-8.

Each WebSocket text message contains exactly one application message.

Binary messages are not supported in protocol version 0.1.

### 2.3 Message size

The initial maximum incoming application message size is 4 KiB.

Messages exceeding this limit must be rejected.

## 3. Message Format

All messages follow a common envelope.

### 3.1 Standard envelope

```json
{
  "type": "message_type",
  "payload": {}
}
```

| Field | Type | Required | Description |
|---|---|---|---|
| `type` | string | Yes | Message identifier |
| `payload` | object | Yes | Message-specific data |

### 3.2 Rules

- Message types are lowercase `snake_case`.
- Payload fields use `snake_case`.
- Unknown message types must produce an `unknown_message_type` error.
- Invalid JSON must be handled without crashing the server.
- Missing or incorrectly typed required fields must be rejected.
- The server must never trust client-supplied authoritative state.

## 4. Connection Lifecycle

The initial connection sequence is:

1. Client opens a WebSocket connection.
2. Server accepts the connection.
3. Server generates a unique player ID.
4. Server creates the player's initial state.
5. Server registers the player in the world.
6. Server sends `welcome`.
7. Server sends `world_state`.
8. Server broadcasts `player_joined` to existing clients.
9. Client enters the connected gameplay state.

The server must register the player before announcing its arrival to other players.

### 4.1 Disconnection

When a connection closes:

1. Server removes the player from the active world.
2. Server releases connection resources.
3. Server broadcasts `player_left` to remaining clients.

The removal operation must be safe if triggered more than once.

## 5. Protocol Messages

### 5.1 Client → Server

| Message | Purpose |
|---|---|
| `move` | Request player movement |

### 5.2 Server → Client

| Message | Purpose |
|---|---|
| `welcome` | Confirm successful connection |
| `world_state` | Provide current world snapshot |
| `player_joined` | Announce a new player |
| `player_left` | Announce a disconnected player |
| `error` | Report an invalid request |

## 6. Message Specifications

### 6.1 `welcome`

**Direction:** Server → Client

Sent once after successful player registration.

```json
{
  "type": "welcome",
  "payload": {
    "player_id": "player-123"
  }
}
```

| Field | Type | Description |
|---|---|---|
| `player_id` | string | Server-generated player identifier |

Rules:

- The server generates the identifier.
- The identifier must be unique among active players.
- The client must use this identifier to identify its local player.
- The identifier is not an authentication credential.

### 6.2 `world_state`

**Direction:** Server → Client

Provides a complete snapshot of the current players in the world instance.

```json
{
  "type": "world_state",
  "payload": {
    "players": [
      {
        "id": "player-123",
        "x": 10.5,
        "y": 0.0,
        "z": 20.0,
        "rotation": 1.57
      },
      {
        "id": "player-456",
        "x": 15.0,
        "y": 0.0,
        "z": 24.0,
        "rotation": 3.14
      }
    ]
  }
}
```

| Field | Type | Description |
|---|---|---|
| `players` | array | Active players in the world |
| `players[].id` | string | Player identifier |
| `players[].x` | number | X coordinate |
| `players[].y` | number | Y coordinate |
| `players[].z` | number | Z coordinate |
| `players[].rotation` | number | Horizontal rotation in radians |

Rules:

- The snapshot contains every active player in the relevant world instance.
- Player IDs must be unique within the snapshot.
- The initial snapshot includes the connecting player.
- An empty player list is valid.
- The client uses player IDs to create or update player representations.

### 6.3 `move`

**Direction:** Client → Server

Requests movement for the player associated with the connection.

```json
{
  "type": "move",
  "payload": {
    "x": 1.0,
    "z": 0.0
  }
}
```

| Field | Type | Description |
|---|---|---|
| `x` | number | Horizontal movement direction |
| `z` | number | Depth movement direction |

Input convention:

| Value | Meaning |
|---|---|
| `x = 1` | Positive X direction |
| `x = -1` | Negative X direction |
| `z = 1` | Positive Z direction |
| `z = -1` | Negative Z direction |
| `x = 0, z = 0` | No movement |

Rules:

- Both values must be finite numbers.
- The server normalizes movement input.
- The server determines movement speed.
- The server calculates the resulting position.
- The client cannot submit a destination position.
- The server must associate the request with the connection's player ID.

Movement input represents the player's current intended direction, not a request to teleport or move a specified distance.

### 6.4 `player_joined`

**Direction:** Server → Client

Announces that a new player has entered the world.

```json
{
  "type": "player_joined",
  "payload": {
    "player": {
      "id": "player-456",
      "x": 15.0,
      "y": 0.0,
      "z": 24.0,
      "rotation": 3.14
    }
  }
}
```

The `player` object uses the same structure as an individual player entry in `world_state`.

Rules:

- Sent to existing clients after registration.
- The newly connected player does not need this notification about itself.
- Duplicate notifications must not create duplicate player representations.

### 6.5 `player_left`

**Direction:** Server → Client

Announces that a player has left the world.

```json
{
  "type": "player_left",
  "payload": {
    "player_id": "player-456"
  }
}
```

| Field | Type | Description |
|---|---|---|
| `player_id` | string | Identifier of the departing player |

Rules:

- Sent after removal from the active world.
- Clients remove the corresponding player representation.
- Receiving an unknown player ID must be harmless.

### 6.6 `error`

**Direction:** Server → Client

Reports a protocol or request error.

```json
{
  "type": "error",
  "payload": {
    "code": "invalid_payload",
    "message": "The movement input is invalid."
  }
}
```

| Field | Type | Description |
|---|---|---|
| `code` | string | Machine-readable error code |
| `message` | string | Human-readable description |

Initial error codes:

| Code | Meaning |
|---|---|
| `invalid_message` | Message could not be parsed |
| `unknown_message_type` | Unsupported message type |
| `invalid_payload` | Payload does not match the specification |
| `rate_limited` | Client exceeded the permitted message rate |

Error messages must not expose internal server details or stack traces.

## 7. Coordinate System

Protocol version 0.1 uses the following coordinate convention:

| Axis | Meaning |
|---|---|
| X | Horizontal world axis |
| Y | Vertical world axis |
| Z | Horizontal depth axis |

Player state:

```json
{
  "id": "player-123",
  "x": 10.5,
  "y": 0.0,
  "z": 20.0,
  "rotation": 1.57
}
```

Rules:

- Position values are world coordinates.
- Rotation is expressed in radians.
- The initial prototype uses a flat world with `y = 0`.
- Coordinate conventions must remain consistent between client and server.

## 8. Movement Synchronization

The initial server simulation runs at a target rate of 20 ticks per second.

The client sends movement input. The server processes the latest input during its simulation updates.

At each update:

1. Validate the latest input.
2. Normalize the movement direction.
3. Calculate displacement using server-controlled movement speed and elapsed simulation time.
4. Update authoritative position.
5. Broadcast the resulting player state.

For protocol version 0.1, the server broadcasts complete `world_state` snapshots at the simulation update rate.

This is intentionally a simple initial synchronization strategy.

### 8.1 Client interpolation

The client may interpolate between received player positions for smooth rendering.

The displayed position may temporarily differ from the latest authoritative position.

The client must not treat its rendered position as authoritative.

## 9. Message Ordering

WebSocket preserves message order within a connection, but this alone does not solve every synchronization problem.

For version 0.1:

- The server sends `welcome` before the initial `world_state`.
- The server sends the initial snapshot before announcing the new player to existing clients.
- The server serializes outgoing messages for each connection.
- Clients process messages in received order.
- Repeated `player_joined` or `player_left` events must be handled safely.

Version 0.1 does not define sequence numbers, acknowledgements or replay protection.

These can be introduced when prediction, reconciliation or more complex synchronization is needed.

## 10. Validation Requirements

The server must validate every client message.

For `move`:

- `x` and `z` must exist.
- Both must be numeric and finite.
- The server must normalize the input.
- The server must apply movement speed limits.
- The server must reject malformed payloads.
- The server must associate movement with the sending connection.

Additional requirements:

- Enforce maximum message size.
- Apply incoming message rate limits.
- Reject unsupported message types.
- Handle malformed JSON safely.
- Do not trust client-supplied player IDs or positions.

## 11. Disconnection and Recovery

When a WebSocket connection closes:

- The associated player is removed from the active world.
- Remaining clients receive `player_left`.
- The server releases connection resources.

If the client loses its connection:

- It must stop treating the session as connected.
- It may attempt to reconnect.
- A reconnection is a new connection in protocol version 0.1.
- The server assigns a new player ID and spawn state.

Session recovery and restoration of persistent characters are outside the initial protocol.

## 12. Protocol Versioning

This document defines protocol version 0.1.

Version 0.1 does not require a version field in individual messages because both client and server are developed together for the initial prototype.

Before public deployment or independent client/server releases, introduce explicit protocol version negotiation.

Breaking changes must not silently alter existing message structures or meanings.

## 13. Initial Message Examples

### Successful connection

Server → Client:

```json
{
  "type": "welcome",
  "payload": {
    "player_id": "player-123"
  }
}
```

Server → Client:

```json
{
  "type": "world_state",
  "payload": {
    "players": [
      {
        "id": "player-123",
        "x": 0.0,
        "y": 0.0,
        "z": 0.0,
        "rotation": 0.0
      }
    ]
  }
}
```

### Movement

Client → Server:

```json
{
  "type": "move",
  "payload": {
    "x": 1.0,
    "z": 0.0
  }
}
```

Server → Client:

```json
{
  "type": "world_state",
  "payload": {
    "players": [
      {
        "id": "player-123",
        "x": 0.25,
        "y": 0.0,
        "z": 0.0,
        "rotation": 0.0
      }
    ]
  }
}
```

The example position assumes a particular simulation interval and movement speed; actual values are calculated by the server.

### Another player joins

Server → Existing Clients:

```json
{
  "type": "player_joined",
  "payload": {
    "player": {
      "id": "player-456",
      "x": 5.0,
      "y": 0.0,
      "z": 5.0,
      "rotation": 0.0
    }
  }
}
```

### Player disconnects

Server → Remaining Clients:

```json
{
  "type": "player_left",
  "payload": {
    "player_id": "player-456"
  }
}
```

## 14. Initial Protocol Acceptance Criteria

Protocol version 0.1 is implemented when:

- [ ] Client and server use the same message envelope.
- [ ] All initial message types follow their documented structures.
- [ ] A client receives `welcome` after connecting.
- [ ] A client receives the initial `world_state`.
- [ ] A client can send valid movement input.
- [ ] The server processes movement authoritatively.
- [ ] Clients receive updated player states.
- [ ] Clients receive player join notifications.
- [ ] Clients receive player departure notifications.
- [ ] Invalid messages are handled safely.
- [ ] Multiple clients can communicate simultaneously.
- [ ] The two-browser multiplayer test succeeds.

---

**Document ownership:** This specification is the shared reference for the initial Freelands network protocol. Both the Go server and TypeScript client must implement the same message definitions.