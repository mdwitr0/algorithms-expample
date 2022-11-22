package search

// LinearSearch имеет время выполнения O(n)
func LinearSearch(list []int, x int) (res int, ok bool) {
	for index, item := range list {
		if item == x {
			return index, true
		}
	}
	return 0, false
}