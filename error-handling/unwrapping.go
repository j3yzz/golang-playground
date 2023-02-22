package main

import (
	"errors"
	"fmt"
)

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error one happened"
}

func main() {
	e1 := errorOne{}
	e2 := fmt.Errorf("e2: %w", e1)
	e3 := fmt.Errorf("e3: %w", e2)
	fmt.Println(errors.Unwrap(e3))
	fmt.Println(errors.Unwrap(e2))
	fmt.Println(errors.Unwrap(e1))
}
