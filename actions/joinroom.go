package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

type JoinRoomAction struct {
	models.BaseAction
	Id   string `json:"id"`
	Pass string `json:"pass"`
}

// JoinRoom lets a player join a room if it is not full and the right password was provided
func JoinRoom(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateRoomList {
		models.SendJsonResponse(false, constants.ActionJoinRoom, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := JoinRoomAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionJoinRoom, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(data.Id)
	if r != nil {
		err = r.AddPlayer(player, data.Pass, false)
		if err != nil {
			if err == constants.ErrRoomNotFound {
				utils.LogToConsole(err.Error())
				models.SendJsonResponse(false, constants.ActionJoinRoom, err.Error(), mt, player)
				return
			}
			utils.LogToConsole(err.Error())
			models.SendJsonResponse(false, constants.ActionJoinRoom, constants.ErrPassWrong.Error(), mt, player)
			return
		}
		models.SendJsonResponse(true, constants.ActionJoinRoom, "joined room", mt, player)
		r.SendToPlayer(true, constants.ActionGetRoom, "", player)
		return
	}
	models.SendJsonResponse(false, constants.ActionJoinRoom, constants.ErrRoomNotFound.Error(), mt, player)
}
