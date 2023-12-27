package search

import (
	"testing"
)

var binarySearchNumbers []int32 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

var binarySearchCheck map[int32]int32 = map[int32]int32{
	1:  0,
	9:  8,
	12: 11,
	3:  2,
	7:  6,
}

func TestBinarySearchLoopShouldBeFound(t *testing.T) {
	for v, i := range binarySearchCheck {
		var want int32 = i
		var find int32 = v

		result, err := BinarySearchLoop(find, binarySearchNumbers)

		if err != nil || result != want {
			t.Fatalf(`BinarySearchLoop(%d) = %d, %d given with error %v`, find, want, result, err)
		}
	}
}

func TestBinarySearchLoopShouldNotBeFound(t *testing.T) {
	var find int32 = 13

	result, err := BinarySearchLoop(find, binarySearchNumbers)

	if err == nil {
		t.Fatalf(`BinarySearchLoop(%d) should not be found, %d given`, find, result)
	}
}

func TestBinarySearchRecurse(t *testing.T) {
	for v, i := range binarySearchCheck {
		var want int32 = i
		var find int32 = v

		result, err := BinarySearchRecurse(find, binarySearchNumbers, 0, int32(len(binarySearchNumbers)))

		if err != nil || result != want {
			t.Fatalf(`BinarySearchRecurse(%d) = %d, %d given with error %v`, find, want, result, err)
		}
	}
}
