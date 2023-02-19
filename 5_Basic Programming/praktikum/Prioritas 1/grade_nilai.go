package main

import "fmt"

func main() {
	var nilai int = 100
	fmt.Print("masukkan nilai : ")
	fmt.Scan(&nilai)
	if nilai >= 80 && nilai <= 100 {
		fmt.Println("A")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Println("B")
	} else if nilai >= 60 && nilai <= 65 {
		fmt.Println("c")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Println("D")
	} else if nilai >= 0 && nilai <= 34 {
		fmt.Println("E")
	} else {
		fmt.Println("invalid")
	}

}
