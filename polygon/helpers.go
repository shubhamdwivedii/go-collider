package polygon

import (
	"math"

	"github.com/shubhamdwivedii/go-collider/shape"
	"github.com/shubhamdwivedii/go-collider/vector"
)

// Create vertex list of coordinate pairs
func ToVertexList(vertices []float64, pairs ...float64) []float64 {
	if len(pairs) <= 1 {
		return vertices
	}
	x, y := pairs[0], pairs[1]
	vertices = append(vertices, x, y)
	return ToVertexList(vertices, pairs[2:]...)
}

// True if three vertices lie on a line
func AreCollinear(p, q, r shape.Point) (collinear bool) {
	return math.Abs(vector.Det(q.X-p.X, q.Y-p.Y, r.X-p.X, r.Y-p.Y)) == 0
}
