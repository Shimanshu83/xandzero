/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	gameengine "xandzero/src/gameEngine"

	"github.com/spf13/cobra"
)

var size int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xandzero",
	Short: "Play a game of x and zero (tic-tac-toe",
	Long: `xandzero is a CLI implementation of the classic game tic-tac-toe, also known as x and zero. 
	It offers a customizable gaming experience, allowing players to choose the size of the game board.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting xandzero game with board size %dx%d\n", size, size)
		gameengine.StartGame(int8(size))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().IntVarP(&size, "size", "s", 3, "Size of the game board (between 3 and 8)")

}
