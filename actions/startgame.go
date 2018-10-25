package actions

import (
	"encoding/json"
	"fmt"

	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

type StartGameAction struct {
	models.BaseAction
	Rounds    int `json:"rounds"`
	Countdown int `json:"countdown"`
}

func StartGame(player *models.Player, message []byte, mt int) {
	if player.State != constants.StateLobby {
		models.SendJsonResponse(false, constants.ActionToggleReady, constants.ErrActionNotPossible.Error(), mt, player)
		return
	}
	data := StartGameAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		utils.LogToConsole(err.Error())
		models.SendJsonResponse(false, constants.ActionJoinRoom, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	// Check if all players are ready
	fmt.Print("Game Rounds: ")
	fmt.Println(data.Rounds)
	fmt.Print("Game Countdown: ")
	fmt.Println(data.Countdown)
	r := models.Rooms.GetRoom(player.GetRoomID())
	if r != nil {
		playerList := []string{}
		for element := range r.Players {
			playerList = append(playerList, element)
		}
		game, _ := models.CreateCoinHunter(playerList, data.Rounds, data.Countdown)
		fmt.Print("Game info ")
		fmt.Println(game.Rounds)
		r.Game = &game
		r.SendToAllPlayers(true, constants.ActionStartGame, "", nil)
		return
	}
	models.SendJsonResponse(false, constants.ActionDeleteRoom, constants.ErrRoomNotFound.Error(), mt, player)
}
