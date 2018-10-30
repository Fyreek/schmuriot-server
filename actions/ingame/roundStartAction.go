package ingameactions

import (
	"fmt"
	"time"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
)

func RoundStart(player *models.Player, mt int) {
	if player.State != constants.StateInGame {
		models.SendJsonResponse(false, constants.ActionStartRound, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		if r.Game != nil {
			r.Game.GenerateField(r.GetPlayerList())
			r.SendToAllPlayers(true, constants.ActionStartRound, "", nil)

			gameTimer := time.NewTimer(time.Duration(r.Game.Countdown) * time.Second)

			<-gameTimer.C
			fmt.Println("gameTimer expired... Triggering round end")
			RoundEnd(player, mt)
			return
		}
		models.SendJsonResponse(false, constants.ActionStartRound, constants.ErrGameNotSet.Error(), mt, player)
		return
	}
	models.SendJsonResponse(false, constants.ActionStartRound, constants.ErrRoomNotFound.Error(), mt, player)
	return
}
