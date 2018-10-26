package lobbyactions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

type LeaveRoomAction struct {
	models.BaseAction
}

func LeaveRoom(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateLobby && player.State != constants.StateInGame {
		models.SendJsonResponse(false, constants.ActionLeaveRoom, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := LeaveRoomAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionLeaveRoom, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		err = r.RemovePlayer(player)
		if err != nil {
			utils.LogToConsole(err.Error())
			models.SendJsonResponse(false, constants.ActionLeaveRoom, err.Error(), mt, player)
			return
		}
		models.SendJsonResponse(true, constants.ActionLeaveRoom, "left room", mt, player)
		models.Players.SendToPlayer(true, constants.ActionGetRooms, "", player)
		return
	}
	models.SendJsonResponse(false, constants.ActionLeaveRoom, constants.ErrRoomNotFound.Error(), mt, player)
}
