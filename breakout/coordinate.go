package breakout

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
