package main

import "fmt"

func main() {
	a := 2
	switch {
	case a <= 1:
		fmt.Println("a is lower than 1")
		
	case a <= 2:
		fmt.Println("a is lower than 2")
		fallthrough
	case a <= 3:
		fmt.Println("a is lower than 3")
	}
}
