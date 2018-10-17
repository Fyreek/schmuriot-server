package models

import (
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"

	"github.com/schmonk.io/schmuriot-server/config"
	"github.com/schmonk.io/schmuriot-server/constants"
)

// Player is struct for a player
type Player struct {
	ID         bson.ObjectId   `json:"id"`
	Name       string          `json:"name"`
	Connection *websocket.Conn `json:"-"`
	State      int             `json:"state"`
	RoomID     *bson.ObjectId  `json:"roomid"`
	PosX       float32         `json:"posx"`
	PosY       float32         `json:"posy"`
	Color      string          `json:"color"`
}

// CreatePlayer creates and returns a new player
func CreatePlayer(con *websocket.Conn) Player {
	p := Player{}
	p.SetID()
	p.SetConnection(con)
	return p
}

// GetID returns the player id as a string
func (p *Player) GetID() string {
	return p.ID.Hex()
}

// SetID sets a random id for the player
func (p *Player) SetID() {
	p.ID = bson.NewObjectId()
}

// SetConnection sets the websocket connection for the player
func (p *Player) SetConnection(con *websocket.Conn) {
	p.Connection = con
}

// SetState sets the state of the player
func (p *Player) SetState(newState int) {
	p.State = newState
}

// SetName sets the name for the player
func (p *Player) SetName(name string) error {
	if len(name) <= config.Config.Room.NameLength {
		p.Name = name
		return nil
	}
	return constants.ErrNameToLong
}

// SetRoom sets the room of the player
func (p *Player) SetRoom(roomid *bson.ObjectId) {
	p.RoomID = roomid
}

func (p *Player) GetRoomID() string {
	return p.RoomID.Hex()
}

func (p *Player) RemoveRoom() {
	p.RoomID = nil
	p.State = constants.StateRoomList
}
