package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("rights.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл")
		return
	}
	file.WriteString("test\n")
	if err := file.Chmod(0444); err != nil {
		fmt.Println(err)
		return
	}
	file.Close()

	fileRead, err := os.Open("rights.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл")
		return
	}
	defer fileRead.Close()

	info, err := fileRead.Stat()
	if err != nil {
		fmt.Println("Ошибка чтения информации о файле")
		return
	}
	buf := make([]byte, info.Size())
	if _, err := io.ReadFull(fileRead, buf); err != nil {
		fmt.Println("Ошибка чтения файла")
		return
	}
	fmt.Printf("%s\n", buf)
	if _, err := fileRead.WriteString("test2\n"); err != nil {
		fmt.Println(err)
		return
	}
}
