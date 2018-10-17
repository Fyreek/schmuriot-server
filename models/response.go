package models

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
)

type StatusResponse struct {
	Status bool   `json:"status"`
	Action string `json:"action"`
}

type StatusResponseMessage struct {
	StatusResponse
	Message string `json:"message"`
}

type StatusResponsePlayerID struct {
	StatusResponse
	PlayerID string `json:"playerid"`
}

type StatusResponseRoomList struct {
	StatusResponse
	Rooms map[string]*Room `json:"rooms"`
}

type StatusResponseRoom struct {
	StatusResponse
	Room *Room `json:"room"`
}

//SendJsonResponse sends a response with a status, the provided action and a custom message
func SendJsonResponse(status bool, action string, message string, mt int, player *Player) {
	resp := StatusResponseMessage{}
	resp.Status = status
	resp.Action = action
	resp.Message = message
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte(constants.ErrSerializing.Error()))
	}
	player.Connection.WriteMessage(mt, bytes)
}

//SendJsonResponsePlayerID sends a response with a status, the provided action and the playerID
func SendJsonResponsePlayerID(status bool, action string, id string, mt int, player *Player) {
	resp := StatusResponsePlayerID{}
	resp.Status = status
	resp.Action = action
	resp.PlayerID = id
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte(constants.ErrSerializing.Error()))
	}
	player.Connection.WriteMessage(mt, bytes)
}

//SendJsonResponseRoomList sends a response with a status, the provided action and the current RoomList
func SendJsonResponseRoomList(status bool, action string, rooms map[string]*Room, mt int, player *Player) {
	resp := StatusResponseRoomList{}
	resp.Status = status
	resp.Action = action
	resp.Rooms = rooms
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte(constants.ErrSerializing.Error()))
	}
	player.Connection.WriteMessage(mt, bytes)
}

func SendJsonResponseRoom(status bool, action string, mt int, player *Player) {
	resp := StatusResponseRoom{}
	resp.Status = status
	resp.Action = action
	resp.Room = Rooms.GetRoom(player.GetRoomID())
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte(constants.ErrSerializing.Error()))
	}
	player.Connection.WriteMessage(mt, bytes)
}
