package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	abs := func(x float64) float64 {
		if x < 0 {
			return -x
		}
		return x
	}

	x = abs(x)

	z := x / 2.0
	for {
		next := z - ((z*z - x) / (z * 2))
		if abs(z-next) < 0.0000000001 {
			return z
		} else {
			z = next
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
