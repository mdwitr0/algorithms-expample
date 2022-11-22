package main

import (
	"algorithms/search"
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res, ok := search.LinearSearch(numbers, 6)

	if ok {
		fmt.Println("Linear search result: ", res)
	}
}
