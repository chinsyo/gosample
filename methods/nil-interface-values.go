package main

import (
	"fmt"
)

type I interface {
	M()
}

func main() {
	var i I
	// output "(<nil>, <nil>)", cuz a nil interface holds neither value nor concrete type
	describe(i)
	// runtime error, cuz no type inside the interface tuple to indicate which concrete method to call
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
