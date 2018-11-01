# Actions

## Overview

- [1 No State](#1-no-state)
	- [1.1 SetUser](#11-setuser)
		- [1.1.1 Request](#111-request)
		- [1.1.2 Responses](#112-responses)
	- [1.2 GetConfig](#12-getconfig)
		- [1.2.1 Request](#121-request)
		- [1.2.2 Responses](#122-responses)
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
	- [3.8 ToggleReady](#38-toggleready)
		- [3.8.1 Request](#381-request)
		- [3.8.2 Responses](#382-responses)
- [4 InGame](#4-ingame)
	- [4.1 StartRound](#41-startround)
		- [4.1.1 Request](#411-request)
		- [4.1.2 Responses](#412-responses)
	- [4.2 MakeMove](#42-makemove)
		- [4.2.1 Request](#421-request)
		- [4.2.2 Responses](#422-responses)
	- [4.3 MoveResult](#43-moveresult)
		- [4.3.1 Request](#431-request)
		- [4.3.2 Responses](#432-responses)
	- [4.4 CoinResult](#44-coinresult)
		- [4.4.1 Request](#441-request)
		- [4.4.2 Responses](#442-responses)
	- [4.5 EndGame](#45-endgame)
		- [4.5.1 Request](#451-request)
		- [4.5.2 Responses](#452-responses)
- [5 Not implemented](#5-not-implemented)
	- [5.1 GetMode](#51-getmode)
	- [5.2 ChangeMode](#52-changemode)
	- [5.3 GetGame](#53-getgame)
	- [5.4 ChangeGame](#54-changegame)

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
- [GetRooms Responses](#212-responses)

### 1.2 GetConfig

Gets the server configuration. Values include e.g. allowed length of room names, etc.

#### 1.2.1 Request

```
{
    "action": "getConfig"
}
```

#### 1.2.2 Responses

```
{
    "status": true,
    "action": "getConfig",
    "config": {
        "player": {
            "minNameLength": 3,
            "maxNameLength": 12
        },
        "room": {
            "minNameLength": 3,
            "maxNameLength": 12,
            "minSlots": 2,
            "maxSlots": 4,
            "minPassLength": 3,
            "maxPassLength": 12
        }
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
   "status":true,
   "action":"getRoom",
   "room":{
      "id":"5bd0ef6fac356640e899d483",
      "name":"TestRoom",
      "protected":false,
      "slots":4,
      "mode":"",
      "game":null,
      "owner":"5bd0ef42ac356640e899d480",
      "players":{
         "5bd0ef38ac356640e899d47f":{
            "id":"5bd0ef38ac356640e899d47f",
            "name":"Peter",
            "state":2,
            "ready":false,
            "roomid":"5bd0ef6fac356640e899d483",
            "posx":0,
            "posy":0,
            "color":""
         },
         "5bd0ef42ac356640e899d480":{
            "id":"5bd0ef42ac356640e899d480",
            "name":"Hans",
            "state":2,
            "ready":false,
            "roomid":"5bd0ef6fac356640e899d483",
            "posx":0,
            "posy":0,
            "color":""
         },
         "5bd0ef4aac356640e899d481":{
            "id":"5bd0ef4aac356640e899d481",
            "name":"Gerd",
            "state":2,
            "ready":false,
            "roomid":"5bd0ef6fac356640e899d483",
            "posx":0,
            "posy":0,
            "color":""
         },
         "5bd0ef4eac356640e899d482":{
            "id":"5bd0ef4eac356640e899d482",
            "name":"Samuel",
            "state":2,
            "ready":false,
            "roomid":"5bd0ef6fac356640e899d483",
            "posx":0,
            "posy":0,
            "color":""
         }
      }
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

### 3.8 ToggleReady

Toggles the player ready state. Game can only be started if all players are ready

#### 3.8.1 Request

```
{
    "action": "toggleReady"
}
```

#### 3.8.2 Responses

```
{
    "status": true,
    "action": "toggleReady",
    "message": "Toggled ready"
}
```
- [GetRoom Responses](#312-responses) -> To all players in the room
- If all players ready: [StartGame Responses](#412-responses)

## 4 InGame

### 4.1 StartRound

Starts a new round

#### 4.1.1 Request

Handled internally => no request, just response

#### 4.1.2 Responses

```
{
   "status":true,
   "action":"startGame",
   "game":{
      "fields":[
         [
            {
               "id":1,
               "player":"",
               "coins":4
            },
            {
               "id":2,
               "player":"5bda3fdfac356611874e3f4d",
               "coins":0
            },
            {
               "id":3,
               "player":"",
               "coins":2
            }
         ],
         [
            {
               "id":4,
               "player":"5bda3fe3ac356611874e3f4e",
               "coins":0
            },
            {
               "id":5,
               "player":"",
               "coins":3
            },
            {
               "id":6,
               "player":"5bda3fe9ac356611874e3f50",
               "coins":0
            }
         ],
         [
            {
               "id":7,
               "player":"",
               "coins":2
            },
            {
               "id":8,
               "player":"5bda3fe6ac356611874e3f4f",
               "coins":0
            },
            {
               "id":9,
               "player":"",
               "coins":1
            }
         ]
      ],
      "rounds":5,
      "countdown":15,
      "currentRound":1,
      "canReach":{
         "5bda3fdfac356611874e3f4d":[
            1,
            3,
            5
         ],
         "5bda3fe3ac356611874e3f4e":[
            1,
            5,
            7
         ],
         "5bda3fe6ac356611874e3f4f":[
            5,
            7,
            9
         ],
         "5bda3fe9ac356611874e3f50":[
            3,
            5,
            9
         ]
      },
      "moves":{

      },
      "coins":{

      },
      "state":0
   }
}
```

### 4.2 MakeMove

Makes a move for the player

#### 4.2.1 Request

```
{
    "action": "makeMove",
    "move": 5
}
```

#### 4.2.2 Responses

```
{
    "status": true,
    "action": "makeMove",
    "message": "Move was saved"
}
```

### 4.3 MoveResult

Sends the movement results of the last round to all players

#### 4.3.1 Request

Handled internally => no request, just response

#### 4.3.2 Responses

```
{
   "status":true,
   "action":"moveResult",
   "players":{
      "5bda3fdfac356611874e3f4d":{
         "player":"5bda3fdfac356611874e3f4d",
         "start":2,
         "move":5,
         "success":true
      },
      "5bda3fe3ac356611874e3f4e":{
         "player":"5bda3fe3ac356611874e3f4e",
         "start":4,
         "move":1,
         "success":true
      },
      "5bda3fe6ac356611874e3f4f":{
         "player":"5bda3fe6ac356611874e3f4f",
         "start":8,
         "move":8,
         "success":false
      },
      "5bda3fe9ac356611874e3f50":{
         "player":"5bda3fe9ac356611874e3f50",
         "start":6,
         "move":6,
         "success":false
      }
   }
}
```

### 4.4 CoinResult

#### 4.4.1 Request

Handled internally => no request, just response

#### 4.4.2 Responses

```
{
   "status":true,
   "action":"coinResult",
   "coins":{
      "5bda3fdfac356611874e3f4d":3,
      "5bda3fe3ac356611874e3f4e":4
   }
}
```

### 4.5 EndGame

Ends the game
Displays winner/s and coins for each player

#### 4.5.1 Request

Handled internally => no request, just response

#### 4.5.2 Responses

```
{
   "status":true,
   "action":"endGame",
   "coins":{
      "5bdb04b0ac12620d806b8505":4
   },
   "winner":[
      "5bdb04b0ac12620d806b8505"
   ]
}
```

## 5 Not implemented

### 5.1 GetMode

Gets a list off all available modes

### 5.2 ChangeMode

Changes the mode and resets player ready bool

### 5.3 GetGame

Gets a list of all available games

### 5.4 ChangeGame

Changes the game and resets player ready bool
