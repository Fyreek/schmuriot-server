package util

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
)

//SendJsonResponse sends a response with a status, the provided action and a custom message
func SendJsonResponse(status bool, action string, message string, mt int, player *models.Player) {
	resp := models.StatusResponseMessage{}
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
func SendJsonResponsePlayerID(status bool, action string, id string, mt int, player *models.Player) {
	resp := models.StatusResponsePlayerID{}
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
func SendJsonResponseRoomList(status bool, action string, rooms map[string]*models.Room, mt int, player *models.Player) {
	resp := models.StatusResponseRoomList{}
	resp.Status = status
	resp.Action = action
	resp.Rooms = rooms
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte(constants.ErrSerializing.Error()))
	}
	player.Connection.WriteMessage(mt, bytes)
}
