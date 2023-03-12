package main

import "fmt"

func main() {
	result := make(chan int, 20)

	go func() {
		for i := 1; i < 30; i++ {
			result <- i
		}
		close(result)
	}()

	for value := range result {
		if value%3 == 0 {
			fmt.Println(value)
		}
	}

	fmt.Println()
}
