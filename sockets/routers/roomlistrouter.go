package socketrouters

import (
	"github.com/schmonk.io/schmuriot-server/actions/roomlist"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// RoomListRouter handles every action from a player with roomlist state
func RoomListRouter(player *models.Player, message []byte, mt int, action string) {
	switch action {
	case constants.ActionGetRooms:
		roomlistactions.GetRooms(player, mt)
	case constants.ActionCreateRoom:
		roomlistactions.CreateRoom(player, message, mt)
	case constants.ActionJoinRoom:
		roomlistactions.JoinRoom(player, message, mt)
	default:
		utils.LogToConsole("Not implemented")
		models.SendJsonResponse(false, constants.ActionNone, constants.ErrNotImplemented.Error(), mt, player)
	}
}
