package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	count := 20
	fmt.Printf("going to create %d goroutines.\n", count)

	flag := true
	if len(os.Args) == 1 {
		flag = false
	}

	var wg sync.WaitGroup

	fmt.Printf("%#v\n", wg)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	if flag {
		wg.Add(1)
	} else {
		wg.Done()
	}

	fmt.Printf("\n%#v\n", wg)
	wg.Wait()
	fmt.Println("\nExiting...")
}
