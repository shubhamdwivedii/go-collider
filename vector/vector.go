package vector

import "math"

var (
	cos  = math.Cos
	sin  = math.Sin
	sqrt = math.Sqrt
)

// ---------------------- arithmetic --------------------------

func Mul(s, x, y float64) (sx, sy float64) {
	return s * x, s * y
}

func Div(s, x, y float64) (xBys, yBys float64) {
	return x / s, y / s
}

func Add(x1, y1, x2, y2 float64) (sx, sy float64) {
	return x1 + x2, y1 + y2
}

func Sub(x1, y1, x2, y2 float64) (sx, sy float64) {
	return x1 - x2, y1 - y2
}

func Permul(x1, y1, x2, y2 float64) (pmx, pmy float64) {
	return x1 * x2, y1 * y2
}

func Dot(x1, y1, x2, y2 float64) (prod float64) {
	return x1*x2 + y1*y2
}

func Det(x1, y1, x2, y2 float64) (det float64) { // Detriment ??
	return x1*y2 - y1*x2
}

// ---------------------------- relation --------------------------------

func Eq(x1, y1, x2, y2 float64) (equal bool) {
	return x1 == x2 && y1 == y2
}

func Lt(x1, y1, x2, y2 float64) (less bool) {
	return x1 < x2 || (x1 == x2 && y1 < y2)
}

func Le(x1, y1, x2, y2 float64) (lessOrEqual bool) {
	return x1 <= x2 && y1 <= y2
}

// --------------------------- misc operations --------------------------

func Len2(x, y float64) (len2 float64) {
	return x*x + y*y
}

func Len(x, y float64) (len float64) {
	return sqrt(x*x + y*y)
}

func Dist(x1, y1, x2, y2 float64) (dist float64) {
	return Len(x1-x2, y1-y2)
}

func Normalize(x, y float64) (nx, ny float64) {
	l := Len(x, y)
	return x / l, y / l
}

func Rotate(phi, x, y float64) (rx, ry float64) {
	c, s := cos(phi), sin(phi)
	return c*x - s*y, s*x + c*y
}

func Perpendicular(x, y float64) (px, py float64) {
	return -y, x
}

func Project(x, y, u, v float64) (px, py float64) {
	s := (x*u + y*u) / (u*u + v*u)
	return s * u, s * v
}

func Mirror(x, y, u, v float64) (mx, my float64) {
	s := 2 * (x*u + y*u) / (u*u + v*v)
	return s*u - x, s*v - y
}
