package search

import (
	"errors"
	"slices"
)

func BinarySearchLoop(needle int, haystack []int) (int, error) {
	var min int = 0
	var max int = int(len(haystack) - 1)

	var mid int
	var curr int

	for {
		mid = (min + max) / 2

		if min > max || mid >= int(len(haystack)) {
			break
		}

		curr = haystack[mid]

		if curr == needle {
			return mid, nil
		}

		if needle < curr {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return 0, errors.New("Not found")
}

func BinarySearchRecurse(needle int, haystack []int, min int, max int) (int, error) {
	var mid int = (min + max) / 2

	if min > max || mid >= int(len(haystack)) {
		return 0, errors.New("Not found")
	}

	var curr int = haystack[mid]

	if curr == needle {
		return mid, nil
	}

	if needle < curr {
		return BinarySearchRecurse(needle, haystack, min, mid-1)
	}

	return BinarySearchRecurse(needle, haystack, mid+1, max)
}

func backtrace(from int, to int, parent map[int]int) []int {
	var path []int = make([]int, 0, len(parent))
	path = append(path, to)

	for path[len(path)-1] != from {
		path = append(path, parent[path[len(path)-1]])
	}

	slices.Reverse(path)
	return path
}

func BreadthFirstSearch(from int, to int, list map[int][]int) []int {
	if len(list) == 0 {
		return make([]int, 0)
	}

	var queue []int = make([]int, 0, len(list))
	var seen map[int]bool = make(map[int]bool)
	var parent map[int]int = make(map[int]int)

	queue = append(queue, from)
	seen[from] = true

	var cursor int

	for len(queue) > 0 {
		cursor, queue = queue[0], queue[1:]

		if cursor == to {
			return backtrace(from, to, parent)
		}

		for _, edge := range list[cursor] {
			if seen[edge] == false {
				queue = append(queue, edge)
				seen[edge] = true
				parent[edge] = cursor
			}
		}
	}

	return make([]int, 0)
}
