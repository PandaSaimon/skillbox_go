package main

import "fmt"

func main() {
	var cnt int
	fmt.Print("Введите количество пар скобок: ")
	fmt.Scanln(&cnt)
	answers := make([]string, cnt)
	for i := 0; i < cnt; i++ {
		if i > 0 {
			answers[i] += "("
		}
		for j := i; j < cnt; j++ {
			answers[i] += addParentheses(i+1, i+1)
		}
		if i > 0 {
			answers[i] += ")"
		}
	}
	fmt.Println(answers)
}

func addParentheses(count int, total int) string {
	var res string
	if count != total || count == 1 {
		res += "("
	}
	if count > 1 {
		res += addParentheses(count-1, total)
	}
	if count != total || count == 1 {
		res += ")"
	}
	return res
}
