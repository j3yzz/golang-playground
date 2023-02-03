package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type student struct {
	person
	studentNo int
	grade     int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   20,
	}

	p2 := person{
		"Miss",
		"Moneypenny",
		23,
	}

	s1 := student{p2, 23330000, 2}

	fmt.Println(p1, p2)
	fmt.Println("---------------")
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last)
	fmt.Println("---------------")
	fmt.Println("p1 firstname:", p1.first)
	fmt.Println("p2 firstname:", p2.first)
	fmt.Println("---------------")
	fmt.Println(s1)
	fmt.Println("---------------")
	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}
	fmt.Println(sa1)
}
