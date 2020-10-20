package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1.6516
	var b float64 = math.Pow(math.Sqrt(a), 2)
	fmt.Println(b - a)
	if math.Abs(b-a) <= 0.0000000000000000001 {
		fmt.Println("a is printing.")
	}
}
