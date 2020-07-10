package main

import (
	"fmt"
	"strconv"
)

// package scope:
var myPackageScopeVariable int = 42

// MyExportedVariable exported so it can be used by all othe source files globally
var MyExportedVariable int = 43

//we have no private scope.
//Naming conventions:
// - pascal or camelCase
// - - Capitalize acronums (HTTP, URL, ID)
// - as short as possible
// - - longer names for longer lives

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

	var m int
	k = 42.5
	//m = k this assignment got you a compiler error, because go is not gonna risk the
	// possibility of losing information through that conversion.
	m = int(k) // this is your decide and everything is ok. (explicit conversion)
	fmt.Printf("%v, %T\n", m, m)
}
