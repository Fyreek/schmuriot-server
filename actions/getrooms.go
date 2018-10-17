package actions

import (
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
)

// GetRooms returns all available rooms to the requesting player
func GetRooms(player *models.Player, mt int) {
	if player.State != constants.StateRoomList {
		models.SendJsonResponse(false, constants.ActionGetRooms, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	models.SendJsonResponseRoomList(true, constants.ActionGetRooms, models.Rooms.GetRooms(), mt, player)
}
