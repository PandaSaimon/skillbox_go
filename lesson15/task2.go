package main

import "fmt"

const arrCount = 10

func inputArray() (data [arrCount]int) {
	for i := range data {
		fmt.Printf("Введите %v элемент: ", i+1)
		fmt.Scanln(&data[i])
	}
	return
}

func reversePrint(data [arrCount]int) {
	fmt.Print("[")
	for i := arrCount - 1; i >= 0; i-- {
		fmt.Print(data[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}

func main() {
	data := inputArray()
	reversePrint(data)
}
