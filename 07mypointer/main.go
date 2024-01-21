package main

import "fmt"

func main() {
	//var ptr *int      //pointer declaration

	a := 20
	var point = &a //location of 20 gets stored in point
	fmt.Println("Value =", *point)
	fmt.Println("Location =", point)
	*point = *point + 2
	fmt.Println("New value =", *point)
	fmt.Println("Value in a =", a)
	fmt.Println("Location of new value =", point)
}
