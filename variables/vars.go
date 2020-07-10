package main

import "fmt"

// we can't use := in the package level.
var i int = 42

// var block
var (
	firstName string = "mostafa"
	lastName  string = "mostafavi"
)

func main() {
	var i int = 42
	fmt.Println(i)

	j := 42
	fmt.Printf("%v, %T", j, j)
}
