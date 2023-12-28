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
