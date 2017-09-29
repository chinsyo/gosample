package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	colors := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		colomn := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			colomn[j] = uint8(i ^ j)
		}
		colors[i] = colomn
	}
	return colors
}

func main() {
	pic.Show(Pic)
}
