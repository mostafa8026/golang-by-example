package main

import (
	"fmt"
	"time"
)

func main() {

	var msg = "Hello anonymously"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
	msg = "Goodbye anonymously"
}

func sayHello() {
	fmt.Println("Hello!")
}
