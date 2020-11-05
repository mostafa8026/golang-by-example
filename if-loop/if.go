package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
	}

	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}
}
