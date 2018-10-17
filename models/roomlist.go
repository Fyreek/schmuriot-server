package models

import (
	"encoding/json"
	"sync"
)

// RoomList is a struct for a list of Rooms
type RoomList struct {
	Rooms map[string]*Room `json:"rooms"`
	Mut   *sync.Mutex      `json:"-"`
}

// Marshal makes json from the object
func (rl *RoomList) Marshal() ([]byte, error) {
	rl.Mut.Lock()
	bytes, err := json.Marshal(rl)
	rl.Mut.Unlock()
	return bytes, err
}

// GetRooms returns all available rooms
func (rl *RoomList) GetRooms() map[string]*Room {
	rl.Mut.Lock()
	rooms := rl.Rooms
	rl.Mut.Unlock()
	return rooms
}

// AddRoom adds a room to the global room list
func (rl *RoomList) AddRoom(room *Room) {
	rl.Mut.Lock()
	rl.Rooms[room.GetID()] = room
	rl.Mut.Unlock()
}

// RemoveRoom removes a room from the global room list
func (rl *RoomList) RemoveRoom(room *Room) {
	rl.Mut.Lock()
	for _, player := range room.Players {
		player.RemoveRoom()
	}
	delete(rl.Rooms, room.GetID())
	rl.Mut.Unlock()
}

// GetRoom returns a room from the global room list
func (rl *RoomList) GetRoom(rID string) *Room {
	rl.Mut.Lock()
	r := rl.Rooms[rID]
	rl.Mut.Unlock()
	return r
}

// ModifyRoom modifies an existing room in the player list
func (rl *RoomList) ModifyRoom(room *Room) {
	rl.Mut.Lock()
	rl.Rooms[room.GetID()] = room
	rl.Mut.Unlock()
}
