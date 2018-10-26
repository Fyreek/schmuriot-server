package socketrouters

import (
	"github.com/schmonk.io/schmuriot-server/actions/undefined"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// UndefinedRouter handles every action from a player with undefined state
func UndefinedRouter(player *models.Player, message []byte, mt int, action string) {
	switch action {
	case constants.ActionSetUser:
		undefinedactions.SetUser(player, message, mt)
	case constants.ActionGetConfig:
		undefinedactions.GetConfig(player, mt)
	default:
		utils.LogToConsole("Not implemented")
		models.SendJsonResponse(false, constants.ActionNone, constants.ErrNotImplemented.Error(), mt, player)
	}
}
