package search

import (
	"container/list"
)

// Linear имеет линейное O(n) время выполнения
func Linear[T comparable](list []T, x T) (res int, ok bool) {
	for index, item := range list {
		if item == x {
			return index, true
		}
	}
	return -1, false
}

// Binary имеет логарифмическое O(log n) время выполнения
func Binary(list []int, x int) (res int, ok bool) {
	first := 0
	last := len(list) - 1

	for first <= last {
		mid := (first + last) / 2
		if list[mid] == x {
			return mid, true
		}
		if list[mid] < x {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return -1, false
}

func BFS(graph map[string][]string, root string, found string) bool {
	var visited []string

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		node := queue.Front()
		nodeId := node.Value.(string)
		if _, ok := Linear(visited, nodeId); !ok {
			for _, item := range graph[nodeId] {
				if item == found {
					return true
				}
				queue.PushBack(item)
			}
		}
		visited = append(visited, nodeId)
		queue.Remove(node)
	}

	return false
}
