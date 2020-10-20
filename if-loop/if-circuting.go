package main

import "fmt"

func main() {
	a := 1

	if a > 2 && ciruit() {
		fmt.Println("a is between 2 and 7")
	}
}

func ciruit() bool {
	fmt.Println("I'm called")
	return true
}
