package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("apakah palindrom?")
	fmt.Print("Masukkan kata: ")
	var kata string
	fmt.Scanln(&kata)

	kata = strings.ToLower(strings.ReplaceAll(kata, " ", ""))

	var balikKata string
	for i := len(kata) - 1; i >= 0; i-- {
		balikKata += string(kata[i])
	}
	if kata == balikKata {
		fmt.Printf("%s  palindrom", kata)
	} else {
		fmt.Printf("%s  bukan palindrome", kata)
	}
}
