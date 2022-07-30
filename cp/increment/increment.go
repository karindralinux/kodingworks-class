package increment

import (
	"fmt"
	"strconv"
	"strings"
)

func Increment(arr []int, mapIndex map[string]interface{}) []int {

	var lastIndex int

	if mapIndex["lastIndex"] == nil {
		lastIndex = len(arr) - 1
	} else {
		lastIndex = mapIndex["lastIndex"].(int)
	}

	if arr[lastIndex] == 9 && lastIndex-1 >= 0 {
		arr[lastIndex] = 0
		return Increment(arr, map[string]interface{}{"lastIndex": lastIndex - 1})
	} else if arr[lastIndex] == 9 && lastIndex-1 < 0 {
		arr[lastIndex] = 0
		return append([]int{1}, arr...)
	}

	arr[lastIndex] += 1
	return arr
}

func CowboyIncrement(arr []int) []int {

	strArr := arrayToString([]int{2, 1, 1}, "")

	number, err := strconv.ParseInt(strArr, 10, 64)

	if err != nil {
		panic(err)
	}

	number += 1

	strNumber := strconv.Itoa(int(number))

	arrStrNumber := strings.Split(strNumber, "")

	var arrNumbers []int

	for _, v := range arrStrNumber {
		numberConvert, _ := strconv.ParseInt(v, 10, 64)
		arrNumbers = append(arrNumbers, int(numberConvert))
	}

	return arrNumbers
}

func arrayToString(arr []int, separator string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), separator), "[]")
}
