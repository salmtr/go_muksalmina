package main

import (
	"fmt"
	"math"
)

func primeNumber(num int) bool {
	result := true
	if num < 2 {
		return false
	}
	akar := int(math.Sqrt(float64(num)))
	for i := 2; i <= akar; i++ {
		if num%i == 0 {
			result = false
			break
		}
	}
	return result
}
func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(1500450271))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
