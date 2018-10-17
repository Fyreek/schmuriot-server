package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/global"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/util"
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
		util.LogToConsole(err.Error())
		util.SendJsonResponse(false, constants.ActionSetUser, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	player.Name = data.Name
	player.SetState(constants.StateRoomList)
	global.Players.AddPlayer(player)
	util.SendJsonResponsePlayerID(true, constants.ActionSetUser, player.GetID(), mt, player)
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
