package main

import (
	"algorithms/math"
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
		var graph = map[string][]string{
			"you":    []string{"alice", "bob", "claire"},
			"bob":    []string{"anuj", "peggy"},
			"alice":  []string{"peggy"},
			"claire": []string{"thom", "jonny"},
			"anuj":   []string{},
			"peggy":  []string{},
			"thom":   []string{},
			"jonny":  []string{},
		}

		res := search.BFS(graph, "you", "thom")

		fmt.Println("BFS result:", res)
	}()

	func() {
		numbers := []int{11, 22, 3, 4, 5, 6, 7, 8, 9, 10}

		list := sort.Selection(numbers)

		fmt.Println("Selection sort result:", list)
	}()

	func() {
		list := sort.Quick(numbers)

		fmt.Println("Quick sort result:", list)
	}()

	func() {
		res := math.Sum(numbers)

		fmt.Println("Sum result:", res)
	}()

	func() {
		res := math.Fact(5)

		fmt.Println("Fact result:", res)
	}()
}
