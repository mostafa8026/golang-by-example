package main

import "fmt"

func main() {
	result := testPointer()
	fmt.Println(*result)
}

func testPointer() *int {
	var i int = 5
	return &i
}
