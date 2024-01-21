package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payementid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T \n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)
}
