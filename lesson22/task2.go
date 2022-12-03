package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func initNArray() (A [n]int) {
	for i := 0; i < n; i++ {
		A[i] = 10*i + rand.Intn(n)
		/*A[i] = i
		if i%2 != 0 {
			A[i]++
		}*/
	}
	fmt.Println(A)
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func findFirstPosition(find int, A [n]int) (position int) {
	position = -1
	first := 0
	middle := n / 2
	last := n
	for first < last {
		if A[middle] == find {
			if middle == first || A[middle-1] > find {
				position = middle
				break
			}
			last = middle - 1
		} else if A[middle] > find {
			last = middle - 1
		} else {
			first = middle + 1
		}
		middle = (last + first) / 2
	}
	return
}

func main() {
	array := initNArray()
	var find int

	//Input data
	fmt.Print("Введите число: ")
	fmt.Scanln(&find)

	//Print result
	fmt.Println("Первая позиция:", findFirstPosition(find, array))
}
