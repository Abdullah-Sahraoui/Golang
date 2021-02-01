package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	content := strings.Split(string(file), "\n")

	fmt.Println(niceString(content))
}

func niceString(s []string) int {
	count := 0
	for _, str := range s {
		if nice1(str) && nice2(str) {
			count++
		}
	}
	return count
}

func nice1(s string) bool {
	for i := 2; i < len(s); i++ {
		j := string(s[i-2]) + string(s[i-1])
		test := strings.Count(s, j)

		if test > 1 {
			return true
		}
	}
	return false
}

func nice2(s string) bool {
	for i := 2; i < len(s); i++ {
		if string(s[i-2]) == string(s[i]) {
			return true
		}
	}
	return false
}
