package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {
	var result string
	for i := 0; i < len(a); i++ {
		for j := i + 1; j <= len(a); j++ {
			if strings.Contains(b, a[i:j]) {
				if len(a[i:j]) > len(result) {
					result = a[i:j]
				}

			}
		}
	}
	return result

}
func main() {
	fmt.Print("Compare :\n")
	fmt.Println(compare("AKA", "AKASHI"))     //AKA
	fmt.Println(compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(compare("KI", "KIJANG"))      //KI
	fmt.Println(compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(compare("ILALANG", "ILA"))    //ILA
}
