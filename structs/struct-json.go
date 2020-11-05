package main

import (
	"encoding/json"
	"fmt"
)

type jsontest struct {
	A int `json:"hashal"`
}

func main() {
	s := "{ \"hashal\": 5 }"
	fmt.Println(s)
	var ss jsontest
	err := json.Unmarshal([]byte(s), &ss)
	fmt.Println(ss)
	fmt.Println(err)
}
