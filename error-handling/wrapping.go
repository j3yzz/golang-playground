package main

import "fmt"

type notPositive struct {
	num int
}

func (p notPositive) Error() string {
	return fmt.Sprintf("checkPositive: given number %d is not a positive number", p.num)
}

type notEven struct {
	num int
}

func (e notEven) Error() string {
	return fmt.Sprintf("checkEven: given number %d is not an even number", e.num)
}

func main() {
	num := 3
	err := checkPositiveAndEven(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("given number is positive and even")
	}
}

func checkPositiveAndEven(num int) interface{} {
	if num > 100 {
		return fmt.Errorf("checkPositiveAndEven: Number %d is greater than 100", num)
	}

	err := checkPositive(num)
	if err != nil {
		return err
	}

	err = checkEven(num)
	if err != nil {
		return err
	}

	return nil
}

func checkEven(num int) error {
	if num%2 == 1 {
		return notEven{num: num}
	}

	return nil
}

func checkPositive(num int) error {
	if num < 0 {
		return notPositive{num: num}
	}

	return nil
}
