package main

import "fmt"

func isEven(a int) bool {
	return a%2 == 0
}

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	if isEven(a) {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}
