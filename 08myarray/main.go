package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"
	fmt.Println("Fruit list is = ", fruitlist)
	fmt.Println("Length of list is:", len(fruitlist))

	var veglist = [5]string{"potato", "onion", "mushroom"}
	fmt.Println("Fruit list is = ", veglist)
	fmt.Println("Length of list is:", len(veglist))
}
