package sockets

import (
	"github.com/schmonk.io/schmuriot-server/actions"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// ActionRouter handles every action per player and calls corresponding functions
func ActionRouter(player *models.Player, message []byte, mt int) {
	baseAction := models.BaseAction{}
	err := baseAction.Unmarshal(message)
	if err != nil {
		models.SendJsonResponse(false, constants.ActionNone, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	utils.LogToConsole(baseAction.Action)
	// switch(player.State) {
	// case constants.StateUndefined:

	// case constants.StateRoomList:

	// case constants.StateLobby:

	// case constants.StateInGame:

	// case constants.StateSpectate:
	// default:

	// }

	switch baseAction.Action {
	case constants.ActionCreateRoom:
		actions.CreateRoom(player, message, mt)
	case constants.ActionGetRooms:
		actions.GetRooms(player, mt)
	case constants.ActionGetRoom:
		actions.GetRoom(player, mt)
	case constants.ActionJoinRoom:
		actions.JoinRoom(player, message, mt)
	case constants.ActionLeaveRoom:
		actions.LeaveRoom(player, message, mt)
	case constants.ActionDeleteRoom:
		actions.DeleteRoom(player, mt)
	case constants.ActionChat:
		actions.Chat(player, message, mt)
	case constants.ActionChangeOwner:
		actions.ChangeOwner(player, message, mt)
	case constants.ActionChangePassword:
		actions.ChangePassword(player, message, mt)
	case constants.ActionKickPlayer:
		actions.KickPlayer(player, message, mt)
	case constants.ActionToggleReady:
		actions.ToggleReady(player, mt)
	case constants.ActionStartGame:
		actions.StartGame(player, message, mt)
	default:
		utils.LogToConsole("Not implemented")
		models.SendJsonResponse(false, constants.ActionNone, constants.ErrNotImplemented.Error(), mt, player)
	}
}
