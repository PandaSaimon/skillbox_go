package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const defName = "default name"

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileBody2, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(fileBody2), nil
}

func main() {
	args := os.Args[1:]
	result := ""
	if len(args) <= 0 {
		fmt.Println("No data input")
		return
	}
	result, err := readFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(args) > 1 {
		result2, err := readFile(args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = strings.Join([]string{result, result2}, "\n")

		if len(args) > 2 {
			fileResultName := args[2]
			file, err := os.Create(fileResultName)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			defer file.Close()
			file.WriteString(result)
			fmt.Println("Результат сохранен в", fileResultName)
		}
	}

	fmt.Println(result)
}
