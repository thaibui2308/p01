package breakout

import "math"

type Coordinate struct {
	X float64
	Y float64
}

// Getter
func (c *Coordinate) GetX() float64 { return c.X }
func (c *Coordinate) GetY() float64 { return c.Y }

// Setter
func (c *Coordinate) SetX(x float64) { c.X = x }
func (c *Coordinate) SetY(y float64) { c.Y = y }

func (c *Coordinate) BounceUpward(angle float64, distance float64) {
	new_x := distance/math.Sqrt(distance*distance+1) + c.GetX()
	new_y := math.Tan(angle)*(c.GetX()-new_x) + c.GetY()

	c.SetX(new_x)
	c.SetY(new_y)
}
