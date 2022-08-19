package breakout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	C         Coordinate
	Height    float64
	Width     float64
	Moving    bool
	Direction int
}

func NewPlayer() *Player {
	player_coordinate := &Coordinate{
		X: PLAYER_X,
		Y: PLAYER_Y,
	}

	return &Player{
		C:         *player_coordinate,
		Height:    PLAYER_HEIGHT,
		Width:     PLAYER_WIDTH,
		Moving:    false,
		Direction: 0,
	}
}

func (p *Player) Location() Coordinate {
	return p.C
}

func (p *Player) H() float64 { return p.Height }

func (p *Player) W() float64 { return p.Width }

func (p *Player) Move() {
	if p.Moving {
		return
	} else {
		p.Moving = true
	}
}

func (p *Player) Stop() {
	if !p.Moving {
		return
	} else {
		p.Moving = false
	}
}

func (p *Player) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, float64(p.C.X), float64(p.C.Y), p.Width, p.Height, color.RGBA{0, 0, 0xff, 0xff})
}

// Update the player's position if there's any keystroke
func (p *Player) Update() error {
	// Now the player only stops when you release the arrow keys
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		p.Direction = 1
		p.Move()

	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		p.Direction = -1
		p.Move()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		p.Stop()
	}

	if p.Moving {
		new_pos := p.C.X + float64(p.Direction)*3
		if new_pos >= SCREEN_WIDTH-PLAYER_WIDTH {
			new_pos = SCREEN_WIDTH - PLAYER_WIDTH
		} else if new_pos <= 0 {
			new_pos = 0
		}
		p.C.SetX(new_pos)
	}

	return nil
}
