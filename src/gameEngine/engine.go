package gameengine

import (
	"fmt"
	"os"
	gamematrix "xandzero/src/GameMatrix"
	"xandzero/src/player"
)

func StartGame(matrixSize int8) {

	// initialize a game matrix and two player
	gameMatrix, err := gamematrix.New(matrixSize)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	player1 := player.New()
	player2 := player.New()

	fmt.Println("Game Start now")

	for gameMatrix.Move < int8(8) {
		gameMatrix.Move += 1
		curPlayer := player1

		if gameMatrix.Move%2 == 1 {
			curPlayer = player2
		}

		gameMatrix.ValidCordPlay(curPlayer)

		gameMatrix.PrintMatrix()

		status := gameMatrix.CheckWinning(curPlayer)

		if status {
			fmt.Printf(" %v won the game \n", curPlayer.Name)
			break
		}

	}

	os.Exit(1)
}
