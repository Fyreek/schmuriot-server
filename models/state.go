package models

import (
	"sync"
)

// Players is a list of all connected players
var Players PlayerList

// Rooms is a list of all available rooms
var Rooms RoomList

// CreatePlayerList initializes an empty player list
func CreatePlayerList() {
	pList := PlayerList{}
	pList.Players = map[string]*Player{}
	pList.Mut = &sync.Mutex{}
	Players = pList
}

// CreateRoomList initializes an empty room list
func CreateRoomList() {
	rList := RoomList{}
	rList.Rooms = map[string]*Room{}
	rList.Mut = &sync.Mutex{}
	Rooms = rList
}
