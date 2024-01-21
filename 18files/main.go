package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - Learncodeonline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data in the file is\n", string(databyte))
}

func checkNilCheckker(err error) {
	if err != nil {
		panic(err)
	}
}
