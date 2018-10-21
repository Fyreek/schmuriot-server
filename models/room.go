package models

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"github.com/schmonk.io/schmuriot-server/config"
	"github.com/schmonk.io/schmuriot-server/constants"
)

// Room is struct for a room
type Room struct {
	ID        bson.ObjectId      `json:"id"`
	Name      string             `json:"name"`
	Pass      string             `json:"-"`
	Protected bool               `json:"protected"`
	Slots     int                `json:"slots"`
	Mode      string             `json:"mode"`
	Game      string             `json:"game"`
	Owner     string             `json:"owner"`
	Players   map[string]*Player `json:"players"`
	Mut       *sync.Mutex        `json:"-"`
}

// CreateRoom creates and returns a new room
func CreateRoom(name, pass, owner string, slots int) (Room, error) {
	r := Room{}
	r.SetID()
	err := r.SetName(name)
	if err != nil {
		return r, err
	}
	err = r.SetPass(pass)
	if err != nil {
		return r, err
	}
	r.SetOwner(owner)
	if slots == 0 {
		slots = config.Config.Room.MaxSlots
	}
	err = r.SetSlots(slots)
	if err != nil {
		return r, err
	}
	r.Players = map[string]*Player{}
	r.Mut = &sync.Mutex{}
	return r, nil
}

// GetID returns the room id as a string
func (r *Room) GetID() string {
	return r.ID.Hex()
}

// GetID returns the room id as a string
func (r *Room) GetBsonID() *bson.ObjectId {
	return &r.ID
}

// SetID sets a random id for the room
func (r *Room) SetID() {
	r.ID = bson.NewObjectId()
}

// SetName sets the name for the room
func (r *Room) SetName(name string) error {
	if len(name) < config.Config.Room.MinNameLength {
		return constants.ErrNameToShort
	} else if len(name) <= config.Config.Room.MaxNameLength {
		r.Name = name
		return nil
	}
	return constants.ErrNameToLong
}

// SetPass sets the password for the room
func (r *Room) SetPass(pass string) error {
	if len(pass) < config.Config.Room.MinPassLength {
		if pass == "" {
			r.Pass = pass
			r.Protected = false
			return nil
		}
		return constants.ErrPassToShort
	} else if len(pass) <= config.Config.Room.MaxPassLength {
		r.Pass = pass
		r.Protected = true
		return nil
	}
	return constants.ErrPassToLong
}

// SetOwner sets the owner id of the room
func (r *Room) SetOwner(oID string) {
	r.Owner = oID
}

// SetSlots sets the number of slots available for the room
func (r *Room) SetSlots(quantity int) error {
	if quantity < config.Config.Room.MinSlots {
		return constants.ErrToLessSlots
	} else if quantity <= config.Config.Room.MaxSlots {
		r.Slots = quantity
		return nil
	}
	return constants.ErrToManySlots
}

// AddPlayer adds a new player to the room
func (r *Room) AddPlayer(player *Player, pass string, first bool) error {
	r.Mut.Lock()
	if len(r.Players) >= r.Slots {
		r.Mut.Unlock()
		return constants.ErrNoSlots
	}
	if r.Pass != "" {
		if r.Pass != pass {
			r.Mut.Unlock()
			return constants.ErrPassWrong
		}
	}
	r.Players[player.GetID()] = player
	player.SetRoom(r.GetBsonID())
	player.SetState(constants.StateLobby)
	r.SendToAllPlayers(true, constants.ActionGetRoom, "", player)
	if !first {
		Players.SendToAllPlayers(true, constants.ActionGetRooms, "", player)
	}
	r.Mut.Unlock()
	return nil
}

// RemovePlayer removes the player from the room
func (r *Room) RemovePlayer(player *Player) error {
	r.Mut.Lock()
	if len(r.Players) > 1 {
		if player.GetID() == r.Owner {
			delete(r.Players, player.GetID())
			for _, p := range r.Players {
				r.Owner = p.GetID()
				break
			}
			player.RoomID = nil
			player.SetState(constants.StateRoomList)
			r.SendToAllPlayers(true, constants.ActionGetRoom, "", player)
			Players.SendToAllPlayers(true, constants.ActionGetRooms, "", player)
			r.Mut.Unlock()
			return nil
		}
	}
	delete(r.Players, player.GetID())
	if len(r.Players) <= 0 {
		player.RoomID = nil
		player.SetState(constants.StateRoomList)
		r.Mut.Unlock()
		Rooms.RemoveRoom(r, true)
		return nil
	}
	player.RoomID = nil
	player.SetState(constants.StateRoomList)
	r.SendToAllPlayers(true, constants.ActionGetRoom, "", player)
	Players.SendToAllPlayers(true, constants.ActionGetRooms, "", player)
	r.Mut.Unlock()
	return nil
}

// GetPlayerCount returns the number of players in the room
func (r *Room) GetPlayerCount() int {
	r.Mut.Lock()
	count := len(r.Players)
	r.Mut.Unlock()
	return count
}

//SendToAllPlayers sends a message to all players
func (r *Room) SendToAllPlayers(status bool, action string, message interface{}, player *Player) {
	for _, p := range r.Players {
		if player != nil {
			if player.GetID() == p.GetID() {
				continue
			}
		}
		r.SendToPlayer(status, action, message, p)
	}
}

//SendToPlayer sends a message to a specific player
func (r *Room) SendToPlayer(status bool, action string, message interface{}, player *Player) {
	if action == constants.ActionGetRoom {
		SendJsonResponseRoom(status, constants.ActionGetRoom, 1, player)
	} else if action == constants.ActionChat {
		str, _ := message.(string)
		SendJsonResponseChat(status, action, str, 1, player)
	} else {
		str, ok := message.(string)
		if ok {
			SendJsonResponse(status, action, str, 1, player)
		} else {
			SendJsonResponse(status, action, constants.ErrUnknownMessageType.Error(), 1, player)
		}
	}
}
