package search

// Linear имеет линейное O(n) время выполнения
func Linear(list []int, x int) (res int, ok bool) {
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
