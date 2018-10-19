package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// ChangeOwnerAction is the action that is used when changing the owner of a room
type KickPlayerAction struct {
	models.BaseAction
	ID string `json:"id"`
}

// ChangeOwner changes the owner of a room
func KickPlayer(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateLobby {
		models.SendJsonResponse(false, constants.ActionKickPlayer, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := KickPlayerAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionKickPlayer, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		if player.GetID() == data.ID {
			models.SendJsonResponse(false, constants.ActionKickPlayer, constants.ErrCanNotKickSelf.Error(), mt, player)
			return
		}
		kickPlayer := r.Players[data.ID]
		if kickPlayer == nil {
			models.SendJsonResponse(false, constants.ActionKickPlayer, constants.ErrPlayerNotFound.Error(), mt, player)
			return
		}
		models.SendJsonResponse(true, constants.ActionKickPlayer, "Player got kicked", mt, player)
		models.SendJsonResponse(true, constants.ActionKickPlayer, "You got kicked by the room owner", mt, kickPlayer)
		r.RemovePlayer(kickPlayer)
		models.Players.SendToPlayer(true, constants.ActionGetRooms, "", kickPlayer)
		return
	}
	models.SendJsonResponse(false, constants.ActionKickPlayer, constants.ErrRoomNotFound.Error(), mt, player)
}
