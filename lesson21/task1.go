package main

import (
	"fmt"
	"math"
)

func calc(x int16, y uint8, z float32) float64 {
	return 2*float64(x) + math.Pow(float64(y), 2) - 3/float64(z)
}

func main() {
	fmt.Println(calc(3, 4, 5))
}
