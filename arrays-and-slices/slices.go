package main

import "fmt"

func main() {
	a := make([]int, 4, 100)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 4)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
