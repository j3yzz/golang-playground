package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("Hello, I am", s.first, s.last, " - the secret agent speak.")
}

func (s person) speak() {
	fmt.Println("Hello, I am", s.first, s.last, " - the person speak.")
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{"amir", "kherad"}

	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	checkIsHuman(p1)
}

func checkIsHuman(h human) {
	fmt.Println("I called human")
}
