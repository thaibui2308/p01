package breakout

type Collidable interface {
	Location() Coordinate
	H() float64
	W() float64
}
