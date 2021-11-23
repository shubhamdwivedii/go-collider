package polygon

import (
	"fmt"
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

// Get index of rightmost vertex (for testing orientation)
func GetIndexOfLeftmost(vertices []shape.Point) int {
	idx := 0
	for i := 1; i < len(vertices); i++ {
		if vertices[i].X < vertices[idx].X {
			idx = i
		}
	}
	return idx
}

// Returns true if three points make a counter clockwise turn
func CCW(p, q, r shape.Point) bool {
	return vector.Det(q.X-p.X, q.Y-p.Y, r.X-p.X, r.Y-p.Y) >= 0
}

// Test wether a and b lie on the same side of line c->d
func OnSameSide(a, b, c, d shape.Point) bool {
	px, py := d.X-c.X, d.Y-c.Y
	l := vector.Det(px, py, a.X-c.X, a.Y-c.Y)
	m := vector.Det(px, py, b.X-c.X, b.Y-c.Y)
	return l*m >= 0
}

func PointInTriangle(p, a, b, c shape.Point) bool {
	return OnSameSide(p, a, b, c) && OnSameSide(p, b, a, c) && OnSameSide(p, c, a, b)

}

func samePoint(a, b shape.Point) bool {
	if a.X == b.X && a.Y == b.Y {
		return true
	}
	return false
}

// Test wheteher any point in vertices (but pqr) lies in the triangle pqr
func AnyPointInTriangle(vertices []shape.Point, p, q, r shape.Point) bool {
	for _, point := range vertices {
		if !samePoint(point, p) && !samePoint(point, q) && !samePoint(point, r) && PointInTriangle(point, p, q, r) {
			return true
		}
	}
	return false
}

// Test is the triangle pqr is an "ear" of the polygon
func IsEar(p, q, r shape.Point, vertices []shape.Point) bool {
	return CCW(p, q, r) && !AnyPointInTriangle(vertices, p, q, r)
}

func SegementsInterset(a, b, p, q shape.Point) bool {
	return !(OnSameSide(a, b, p, q) || (OnSameSide(p, q, a, b)))
}

// Returns starting/ending indices of shared edge,
// ie: if p and q share the edge with indices p1, p2 of p and q1,q2 of q, the return value is p1, q2
func GetSharedEdge(p, q []shape.Point) (e1, e2 int) {
	pIndices := make(map[string]int)

	// Record indices of vertices in p by their corrdinates
	for i := 0; i < len(p); i++ {
		pIndices[fmt.Sprintf("%v-%v", p[i].X, p[i].Y)] = i
	}

	// Iterate over all edges in q. If both endpoints of that edge are in p as well,
	// return the indices of the starting vertex
	idx := len(q) - 1
	for k := 0; k < len(q); k++ {
		v, w := q[idx], q[k]
		if _, ok := pIndices[fmt.Sprintf("%v-%v", v.X, v.Y)]; ok {
			if val, ok := pIndices[fmt.Sprintf("%v-%v", w.X, w.Y)]; ok {
				return val, k
			}
		}
		idx = k
	}

	return -1, -1
}

// if point x,y lies on ray pq
func CutRay(p, q shape.Point, x, y float64) bool {
	possibleCut := (p.Y > y && q.Y < y) || (p.Y < y && q.Y > y) // possible cut
	maybeCut := x-p.X < ((y - p.Y) * (q.X - p.X) / (q.Y - p.Y))
	// x < cut.X
	return possibleCut && maybeCut
}

// test if the ray crosses boundary from interior to exterior.
// 	-- this is needed due to edge cases, when the ray passes through
// 	-- polygon corners
func CrossBoundary(p, q shape.Point, x, y float64) bool {
	c1 := (p.Y == y && p.X > x && q.Y < y)
	c2 := (q.Y == y && q.X > x && p.Y < y)
	return c1 || c2
}
