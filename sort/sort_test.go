package sort

import (
	"reflect"
	"testing"
)

var unsortedNumbers []int = []int{5, 12, 1, 6, 3, 4, 7, 7, 11, 8, 9, 2, 10, 1}

var wantDesc []int = []int{12, 11, 10, 9, 8, 7, 7, 6, 5, 4, 3, 2, 1, 1}

func TestSimpleSortDesc(t *testing.T) {
	result := SimpleSortDesc(unsortedNumbers)

	if reflect.DeepEqual(result, wantDesc) == false {
		t.Fatalf("SimpleSortDesc(%v) = %v, %v given", unsortedNumbers, wantDesc, result)
	}
}
