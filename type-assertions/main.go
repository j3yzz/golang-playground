package main

import "fmt"

func main() {
	var a interface{} = "Hello"

	fmt.Printf("%T\n", a)

	s, ok := a.(string)
	fmt.Println(s, ok)

	f, ok := a.(float64)
	fmt.Println(f, ok)
}
