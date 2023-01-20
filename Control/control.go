package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}
	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("zero")
	case "1":
		fmt.Println("One")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("Value:", argument)
	}

	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int:", argument)
		return
	}
	switch {
	case value == 0:
		fmt.Println("zero!")
	case value > 0:
		fmt.Println("positive integer")
	case value < 0:
		fmt.Println("negative integer")
	default:
		fmt.Println("this should not happen:", value)
	}
}