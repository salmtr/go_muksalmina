package main

import "fmt"

func romawi(num int) string {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	symbol := ""
	for i := 0; i < len(value); i++ {
		for num >= value[i] {
			symbol += romawi[i]
			num -= value[i]
		}
	}
	return symbol
}

func main() {
	fmt.Println(romawi(4))    // IV
	fmt.Println(romawi(9))    // IX
	fmt.Println(romawi(23))   //XXIII
	fmt.Println(romawi(2021)) //MMXXI
	fmt.Println(romawi(1646)) //MDCXLVI
}
