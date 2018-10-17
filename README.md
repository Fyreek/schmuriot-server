# Schmuriot Server

### Example Config

```
[SERVER]
IP = "localhost"
Port = 8080
Slots = 20
CORS = false
Debug = false

[Room]
NameLength = 16
Slots = 10
```

### Actions

#### SetUser

- `{"action":"setUser","name":"..."}`

#### CreateRoom

Creates a new room 

- `{"action":"createRoom","name":"...","pass":"","map":"","slots":4}`

#### GetRooms

Returns all rooms on the server if the player is in the roomlist

- `{"action":"getRooms"}`

#### GetRoom

Returns the current room the player is in if the player is in a room

- `{"action":"getRoom"}`

#### JoinRoom:

Joins the room with the provided id an optional password

- `{"action":"joinRoom","id":"...","pass":""}`

#### LeaveRoom:

Leaves the current room

- `{"action":"leaveRoom"}`
