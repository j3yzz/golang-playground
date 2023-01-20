package main

import "fmt"

func main() {
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}

	fmt.Println()

	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v:= range aSlice {
		fmt.Println("index:", i, "value:", v)
	}
}
