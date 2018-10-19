package actions

import (
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
)

// DeleteRoom is used to delete the current room if you are the owner
func DeleteRoom(player *models.Player, mt int) {
	if player.State != constants.StateLobby {
		models.SendJsonResponse(false, constants.ActionDeleteRoom, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		if r.Owner == player.GetID() {
			models.Rooms.RemoveRoom(r, false)
			for _, p := range r.Players {
				p.RemoveRoom()
				models.SendJsonResponse(true, constants.ActionDeleteRoom, "closed by room admin", mt, p)
			}
			models.Players.SendToAllPlayers(true, constants.ActionGetRooms, "", nil)
			return
		}
		models.SendJsonResponse(false, constants.ActionDeleteRoom, constants.ErrNotAdmin.Error(), mt, player)
		return
	}
	models.SendJsonResponse(false, constants.ActionDeleteRoom, constants.ErrRoomNotFound.Error(), mt, player)
}
