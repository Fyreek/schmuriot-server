package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// ChatAction is the action that is used when sending a chat message
type ChatAction struct {
	models.BaseAction
	Message string `json:"message"`
}

// Chat sends a chat message to all room players
func Chat(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateLobby && player.State != constants.StateInGame {
		models.SendJsonResponse(false, constants.ActionChat, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := ChatAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionChat, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		r.SendToAllPlayers(true, constants.ActionChat, data.Message, nil)
		return
	}
	models.SendJsonResponse(false, constants.ActionChat, constants.ErrRoomNotFound.Error(), mt, player)
}
