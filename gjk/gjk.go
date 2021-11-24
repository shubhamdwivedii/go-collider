package gjk

// Gilbert–Johnson–Keerthi distance algorithm

import (
	"math"

	"github.com/shubhamdwivedii/go-collider/shape"
	"github.com/shubhamdwivedii/go-collider/vector"
)

var (
	huge = math.MaxFloat64
	abs  = math.Abs
)

func Support(shapeA, shapeB shape.Shape, dx, dy float64) (sx, sy float64) {
	ax, ay := shapeA.Support(dx, dy)
	bx, by := shapeB.Support(-dx, -dy)
	return vector.Sub(ax, ay, bx, by)
}

// // Return closest edge to the origin
// func ClosestEdge(n) {

// }
