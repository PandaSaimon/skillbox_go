package main

import (
	"flag"
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
	var fileName1, fileName2, fileNameResult string
	flag.StringVar(&fileName1, "file1", defName, "input file1")
	flag.StringVar(&fileName2, "file2", defName, "input file2")
	flag.StringVar(&fileNameResult, "fileResult", defName, "input fileResult")
	flag.Parse()

	result := ""
	if fileName1 == defName {
		log.Fatalln("No data input")
		return
	}
	result = readFile(fileName1)

	if fileName2 != defName {
		result = strings.Join([]string{result, readFile(fileName2)}, "\n")
	}

	if fileNameResult != defName {
		file, err := os.Create(fileNameResult)
		if err != nil {
			log.Fatalln(err)
			return
		}
		defer file.Close()
		file.WriteString(result)
		log.Println("Результат сохранен в", fileNameResult)
	}

	fmt.Println(result)
}
