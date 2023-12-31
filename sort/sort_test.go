package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var unsorted []int = []int{9, 3, 7, 4, 69, 420, 0, 42, 7}

var sorted []int = []int{420, 69, 42, 9, 7, 7, 4, 3, 0}

func TestSimpleSortDesc(t *testing.T) {
	var assert *assert.Assertions = assert.New(t)

	result := SimpleSortDesc(unsorted)

	assert.Equal(sorted, result, "should be equal")
}

func TestQuickSortDesc(t *testing.T) {
	var assert *assert.Assertions = assert.New(t)

	result := QuickSortDesc(unsorted)

	assert.Equal(sorted, result, "should be equal")
}
