package player

import (
	"fmt"
)

var playerNo int = 0

type Player struct {
	Name  string
	Move  rune
	Score int8
}

func New() *Player {
	playerNo += 1
	player := new(Player)
	fmt.Printf("Enter %v player name : ", playerNo)
	fmt.Scanf("%v", &player.Name)
	if playerNo == 1 {
		player.Move = 'O'
	} else {
		player.Move = 'X'
	}

	return player
}
