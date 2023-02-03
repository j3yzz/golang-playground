package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pixel := make([][]uint8, dy)
	data := make([]uint8, dx)

	for i := range pixel {
		for j := range data {
			data[j] = uint8((i + j) / 2)
		}
		pixel[i] = data
	}
	return pixel
}

func main() {
	pic.Show(Pic)
}
