package main

import (
	"fmt"
	"strconv"
)

// we can't use := in the package level.
var i int = 27

// var block
var (
	firstName string = "mostafa"
	lastName  string = "mostafavi"
)

func main() {
	fmt.Println("Program started ... let's talk about variables.")
	fmt.Println(i)
	var i int = 42
	fmt.Println(i)

	j := 42
	fmt.Printf("%v, %T\n", j, j)

	var k float32
	k = float32(j)
	fmt.Printf("%v, %T\n", k, k)

	var l string
	l = string(i)
	fmt.Printf("%v, %T\n", l, l)
	l = strconv.Itoa(i)
	fmt.Printf("%v, %T", l, l)
}
