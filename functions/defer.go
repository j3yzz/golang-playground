package main

import "fmt"

func main() {
	newFoo()
	defer newBar()
}

func newFoo() {
	fmt.Println("Hello From Foo")
}

func newBar() {
	fmt.Println("Hello From Bar")
}
