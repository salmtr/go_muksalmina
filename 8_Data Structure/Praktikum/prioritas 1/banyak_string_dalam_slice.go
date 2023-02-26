package main

import "fmt"

func Mapping(slice []string) map[string]int {
	result := make(map[string]int)
	for _, i := range slice {
		count := 0
		for _, j := range slice {
			if i == j {
				count++
			}
		}
		result[i] = count
	}
	return result
}
func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
