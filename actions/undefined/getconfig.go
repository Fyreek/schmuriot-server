package undefinedactions

import (
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
)

func GetConfig(player *models.Player, mt int) {
	if player.State != constants.StateUndefined {
		models.SendJsonResponse(false, constants.ActionDeleteRoom, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	models.SendJsonResponseConfig(mt, player)
}
