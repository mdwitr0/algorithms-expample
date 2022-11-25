package math

func Sum(list []int) int {
	first := list[0]

	if len(list) <= 1 {
		return first
	}

	list = append(list[:0], list[0+1:]...)
	return first + Sum(list)
}
