package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func initNArray() (A [size]int) {
	for i := 0; i < size; i++ {
		A[i] = rand.Intn(10 * size)
	}
	fmt.Println(A)
	return
}

func evenAndOdd(A [size]int) (even []int, odd []int) {
	for _, a := range A {
		if a%2 == 0 {
			even = append(even, a)
		} else {
			odd = append(odd, a)
		}
	}
	return
}

func main() {
	array := initNArray()
	even, odd := evenAndOdd(array)

	fmt.Println("Четные:")
	fmt.Println(even)
	fmt.Println("Нечетные:")
	fmt.Println(odd)
}
