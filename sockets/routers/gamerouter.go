package socketrouters

import (
	"github.com/schmonk.io/schmuriot-server/actions"
	"github.com/schmonk.io/schmuriot-server/actions/ingame"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// GameRouter handles every action from a player with inGame state
func GameRouter(player *models.Player, message []byte, mt int, action string) {
	switch action {
	case constants.ActionChat:
		actions.Chat(player, message, mt)
	case constants.ActionMakeMove:
		ingameactions.MakeMove(player, message, mt)
	default:
		utils.LogToConsole("Not implemented")
		models.SendJsonResponse(false, constants.ActionNone, constants.ErrNotImplemented.Error(), mt, player)
	}
}
