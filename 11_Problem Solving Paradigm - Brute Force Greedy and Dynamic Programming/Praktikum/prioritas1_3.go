package main

import "fmt"

func fibonaciTopDown(number int, listFibo map[int]int) int {
	if number <= 1 {
		return number
	}

	if value, ok := listFibo[number]; ok {
		return value
	}

	listFibo[number] = fibonaciTopDown(number-1, listFibo) + fibonaciTopDown(number-2, listFibo)
	return listFibo[number]
}

func main() {
	list := make(map[int]int)
	fmt.Println(fibonaciTopDown(0, list))  // 0
	fmt.Println(fibonaciTopDown(1, list))  // 1
	fmt.Println(fibonaciTopDown(2, list))  // 1
	fmt.Println(fibonaciTopDown(3, list))  // 2
	fmt.Println(fibonaciTopDown(5, list))  // 5
	fmt.Println(fibonaciTopDown(6, list))  // 8
	fmt.Println(fibonaciTopDown(7, list))  // 13
	fmt.Println(fibonaciTopDown(9, list))  // 34
	fmt.Println(fibonaciTopDown(10, list)) // 55

}
