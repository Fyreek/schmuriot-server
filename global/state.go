package global

import (
	"sync"

	"github.com/schmonk.io/schmuriot-server/models"
)

// Players is a list of all connected players
var Players models.PlayerList

// Rooms is a list of all available rooms
var Rooms models.RoomList

// CreatePlayerList initializes an empty player list
func CreatePlayerList() {
	pList := models.PlayerList{}
	pList.Players = map[string]*models.Player{}
	pList.Mut = &sync.Mutex{}
	Players = pList
}

// CreateRoomList initializes an empty room list
func CreateRoomList() {
	rList := models.RoomList{}
	rList.Rooms = map[string]*models.Room{}
	rList.Mut = &sync.Mutex{}
	Rooms = rList
}
