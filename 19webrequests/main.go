package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Respomse is of type: %T \n", response)

	defer response.Body.Close() //caller's resposibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}
