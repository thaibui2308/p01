package breakout

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Start          bool
	GameOver       bool
	Score          int
	LifeLeft       int
	Player         *Player
	Bullet         *Bullet
	TargetSprite   *TargetSprite
	ParticleSystem []*ParticleSystem
}

func NewGame() *Game {
	new_player := NewPlayer()
	new_bullet := NewBullet()
	new_sprite := NewTargetSprite()
	p_system := make([]*ParticleSystem, 0)

	return &Game{
		Start:          false,
		GameOver:       false,
		LifeLeft:       4,
		Score:          0,
		Player:         new_player,
		Bullet:         new_bullet,
		TargetSprite:   new_sprite,
		ParticleSystem: p_system,
	}
}

func (g *Game) Update() error {
	if g.GameOver {
		g.Reset()
	}
	err := g.Player.Update()
	_ = g.Bullet.Update(g)
	for _, v := range g.ParticleSystem {
		v.Update()
	}
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
	g.TargetSprite.Draw(screen)
	g.Bullet.Draw(screen)
	for _, v := range g.ParticleSystem {
		v.Draw(screen)
	}
	g.Score = g.TargetSprite.HitSprite * 2
	toScreen := "Score: " + strconv.Itoa(g.Score) + "\n" +
		"Heart: " + strconv.Itoa(g.LifeLeft)
	ebitenutil.DebugPrint(screen, toScreen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) AppendSystem(s *ParticleSystem) {
	g.ParticleSystem = append(g.ParticleSystem, s)
}

func (g *Game) Reset() {
	g.Start = false
	g.GameOver = false
	g.Score = 0
	g.LifeLeft = 4
	g.Player = NewPlayer()
	g.Bullet = NewBullet()
	g.TargetSprite = NewTargetSprite()

}
