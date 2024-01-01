package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var numbers []int = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

func TestBinarySearchLoop(t *testing.T) {
	var assert *assert.Assertions = assert.New(t)

	val1, _ := BinarySearchLoop(69, numbers)
	assert.Equal(3, val1, "should be equal")

	_, err1 := BinarySearchLoop(1336, numbers)
	assert.NotNil(err1, "1336 should not be found")

	val2, _ := BinarySearchLoop(69420, numbers)
	assert.Equal(10, val2, "should be equal")

	_, err2 := BinarySearchLoop(69421, numbers)
	assert.NotNil(err2, "69421 should not be found")

	val3, _ := BinarySearchLoop(1, numbers)
	assert.Equal(0, val3, "should be equal")

	_, err3 := BinarySearchLoop(0, numbers)
	assert.NotNil(err3, "0 should not be found")
}

func TestBinarySearchRecurse(t *testing.T) {
	var assert *assert.Assertions = assert.New(t)

	val1, _ := BinarySearchRecurse(69, numbers, 0, len(numbers))
	assert.Equal(3, val1, "should be equal")

	_, err1 := BinarySearchRecurse(1336, numbers, 0, len(numbers))
	assert.NotNil(err1, "1336 should not be found")

	val2, _ := BinarySearchRecurse(69420, numbers, 0, len(numbers))
	assert.Equal(10, val2, "should be equal")

	_, err2 := BinarySearchRecurse(69421, numbers, 0, len(numbers))
	assert.NotNil(err2, "69421 should not be found")

	val3, _ := BinarySearchRecurse(1, numbers, 0, len(numbers))
	assert.Equal(0, val3, "should be equal")

	_, err3 := BinarySearchRecurse(0, numbers, 0, len(numbers))
	assert.NotNil(err3, "0 should not be found")
}

func TestBreadthFirstSearch(t *testing.T) {
	var graphList map[int][]int = make(map[int][]int)
	graphList[0] = []int{1, 2}
	graphList[1] = []int{4}
	graphList[2] = []int{3}
	graphList[3] = []int{}
	graphList[4] = []int{1, 3, 5}
	graphList[5] = []int{2, 6}
	graphList[6] = []int{3}

	var assert *assert.Assertions = assert.New(t)

	val1 := BreadthFirstSearch(0, 6, graphList)
	assert.Equal([]int{0, 1, 4, 5, 6}, val1, "should be equal")
}
