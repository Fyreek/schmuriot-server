package models

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"github.com/schmonk.io/schmuriot-server/config"
	"github.com/schmonk.io/schmuriot-server/constants"
)

// Room is struct for a room
type Room struct {
	ID        bson.ObjectId      `json:"_id"`
	Name      string             `json:"name"`
	Pass      string             `json:"-"`
	Protected bool               `json:"protected"`
	Map       string             `json:"map"`
	Slots     int                `json:"slots"`
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
	r.SetPass(pass)
	r.SetOwner(owner)
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

// SetID sets a random id for the room
func (r *Room) SetID() {
	r.ID = bson.NewObjectId()
}

// SetName sets the name for the room
func (r *Room) SetName(name string) error {
	if len(name) < 3 {
		return constants.ErrNameToShort
	} else if len(name) <= config.Config.Game.NameLength {
		r.Name = name
		return nil
	}
	return constants.ErrNameToLong
}

// SetPass sets the password for the room
func (r *Room) SetPass(pass string) {
	r.Pass = pass
	if pass == "" {
		r.Protected = false
		return
	}
	r.Protected = true
}

// SetOwner sets the owner id of the room
func (r *Room) SetOwner(oID string) {
	r.Owner = oID
}

// SetSlots sets the number of slots available for the room
func (r *Room) SetSlots(quantity int) error {
	if quantity < 2 {
		return constants.ErrToLessSlots
	} else if quantity <= config.Config.Game.SlotsPerRoom {
		r.Slots = quantity
		return nil
	}
	return constants.ErrToManySlots
}

// AddPlayer adds a new player to the room
func (r *Room) AddPlayer(player *Player, pass string) error {
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
	r.Mut.Unlock()
	return nil
}

// RemovePlayer removes the player from the room
func (r *Room) RemovePlayer(player *Player) error {
	r.Mut.Lock()
	if player.GetID() == r.Owner {
		delete(r.Players, player.GetID())
		for _, p := range r.Players {
			r.Owner = p.GetID()
			break
		}
		r.Mut.Unlock()
		return nil
	}
	delete(r.Players, player.GetID())
	if len(r.Players) <= 0 {
		r.Mut.Unlock()
		//delete the room (difficult because i have to use a global variable, which references models and creates a circular reference)
		//probably do it from the action
		return constants.ErrNoPlayer
	}
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

func (r *Room) Lock() {
	r.Mut.Lock()
}

func (r *Room) Unlock() {
	r.Mut.Unlock()
}
