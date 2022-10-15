package main

import (
	"crypto/rand"
	"fmt"
)

type game struct {
	gameId         string
	playerOneName  string
	playerTwoName  string
	playerOneScore int
	playerTwoScore int
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
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
	fmt.Printf("I'm running this")
	a := tokenGenerator()
	newGame := newGame(a)
	fmt.Printf("%v", newGame)
}
