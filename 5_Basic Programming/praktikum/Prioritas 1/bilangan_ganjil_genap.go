package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		fmt.Print("masukkan bilangan : ")
		fmt.Scan(&i)
		if i <= 100 {
			if i%2 == 0 {
				fmt.Print(i, "adalah bilangan genap \n")
			} else {
				fmt.Print(i, "adalah bilangan ganjil \n")
			}
		} else {
			fmt.Print("error sudah mencapai masa tenggang")
		}

	}
}
