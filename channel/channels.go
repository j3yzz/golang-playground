package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	fmt.Println("Read:", <-c)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("channel is closed!")
	}
	wg.Wait()
	var ch chan bool = make(chan bool)
	for i := 0; i < 5; i++ {
		go printer(ch)
	}

	n := 0
	for i := range ch {
		fmt.Println(i)
		if i {
			n++
		}
		if n > 2 {
			fmt.Println("n:", n)
			close(ch)
			break
		}

	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func printer(ch chan bool) {
	ch <- true
}
