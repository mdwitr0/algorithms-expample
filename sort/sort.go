package sort

func findSmallest(list []int) int {
	smallestIndex := 0
	smallest := list[smallestIndex]

	for i, item := range list {
		if item < smallest {
			smallestIndex = i
			smallest = item
		}
	}

	return smallestIndex
}

func Selection(list []int) []int {
	size := len(list)
	sortedList := make([]int, size)

	for i := range list {
		position := findSmallest(list)
		sortedList[i] = list[position]

		list = append(list[:position], list[position+1:]...)
	}

	return sortedList
}

func Quick(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[len(list)/2]

	var less []int
	var greater []int

	for _, item := range list {
		if item < pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	less = append(Quick(less), pivot)
	greater = append(greater)

	return append(less, greater...)
}
