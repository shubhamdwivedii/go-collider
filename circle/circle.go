package circle

type Circle struct {
	x  float64
	y  float64
	r  float64
	rt float64
}

func New(x, y, r float64) *Circle {
	return &Circle{x: x, y: y, r: r}
}

func (c *Circle) Center() (x, y float64) {
	return c.x, c.y
}

func (c *Circle) Move(x, y float64) {
	c.x += x
	c.y += y
}

func (c *Circle) MoveTo(x, y float64) {
	cx, cy := c.Center()
	c.Move(x-cx, y-cy)
}

func (c *Circle) Rotation() float64 {
	return c.rt
}

func (c *Circle) Rotate(angle float64) {
	c.rt += angle
}

func (c *Circle) SetRotation(angle, x, y float64) {
	c.rt = angle
}

func (c *Circle) OutCircle() (cx, cy, r float64) {
	return c.x, c.y, c.r
}

func (c *Circle) BoundingBox() (x1, y1, x2, y2 float64) {
	return c.x - c.r, c.y - c.r, c.x + c.r, c.y + c.r
}
