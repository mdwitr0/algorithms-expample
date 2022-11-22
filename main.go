package main

import (
	"algorithms/search"
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	func() {
		res, ok := search.Linear(numbers, 6)

		if ok {
			fmt.Println("Linear search result:", res)
		}
	}()

	func() {
		res, ok := search.Binary(numbers, 5)

		if ok {
			fmt.Println("Binary search result:", res)
		}
	}()
}
