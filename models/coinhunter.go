package models

import (
	"fmt"
	"math/rand"

	"github.com/schmonk.io/schmuriot-server/constants"
)

type CoinHunter struct {
	Fields       [][]CoinHunterField        `json:"fields"`
	Rounds       int                        `json:"rounds"`
	Countdown    int                        `json:"countdown"`
	CurrentRound int                        `json:"currentRound"`
	CanReach     map[string][]int           `json:"canReach"`
	Moves        map[string]CoinHunterMoves `json:"moves"`
	Coins        map[string][]int           `json:"coins"`
	State        int                        `json:"state"`
}

type CoinHunterField struct {
	ID     int    `json:"id"`
	Player string `json:"player"`
	Coins  int    `json:"coins"`
}

type CoinHunterMoves struct {
	Player string `json:"player"`
	Field  int    `json:"field"`
}

type CoinHunterMovement struct {
	Player  string `json:"player"`
	Start   int    `json:"start"`
	Move    int    `json:"move"`
	Success bool   `json:"success"`
}

func CreateCoinHunter(round, countdown int) (CoinHunter, error) {
	ch := CoinHunter{}
	ch.Rounds = round
	ch.CurrentRound = 0
	ch.Countdown = countdown
	ch.State = constants.GameStateNotStarted
	return ch, nil
}

func (ch *CoinHunter) GenerateField(players []string) {
	fmt.Print("Anzahl Spieler: ")
	fmt.Println(len(players))
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})
	fields := [][]CoinHunterField{}
	row1 := []CoinHunterField{}
	row2 := []CoinHunterField{}
	row3 := []CoinHunterField{}

	reaching := map[string][]int{}

	c := 0
	for i := 0; i < 3; i++ {
		fmt.Print("Loop 1 durchlauf ")
		fmt.Println(i)
		field := CoinHunterField{}
		field.ID = i + 1
		if i != 1 {
			field.SetCoins()
		} else {
			field.SetPlayer(players[c])
			reaching[players[c]] = ch.CalcReach(field.ID)
			c = c + 1
		}
		fmt.Println(field.ID)
		row1 = append(row1, field)
	}
	fmt.Println(row1)
	fields = append(fields, row1)
	for i := 0; i < 3; i++ {
		fmt.Print("Loop 2 durchlauf ")
		fmt.Println(i)
		field := CoinHunterField{}
		field.ID = 4 + i
		if i == 1 {
			field.SetCoins()
		} else {
			field.SetPlayer(players[c])
			reaching[players[c]] = ch.CalcReach(field.ID)
			c = c + 1
		}
		fmt.Println(field)
		row2 = append(row2, field)
	}
	fmt.Println(row2)
	fields = append(fields, row2)
	for i := 0; i < 3; i++ {
		fmt.Print("Loop 3 durchlauf ")
		fmt.Println(i)
		field := CoinHunterField{}
		field.ID = 7 + i
		if i != 1 {
			field.SetCoins()
		} else {
			field.SetPlayer(players[c])
			reaching[players[c]] = ch.CalcReach(field.ID)
			c = c + 1
		}
		fmt.Println(field)
		row3 = append(row3, field)
	}
	fmt.Println(row3)
	fields = append(fields, row3)
	ch.Fields = fields
	ch.CanReach = reaching
}

func (ch *CoinHunter) CalcReach(pos int) []int {
	reach := []int{}
	switch pos {
	case 1:
		reach = []int{2, 4}
		break
	case 2:
		reach = []int{1, 3, 5}
		break
	case 3:
		reach = []int{2, 6}
		break
	case 4:
		reach = []int{1, 5, 7}
		break
	case 5:
		reach = []int{2, 4, 6, 8}
		break
	case 6:
		reach = []int{3, 5, 9}
		break
	case 7:
		reach = []int{4, 8}
		break
	case 8:
		reach = []int{5, 7, 9}
		break
	case 9:
	default:
		reach = []int{6, 8}
		break
	}
	return reach
}

func (chf *CoinHunterField) SetCoins() {
	chf.Coins = rand.Intn(4) + 1
}

func (chf *CoinHunterField) SetPlayer(player string) {
	chf.Player = player
}
