package main

import "fmt"

func main() {
	arrayOne := []int{1, 2, 3, 4}
	arrayTwo := []int{5, 6, 7, 8}
	s := conc(arrayOne, arrayTwo)

	fmt.Println(s)
}

func conc(a, b []int) []int {
	newSlice := make([]int, len(a), len(b))
	copy(newSlice, a)
	copy(newSlice[len(a):], b)
	return newSlice
}
