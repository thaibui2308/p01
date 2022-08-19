package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thaibui2308/p01/breakout"
)

func main() {
	game := breakout.NewGame()
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Breakout")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
