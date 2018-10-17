package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmonk-server/util"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

type LeaveRoomAction struct {
	models.BaseAction
}

func LeaveRoom(player *models.Player, message []byte, mt int) {
	if player.State != util.StateLobby && player.State != util.StateReady && player.State != util.StatePlaying {
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
		return
	}
	models.SendJsonResponse(false, constants.ActionLeaveRoom, constants.ErrRoomNotFound.Error(), mt, player)

	// err = global.Rooms.RemovePlayer(player)
	// if err != nil {
	// 	util.LogToConsole(err.Error())
	// 	models.SendJsonResponse(false, util.ActionLeaveRoom, err.Error(), mt, &player.BasePlayer)
	// 	return
	// }
	// player.SetState(util.StateRoomList)
	// models.SendJsonResponse(true, util.ActionLeaveRoom, "left room", mt, &player.BasePlayer)
}
