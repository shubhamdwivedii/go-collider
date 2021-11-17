package circle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCircle(t *testing.T) {

	c1 := New(10, 10, 5)
	c2 := New(7, 7, 3)

	collides, dx, dy := c1.CollidesWith(c2)
	collides2, dx2, dy2 := c2.CollidesWith(c1)
	require.True(t, collides)
	require.True(t, collides2)
	require.Equal(t, dx, -dx2)
	require.Equal(t, dy, -dy2)

	x, y := c2.Center()
	contains := c1.Contains(x, y)
	require.True(t, contains)

	contains2 := c1.Contains(0, 0)
	require.False(t, contains2)

}
