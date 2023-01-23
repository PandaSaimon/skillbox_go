package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const defName = "default name"

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	defer file.Close()
	fileBody2, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return string(fileBody2)
}

func main() {
	args := os.Args[1:]
	result := ""
	if len(args) <= 0 {
		log.Fatalln("No data input")
		return
	}
	result = readFile(args[0])

	if len(args) > 1 {
		result = strings.Join([]string{result, readFile(args[1])}, "\n")

		if len(args) > 2 {
			fileResultName := args[2]
			file, err := os.Create(fileResultName)
			if err != nil {
				log.Fatalln(err)
				return
			}
			defer file.Close()
			file.WriteString(result)
			log.Println("Результат сохранен в", fileResultName)
		}
	}

	fmt.Println(result)
}
