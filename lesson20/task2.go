package main

import "fmt"

const (
	row     = 3
	row_col = 5
	column  = 4
)

func multiplicationMatrix(a [row][row_col]int, b [row_col][column]int) (c [row][column]int) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			for k := 0; k < row_col; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return
}

func main() {
	a := [row][row_col]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{1, 4, 6, 7, 8},
	}

	b := [row_col][column]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 1, 2},
		{3, 4, 5, 6},
		{7, 8, 9, 10},
	}

	fmt.Println(multiplicationMatrix(a, b))
}
