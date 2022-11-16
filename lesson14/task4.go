package main

import "fmt"

var a, b, c = 2, 3, 4

func first(x int) int {
	return x + a
}

func second(x int) int {
	return x + b
}

func third(x int) int {
	return x + c
}

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	a = first(second(third(a)))

	fmt.Println("a:", a)
}
