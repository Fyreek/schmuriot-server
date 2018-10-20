# Actions

## Overview

- [1 No State](#1-no-state)
    - [1.1 SetUser](#11-setuser)
        - [1.1.1 Request](#111-request)
        - [1.1.2 Responses](#112-responses)
- [2 RoomList](#2-roomlist)
    - [2.1 GetRooms](#21-getrooms)
        - [2.1.1 Request](#211-request)
        - [2.1.2 Responses](#212-responses)
    - [2.2 CreateRoom](#22-createroom)
        - [2.2.1 Request](#221-request)
        - [2.2.2 Responses](#222-responses)
    - [2.3 JoinRoom](#23-joinroom)
        - [2.3.1 Request](#231-request)
        - [2.3.2 Responses](#232-responses)
- [3 Room](#3-room)
    - [3.1 GetRoom](#31-getroom)
        - [3.1.1 Request](#311-request)
        - [3.1.2 Responses](#312-responses)
    - [3.2 LeaveRoom](#32-leaveroom)
        - [3.2.1 Request](#321-request)
        - [3.2.2 Responses](#322-responses)
    - [3.3 DeleteRoom](#33-deleteroom)
        - [3.3.1 Request](#331-request)
        - [3.3.2 Responses](#332-responses)
    - [3.4 Chat](#34-chat)
        - [3.4.1 Request](#341-request)
        - [3.4.2 Responses](#342-responses)
    - [3.5 ChangeOwner](#35-changeowner)
        - [3.5.1 Request](#351-request)
        - [3.5.2 Responses](#352-responses)
    - [3.6 ChangePassword](#36-changepassword)
        - [3.6.1 Request](#361-request)
        - [3.6.2 Responses](#362-responses)
    - [3.7 KickPlayer](#37-kickplayer)
        - [3.7.1 Request](#371-request)
        - [3.7.2 Responses](#372-responses)
- [4 Not implemented](#4-not-implemented)
    - [4.1 GetMode](#41-getmode)
    - [4.2 ChangeMode](#42-changemode)
    - [4.3 GetGame](#43-getgame)
    - [4.4 ChangeGame](#44-changegame)

## 1 No State

### 1.1 SetUser

Sets the name of the user and registers him on the server

#### 1.1.1 Request

```
{
    "action": "setUser",
    "name": "..."
}
```

#### 1.1.2 Responses

```
{
    "status": true,
    "action": "setUser",
    "playerid": "5678"
}
```
```
{
    "status": true,
    "action": "getRooms",
    "rooms": {
        ...
    }
}
```

## 2 RoomList

### 2.1 GetRooms

Gets a list off all available rooms

#### 2.1.1 Request

```
{
    "action": "getRooms"
}
```

#### 2.1.2 Responses

```
{
    "status": true,
    "action": "getRooms",
    "rooms": {
        "1234": {
            "id": "1234",
            "name": "RaumName",
            "protected": false,
            "map": "",
            "slots": 4,
            "owner": "5678",
            "players": {
                "5678": {
                    "id": "5678",
                    "name": "SpielerName",
                    "state": 2,
                    "roomid": "1234",
                    "posx": 0,
                    "posy": 0,
                    "color": ""
                }
            }
        }
    }
}
```

### 2.2 CreateRoom

Creates a new room. Then joins it as the owner of the room

#### 2.2.1 Request

```
{
    "action": "createRoom",
    "name": "...",
    "pass": ""
}
```

#### 2.2.2 Responses

```
{
    "status": true,
    "action": "createRoom",
    "message": "created room"
}
```

- [GetRoom Responses](#312-responses) -> To the player that created the room
- [GetRooms Responses](#212-responses) -> To all other players still in roomlist

### 2.3 JoinRoom

Joines a room. Password is optional

#### 2.3.1 Request

```
{
    "action": "joinRoom",
    "id": "1234",
    "pass": ""
}
```

#### 2.3.2 Responses

```
{
    "status": true,
    "action": "joinRoom",
    "message": "joined room"
}
```

- [GetRoom Responses](#312-responses) -> To all players in the room
- [GetRooms Responses](#212-responses) -> To all other players still in roomlist

## 3 Room

### 3.1 GetRoom

Gets the room info.

#### 3.1.1 Request

```
{
    "action": "getRoom"
}
```

#### 3.1.2 Responses

```
{
    "status": true,
    "action": "getRoom",
    "room": {
        "id": "1234",
        "name": "TestRaum",
        "protected": false,
        "slots": 4,
        "owner": "56789",
        "players": {
            "56789":{
                "id": "56789",
                "name": "TestSpieler",
                "state": 2,
                "roomid": "1234",
                "posx": 0,
                "posy": 0,
                "color": ""
            },
            "734732": {
                ...
            }
        },
        "mode":"",
        "game":""
    }
}
```

### 3.2 LeaveRoom

Leaves the current room and gets back to the roomlist

#### 3.2.1 Request

```
{
    "action": "leaveRoom"
}
```

#### 3.2.2 Responses

```
{
    "status": true,
    "action": "leaveRoom",
    "message": "left room"
}
```

- [GetRoom Responses](#312-responses) -> To all players in the room
- [GetRooms Responses](#212-responses) -> To all other players still in roomlist

### 3.3 DeleteRoom

Deletes the current room and puts all players back to the roomlist. 
Can only be run by the owner.

#### 3.3.1 Request

```
{
    "action": "deleteRoom"
}
```

#### 3.3.2 Responses

```
{
    "status": true,
    "action": "deleteRoom",
    "message": "closed by room admin"
}
```
- -> To all players
- [GetRooms Responses](#212-responses) -> To all other players still in roomlist

### 3.4 Chat

Sends a chat message to all players in a room

#### 3.4.1 Request

```
{
    "action": "chat",
    "message": "..."
}
```

#### 3.4.2 Responses

```
{
    "status": true,
    "action": "chat",
    "message": "...",
    "playerid": "1234"
}
```
- -> To all players

### 3.5 ChangeOwner

Sets a new player from the room as owner.
Can only be run by the owner.

#### 3.5.1 Request

```
{
    "action": "changeOwner",
    "id": "5678"
}
```

#### 3.5.2 Responses

```
{
    "status": true,
    "action": "changeOwner",
    "message": "owner changed"
}
```

- [GetRoom Responses](#312-responses) -> To all players in the room

### 3.6 ChangePassword

Changes the password of a room.
"pass" can also be empty to remove the room password.
Can only be run by the owner.

#### 3.6.1 Request

```
{
    "action": "changePassword",
    "pass": "..."
}
```

#### 3.6.2 Responses

```
{
    "status": true,
    "action": "changePassword",
    "message": "password changed"
}
```

### 3.7 KickPlayer

Kicks a player from the room.
Can only be run by the admin.

#### 3.7.1 Request

```
{
    "action": "kickPlayer",
    "id": "5678"
}
```

#### 3.7.2 Responses

```
{
    "status": true,
    "action": "kickPlayer",
    "message": "Player got kicked"
}
```
- -> To the owner
```
{
    "status": true,
    "action": "kickPlayer",
    "message": "You got kicked by the room owner"
}
```
- -> To the kicked player
- [GetRoom Responses](#312-responses) -> To all players in the room
- [GetRooms Responses](#212-responses) -> To all other players still in roomlist

## 4 Not implemented

### 4.1 GetMode

### 4.2 ChangeMode

### 4.3 GetGame

### 4.4 ChangeGame
