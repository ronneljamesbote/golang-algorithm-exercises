package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var numbers []int = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

func TestBinarySearchLoop(t *testing.T) {
	var assert *assert.Assertions = assert.New(t)

	val1, _ := BinarySearchLoop(69, numbers)
	assert.Equal(val1, 3, "should be equal")

	_, err1 := BinarySearchLoop(1336, numbers)
	assert.NotNil(err1, "1336 should not be found")

	val2, _ := BinarySearchLoop(69420, numbers)
	assert.Equal(val2, 10, "should be equal")

	_, err2 := BinarySearchLoop(69421, numbers)
	assert.NotNil(err2, "69421 should not be found")

	val3, _ := BinarySearchLoop(1, numbers)
	assert.Equal(val3, 0, "should be equal")

	_, err3 := BinarySearchLoop(0, numbers)
	assert.NotNil(err3, "0 should not be found")
}

func TestBinarySearchRecurse(t *testing.T) {
	var assert *assert.Assertions = assert.New(t)

	val1, _ := BinarySearchRecurse(69, numbers, 0, len(numbers))
	assert.Equal(val1, 3, "should be equal")

	_, err1 := BinarySearchRecurse(1336, numbers, 0, len(numbers))
	assert.NotNil(err1, "1336 should not be found")

	val2, _ := BinarySearchRecurse(69420, numbers, 0, len(numbers))
	assert.Equal(val2, 10, "should be equal")

	_, err2 := BinarySearchRecurse(69421, numbers, 0, len(numbers))
	assert.NotNil(err2, "69421 should not be found")

	val3, _ := BinarySearchRecurse(1, numbers, 0, len(numbers))
	assert.Equal(val3, 0, "should be equal")

	_, err3 := BinarySearchRecurse(0, numbers, 0, len(numbers))
	assert.NotNil(err3, "0 should not be found")
}
