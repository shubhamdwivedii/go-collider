package main

import (
	"fmt"

	cr "github.com/shubhamdwivedii/go-collider/circle"
)

func main() {
	c1 := cr.New(10, 10, 5)
	c2 := cr.New(7, 7, 3)

	collides, dx, dy := c1.CollidesWith(c2)
	collides2, dx2, dy2 := c2.CollidesWith(c1)

	c3 := cr.New(0, 0, 3)

	collides3, dx3, dy3 := c1.CollidesWith(c3)

	fmt.Println(collides, dx, dy)
	fmt.Println(collides2, dx2, dy2)
	fmt.Println(collides3, dx3, dy3)
}
