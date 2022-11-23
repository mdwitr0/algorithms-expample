package main

import (
	"algorithms/search"
	"algorithms/sort"
	"fmt"
)

func main() {
	numbers := []int{11, 22, 3, 4, 5, 6, 7, 8, 9, 10}
	func() {
		res, ok := search.Linear(numbers, 6)

		if ok {
			fmt.Println("Linear search result:", res)
		}
	}()

	func() {
		res, ok := search.Binary(numbers, 9)

		if ok {
			fmt.Println("Binary search result:", res)
		}
	}()

	func() {
		list := sort.SelectionSearch(numbers)

		fmt.Println("Selection search result:", list)
	}()
}
