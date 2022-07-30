package increment

// Given array representing a digit (123 = [1, 2, 3])
// write a function to increment the number

// [1, 2, 3] => [1, 2, 4]
// [2, 9, 0] => [2, 9, 1]

func Increment(arr []int, mapIndex map[string]int) []int {

	var lastIndex int

	if mapIndex == nil {
		lastIndex = len(arr) - 1
	} else {
		lastIndex = mapIndex["lastIndex"]
	}

	if arr[lastIndex] == 9 && lastIndex-1 >= 0 {
		arr[lastIndex] = 0
		return Increment(arr, map[string]int{"lastIndex": lastIndex - 1})
	} else if arr[lastIndex] == 9 && lastIndex-1 < 0 {
		arr[lastIndex] = 0
		return append([]int{1}, arr...)
	}

	arr[lastIndex] += 1

	return arr
}
