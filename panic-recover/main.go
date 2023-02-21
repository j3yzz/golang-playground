package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	hello := []string{"salam", "hello"}
	checkAndPrint(hello, 2)
	fmt.Println("Exiting normally...")
}

func checkAndPrint(hello []string, i int) {
	defer handleOutOfBound()
	if len(hello) - 1 < i {
		panic("out of bound access for slice")
	}
	fmt.Println(hello[i])
}

func handleOutOfBound() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
		fmt.Println("Stack trace:")
		debug.PrintStack()
	}
}