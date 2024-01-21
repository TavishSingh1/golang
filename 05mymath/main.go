package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdtime := time.Date(2023, time.April, 14, 12, 12, 12, 0, time.UTC)
	fmt.Println(createdtime)
	fmt.Println(createdtime.Format("01-02-2006"))
}
