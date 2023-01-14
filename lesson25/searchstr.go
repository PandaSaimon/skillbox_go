package main

import (
	"flag"
	"fmt"
)

func find(str, substr string) bool {
	result := false
	if len(str) == 0 || len(substr) == 0 {
		return result
	}

	subrune := []rune(substr)

	findFlag := 0
	for _, v := range []rune(str) {
		if findFlag == len(subrune) {
			break
		}
		if v == subrune[findFlag] {
			findFlag++
		} else if findFlag != 0 && v != subrune[findFlag] {
			break
		}
	}

	if len(subrune) == findFlag {
		result = true
	}

	return result
}

func main() {
	var str, substr string
	flag.StringVar(&str, "str", "default string", "input string")
	flag.StringVar(&substr, "substr", "default string", "input substring")
	flag.Parse()

	isFind := find(str, substr)
	fmt.Println("Result is:", isFind)
}
