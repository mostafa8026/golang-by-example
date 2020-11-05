package main

import "fmt"

func main() {
	var i interface{}
	i = 24
	switch i.(type) {
	case int:
		fmt.Println("i is an interger")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is default")
	}
}
