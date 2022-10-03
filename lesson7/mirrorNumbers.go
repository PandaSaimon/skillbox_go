package main

import (
	"fmt"
	"strconv"
)

func main() {
	minNumber := 100000
	maxNumber := 999999
	mirrorCount := 0

	for i := minNumber; i <= maxNumber; i++ {
		if isMirror(i) {
			mirrorCount++
		}
	}

	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999:", mirrorCount)
}

func isMirror(digit int) bool {
	result := true
	strDigit := strconv.Itoa(digit)
	for i := 0; i < len(strDigit)/2; i++ {
		if strDigit[i] != strDigit[len(strDigit)-i-1] {
			result = false
			break
		}
	}
	return result
}
