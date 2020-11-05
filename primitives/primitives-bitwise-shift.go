package main

import "fmt"

func main() {
	a := 10 // 001010

	fmt.Println(a << 1) // 010100
	fmt.Println(a << 2) // 0101000
	fmt.Println(a << 3) // 1010000

	a = 24              // 11000
	fmt.Println(a >> 1) // 01100
	fmt.Println(a >> 2) // 00110
	fmt.Println(a >> 3) // 00011
}
