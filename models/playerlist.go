package models

import (
	"sync"

	"github.com/schmonk.io/schmuriot-server/constants"
)

// PlayerList is a struct for a list of BasePlayers
type PlayerList struct {
	Players map[string]*Player `json:"players"`
	Mut     *sync.Mutex        `json:"-"`
}

// GetPlayerCount returns the number of players on the server
func (pl *PlayerList) GetPlayerCount() int {
	pl.Mut.Lock()
	count := len(pl.Players)
	pl.Mut.Unlock()
	return count
}

// AddPlayer adds a new player to the global player list
func (pl *PlayerList) AddPlayer(player *Player) {
	pl.Mut.Lock()
	_, ok := pl.Players[player.GetID()]
	if !ok {
		pl.Players[player.GetID()] = player
	}
	pl.Mut.Unlock()
}

// RemovePlayer removes a player from the global player list
func (pl *PlayerList) RemovePlayer(player *Player) {
	pl.Mut.Lock()
	delete(pl.Players, player.GetID())
	pl.Mut.Unlock()
}

// GetPlayer returns a player from the global player list
func (pl *PlayerList) GetPlayer(pID string) *Player {
	pl.Mut.Lock()
	bp := pl.Players[pID]
	pl.Mut.Unlock()
	return bp
}

// ModifyPlayer modifies an existing player in the player list
func (pl *PlayerList) ModifyPlayer(player *Player) {
	pl.Mut.Lock()
	pl.Players[player.GetID()] = player
	pl.Mut.Unlock()
}

//SendToAllPlayers sends a message to all players
func (pl *PlayerList) SendToAllPlayers(status bool, action string, message interface{}, player *Player) {
	for _, p := range pl.Players {
		if player != nil {
			if player.GetID() == p.GetID() {
				continue
			}
		}
		pl.SendToPlayer(status, action, message, p)
	}
}

//SendToPlayer sends a message to a specific player
func (pl *PlayerList) SendToPlayer(status bool, action string, message interface{}, player *Player) {
	if action == constants.ActionGetRooms && player.State == constants.StateRoomList {
		SendJsonResponseRoomList(status, action, Rooms.Rooms, 1, player)
		return
	}
	return
}
