package actions

import (
	"github.com/schmonk.io/schmuriot-server/global"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/util"
)

// Disconnect disconnects a player from the server
func Disconnect(mt int, player *models.Player) {
	if player.RoomID != nil {
		r := global.Rooms.GetRoom(player.RoomID.Hex())
		r.RemovePlayer(player)
	}
	global.Players.RemovePlayer(player)
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
