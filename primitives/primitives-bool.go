package main

import "fmt"

func mainMain() {
	var a bool = false
	var b bool = true
	var c bool = !b

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
}
