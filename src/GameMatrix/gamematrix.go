package gamematrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"xandzero/src/player"
)

type Cordinate struct {
	X int8
	Y int8
}

const (
	Blank = '-'
)

type GameMatrix struct {
	Size   int8
	Matrix [][]rune
	Move   int8
}

func New(len int8) (*GameMatrix, error) {
	gameMatrix := new(GameMatrix)

	if len < 3 || len > 60 {
		return gameMatrix, errors.New("length of matrix should be between 3 and 6")
	}
	gameMatrix.Size = len
	gameMatrix.Matrix = generateMatrix(len)
	gameMatrix.Move = 0

	return gameMatrix, nil
}

func generateMatrix(len int8) [][]rune {
	var matrix [][]rune = make([][]rune, len)

	for i := int8(0); i < len; i++ {
		var inSlice []rune = make([]rune, len)
		for j := int8(0); j < len; j++ {
			inSlice[j] = '-'
		}
		matrix[i] = inSlice
	}

	return matrix

}

func (gameMatrix *GameMatrix) MarkMatrix(cord *Cordinate, move rune) bool {
	curVal := gameMatrix.Matrix[cord.X-1][cord.Y-1]

	fmt.Println(curVal)

	if curVal == Blank {
		gameMatrix.Matrix[cord.X-1][cord.Y-1] = move
		return true
	}

	return false
}

/*
I want to do something like this printing from my own
box like structure

-----------------------
|   X  |   O   |  X   |
-----------------------
|   X  |   O   |  X   |
-----------------------
|   0  |   X   |  O   |
-----------------------

this is the problem statement for me and I don't want any help from any api
because this will be fun to solve.
*/
func (gameMatrix *GameMatrix) PrintMatrix() {
	const padding int8 = 6
	const borderChar string = "-"

	lineLen := (gameMatrix.Size * padding) + 1
	border := ""

	for i := 0; i < int(lineLen); i++ {
		border += borderChar
	}
	border += "\n"

	resultString := ""

	for _, row := range gameMatrix.Matrix {
		resultString += border
		colString := ""

		for _, elem := range row {
			colString += fmt.Sprintf("|  %c  ", elem)
		}

		colString += "|\n"

		resultString += colString
	}

	resultString += border

	fmt.Println(resultString)
}

// valid cordinate will be in this format
// X,Y  <- so I have to extract a string once
// string got extracted then need to get X and Y
// value and check if value are with int the range or not
func (gameMatrix *GameMatrix) ValidCordPlay(player *player.Player) *Cordinate {
	var cordStr string
	inValid := false
	cord := new(Cordinate)

	for !inValid {
		fmt.Printf("Please Enter Cordinate in X,Y format %v : ", player.Name)
		fmt.Scanf("%s", &cordStr)
		inValid = checkValidRange(player, cordStr, cord, gameMatrix)
	}

	return cord

}

func (gameMatrix *GameMatrix) CheckWinning(player *player.Player) bool {
	// Minimum moves required to win
	const minMovesToWin = 5

	if gameMatrix.Move < minMovesToWin {
		return false
	}

	// Check rows and columns in a single loop
	for i := int8(0); i < gameMatrix.Size; i++ {
		rowCorrVal, colCorrVal := true, true

		for j := int8(0); j < gameMatrix.Size; j++ {
			cellVal := gameMatrix.Matrix[i][j]

			// Check row
			if cellVal != player.Move {
				rowCorrVal = false
			}

			// Check column (use j as the row index)
			if gameMatrix.Matrix[j][i] != player.Move {
				colCorrVal = false
			}

			// Early exit if both row and column are not winning for efficiency
			if !rowCorrVal && !colCorrVal {
				break
			}
		}

		// Winning condition met for either row or column
		if rowCorrVal || colCorrVal {
			return true
		}
	}

	// Check diagonals using separate loops for clarity
	corrVal := true
	for i := int8(0); i < gameMatrix.Size; i++ {
		if gameMatrix.Matrix[i][i] != player.Move {
			corrVal = false
			break
		}
	}
	if corrVal {
		return true
	}

	corrVal = true
	for i := int8(0); i < gameMatrix.Size; i++ {
		if gameMatrix.Matrix[i][gameMatrix.Size-1-i] != player.Move {
			corrVal = false
			break
		}
	}
	if corrVal {
		return true
	}

	return false
}

func checkValidRange(player *player.Player, cordStr string, cord *Cordinate, gameMatrix *GameMatrix) bool {
	// let's get this value using string manupulation and other technique.
	parts := strings.Split(cordStr, ",")

	// there will be at most two parth first one is x and other one is y
	x, err := strconv.Atoi(parts[0])

	if err != nil {
		fmt.Println("Unablt to parse string in cord ")
		return false
	}

	y, err := strconv.Atoi(parts[1])

	if err != nil {
		fmt.Println("Unablt to parse string in cord ")
		return false
	}

	cord.X = int8(x)
	cord.Y = int8(y)

	if err != nil {
		fmt.Println("Unable to parse string in cord 😭")
		return false
	}

	inRange := func(len int8, val int8) bool {
		return val > 0 && val <= len
	}

	xInRange := inRange(gameMatrix.Size, cord.X)
	yInRange := inRange(gameMatrix.Size, cord.Y)

	isValidRange := xInRange && yInRange

	if !isValidRange {
		fmt.Printf("X and Y should be in range from 1 to %d \n", gameMatrix.Size)
		return false
	}

	// check if no other element is present there

	validMatrixPos := gameMatrix.MarkMatrix(cord, player.Move)

	if !validMatrixPos {
		fmt.Println("Postion is already filled choose another postion please ")
		return false
	}

	return true

}
