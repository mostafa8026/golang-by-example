package main

import "fmt"

func main() {
	greetingV("mostafa", "mostafavi")
}

func greeting(name, family string) {
	fmt.Printf("hello, %v %v", name, family)
}

func greetingV(s int, variadic ...string) {
	fmt.Printf("%T %v\n", variadic, variadic)
	for k, v := range variadic {
		fmt.Println(k, v)
	}
}
