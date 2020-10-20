package main

import "fmt"

func main() {
	i, j := 1, 1
	for {
		fmt.Println(i)
		i, j = i+1, j+1
		if i < 10 {
			break
		}
	}
}
