package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	// for i:= range days{
	// 	fmt.Println(days)
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	r := 1
	for r < 10 {
		fmt.Println("Value of r is\n", r)
	}
}
