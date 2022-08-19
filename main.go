package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thaibui2308/p01/breakout"
	"github.com/thaibui2308/p01/breakout/ps"
	// "github.com/thaibui2308/p01/breakout"
)

func main() {
	game := ps.NewParticleSystem(breakout.NewPlayer())
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Breakout")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
