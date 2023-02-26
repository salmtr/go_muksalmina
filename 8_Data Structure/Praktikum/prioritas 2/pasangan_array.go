package main

import "fmt"

func PairSum(array []int, target int) []int {
	var result []int
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i]+array[j] == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // Output : [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // Output : [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // Output : [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // Output : [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // Output : [0, 1]
}
