package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	mapItem := make(map[string]int)
	hasil := []pair{}

	for _, s := range items {
		mapItem[s]++
	}
	for nama, jumlah := range mapItem {
		hasil = append(hasil, pair{nama + "", jumlah})
	}

	sort.Slice(hasil, func(low, high int) bool {
		sort := hasil[low].count < hasil[high].count
		return sort
	})
	return hasil
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) // golang ->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))                // football->1 basketball->1 tenis->1
}
