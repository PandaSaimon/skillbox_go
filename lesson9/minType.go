package main

import (
	"fmt"
	"math"
)

func main() {
	var a1, a2 int16
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a1)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&a2)

	var result int32 = int32(a1) * int32(a2)
	if result < 0 {
		fmt.Println(findNegativeType(result))
	} else {
		fmt.Println(findPositiveType(result))
	}
}

func findPositiveType(digit int32) string {
	result := "uint32"
	switch {
	case digit <= math.MaxUint8:
		result = "uint8"
	case digit <= math.MaxUint16:
		result = "uint16"
	}
	return result
}

func findNegativeType(digit int32) string {
	result := "int32"
	switch {
	case digit >= math.MinInt8:
		result = "int8"
	case digit >= math.MinInt16:
		result = "int16"
	}
	return result
}
