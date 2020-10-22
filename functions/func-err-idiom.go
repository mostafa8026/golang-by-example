package main

import "fmt"

func main() {
	result, err := divide(5, 0)
	if err != nil {
		fmt.Println("not a good number")
		return
	}
	fmt.Println("result is: ", result)
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("b must not be 0")
	}
	result := a / b
	return result, nil
}
