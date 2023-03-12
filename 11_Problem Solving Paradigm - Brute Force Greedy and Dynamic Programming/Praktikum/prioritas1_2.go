package main

import "fmt"

func pascalTriangle(numsRows int) [][]int {
	segiTiga := make([][]int, numsRows)

	for i := range segiTiga {
		segiTiga[i] = make([]int, i+1)
		segiTiga[i][0], segiTiga[i][i] = 1, 1

		for j := 1; j < i; j++ {
			segiTiga[i][j] = segiTiga[i-1][j-1] + segiTiga[i-1][j]
		}
	}
	return segiTiga
}

func main() {
	fmt.Println(pascalTriangle(5))

}
