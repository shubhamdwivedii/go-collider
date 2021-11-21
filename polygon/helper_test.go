package polygon

import (
	"testing"

	"github.com/shubhamdwivedii/go-collider/shape"
	"github.com/stretchr/testify/require"
)

func TestAreCollinear(t *testing.T) {

	p1 := shape.Point{
		X: 1,
		Y: 1,
	}

	p2 := shape.Point{
		X: 1,
		Y: 2,
	}

	p3 := shape.Point{
		X: 1,
		Y: 3,
	}

	p4 := shape.Point{
		X: 2,
		Y: 1,
	}

	p5 := shape.Point{
		X: 3,
		Y: 1,
	}

	res1 := AreCollinear(p1, p2, p3)
	require.True(t, res1)

	res2 := AreCollinear(p1, p4, p5)
	require.True(t, res2)

	res3 := AreCollinear(p2, p3, p4)
	require.False(t, res3)
}
