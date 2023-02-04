package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	fullName, ok := getFullName("Ian", "Fleming")
	fmt.Println(fullName, ok)
	printAllX(2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo,", s)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func getFullName(fn, ln string) (string, bool) {
	if fn == "" {
		panic("panic in fn")
		return "", false
	}
	if ln == "" {
		panic("panic in ln")
		return "", false
	}

	return fmt.Sprint(fn, " ", ln), true
}

func printAllX(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		fmt.Println("for item in index position", i, "we are now adding,", v, "to the total and current total is:", sum)
		sum += v
	}

	fmt.Println("total:", sum)

}
