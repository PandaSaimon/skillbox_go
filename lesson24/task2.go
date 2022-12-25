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

func lastIndexRune(sentence []rune, find rune) int {
	index := -1
	for i := len(sentence) - 1; i >= 0; i-- {
		log.Println(strings.ToUpper(string(sentence[i])), strings.ToUpper(string(find)))
		if strings.ToUpper(string(sentence[i])) == strings.ToUpper(string(find)) {
			index = i
			break
		}
	}
	log.Println(" ")
	return index
}

func parse(sentences [sentenceCount]string, chars [runeCount]rune) (result [sentenceCount][runeCount]int) {
	if len(sentences) == 0 {
		log.Fatal("Массив строк пустой")
	}
	if len(chars) == 0 {
		log.Fatal("Массив рун пустой")
	}
	for s, sentence := range sentences {
		for i2, char := range chars {
			result[s][i2] = -1
			if strings.Index(strings.ToUpper(sentence), strings.ToUpper(string(char))) != -1 {
				result[s][i2] = lastIndexRune([]rune(sentence), char)
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
