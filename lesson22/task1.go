package main

import (
	"fmt"
	"math/rand"
	"time"
)

const length = 10

func initArray() (A [length]int) {
	for i := 0; i < length; i++ {
		A[i] = rand.Intn(length)
	}
	fmt.Println(A)
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func findPosition(find int, A [length]int) (result int) {
	result = -1
	for i, a := range A {
		if a == find {
			result = i + 1
			break
		}
	}
	return
}

func findLength(find int, A [length]int) int {
	pos := findPosition(find, A)
	if pos == -1 {
		return pos
	}
	return length - pos
}

func main() {
	//Initialize variables
	array := initArray()
	var find int

	//Input data
	fmt.Print("Введите число: ")
	fmt.Scanln(&find)

	//Print result
	fmt.Println("Количество символов до конца массива:", findLength(find, array))
}
