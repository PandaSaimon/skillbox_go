package main

import "fmt"

const arrCount = 6

func inputArray() (data [arrCount]int) {
	for i := range data {
		fmt.Printf("Введите %v элемент: ", i+1)
		fmt.Scanln(&data[i])
	}
	return
}

func sortArray(data [arrCount]int) [arrCount]int {
	for i := 0; i < arrCount; i++ {
		for j := 0; j < arrCount-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	return data
}

func main() {
	data := sortArray(inputArray())
	fmt.Println(data)
}
