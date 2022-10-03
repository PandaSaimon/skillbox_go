package main

import "fmt"

func main() {
	result := true
	cost := 5
	bills := []int{5, 5, 10, 10}
	billsCount := len(bills)

	change := make(map[int]int)

	for i := 0; i < billsCount; i++ {
		if bills[i] == cost {
			change[cost]++
		} else {
			result = findChange(&change, bills[i], cost)
			change[bills[i]]++
			if !result {
				break
			}
		}
	}
	fmt.Println(result)
}

func findChange(change *map[int]int, sum int, goal int) bool {
	maxValue := 0
	for key, value := range *change {
		if key > maxValue && key < sum && value != 0 {
			maxValue = key
		}
	}
	if maxValue == 0 {
		return false
	}
	sum -= maxValue
	(*change)[maxValue]--
	if sum == goal {
		return true
	}
	return findChange(change, sum, goal)
}
