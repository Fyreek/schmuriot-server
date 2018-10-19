package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// ChangeOwnerAction is the action that is used when changing the owner of a room
type ChangeOwnerAction struct {
	models.BaseAction
	ID string `json:"id"`
}

// ChangeOwner changes the owner of a room
func ChangeOwner(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateLobby {
		models.SendJsonResponse(false, constants.ActionCreateRoom, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := ChangeOwnerAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionCreateRoom, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		if r.Owner != player.GetID() {
			models.SendJsonResponse(false, constants.ActionChangeOwner, constants.ErrNotAdmin.Error(), mt, player)
			return
		}
		newOwner := r.Players[data.ID]
		if newOwner == nil {
			models.SendJsonResponse(false, constants.ActionChangeOwner, constants.ErrPlayerNotFound.Error(), mt, player)
			return
		}
		r.SetOwner(data.ID)
		models.SendJsonResponse(true, constants.ActionChangeOwner, "owner changed", mt, player)
		r.SendToAllPlayers(true, constants.ActionGetRoom, "", player)
		r.SendToPlayer(true, constants.ActionGetRoom, "", player)
		return
	}
	models.SendJsonResponse(false, constants.ActionJoinRoom, constants.ErrRoomNotFound.Error(), mt, player)
}
