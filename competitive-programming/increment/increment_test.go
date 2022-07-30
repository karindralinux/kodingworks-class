package increment

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrement(t *testing.T) {

	t.Run("one digit case", func(t *testing.T) {
		inputArray := []int{5}
		expected := []int{6}

		actual := Increment(inputArray, nil)

		fmt.Println("Actual One Digit : ", actual)

		assert.Equal(t, expected, actual)
	})

	t.Run("two digit case", func(t *testing.T) {
		inputArray := []int{5, 2}
		expected := []int{5, 3}

		actual := Increment(inputArray, nil)

		fmt.Println("Actual Two Digit : ", actual)

		assert.Equal(t, expected, actual)

	})

	t.Run("two digit with last digit is 9", func(t *testing.T) {
		inputArray := []int{1, 9}
		expected := []int{2, 0}

		actual := Increment(inputArray, nil)

		fmt.Println("Actual Two Digit With Last Digit is 9 : ", actual)

		assert.Equal(t, expected, actual)
	})

	t.Run("more than two digit with 2 digit is 9", func(t *testing.T) {
		inputArray := []int{1, 9, 9}
		expected := []int{2, 0, 0}

		actual := Increment(inputArray, nil)

		fmt.Println("Actual Two Digit With Two Digit is 9 : ", actual)

		assert.Equal(t, expected, actual)
	})

	t.Run("all digit is 9", func(t *testing.T) {
		inputArray := []int{9, 9, 9}
		expected := []int{1, 0, 0, 0}

		actual := Increment(inputArray, nil)

		fmt.Println("Actual All Digit is 9 : ", actual)

		assert.Equal(t, expected, actual)
	})

}
