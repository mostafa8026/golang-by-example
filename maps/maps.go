package main

import "fmt"

func main() {
	var s = map[string]int{
		"akbar": 1,
	}
	var m = make(map[string]int)
	fmt.Println(m)

	s["asghar"] = 2
	b, ok := s["asghar"]

	if ok {
		fmt.Println(b)
	} else {
		fmt.Println("not ok")
	}

	m = s

	fmt.Println(s)
	fmt.Println(m)

	s["asghar"] = 5

	fmt.Println(s)
	fmt.Println(m)

	delete(s, "asghar")

	fmt.Println(s)
	fmt.Println(m)

	var i = map[interface{}]interface{}{}
	i["name"] = "ali"
	i["age"] = 35
	i[55] = "asdf"
	fmt.Println(i)
}
