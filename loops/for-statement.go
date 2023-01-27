package main

import "fmt"

func main() {
	x := 1
	for {
		if x >= 10 {
			break
		}

		fmt.Println(x)
		x++
	}

}
