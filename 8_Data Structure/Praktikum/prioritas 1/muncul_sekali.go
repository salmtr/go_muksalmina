package main

import "fmt"

func munculSekali(angka string) []int {
	count := make(map[int]int)

	for _, i := range angka {
		input := int(i - '0')
		count[input]++
	}

	result := []int{}
	for input, freq := range count {
		if freq == 1 {
			result = append(result, input)
		}
	}

	return result
}
func main() {
	fmt.Println(munculSekali("1234123"))    // Output :[4]
	fmt.Println(munculSekali("76523752"))   // Output :[6 3]
	fmt.Println(munculSekali("12345"))      // Output :[1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // Output :[]
	fmt.Println(munculSekali("0872504"))    // Output :[8 7 2 5 4]
}
