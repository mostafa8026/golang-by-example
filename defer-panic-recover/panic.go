package main

import "fmt"

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	panic("Something bad is happened!")
	fmt.Println("end panic")
}
