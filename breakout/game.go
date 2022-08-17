package breakout

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Start        bool
	GameOver     bool
	Score        int
	Player       *Player
	Bullet       *Bullet
	TargetSprite *TargetSprite
}

func NewGame() *Game {
	new_player := NewPlayer()
	new_bullet := NewBullet()
	new_sprite := NewTargetSprite()

	return &Game{
		Start:        false,
		GameOver:     false,
		Score:        0,
		Player:       new_player,
		Bullet:       new_bullet,
		TargetSprite: new_sprite,
	}
}

func (g *Game) Update() error {
	if g.GameOver {
		g.Reset()
	}
	err := g.Player.Update()
	_ = g.Bullet.Update(g)
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
	g.TargetSprite.Draw(screen)
	g.Bullet.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Reset() {
	g.Start = false
	g.GameOver = false
	g.Score = 0
	g.Player = NewPlayer()
	g.Bullet = NewBullet()
	g.TargetSprite = NewTargetSprite()

}
