package main

import "fmt"

type shape struct {
	x int
	y int
}

type circle struct {
	shape
	radious int
}

func main() {
	c := circle{}
	c.x = 0
	c.y = 0
	c.radious = 2

	fmt.Println(c)

	b := circle{
		shape:   shape{x: 0, y: 0},
		radious: 2,
	}

	fmt.Println(b)
}
