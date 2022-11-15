package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var message string
	var b bytes.Buffer
	fileName := "log2.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Не удалось создать файл")
	}
	for {
		fmt.Print("Введите сообщение: ")
		fmt.Scanln(&message)
		if message == "exit" {
			break
		}
		b.WriteString(time.Now().Format("2006-01-02 15:04:05") + ": " + message + "\n")
	}
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		fmt.Println(err)
	}
	file.Close()

	fileRead, err := os.Open("log2.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл")
		return
	}
	defer fileRead.Close()

	fileBody, err := ioutil.ReadAll(fileRead)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nЛог: ")
	fmt.Println(string(fileBody))
}
