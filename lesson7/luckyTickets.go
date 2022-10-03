package main

import "fmt"

func main() {
	min := 100000
	max := 999999

	maxTickets := 0
	lastTicket := min

	for i := min; i <= max; i++ {
		if isLuckyTicket(i) {
			maxTickets = i - lastTicket
			lastTicket = i
		}
	}

	fmt.Println("Минимальное количество билетов, которое нужно купить, чтобы среди них оказался счастливый:", maxTickets)
}

func isLuckyTicket(ticketNumber int) bool {
	firstDigit := ticketNumber / 100000
	secondDigit := ticketNumber / 10000 % 10
	thirdDigit := ticketNumber / 1000 % 10
	fourthDigit := ticketNumber / 100 % 10
	fifthDigit := ticketNumber / 10 % 10
	sixDigit := ticketNumber % 10

	if firstDigit+secondDigit+thirdDigit == fourthDigit+fifthDigit+sixDigit {
		return true
	}

	return false
}
