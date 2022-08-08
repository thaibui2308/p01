package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thaibui2308/p01/breakout"
)

// testing variables
var (
	test_player       *breakout.Player
	test_targetSprite *breakout.TargetSprite
)

func init() {
	// test player
	test_player = breakout.NewPlayer()

	// test target sprite
	test_targetSprite = breakout.NewTargetSprite()
}

type Game struct{}

func (g *Game) Update() error {
	err := test_player.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	test_player.Draw(screen)
	test_targetSprite.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return breakout.SCREEN_WIDTH, breakout.SCREEN_HEIGHT
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Breakout")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
