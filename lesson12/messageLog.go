package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var message string
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл")
	}
	defer file.Close()
	for {
		fmt.Print("Введите сообщение: ")
		fmt.Scanln(&message)
		if message == "exit" {
			break
		}
		file.WriteString(time.Now().Format("2006-01-02 15:04:05") + ": " + message + "\n")
	}
}
