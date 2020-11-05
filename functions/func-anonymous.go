package main

import "fmt"

func main() {
	a := func(i int) int {
		return 5 + i
	}

	fmt.Println(a(9))
}
