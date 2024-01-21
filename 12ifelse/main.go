package main

import "fmt"

func main() {
	logincount := 23
	var result string

	if logincount > 10 {
		result = "Regular user"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Not less than 10")
	}
}
