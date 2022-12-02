package main

import "fmt"

const size = 3

func matrixDeterminate(A [][]int) int {
	if len(A) == 1 {
		return A[0][0]
	}
	sign, d := 1, 0
	for i, x := range A[0] {
		d += sign * x * matrixDeterminate(excludeColumn(A[1:], i))
		sign *= -1
	}
	return d
}

func excludeColumn(a [][]int, i int) [][]int {
	b := make([][]int, len(a))
	n := len(a[0]) - 1
	for j, row := range a {
		r := make([]int, n)
		copy(r[:i], row[:i])
		copy(r[i:], row[i+1:])
		b[j] = r
	}
	return b
}

func main() {
	A := [][]int{
		{9, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrixDeterminate(A))
}
