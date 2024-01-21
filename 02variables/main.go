package main

import "fmt"

const LoginToken string = "abcdefg" //if the first letter of constant name in captial then it is a public constant

func main() {
	var username string = "Tavish"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isloggedin bool = false
	fmt.Println(isloggedin)
	fmt.Printf("Variable is of type: %T \n", isloggedin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallFloat float32 = 34.4958594892389
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var largeFloat float64 = 34.4958594892389
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	var number int
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

	var s string
	fmt.Println(s)
	fmt.Printf("Variable is of type: %T \n", s)

	//implicit type
	var name = "Tavish"
	fmt.Println("Implicit type declaration:", name)

	//no var style(allowed only inside a function)
	numberOfuser := 3000
	fmt.Println(numberOfuser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
