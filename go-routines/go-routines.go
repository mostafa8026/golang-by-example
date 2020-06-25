package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	var msg = "Hello anonymously"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye anonymously"
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello!")
}
