package main

import "fmt"

func main() {
	var myrune rune = 'A'

	fmt.Println("myrune", myrune)

	var str string = "ABCD"
	array := []rune(str)
	fmt.Printf("Array of rune values for A, B, C and D: %U\n", array)

	emoji := '😀'

	// %T represents the type of the variable num
	fmt.Printf("Data type of %c is %T and the rune value is %U \n", emoji, emoji, emoji)

}
