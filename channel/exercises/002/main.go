package main

import "fmt"

func main() {
	ch := generator()
	receiver(ch)
}

func receiver(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func generator() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
