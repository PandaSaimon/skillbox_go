package main

import "fmt"

const arrCount = 10

func main() {
	var data [arrCount]int

	for i := range data {
		fmt.Printf("Введите %v элемент: ", i+1)
		fmt.Scanln(&data[i])
	}

	evenCount := 0
	oddCount := 0

	for _, val := range data {
		if val%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	fmt.Println("Сумма четных чисел: ", evenCount)
	fmt.Println("Сумма нечетных чисел: ", oddCount)
}
