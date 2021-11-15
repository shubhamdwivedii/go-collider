package shape

type Point struct {
	X float64
	Y float64
}

type Shape interface {
	MoveTo(x, y float64)
	MoveBy(dx, dy float64)
	Center() (cx, cy float64)

	Rotation() (rot float64)    // returns rotation
	Rotate(r float64)           // rotates by centre
	RotateAt(r, cx, cy float64) // cx,cy origin of rotation
	SetRotation(r, cx, cy float64)

	Scale(f int) // scales shape by factor

	OutCircle() (x, y, r float64)          // A Circle that fully encloses the shape
	BoundingBox() (x1, y1, x2, y2 float64) // Bounding Box of Shape

	Draw(mode string)                      // "fill" or "line"
	Support(dx, dy float64) (x, y float64) // Get furthest vertex of the Shape with respect to the direction dx,dy
	// Used for collision detection, can be used for shadow/lighting too

	CollidesWith(other Shape) (collide bool, dx, dy float64) // Checks collision with other shape.
	// Also returns separating vector

	Contains(x, y float64) (contains bool) // Checks if shape contains point.

	IntersectionsWithRay(x, y, dx, dy float64) []Point // Checks if shape intersects with ray. Points of intersecions are returned
}

// Custom Shapes must implement these Move, Rotate, Scale, BoundingBox, CollidesWith
