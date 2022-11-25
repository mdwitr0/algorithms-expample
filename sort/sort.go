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

func SelectionSort(list []int) []int {
	size := len(list)
	sortedList := make([]int, size)

	for i := 0; i < size; i++ {
		position := findSmallest(list)
		sortedList[i] = list[position]

		list = append(list[:position], list[position+1:]...)
	}

	return sortedList
}
