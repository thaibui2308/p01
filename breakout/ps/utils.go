package ps

import (
	"math/rand"
	"time"

	"github.com/thaibui2308/p01/breakout"
)

func RandVelocity() *breakout.Coordinate {
	rand.Seed(time.Now().UnixNano())

	scaleX := rand.Float64()*2 - 1.0
	scaleY := rand.Float64()*2 - 1.0

	return &breakout.Coordinate{
		X: breakout.INITIAL_BULLET_DIRECTION.X * scaleX,
		Y: breakout.INITIAL_BULLET_DIRECTION.Y * scaleY,
	}
}
