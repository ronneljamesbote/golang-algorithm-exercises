package sort

func SimpleSortDesc(numbers []int) []int {
	if len(numbers) == 0 || len(numbers) == 1 {
		return numbers
	}

	sorted := make([]int, len(numbers))
	copy(sorted, numbers)

	var tmp int

	for curr := 0; curr < len(sorted); curr++ {
		for next := 0; next < len(sorted); next++ {
			if sorted[curr] > sorted[next] {
				tmp = sorted[next]
				sorted[next] = sorted[curr]
				sorted[curr] = tmp
			}
		}
	}

	return sorted
}

func QuickSortDesc(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	var higher []int = make([]int, 0, len(numbers))
	var lower []int = make([]int, 0, len(numbers))

	pivot := numbers[0]

	for _, value := range numbers[1:] {
		if pivot > value {
			lower = append(lower, value)
			continue
		}

		higher = append(higher, value)
	}

	leftPartition := QuickSortDesc(higher)
	rightPartition := QuickSortDesc(lower)

	var sorted []int
	sorted = append(sorted, leftPartition...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, rightPartition...)

	return sorted
}
