package main

import "fmt"

func main() {
	var alas_m, alas_n, tinggi, luas float64
	fmt.Print("masukkan panjang alas m:")
	fmt.Scan(&alas_m)
	fmt.Print("masukkan panjang alas n:")
	fmt.Scan(&alas_n)
	fmt.Print("masukkan tinggi:")
	fmt.Scan(&tinggi)
	luas = (alas_m + alas_n) * tinggi / 2
	fmt.Println("luas trapesium adalah", luas)
}
