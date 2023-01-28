package main

import "fmt"

func main() {
	x := []int{1, 4, 5, 12, 51, 123}
	/// print last item in slice
	fmt.Println(x[len(x)-1:])

	for i, v := range x {
		fmt.Println(i, "->", v)
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i, x[i])
	}
	//// append elements to slice
	x = append(x, 256, 512, 1024)
	fmt.Println(x)
	y := []int{2020, 2021, 2022, 2023}
	x = append(x, y...)
	fmt.Println(x)

	//// delete element from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
	fmt.Println("------------")
	fmt.Println("//// MAKE ////")
	fmt.Println("------------")
	/// make | built-in function
	m := make([]int, 10, 12)
	fmt.Println(m)
	fmt.Println("len:", len(m))
	fmt.Println("capacity:", cap(m))
	m[0] = 23
	m[9] = 123

	m = append(m, 322)
	fmt.Println(m)
	fmt.Println("len:", len(m))
	fmt.Println("capacity:", cap(m))

	m = append(m, 324)
	m = append(m, 325)
	fmt.Println(m)
	fmt.Println("len:", len(m))
	fmt.Println("capacity:", cap(m))

	jb := []string{
		"james",
		"bond",
		"chocolate",
		"martini",
	}
	fmt.Println(jb)

	mp := []string{
		"miss",
		"Moneypenny",
		"Strawberry",
		"Hazelnut",
	}
	fmt.Println(mp)
	var xp [][]string
	xp = append(xp, jb, mp)
	fmt.Println(xp)
}
