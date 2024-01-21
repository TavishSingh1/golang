package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to functions")
	greeter()
	greetertwo()
	result := sum(3, 5)
	fmt.Println("Addtion is", result)
	proresult := prosum(2, 3, 4, 5, 5, 3, 22, 2, 2)
	fmt.Println("Proaddition is", proresult)
}

func greeter() {
	fmt.Println("Namaste from Golang")

}

func greetertwo() {
	fmt.Println("Good night")
}

func sum(a int, b int) int {
	s := a + b
	return s
}

func prosum(values ...int) int {
	total := 0
	var val int
	for _, val = range values {
		total += val
	}
	return total
}
