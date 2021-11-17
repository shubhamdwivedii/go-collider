package circle

import (
	"math"

	. "github.com/shubhamdwivedii/go-collider/shape"
	vector "github.com/shubhamdwivedii/go-collider/vector"
)

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

func (c *Circle) MoveBy(dx, dy float64) {
	c.MoveTo(c.x+dx, c.y+dy)
}

func (c *Circle) Rotation() float64 {
	return c.rt
}

func (c *Circle) Rotate(angle float64) {
	c.rt += angle
}

func (c *Circle) RotateAt(angle, cx, cy float64) {
	c.Rotate(angle)
	rx, ry := vector.Rotate(angle, c.x-cx, c.y-cy)
	c.x, c.y = vector.Add(cx, cy, rx, ry)
}

func (c *Circle) SetRotation(angle, x, y float64) {
	c.rt = angle
}

func (c *Circle) Scale(s float64) {
	c.r = c.r * s
}

func (c *Circle) OutCircle() (cx, cy, r float64) {
	return c.x, c.y, c.r
}

func (c *Circle) BoundingBox() (x1, y1, x2, y2 float64) {
	return c.x - c.r, c.y - c.r, c.x + c.r, c.y + c.r
}

func (c *Circle) Support(dx, dy float64) (x, y float64) {
	ndx, ndy := vector.Normalize(dx, dy)
	mx, my := vector.Mul(c.r, ndx, ndy)
	return vector.Add(c.x, c.y, mx, my)
}

func (c *Circle) CollidesWith(other Shape) (collide bool, dx, dy float64) {
	if c == other {
		return false, 0, 0
	}
	if otherCirc, ok := other.(*Circle); ok {
		ox, oy := other.Center()
		px, py := c.x-ox, c.y-oy
		d := vector.Len2(px, py)
		radii := c.r + otherCirc.r
		if d < radii*radii {
			// if circle overlap, push it out upwards
			if d == 0 {
				return true, 0, radii
			}
			// otherwise push out in best direction
			nx, ny := vector.Normalize(px, py)
			mx, my := vector.Mul(radii-math.Sqrt(d), nx, ny)
			return true, mx, my
		}
		return false, 0, 0
	}
	// else let other shape decide
	return other.CollidesWith(c)
}

func (c *Circle) Contains(x, y float64) (contains bool) {
	return vector.Len2(x-c.x, y-c.y) < c.r*c.r
}

// func (c *Circle) IntersectionsWithRay(x, y, dx, dy float64) (points []Point) {
// 	pcx, pcy := x-c.x, y-c.y

// 	a := vector.Len2(dx, dy)
// 	b := 2 * vector.Dot(dx, dy, pcx, pcy)
// 	d := vector.Len2(pcx, pcy) - c.r*c.r
// 	discr := b*b - 4*a*d

// 	if discr < 0 {
// 		return nil
// 	}

// 	discr = math.Sqrt(discr)
// 	t1, t2 := discr-b, -discr-b
// 	var ts []Point
// 	if t1 >= 0 {
// 		ts = append(ts, t1/(2*a))
// 	}
// 	if t2 >= 0 {
// 		ts = append(ts, t2/(2*a))
// 	}
// }

// func (c *Circle) InterSectsRay(x, y, dx, dy float64) (intersects bool, t float64) {
// 	px, py := c.x
// }

func (c *Circle) Draw(mode string) {
	// yet to be implemented
}
