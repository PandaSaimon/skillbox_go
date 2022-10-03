package main

import (
	"fmt"
	"math"
)

func main() {
	var x, accuracy int
	fmt.Print("Введите х: ")
	fmt.Scanln(&x)
	fmt.Print("Введите порядок точности: ")
	fmt.Scanln(&accuracy)

	epsilon := 1.0 / math.Pow(10, float64(accuracy))
	ex, lastEx, n := 0.0, 0.0, 0
	fact := 1

	for {
		if n > 0 {
			fact *= n
		}
		lastEx = ex
		ex += math.Pow(float64(x), float64(n)) / float64(fact)

		if math.Abs(ex-lastEx) < epsilon {
			break
		}
		n++
	}
	fmt.Println("ex = ", ex, "контрольное значение =", math.Exp(float64(x)))
}
