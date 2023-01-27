package main

import "fmt"

var y string
var z int

func main() {
	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println("printing the value of z", z, "ending")
	fmt.Printf("%T\n", z)

	y = "Bond, James Bond"

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	z = 12

	fmt.Println("printing the value of z", z, "ending")
	fmt.Printf("%T\n", z)

}
