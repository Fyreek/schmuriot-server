package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/schmonk.io/schmuriot-server/config"
	"github.com/schmonk.io/schmuriot-server/models"
	"github.com/schmonk.io/schmuriot-server/sockets"
)

func main() {
	setup()
	sAddress := config.Config.Server.IP + ":" + strconv.Itoa(config.Config.Server.Port)
	router := gin.New()
	if config.Config.Server.Debug {
		router = gin.Default()
	}
	router.GET("/ws", func(c *gin.Context) {
		sockets.InitSocket(c)
		fmt.Println("ws connect")
	})
	models.CreatePlayerList()
	models.CreateRoomList()
	log.Fatal(router.Run(sAddress))
}

func setup() {
	err := config.ReadConfig("server.conf")
	if err != nil {
		fmt.Println("[Failure] Could not read config file")
		os.Exit(1)
	}
	log.SetFlags(0)
}
