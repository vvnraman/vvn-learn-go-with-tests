package main

import (
	"fmt"
	"reflect"
)

func walk_v1(x interface{}, fn func(input string)) {
	fn("Germany beat Brazil 7-0")
}

func walk_v2(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	// this would panic if the struct has no fields in it.
	field := val.Field(0)

	// this would be wrong if the type of the field is not string
	fn(field.String())
}

func walk_v3(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fn(field.String())
	}

}

func main() {
	fmt.Println("vim-go")
}
