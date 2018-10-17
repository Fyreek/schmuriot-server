package sockets

import (
	"strconv"

	"github.com/schmonk.io/schmuriot-server/actions"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/util"
)

// PlayerLoop handles every websocket message and calls appropriate functions
func PlayerLoop(player *models.Player) {
	for {
		mt, message, err := player.Connection.ReadMessage()
		util.LogToConsole("MT: " + strconv.Itoa(mt))
		util.LogToConsole("Message: " + string(message))
		if err != nil {
			if mt == -1 {
				actions.Disconnect(mt, player)
				util.LogToConsole("disconnect player:", err)
			} else {
				util.LogToConsole("message read error:", err)
			}
			break
		}
		if player.State == constants.StateUndefined {
			baseAction := models.BaseAction{}
			err := baseAction.Unmarshal(message)
			if err != nil {
				util.SendJsonResponse(false, constants.ActionNone, constants.ErrInvalidJSON.Error(), mt, player)
				continue
			}
			if !baseAction.Check(constants.ActionSetUser) {
				util.SendJsonResponse(false, constants.ActionNone, constants.ErrNameNotSet.Error(), mt, player)
				continue
			}
			actions.SetUser(player, message, mt)
		} else if player.State != constants.StateUndefined {
			ActionRouter(player, message, mt)
		} else {
			util.SendJsonResponse(false, constants.ActionNone, constants.ErrInvalidPlayerState.Error(), mt, player)
			continue
		}
	}
}
