package main

import "fmt"

func sum(a1, a2 int) int {
	if a1 > a2 {
		return sum(a2, a1)
	}
	s := 0
	for i := a1; i <= a2; i++ {
		if i%2 == 0 {
			s += i
		}
	}
	return s
}

func main() {
	var a1, a2 int
	fmt.Print("Введите 1-е число: ")
	fmt.Scanln(&a1)
	fmt.Print("Введите 2-е число: ")
	fmt.Scanln(&a2)
	fmt.Println(sum(a1, a2))
}
