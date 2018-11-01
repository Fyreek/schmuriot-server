package socketrouters

import (
	"github.com/schmonk.io/schmuriot-server/actions"
	"github.com/schmonk.io/schmuriot-server/actions/lobby"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// LobbyRouter handles every action from a player with lobby state
func LobbyRouter(player *models.Player, message []byte, mt int, action string) {
	switch action {
	case constants.ActionGetRoom:
		lobbyactions.GetRoom(player, mt)
	case constants.ActionLeaveRoom:
		lobbyactions.LeaveRoom(player, message, mt)
	case constants.ActionDeleteRoom:
		lobbyactions.DeleteRoom(player, mt)
	case constants.ActionChat:
		actions.Chat(player, message, mt)
	case constants.ActionChangeOwner:
		lobbyactions.ChangeOwner(player, message, mt)
	case constants.ActionChangePassword:
		lobbyactions.ChangePassword(player, message, mt)
	case constants.ActionKickPlayer:
		lobbyactions.KickPlayer(player, message, mt)
	case constants.ActionToggleReady:
		lobbyactions.ToggleReady(player, mt)
	// case constants.ActionChangeGame:
	// case constants.ActionChangeMode:
	// case constants.ActionGetGame:
	// case constants.ActionGetMode:
	default:
		utils.LogToConsole("Not implemented")
		models.SendJsonResponse(false, constants.ActionNone, constants.ErrNotImplemented.Error(), mt, player)
	}
}
