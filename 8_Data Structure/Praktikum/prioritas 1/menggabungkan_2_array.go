package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// Menggabungkan 2 array
	merged := append(arrayA, arrayB...)
	// Membuat array kosong
	result := make([]string, 0)
	// Membuat map kosong
	isExist := make(map[string]bool)

	// Mengisi array yang kosong menggunakan array yang sudah digabung
	for _, input := range merged {
		// Jika map kosong, maka  map akan di isi dengan nilai true
		if !isExist[input] {
			isExist[input] = true
			// mengisi array yang kosong menggunakan nilai dari array yang sudah digabung tadi
			result = append(result, input)
		}
	}
	return result
}
func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}
