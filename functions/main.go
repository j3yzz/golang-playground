package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	fullName, ok := getFullName("Ian", "Fleming")
	fmt.Println(fullName, ok)
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
