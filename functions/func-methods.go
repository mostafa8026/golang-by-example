package main

import "fmt"

func main() {
	var r rectangle = rectangle{1, 2}
	fmt.Println(r.Area())
}

type rectangle struct {
	x, y int
}

func (s *rectangle) Area() int {
	return s.x * s.y
}
