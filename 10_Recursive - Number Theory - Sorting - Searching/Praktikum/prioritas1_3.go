package main

import (
	"fmt"
)

func isPrime(num, i int) bool {
	if num == i {
		return true
	} else if num%i == 0 {
		return false
	} else {
		return isPrime(num, i+1)
	}
}
func primeX(num int) int {
	result := []int{}
	if num < 1 {
		return 0
	} else {
		for i := 2; i <= num*100; i++ {
			if isPrime(i, 2) == true {
				result = append(result, i)
				if len(result) == num {
					break
				}
			}
		}
	}

	return result[num-1]
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
