package socketrouters

import (
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// BaseRouter handles every action and forwards it to the corresponding actions router
func BaseRouter(player *models.Player, message []byte, mt int) {
	baseAction := models.BaseAction{}
	err := baseAction.Unmarshal(message)
	if err != nil {
		models.SendJsonResponse(false, constants.ActionNone, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	utils.LogToConsole(baseAction.Action)
	switch player.State {
	case constants.StateUndefined:
		UndefinedRouter(player, message, mt, baseAction.Action)
	case constants.StateRoomList:
		RoomListRouter(player, message, mt, baseAction.Action)
	case constants.StateLobby:
		LobbyRouter(player, message, mt, baseAction.Action)
	case constants.StateInGame:
		GameRouter(player, message, mt, baseAction.Action)
	case constants.StateSpectate:
	default:
		utils.LogToConsole("Not implemented")
		models.SendJsonResponse(false, constants.ActionNone, constants.ErrNotImplemented.Error(), mt, player)
	}
}
