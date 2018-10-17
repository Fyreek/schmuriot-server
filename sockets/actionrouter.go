package sockets

import (
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/util"
)

// ActionRouter handles every action per player and calls corresponding functions
func ActionRouter(player *models.Player, message []byte, mt int) {
	baseAction := models.BaseAction{}
	err := baseAction.Unmarshal(message)
	if err != nil {
		util.SendJsonResponse(false, constants.ActionNone, constants.ErrInvalidJSON.Error(), mt, player)
		return
	}
	switch baseAction.Action {
	default:
		util.LogToConsole("Not implemented")
		util.SendJsonResponse(false, constants.ActionNone, constants.ErrNotImplemented.Error(), mt, player)
	}
}
