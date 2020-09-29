package main

import (
	"fmt"
	"reflect"
)

type tag struct {
	a int `length max: 100`
}

func main() {
	t := tag{a: 5}

	fmt.Println(t)

	tt := reflect.TypeOf(tag{})
	field, err := tt.FieldByName("a")
	fmt.Println(field, err)
	fmt.Println(field.Tag)
}
