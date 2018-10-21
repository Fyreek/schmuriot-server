package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// ChangeOwnerAction is the action that is used when changing the owner of a room
type ChangePasswordAction struct {
	models.BaseAction
	Pass string `json:"pass"`
}

// ChangeOwner changes the owner of a room
func ChangePassword(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateLobby {
		models.SendJsonResponse(false, constants.ActionChangePassword, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := ChangePasswordAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionChangePassword, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		if r.Owner != player.GetID() {
			models.SendJsonResponse(false, constants.ActionChangePassword, constants.ErrNotAdmin.Error(), mt, player)
			return
		}
		err = r.SetPass(data.Pass)
		if err != nil {
			models.SendJsonResponse(false, constants.ActionChangePassword, err.Error(), mt, player)
			return
		}
		models.SendJsonResponse(true, constants.ActionChangePassword, "password changed", mt, player)
		models.Players.SendToAllPlayers(true, constants.ActionGetRooms, "", nil)
		return
	}
	models.SendJsonResponse(false, constants.ActionChangePassword, constants.ErrRoomNotFound.Error(), mt, player)
}
