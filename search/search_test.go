package search

import (
	"testing"
)

func TestBinarySearchLoop(t *testing.T) {
	var numbers []int32 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	var check map[int32]int32 = map[int32]int32{
		1:  0,
		9:  8,
		12: 11,
		3:  2,
		7:  6,
	}

	for v, i := range check {
		var want int32 = i
		var find int32 = v

		result, err := BinarySearchLoop(find, numbers)

		if err != nil || result != want {
			t.Fatalf(`BinarySearch(%d) = %d, %d given with error %v`, find, want, result, err)
		}
	}
}
