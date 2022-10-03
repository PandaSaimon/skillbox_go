package main

import "fmt"

func main() {
	var rowCount int
	fmt.Print("Введите количество строк: ")
	fmt.Scanln(&rowCount)

	columnCount := 2*rowCount - 1

	for row := 1; row <= rowCount; row++ {
		treeSimbolCount := 2*row - 1
		for column := 1; column <= columnCount; column++ {
			if (columnCount-treeSimbolCount)/2 >= column || (columnCount+treeSimbolCount)/2 < column {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
