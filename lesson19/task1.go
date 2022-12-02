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
	i, j, k := 0, 0, 0
	for ; i < firstLen && j < secondLen; k++ {
		if data1[i] < data2[j] {
			result[k] = data1[i]
			i++
		} else {
			result[k] = data2[j]
			j++
		}
	}

	for ; i < firstLen; i++ {
		result[k] = data1[i]
		k++
	}

	for ; j < secondLen; j++ {
		result[k] = data2[j]
		k++
	}
	return
}

func main() {
	data1 := inputFirstArray()
	data2 := inputSecondArray()
	fmt.Println(concatArrays(data1, data2))
}
