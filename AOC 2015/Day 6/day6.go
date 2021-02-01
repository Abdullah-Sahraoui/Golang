package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	content := string(file)
	fmt.Println(content)
}
