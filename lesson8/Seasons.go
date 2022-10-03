package main

import "fmt"

func main() {
	var mounth, result string

	fmt.Print("Введите месяц: ")
	fmt.Scanln(&mounth)

	switch mounth {
	case "декабрь", "январь", "февраль":
		result = "зима"
	case "март", "апрель", "май":
		result = "весна"
	case "июнь", "июль", "август":
		result = "лето"
	case "сентябрь", "октябрь", "ноябрь":
		result = "осень"
	default:
		result = "Такого месяца нет."
	}

	fmt.Println(result)
}
