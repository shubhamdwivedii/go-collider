package polygon

import (
	"errors"
	"math"

	"github.com/shubhamdwivedii/go-collider/shape"
	"github.com/shubhamdwivedii/go-collider/vector"
)

type Polygon struct {
	vertices []shape.Point
	// rotation float64
	area     float64
	centroid shape.Point
	isConvex bool
	radius   float64
}

func New(coordinates ...float64) (*Polygon, error) {
	vertexList := ToVertexList([]shape.Point{}, coordinates...)
	vertices := RemoveCollinear(vertexList)

	if len(vertices) < 3 {
		return nil, errors.New("need atleast 3 non collinear points to build a polygon")
	}

	// Assert polygon is oriented counter clockwise
	r := GetIndexOfLeftmost(vertices)

	q := len(vertices)
	if r > 1 {
		q = r - 1
	}

	s := 1
	if r < len(vertices) {
		s = r + 1
	}

	if !CCW(vertices[q], vertices[r], vertices[s]) {
		var tmp []shape.Point
		for i := len(vertices) - 1; i >= 0; i-- {
			tmp = append(tmp, vertices[i])
		}
		vertices = tmp
	}

	// Assert polygon is not self-intersecting
	// outer: only need to check segemts len(vertices);1, 1;2, ....,
	// inner: only need to check unconnected segements
	a := vertices[len(vertices)-1]
	b := a
	for i := 0; i < len(vertices)-2; i++ {
		b, a = a, vertices[i]
		for k := i + 1; k < len(vertices)-1; k++ {
			c, d := vertices[k], vertices[k+1]
			if SegementsInterset(b, a, c, d) {
				return nil, errors.New("Polygon may not intersect with itself")
			}
		}
	}

	// Make vertices immutable ?
	polygon := &Polygon{
		vertices: vertices,
	}

	// Compute polygon area
	pp, qq := vertices[len(vertices)-1], vertices[0]
	det := vector.Det(pp.X, pp.Y, qq.X, qq.Y) // also used below
	polygon.area = det

	for i := 1; i < len(vertices); i++ {
		pp, qq = qq, vertices[i]
		polygon.area = polygon.area + vector.Det(pp.X, pp.Y, qq.X, qq.Y)
	}
	polygon.area = polygon.area / 2

	// Compute polygon centroid

	pp, qq = vertices[len(vertices)-1], vertices[0]
	polygon.centroid = shape.Point{
		X: (pp.X + qq.X) * det,
		Y: (pp.Y + qq.Y) * det,
	}

	for i := 1; i < len(vertices); i++ {
		pp, qq = qq, vertices[i]
		det = vector.Det(pp.X, pp.Y, qq.X, qq.Y)
		polygon.centroid.X = polygon.centroid.X + (pp.X+qq.X)*det
		polygon.centroid.Y = polygon.centroid.Y + (pp.Y+qq.Y)*det
	}
	polygon.centroid.X = polygon.centroid.X / (6 * polygon.area)
	polygon.centroid.Y = polygon.centroid.Y / (6 * polygon.area)

	// Get outcircle
	polygon.radius = 0
	for _, point := range vertices {
		polygon.radius = math.Max(
			polygon.radius,
			vector.Dist(point.X, point.Y, polygon.centroid.X, polygon.centroid.Y),
		)
	}

	// Checking if polygon is convex (all edges are oriented)
	polygon.isConvex = CheckConvex(polygon.vertices)

	return polygon, nil
}

// return vertices as x1,y1,x2,y2, ..., xn,yn
func (polygon *Polygon) Unpack() []float64 {
	var res []float64
	for _, point := range polygon.vertices {
		res = append(res, point.X, point.Y)
	}
	return res
}

// Deep clone of the polygon
func Clone(polygon Polygon) (*Polygon, error) {
	coordinates := polygon.Unpack()
	return New(coordinates...)
}

func (polygon *Polygon) BoundingBox() (x1, y1, x2, y2 float64) {
	ulx, uly := polygon.vertices[0].X, polygon.vertices[0].Y
	lrx, lry := ulx, uly
	for i := 1; i < len(polygon.vertices); i++ {
		p := polygon.vertices[i]
		if ulx > p.X {
			ulx = p.X
		}
		if uly > p.Y {
			uly = p.Y
		}
		if lrx < p.X {
			lrx = p.X
		}
		if lry < p.Y {
			lry = p.Y
		}
	}
	return ulx, uly, lrx, lry
}

// Only calculated once
func (polygon *Polygon) IsConvex() bool {
	return polygon.isConvex
}

// Checking if polygon is convex (all edges are oriented)
func CheckConvex(v []shape.Point) bool {
	if len(v) == 3 {
		return true
	}

	if !CCW(v[len(v)-1], v[0], v[1]) {
		return false
	}

	for i := 1; i < len(v)-2; i++ {
		if !CCW(v[i-1], v[i], v[i+1]) {
			return false
		}
	}

	return CCW(v[len(v)-2], v[len(v)-1], v[0])
}

func (polygon *Polygon) Move(dx, dy float64) {
	for _, point := range polygon.vertices {
		point.X += dx
		point.Y += dy
	}

	polygon.centroid.X += dx
	polygon.centroid.Y += dy
}

func (polygon *Polygon) Rotate(angle float64) {
	cx, cy := polygon.centroid.X, polygon.centroid.Y
	polygon.RotateAt(angle, cx, cy)
}

func (polygon *Polygon) RotateAt(angle, cx, cy float64) {
	for _, point := range polygon.vertices {
		rx, ry := vector.Rotate(angle, point.X-cx, point.Y-cy)
		point.X, point.Y = vector.Add(cx, cy, rx, ry)
	}
	vx, vy := polygon.centroid.X, polygon.centroid.Y
	rx, ry := vector.Rotate(angle, vx-cx, vy-cy)
	polygon.centroid.X, polygon.centroid.Y = vector.Add(cx, cy, rx, ry)
}

func (polygon *Polygon) ScaleAt(s, cx, cy float64) {
	for _, point := range polygon.vertices {
		// point = (point - center) * s + center
		mx, my := vector.Mul(s, point.X-cx, point.Y-cy)
		point.X, point.Y = vector.Add(cx, cy, mx, my)
	}

	polygon.radius *= s
}

func (polygon *Polygon) Scale(s float64) {
	cx, cy := polygon.centroid.X, polygon.centroid.Y
	polygon.ScaleAt(s, cx, cy)
}

// Polygon:Triangulate & Polygon:SplitConvex & Polygon:MergeWith pending

func (polygon *Polygon) Contains(x, y float64) bool {
	// test if an edge cuts the ray
	v := polygon.vertices
	inPolygon := false

	p, q := v[len(v)-1], v[len(v)-1]
	for i := 0; i < len(v); i++ {
		p, q = q, v[i]
		if CutRay(p, q, x, y) || CrossBoundary(p, q, x, y) {
			inPolygon = !inPolygon
		}
	}
	return inPolygon
}
