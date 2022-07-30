package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	strArr := arrayToString([]int{2, 1, 1}, "")

	number, err := strconv.ParseInt(strArr, 10, 64)

	if err != nil {
		panic(err)
	}

	number += 1

	strNumber := strconv.Itoa(int(number))

	arrStrNumber := strings.Split(strNumber, "")

	var arrNumbers []int64

	for _, v := range arrStrNumber {
		numberConvert, _ := strconv.ParseInt(v, 10, 64)
		arrNumbers = append(arrNumbers, numberConvert)
	}

	fmt.Println(arrNumbers)
}

func Add(a float64, b float64) float64 {
	return a + b
}

func ArrIncrement(arr []int) {

	strArr := arrayToString(arr, "")
	fmt.Println(strArr)
}

func arrayToString(arr []int, separator string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), separator), "[]")
}
