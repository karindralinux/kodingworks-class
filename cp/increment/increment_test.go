package increment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Given array representing a digit (123 = [1, 2, 3])
	write a function to increment the number
*/

func TestIncrement(t *testing.T) {

	t.Run("one digit test case", func(t *testing.T) {
		inputArray := []int{3}
		expectedArray := []int{4}

		actualArray := Increment(inputArray, nil)

		assert.Equal(t, expectedArray, actualArray)
	})

	t.Run("two digit test case", func(t *testing.T) {
		inputArray := []int{3, 2}
		expectedArray := []int{3, 3}

		actualArray := Increment(inputArray, nil)

		assert.Equal(t, expectedArray, actualArray)
	})

	t.Run("two digit with last number is 9", func(t *testing.T) {
		inputArray := []int{1, 9}
		expectedArray := []int{2, 0}

		actualArray := Increment(inputArray, nil)

		assert.Equal(t, expectedArray, actualArray)
	})

	t.Run("more than one number 9", func(t *testing.T) {
		inputArray := []int{9, 9}
		expectedArray := []int{1, 0, 0}

		actualArray := Increment(inputArray, nil)

		assert.Equal(t, expectedArray, actualArray)
	})
}

var total = 1000000
var bestCase []int = []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}

func BenchmarkIncrementBestCase(b *testing.B) {
	for i := 0; i < total; i++ {
		Increment(bestCase, nil)
	}
}

func BenchmarkCowboyIncrementBestCase(b *testing.B) {
	for i := 0; i < total; i++ {
		CowboyIncrement(bestCase)
	}
}
