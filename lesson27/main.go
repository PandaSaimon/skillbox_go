package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"students/students"
)

func Scan() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	inputBuffer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(inputBuffer, "\r\n"), nil
}

func main() {
	studentsMap := make(map[string]*students.Student)
	defer func() {
		for _, v := range studentsMap {
			v.Print()
		}
	}()

	for {
		fmt.Println("Введите имя студента, возраст и оценку")
		input, e := Scan()
		if e != nil {
			return
		}
		s, err := students.SetValuesFromString(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		studentsMap[s.GetName()] = &s
	}
}
