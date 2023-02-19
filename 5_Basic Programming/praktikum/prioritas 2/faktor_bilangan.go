package main

import "fmt"

func main() {
	angka := 26
	fmt.Println("angka : ", angka)
	for i := 1; i <= angka; i++ {
		factorial := angka % i
		if factorial == 0 {
			fmt.Println(i)
		}
	}
}
