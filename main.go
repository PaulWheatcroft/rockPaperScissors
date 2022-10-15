package main

import "fmt"

type game struct {
	gameId         string
	playerOneName  string
	playerTwoName  string
	playerOneScore int
	playerTwoScore int
}

func newGame(gameId string) *game {

	g := game{gameId: gameId}
	g.playerOneName = "Paul"
	g.playerTwoName = ""
	g.playerOneScore = 0
	g.playerTwoScore = 0
	return &g
}

func main() {
	fmt.Printf("I'm running")
	newGame := newGame("abc")
	fmt.Printf("%v", newGame)
}
