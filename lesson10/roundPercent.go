package main

import (
	"fmt"
	"math"
)

func main() {
	var incomeSum, percent float64
	var years int

	fmt.Print("Введите сумму: ")
	fmt.Scanln(&incomeSum)
	fmt.Print("Введите ежемесячный процент: ")
	fmt.Scanln(&percent)
	fmt.Print("Введите срок вклада (лет): ")
	fmt.Scanln(&years)

	totalMonth := years * 12
	sum := incomeSum
	bankMargin := 0.0
	monthProfit := 0.0
	monthProfitRound := 0.0

	for i := 1; i <= totalMonth; i++ {
		monthProfit = sum * percent / 100
		monthProfitRound = math.Trunc(monthProfit*100) / 100
		bankMargin += monthProfit - monthProfitRound
		sum += monthProfitRound
	}
	fmt.Println("Сумма итого =", sum, "выгода банка = ", bankMargin)
}
