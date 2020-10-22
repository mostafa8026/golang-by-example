package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	// we can't use arithmetic operation on pointers
	// (if you really want to do use the unsafe package)
	fmt.Printf("%v %p %p\n", a, b, c)
}
