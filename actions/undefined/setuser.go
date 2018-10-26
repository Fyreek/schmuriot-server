package undefinedactions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

// SetUserAction is the action to set a users name and register the user at the server
type SetUserAction struct {
	models.BaseAction
	Name string `json:"name"`
}

// SetUser gets called to register a new user
func SetUser(player *models.Player, message []byte, mt int) {
	data := SetUserAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionSetUser, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	err = player.SetName(data.Name)
	if err != nil {
		models.SendJsonResponse(false, constants.ActionSetUser, err.Error(), mt, player)
		return
	}
	player.SetState(constants.StateRoomList)
	models.Players.AddPlayer(player)
	models.SendJsonResponsePlayerID(true, constants.ActionSetUser, player.GetID(), mt, player)
	models.Players.SendToPlayer(true, constants.ActionGetRooms, "", player)
	utils.LogToConsole("Connected Players:", models.Players.GetPlayerCount())
}
