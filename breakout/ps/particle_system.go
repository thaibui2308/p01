package ps

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thaibui2308/p01/breakout"
)

type ParticleSystem struct {
	Emitter  breakout.Collidable
	Paricles []Particle
}

func NewParticleSystem(emitter breakout.Collidable) *ParticleSystem {
	particles := make([]Particle, 0)
	for i := 0; i < 4; i++ {
		particles = append(particles, *NewParticle(emitter))
	}
	fmt.Println(particles)
	return &ParticleSystem{
		Emitter:  emitter,
		Paricles: particles,
	}
}

func (s *ParticleSystem) Update() error {
	for _, v := range s.Paricles {
		_ = v.Update()
	}
	return nil
}

func (s *ParticleSystem) Draw(dst *ebiten.Image) {
	for _, v := range s.Paricles {
		v.Draw(dst)
	}
}

func (s *ParticleSystem) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return breakout.SCREEN_WIDTH, breakout.SCREEN_HEIGHT
}
