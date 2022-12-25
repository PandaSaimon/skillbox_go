package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func initNArray() (A [n]int) {
	for i := 0; i < n; i++ {
		A[i] = rand.Intn(n)
	}
	fmt.Println(A)
	return
}

func sort(array [n]int) [n]int {
	for i := 1; i < n-1; i++ {
		tmp := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > tmp; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = tmp
	}
	return array
}

func main() {
	array := initNArray()

	//Print result
	fmt.Println("Отсортированный массив:", sort(array))
}
