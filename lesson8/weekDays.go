package main

import "fmt"

func main() {
	var day string

	fmt.Print("Введите день недели: ")
	fmt.Scanln(&day)

	switch day {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	default:
		fmt.Println("Такого дня не существует")
	}
}
