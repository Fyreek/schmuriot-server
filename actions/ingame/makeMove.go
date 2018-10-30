package ingameactions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

type MakeMoveAction struct {
	models.BaseAction
	Move int `json:"move"`
}

func MakeMove(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateInGame {
		models.SendJsonResponse(false, constants.ActionMakeMove, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := MakeMoveAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionMakeMove, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		if r.Game != nil {
			reachArray := r.Game.CanReach[player.GetID()]

			found := false
			// Crash because of index error
			for _, i := range reachArray {
				if reachArray[i] == data.Move {
					found = true
				}
			}
			if !found {
				models.SendJsonResponse(false, constants.ActionMakeMove, constants.ErrFieldNotReachable.Error(), mt, player)
				return
			}

			pMove := models.CoinHunterMoves{}
			pMove.Player = player.GetID()
			pMove.Field = data.Move

			r.Game.Moves[player.GetID()] = pMove

			models.SendJsonResponse(true, constants.ActionMakeMove, "Move was saved", mt, player)

			// Later: Check if all players made move. If true, cancel countdown and end round immediately

			return

		}
		models.SendJsonResponse(false, constants.ActionMakeMove, constants.ErrGameNotSet.Error(), mt, player)
		return
	}
	models.SendJsonResponse(false, constants.ActionMakeMove, constants.ErrRoomNotFound.Error(), mt, player)
	return
}
