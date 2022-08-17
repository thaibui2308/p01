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

func (c *Coordinate) Translate(vector *Coordinate) {
	c.SetX(c.GetX() + vector.X)
	c.SetY(c.GetY() + vector.Y)
}

func (c *Coordinate) ReflectOx() {
	c.SetY((-1) * c.GetY())
}

func (c *Coordinate) ReflectOy() {
	c.SetX((-1) * c.GetX())
}

func (c *Coordinate) Add(vector *Coordinate) *Coordinate {
	return &Coordinate{
		c.X + vector.X,
		c.Y + vector.Y,
	}
}
