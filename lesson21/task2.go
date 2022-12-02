package main

import "fmt"

func deferPrint(A func()) {
	defer A()
	fmt.Println("function run")
}

func main() {
	deferPrint(func() { fmt.Println("Test 1") })
	deferPrint(func() { fmt.Println("Test 2") })
	deferPrint(func() { fmt.Println("Test 3") })
}
