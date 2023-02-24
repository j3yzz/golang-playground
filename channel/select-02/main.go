package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("exiting...")
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("from even channel:", v)
		case v := <-odd:
			fmt.Println("from odd channel:", v)
		case v := <-quit:
			fmt.Println("from quit channel:", v)
			return
		}
	}
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- true
}
