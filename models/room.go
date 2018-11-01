package models

import (
	"fmt"
	"sync"
	"time"

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
	Game      *CoinHunter        `json:"game"`
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
	} else if action == constants.ActionStartRound {
		SendJsonResponseGame(status, action, 1, player)
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

func (r *Room) GetPlayerList() []string {
	playerList := []string{}
	for element := range r.Players {
		playerList = append(playerList, element)
	}
	return playerList
}

func (r *Room) CheckAllReady() bool {
	playerCount := len(r.Players)
	if playerCount < r.Slots {
		return false
	}
	for element := range r.Players {
		player := r.Players[element]
		if !player.Ready {
			return false
		}
	}
	return true
}

func (r *Room) StartGame() {
	// ready := r.CheckAllReady()
	for element := range r.Players {
		p := r.Players[element]
		p.SetState(constants.StateInGame)
	}
	countdown := 15
	rounds := 5
	game, _ := CreateCoinHunter(rounds, countdown)
	fmt.Print("Game info ")
	fmt.Println(game.Rounds)
	r.Game = &game
	r.Game.CurrentRound = 1
	r.StartRound()
}

func (r *Room) StartRound() {
	if r.Game != nil {
		r.Game.State = constants.GameStatePlaying
		r.Game.Moves = map[string]CoinHunterMoves{}
		r.Game.GenerateField(r.GetPlayerList())
		r.SendToAllPlayers(true, constants.ActionStartRound, "", nil)
		gameTimer := time.NewTimer(time.Duration(r.Game.Countdown) * time.Second)
		<-gameTimer.C
		fmt.Println("gameTimer expired... Triggering round end")
		r.EndRound()
	}
}

func (r *Room) EndRound() {
	if r.Game != nil {
		r.Game.State = constants.GameStateEnd

		// Calculate actual movement
		// If two players want to get to the same field, cancel
		// Collect coins
		// Send response

		validMovements := map[string]int{}
		for pID := range r.Game.Moves {
			pMove := r.Game.Moves[pID]
			collision := false
			for innerPID := range r.Game.Moves {
				innerPMove := r.Game.Moves[innerPID]
				if innerPMove.Player != pMove.Player {
					if innerPMove.Field == pMove.Field {
						collision = true
					}
				}
			}
			if !collision {
				validMovements[pMove.Player] = pMove.Field
			}
		}
		startPos := map[string]int{}
		coinPos := map[int]int{}
		for _, outerField := range r.Game.Fields {
			for _, innerField := range outerField {
				if innerField.Player != "" {
					startPos[innerField.Player] = innerField.ID
				}
				if innerField.Coins != 0 {
					coinPos[innerField.ID] = innerField.Coins
				}
			}
		}
		movements := map[string]CoinHunterMovement{}
		for p := range r.Players {
			m := CoinHunterMovement{}
			m.Player = r.Players[p].GetID()
			m.Start = startPos[p]
			if vMove, ok := validMovements[p]; ok {
				m.Move = vMove
				m.Success = true
				curCoins := r.Game.Coins[m.Player]
				r.Game.Coins[m.Player] = curCoins + coinPos[vMove]
			} else {
				if r.Game.Moves[p].Field != 0 {
					m.Move = r.Game.Moves[p].Field
				} else {
					m.Move = startPos[p]
				}
				m.Success = false
			}
			movements[p] = m
		}
		for p := range r.Players {
			SendJsonResponseMovement(movements, 1, r.Players[p])
			SendJsonResponseCoins(r.Game.Coins, 1, r.Players[p])
		}

		// Add coins
		// Leaderboards response

		if r.Game.CurrentRound == r.Game.Rounds {
			// Game End
			// Show end screen
			waitTimer := time.NewTimer(time.Duration(5) * time.Second)
			<-waitTimer.C
			// Go back to lobby
		} else {
			// Next Round
			r.Game.CurrentRound = r.Game.CurrentRound + 1
			waitTimer := time.NewTimer(time.Duration(5) * time.Second)
			<-waitTimer.C
			r.StartRound()
		}
	}
}
