package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Balance struct {
	Period     string
	Sector     string
	Descriptor string
}

func createBalance(record Balance) {
	time.Sleep(10 * time.Millisecond)
}

func readData(balanceChan chan []Balance) {
	var balances []Balance

	csvFile, err := os.Open("file.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range csvLines {
		balances = append(balances, Balance{
			Period:     line[0],
			Sector:     line[1],
			Descriptor: line[3],
		})
	}

	balanceChan <- balances
}

func worker(balanceChan chan Balance) {
	for val := range balanceChan {
		createBalance(val)
	}
}

func main() {
	startTime := time.Now()

	balances := make(chan []Balance)

	go readData(balances)

	const workers = 5
	jobs := make(chan Balance, 1000)

	for w := 1; w <= workers; w++ {
		go worker(jobs)
	}

	counter := 0
	for _, val := range <-balances {
		counter++
		jobs <- val
	}

	fmt.Println("records saved:", counter)
	fmt.Println("total time:", time.Since(startTime))
}
