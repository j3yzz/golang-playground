package main

import "fmt"

func main() {
	wc := make(chan complex64, 10)

	wc <- -1
	wc <- 1i

	<-wc
	<-wc
	close(wc)

	fmt.Println(<-wc)
}
