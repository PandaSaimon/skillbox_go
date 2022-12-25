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

func main() {
	array := initNArray()
	f := func(array [n]int) [n]int {
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				if array[i] < array[j] {
					array[i], array[j] = array[j], array[i]
				}
			}
		}
		return array
	}
	//Print result
	fmt.Println("Отсортированный массив:", f(array))
}
