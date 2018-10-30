package ingameactions

import (
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
)

func RoundEnd(player *models.Player, mt int) {
	if player.State != constants.StateInGame {
		models.SendJsonResponse(false, constants.ActionEndRound, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		if r.Game != nil {
			r.Game.State = constants.GameStateEnd

			// Calculate actual movement
			// If two players want to get to the same field, cancel
			// Collect coins
			// Send response

			validMovements := map[string]int{}
			for pID := range r.Game.Moves {
				pMove := r.Game.Moves[pID]
				collision := false
				for innerPID := range r.Game.Moves {
					innerPMove := r.Game.Moves[innerPID]
					if innerPMove.Player != pMove.Player {
						if innerPMove.Field == pMove.Field {
							collision = true
						}
					}
				}
				if !collision {
					validMovements[pMove.Player] = pMove.Field
				}
			}
			movements := map[string]models.CoinHunterMovement{}
			for p := range r.Players {
				m := models.CoinHunterMovement{}
				m.Player = r.Players[p].GetID()
				// m.Start = r.Game.Moves[p].Field
				if vMove, ok := validMovements[p]; ok {
					m.Move = vMove
					m.Success = true
				} else {
					m.Move = r.Game.Moves[p].Field
					m.Success = false
				}
				movements[p] = m
			}
			for p := range r.Players {
				models.SendJsonResponseMovement(movements, mt, r.Players[p])
			}

			// test if round limit reached
			// increase round
			// if round limit send leaderboard

			return
		}
		models.SendJsonResponse(false, constants.ActionEndRound, constants.ErrGameNotSet.Error(), mt, player)
		return
	}
	models.SendJsonResponse(false, constants.ActionEndRound, constants.ErrRoomNotFound.Error(), mt, player)
	return
}
