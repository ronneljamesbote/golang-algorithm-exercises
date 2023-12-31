package search

import (
	"errors"
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
