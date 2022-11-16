package main

import "fmt"

func add(a int) (c int) {
	c = a + 5
	return
}

func multiply(a int) (c int) {
	c = a * 2
	return
}

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	a = add(a)
	a = multiply(a)

	fmt.Println("a:", a)
}
