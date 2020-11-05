package main

import "fmt"

func main() {
Loop:
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j > 3 {
				break Loop
			}
		}
		fmt.Println("out of j loop")
	}
	fmt.Println("out of loop")
}
