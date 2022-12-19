package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	sentenceCount = 4
	runeCount     = 5
)

func parse(sentences [sentenceCount]string, chars [runeCount]rune) (result [sentenceCount][runeCount]int) {
	if len(sentences) == 0 {
		log.Fatal("Массив строк пустой")
	}
	if len(chars) == 0 {
		log.Fatal("Массив рун пустой")
	}
	for s, sentence := range sentences {
		words := strings.Split(sentence, " ")
		for i2, char := range chars {
			result[s][i2] = -1
			for i := len(words) - 1; i >= 0; i-- {
				index := strings.Index(strings.ToUpper(words[i]), strings.ToUpper(string(char)))
				if index != -1 {
					result[s][i2] = strings.Index(strings.ToUpper(sentence), strings.ToUpper(words[i])) + index
					break
				}
			}
		}
	}
	return
}

func printArray(array [sentenceCount][runeCount]int) {
	for i := 0; i < sentenceCount; i++ {
		fmt.Println(array[i])
	}
}

func main() {
	sentences := [sentenceCount]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	rune := [runeCount]rune{'H', 'E', 'L', 'П', 'М'}

	printArray(parse(sentences, rune))
}
