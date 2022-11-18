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

func reverseArray(data [arrCount]int) [arrCount]int {
	for i := 0; i < arrCount/2; i++ {
		data[i], data[arrCount-i-1] = data[arrCount-i-1], data[i]
	}
	return data
}

func main() {
	data := inputArray()
	data = reverseArray(data)
	fmt.Println(data)
}
