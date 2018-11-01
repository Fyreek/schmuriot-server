package lobbyactions

import (
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
)

// ToggleReady is used to toggle the ready attribute of the player
func ToggleReady(player *models.Player, mt int) {
	if player.State != constants.StateLobby {
		models.SendJsonResponse(false, constants.ActionToggleReady, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		player.ToggleReady()
		models.SendJsonResponse(true, constants.ActionToggleReady, "Toggled ready", mt, player)
		r.SendToAllPlayers(true, constants.ActionGetRoom, "", nil)
		start := r.CheckAllReady()
		if start {
			// startGameAction := StartGameAction{}
			// startGameAction.Action = constants.ActionStartGame
			// startGameAction.Countdown = 15
			// startGameAction.Rounds = 5
			// bytes, _ := startGameAction.Marshal()
			// StartGame(player, bytes, mt)
			go r.StartGame()
		}
		return
	}
	models.SendJsonResponse(false, constants.ActionDeleteRoom, constants.ErrRoomNotFound.Error(), mt, player)
}
