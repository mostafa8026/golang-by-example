package main

import "fmt"

func main() {
	a := 42
	var b *int = &a
	fmt.Printf("%v %T\n", *b, b)
	fmt.Println(a, *b)
	*b = 27
	fmt.Println(a, *b)
}
