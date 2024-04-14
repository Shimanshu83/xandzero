package util

import "fmt"

func GetInput(mesasge string, format string, valPtr *string) {
	fmt.Print("Enter first player 1 : ")
	fmt.Scanf(format, valPtr)
}
