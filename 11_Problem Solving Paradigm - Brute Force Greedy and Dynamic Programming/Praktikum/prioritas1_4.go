package main

import (
	"fmt"
)

func SimpleEquation(a, b, c int) {
	for x := 1; x <= a/3; x++ {
		y := (a - x) / 2
		z := a - x - y
		if x*y*z == b && x*x+y*y+z*z == c {
			fmt.Println(x, y, z)
			return
		}
	}
	fmt.Println("No solution.")
}

func main() {
	SimpleEquation(1, 2, 3)  // No solution.
	SimpleEquation(6, 6, 14) // 1 2 3
}
