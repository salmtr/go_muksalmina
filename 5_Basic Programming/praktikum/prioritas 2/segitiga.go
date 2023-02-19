package main

import "fmt"

func main() {
	nilai := 5

	for i := 1; i <= nilai; i++ {
		for j := i; j < nilai; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}
