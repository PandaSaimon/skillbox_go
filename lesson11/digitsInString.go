package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "a10 10 20b 20 30c30 30 dd"
	s = strings.Trim(s, " ")
	spaceIndex := strings.Index(s, " ")

	for spaceIndex > 0 {
		firstWorld := s[:spaceIndex]
		i, err := strconv.Atoi(firstWorld)
		if err == nil {
			fmt.Println(i)
		}
		s = s[spaceIndex+1:]
		s = strings.Trim(s, " ")
		spaceIndex = strings.Index(s, " ")
	}
}
