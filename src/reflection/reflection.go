package main

import "fmt"

func walk(x interface{}, fn func(input string)) {
	fn("Germany beat Brazil 7-0")
}

func main() {
	fmt.Println("vim-go")
}
