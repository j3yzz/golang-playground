package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args
	if len(args) != 5 {
		fmt.Println("Usage: randomFiles start end filename directory")
	}

	start, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	if end < start {
		fmt.Println(end, " < ", start)
		return
	}

	filename := args[3]
	path := args[4]
	_, err = os.Open(path)
	if err != nil {
		fmt.Println(path, "doesn't exist.")
		return
	}
	var wg sync.WaitGroup

	for i := start; i <= end; i++ {
		wg.Add(1)
		filepath := fmt.Sprintf("%s/%s%d", path, filename, i)
		go func(f string) {
			defer wg.Done()
			createFile(f)
		}(filepath)
	}

	wg.Wait()
}

func createFile(file string) {
	_, err := os.Stat(file)
	if err == nil {
		fmt.Printf("%s already exist!\n", file)
		return
	}

	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	lines := random(10, 30)
	for i := 0; i < lines; i++ {
		data := random(0, 20)
		_, err := fmt.Fprintf(f, "%d\n", data)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%s created.\n", file)
}

func random(i int, i2 int) int {
	return rand.Intn(i2-i) + i
}
