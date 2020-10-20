package main

import "fmt"

func main() {
	var a interface{}
	a = 1.2
	switch a.(type) {
	case int:
		fmt.Println("a is an integer")
	case float64:
		fmt.Println("a is a float")
	}
}
