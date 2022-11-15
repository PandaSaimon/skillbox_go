package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл")
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка чтения информации о файле")
		return
	}
	buf := make([]byte, info.Size())
	if _, err := io.ReadFull(file, buf); err != nil {
		fmt.Println("Ошибка чтения файла")
		return
	}
	fmt.Printf("%s\n", buf)
}
