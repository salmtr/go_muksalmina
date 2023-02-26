package main

import (
	"fmt"
	"math"
)

func main() {
	diagonalKiri := 0
	diagonalKanan := 0
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				diagonalKiri += matrix[i][j]
			}
			if i+j == 2 {
				diagonalKanan += matrix[i][j]
			}
		}
	}
	fmt.Println("Diagonal kiri :", diagonalKiri)
	fmt.Println("Diagonal kanan:", diagonalKanan)
	fmt.Println("Perbedaan Mutlak :", int(math.Abs(float64(diagonalKiri-diagonalKanan))))
}
