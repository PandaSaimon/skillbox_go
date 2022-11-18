package main

import "fmt"

const arrCount = 10

var (
	evenCount int
	oddCount  int
)

func inputArray() (data [arrCount]int) {
	for i := range data {
		fmt.Printf("Введите %v элемент: ", i+1)
		fmt.Scanln(&data[i])
	}
	return
}

func processElement(value int) {
	if value%2 == 0 {
		evenCount++
	} else {
		oddCount++
	}
}

func main() {
	data := inputArray()

	for _, val := range data {
		processElement(val)
	}

	fmt.Println("Сумма четных чисел: ", evenCount)
	fmt.Println("Сумма нечетных чисел: ", oddCount)
}
