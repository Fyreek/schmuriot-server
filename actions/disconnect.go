package actions

import (
	"fmt"

	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// Disconnect disconnects a player from the server
func Disconnect(mt int, player *models.Player) {
	if player.RoomID != nil {
		r := models.Rooms.GetRoom(player.RoomID.Hex())
		err := r.RemovePlayer(player)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	models.Players.RemovePlayer(player)
	utils.LogToConsole("Connected Players:", models.Players.GetPlayerCount())
}
