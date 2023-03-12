package main

import "fmt"

func biner(num int) {
	for i := 0; i <= num; i++ {
		fmt.Printf("%b", i)
	}
}
func main() {
	biner(2)
	fmt.Println()
	biner(3)
}
