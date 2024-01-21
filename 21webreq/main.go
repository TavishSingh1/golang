package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PerformRequest()
}

func PerformRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content lenght", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	Bytecount, _ := responseString.Write(content)
	fmt.Println("Bytecount is: ", Bytecount)
	fmt.Println(responseString.String())

	//fmt.Println(string(content))
}
