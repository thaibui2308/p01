package breakout

import (
	"math/rand"
	"time"
)

func RandVelocity() *Coordinate {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	scaleX := r.Float64()*2 - 1.0
	scaleY := r.Float64()*2 - 1.0

	return &Coordinate{
		X: INITIAL_BULLET_DIRECTION.X * scaleX * .4,
		Y: INITIAL_BULLET_DIRECTION.Y * scaleY * .4,
	}
}
