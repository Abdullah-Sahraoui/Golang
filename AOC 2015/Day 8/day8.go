package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	file := strings.Split(string(input), "\n")

	var cc int
	var mc int

	for _, l := range file {
		cc += len(l)
		unquoted, err := strconv.Unquote(l)
		if err != nil {
			panic(err)
		}
		mc += len(unquoted)
	}
	fmt.Println(cc - mc)
}
