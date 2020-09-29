package main

import "fmt"

func main() {
	type akbar struct {
		a interface{}
		b string
		y int
		c float32
	}

	s := akbar{
		a: 1,
		b: "asdf",
		c: 1231.12,
	}

	fmt.Println(s.c)

	var t = struct{ name string }{"asdfasdf"}
	fmt.Println(t)
	tt := &t
	fmt.Println(tt)
	t.name = "test"
	fmt.Println(t)
	fmt.Println(tt)
}
