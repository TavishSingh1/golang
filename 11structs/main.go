package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang,  no parents
	hitesh := User{"Hitesh", "hitesh@gmail.com", true, 18}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	hitesh.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is: ", u.Email)
}
