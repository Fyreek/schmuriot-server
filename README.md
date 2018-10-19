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

Sets the username for the connected user

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

Leave the current room

- `{"action":"leaveRoom"}`

#### Chat:

Posts a message into the room

- `{"action":"chat", "message":"..."}`

#### ChangeOwner:

Change the owner of a room

- `{"action":"changeOwner", "id":"..."}`
