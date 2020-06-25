package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 4)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
