package main

import (
	"fmt"
	"math"
)

var Global int = 1276
var AnotherGlobal = -5678

func main() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("Init j Value:", j)
	j = Global
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d, k=%.2f.\n", Global, i, j, k)
}
