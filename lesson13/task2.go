package main

import "fmt"

func swap(a1, a2 *int) {
	*a1 += *a2
	*a2 = *a1 - *a2
	*a1 = *a1 - *a2
}

func main() {
	var a1, a2 int
	fmt.Print("Введите 1-е число: ")
	fmt.Scanln(&a1)
	fmt.Print("Введите 2-е число: ")
	fmt.Scanln(&a2)
	swap(&a1, &a2)
	fmt.Println(a1, a2)
}
