package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	s = strings.Trim(s, " ")
	count := 0
	spaceIndex := strings.Index(s, " ")

	for spaceIndex > 0 {
		firstWorld := s[:spaceIndex]
		if firstWorld == strings.Title(firstWorld) {
			count++
		}
		s = s[spaceIndex+1:]
		s = strings.Trim(s, " ")
		spaceIndex = strings.Index(s, " ")
	}

	if s == strings.Title(s) {
		count++
	}
	fmt.Println("Количество слов с большой буквы:", count)
}
