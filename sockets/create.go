package sockets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/schmonk.io/schmuriot-server/config"
	"github.com/schmonk.io/schmuriot-server/constants"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/utils"
)

var upgrader = websocket.Upgrader{}

//InitSocket upgrades and initializes the socket connection
func InitSocket(c *gin.Context) {
	if !config.Config.Server.CORS {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	}

	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	utils.LogToConsole("socket upgrade request")
	if err != nil {
		utils.LogToConsole("socket upgrade failed:", err)
		return
	}

	gSlots := models.Players.GetPlayerCount()
	if gSlots < config.Config.Server.Slots {
		createSocketPlayer(con)
	} else {
		con.WriteMessage(1, []byte("Slots exceeded"))
		con.Close()
		return
	}
}

// createSocketPlayer creates the player for the websocket connection
func createSocketPlayer(con *websocket.Conn) {
	utils.LogToConsole("Connected Players:", models.Players.GetPlayerCount())
	player := models.CreatePlayer(con)
	player.State = constants.StateUndefined
	defer con.Close()
	PlayerLoop(&player)
}
