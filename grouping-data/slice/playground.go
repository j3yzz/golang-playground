package main

import "fmt"

func main() {
	x := []int{1, 4, 5, 12, 51, 123}
	fmt.Println(x)
	fmt.Println("----------------")
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("----------------")
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
	fmt.Println("----------------")
	i := 0
	for i < len(x) {
		fmt.Println(i, x[i])
		i++
	}
	fmt.Println("----------------")
	i = 0
	for {
		if i >= len(x) {
			break
		}
		fmt.Println(i, x[i])
		i++
	}
}
