package breakout

func Velocity() *Coordinate {

	return &Coordinate{
		X: INITIAL_BULLET_DIRECTION.X * .4,
		Y: INITIAL_BULLET_DIRECTION.Y * .4,
	}
}
