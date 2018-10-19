package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// CreateRoomAction is the action that is used when creating rooms
type CreateRoomAction struct {
	models.BaseAction
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Slots int    `json:"slots,omitempty"`
}

// CreateRoom gets called to create a new room
func CreateRoom(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateRoomList {
		models.SendJsonResponse(false, constants.ActionCreateRoom, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := CreateRoomAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionCreateRoom, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	r, err := models.CreateRoom(data.Name, data.Pass, player.GetID(), data.Slots)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionCreateRoom, err.Error(), mt, player)
		return
	}
	err = r.AddPlayer(player, data.Pass, true)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionCreateRoom, err.Error(), mt, player)
		return
	}
	models.Rooms.AddRoom(&r)
	player.SetState(constants.StateLobby)
	player.SetRoom(r.GetBsonID())
	models.SendJsonResponse(true, constants.ActionCreateRoom, "created room", mt, player)
	r.SendToPlayer(true, constants.ActionGetRoom, "", player)
}
