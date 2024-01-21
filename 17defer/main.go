package main

import "fmt"

func main() {
	defer fmt.Println("one")
	fmt.Println("Hello")
	defer fmt.Println("two")
	defer fmt.Println("three")
	MyDefer()
}

func MyDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
