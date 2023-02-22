package main

import (
	"fmt"
)

type animal interface {
	walk()
}

type lion struct {
	color string
	age int
}

type dog struct {
	color string
	age int
}

func (l lion) walk()  {
	fmt.Println("Lion can walk.")
}

func (d dog) walk()  {
	fmt.Println("Dog can walk.")
}

func main() {
	//d := dog{age: 12, color: "red"}
	//l := lion{age: 22, color: "yellow"}
	//
	//d.walk()
	//l.walk()

	var a animal
	a = lion{age: 42, color: "black"}
	printInt(a)
}

func printInt(a animal) {
	val, ok := a.(lion)
	if ok {
		fmt.Println("age", val.age)
	}
}
