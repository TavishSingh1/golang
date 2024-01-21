package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["py"] = "Python"
	languages["go"] = "Golang"
	fmt.Println("List of all languages", languages)
	fmt.Println("Js is shortform of :", languages["JS"])
	delete(languages, "py")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is, %v\n", key, value)
	}
}
