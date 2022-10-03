package main

import "fmt"

func main() {
	blackCell := " "
	whiteCell := "*"
	var width, height int

	fmt.Print("Введите ширину: ")
	fmt.Scanln(&width)
	fmt.Print("Введите высоту: ")
	fmt.Scanln(&height)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(blackCell)
			} else {
				fmt.Print(whiteCell)
			}
		}
		fmt.Println()
	}
}
