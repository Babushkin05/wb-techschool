package main

import (
	"fmt"
	"reflect"
)

func guessType(a interface{}) string {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "other"
	}
}

func main() {
	data := []interface{}{
		42, "hello", true, make(chan int), func() {},
	}
	for _, value := range data {
		fmt.Printf("%v: %s\n", value, guessType(value))
	}
}
