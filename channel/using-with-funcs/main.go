package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c)

	fmt.Println("about to exit")
}

// send
// when we want sending channel => chan<-
func foo(c chan<- int) {
	c <- 42
}

// receive
// when we want to receive channel => <-chan
func bar(c <-chan int) {
	fmt.Println(<-c)
}
