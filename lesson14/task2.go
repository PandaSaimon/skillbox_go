package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 100

func getRandPoint() (int, int) {
	return rand.Intn(n), rand.Intn(n)
}

func processPoint(x *int, y *int) {
	*x = 2*(*x) + 10
	*y = -3*(*y) - 5
}

func main() {
	rand.Seed(time.Now().UnixNano())

	x1, y1 := getRandPoint()
	x2, y2 := getRandPoint()
	x3, y3 := getRandPoint()

	fmt.Println("x1:", x1, "y1:", y1)
	fmt.Println("x2:", x2, "y2:", y2)
	fmt.Println("x3:", x3, "y3:", y3)

	processPoint(&x1, &y1)
	processPoint(&x2, &y2)
	processPoint(&x3, &y3)

	fmt.Println("x1:", x1, "y1:", y1)
	fmt.Println("x2:", x2, "y2:", y2)
	fmt.Println("x3:", x3, "y3:", y3)
}
