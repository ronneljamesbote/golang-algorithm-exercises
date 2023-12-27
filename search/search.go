package search

import (
	"errors"
)

func BinarySearchLoop(needle int32, haystack []int32) (int32, error) {
	var min int32 = 0
	var max int32 = int32(len(haystack))
	var mid int32 = (min + max) / 2

	var curr int32 = haystack[mid]

	for min <= max {
		if curr == needle {
			return int32(mid), nil
		}

		if needle < curr {
			max = mid - 1
		} else {
			min = mid + 1
		}

		mid = (min + max) / 2
		curr = haystack[mid]
	}

	return 0, errors.New("Needle not found")
}
