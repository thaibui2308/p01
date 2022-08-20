package breakout

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleSystem struct {
	Emitter  Collidable
	Paricles []*Particle
}

func NewParticleSystem(emitter Collidable) *ParticleSystem {
	particles := make([]*Particle, 0)
	for i := 0; i < 12; i++ {
		tmp := NewParticle(emitter)
		particles = append(particles, tmp)
		if i >= 1 {
			if i%2 == 1 {
				particles[i].Velocity.SetY(particles[i-1].Velocity.GetY() * (-1))
			} else {
				particles[i].Velocity.SetX(particles[i-1].Velocity.GetX() * (-1))
			}
		}
	}
	spacingX := emitter.W() / 9
	spacingY := emitter.H() / 3
	for i := 0; i < len(particles); i++ {
		new_X := particles[i].Location.GetX() + float64(i%12)*spacingX
		new_Y := particles[i].Location.GetY() + float64(i%3)*spacingY
		particles[i].Location.SetX(new_X)
		particles[i].Location.SetY(new_Y)

	}
	return &ParticleSystem{
		Emitter:  emitter,
		Paricles: particles,
	}
}

func (s *ParticleSystem) Update() {
	for _, v := range s.Paricles {
		_ = v.Update()

	}

}

func (s *ParticleSystem) Draw(dst *ebiten.Image) {
	for _, v := range s.Paricles {
		v.Draw(dst)
	}
}

func (s *ParticleSystem) Remove(index int) {
	s.Paricles[index] = s.Paricles[len(s.Paricles)-1]
	s.Paricles = s.Paricles[:len(s.Paricles)-1]
}
