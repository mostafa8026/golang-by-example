package main

import (
	"fmt"
	"time"
)

func main() {
	message := "mostafa"
	go func(message string) {
		fmt.Println(message)
	}(message)
	message = "hoseiny"
	time.Sleep(1 * time.Second)
}
