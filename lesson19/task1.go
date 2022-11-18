package main

import "fmt"

const firstLen, secondLen, totalLen = 4, 5, 9

func inputFirstArray() (data [firstLen]int) {
	for i := range data {
		fmt.Printf("Введите %v элемент: ", i+1)
		fmt.Scanln(&data[i])
	}
	return
}

func inputSecondArray() (data [secondLen]int) {
	for i := range data {
		fmt.Printf("Введите %v элемент: ", i+1)
		fmt.Scanln(&data[i])
	}
	return
}

func concatArrays(data1 [firstLen]int, data2 [secondLen]int) (result [totalLen]int) {
	for i, el := range data1 {
		result[i] = el
	}

	for i, el := range data2 {
		result[i+firstLen] = el
	}
	return
}

func main() {
	data1 := inputFirstArray()
	data2 := inputSecondArray()
	fmt.Println(concatArrays(data1, data2))
}
