package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- i
			}
		}(i)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-ch)
	}
}
