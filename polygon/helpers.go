package polygon

import (
	"math"

	"github.com/shubhamdwivedii/go-collider/shape"
	"github.com/shubhamdwivedii/go-collider/vector"
)

// Create vertex list of coordinate pairs
func ToVertexList(vertices []shape.Point, pairs ...float64) []shape.Point {
	if len(pairs) <= 1 {
		return vertices
	}
	x, y := pairs[0], pairs[1]
	vertices = append(vertices, shape.Point{X: x, Y: y})
	return ToVertexList(vertices, pairs[2:]...)
}

// True if three vertices lie on a line
func AreCollinear(p, q, r shape.Point) (collinear bool) {
	return math.Abs(vector.Det(q.X-p.X, q.Y-p.Y, r.X-p.X, r.Y-p.Y)) == 0
}

// Remove vertices that lie on a line
func RemoveCollinear(vertices []shape.Point) []shape.Point {
	var filtered []shape.Point
	i, k := vertices[len(vertices)-2], vertices[len(vertices)-1]
	for _, l := range vertices[:len(vertices)-2] {
		if !AreCollinear(i, k, l) {
			filtered = append(filtered, k)
		}
		i, k = k, l
	}
	return filtered
}
