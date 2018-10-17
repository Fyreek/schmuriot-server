package sockets

import (
	"strconv"

	"github.com/schmonk.io/schmuriot-server/actions"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// PlayerLoop handles every websocket message and calls appropriate functions
func PlayerLoop(player *models.Player) {
	for {
		mt, message, err := player.Connection.ReadMessage()
		utils.LogToConsole("MT: " + strconv.Itoa(mt))
		utils.LogToConsole("Message: " + string(message))
		if err != nil {
			if mt == -1 {
				actions.Disconnect(mt, player)
				utils.LogToConsole("disconnect player:", err)
			} else {
				utils.LogToConsole("message read error:", err)
			}
			break
		}
		if player.State == constants.StateUndefined {
			baseAction := models.BaseAction{}
			err := baseAction.Unmarshal(message)
			if err != nil {
				models.SendJsonResponse(false, constants.ActionNone, constants.ErrInvalidJSON.Error(), mt, player)
				continue
			}
			if !baseAction.Check(constants.ActionSetUser) {
				models.SendJsonResponse(false, constants.ActionNone, constants.ErrNameNotSet.Error(), mt, player)
				continue
			}
			actions.SetUser(player, message, mt)
		} else if player.State != constants.StateUndefined {
			ActionRouter(player, message, mt)
		} else {
			models.SendJsonResponse(false, constants.ActionNone, constants.ErrInvalidPlayerState.Error(), mt, player)
			continue
		}
	}
}
