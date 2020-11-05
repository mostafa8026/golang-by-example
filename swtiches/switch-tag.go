package main

import "fmt"

func main() {
	a := 3
	switch a {
	case 1, 5:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is default")
	}
}
