package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Value of dice number is", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can mobe 2 spots")
	case 3:
		fmt.Println("You can mobe 3 spots")
	case 4:
		fmt.Println("You can mobe 4 spots")
	case 5:
		fmt.Println("You can mobe 5 spots")
	case 6:
		fmt.Println("Congrats you can move 6 spots and get an extra chance")
	default:
		fmt.Println("What was that!")
	}
}
