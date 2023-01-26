package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"students_v2/students_v2"
)

type App struct {
	input IInput
	studs students_v2.Students
}

type IInput interface {
	GetString() (string, error)
}

type Input struct{}

func (input *Input) GetString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	inputBuffer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(inputBuffer, "\r\n"), nil
}

func (a *App) Run() {
	defer a.studs.Print()
	for {
		fmt.Println("Введите имя студента, возраст и оценку")
		if s, e := a.input.GetString(); e == nil {
			a.studs.Add(s)
		} else {
			break
		}
	}
}

func main() {
	studs := students_v2.NewStudents()
	input := &Input{}
	app := &App{input: input, studs: studs}
	app.Run()
}
