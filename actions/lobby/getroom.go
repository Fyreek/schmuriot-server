package lobbyactions

import (
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
)

func GetRoom(player *models.Player, mt int) {
	if player.State != constants.StateLobby && player.State != constants.StateInGame {
		models.SendJsonResponse(false, constants.ActionGetRoom, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	models.SendJsonResponseRoom(true, constants.ActionGetRoom, mt, player)
}
