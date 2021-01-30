package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// For keeping count of number of nice strings
var theseAreNice int

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(file)
	slice := strings.Split(content, "\n")

	for _, str := range slice {
		checkNice(string(str))
	}

	fmt.Println(theseAreNice)
}

func nice1(s string) bool {
	// This is the count - we're checking for 3 vowels
	vowels := [5]string{"a", "e", "i", "o", "u"}
	c := 0
	for _, v := range s {
		for _, n := range vowels {
			if string(v) == n {
				c++
			}
		}
	}

	if c >= 3 {
		return true
	}

	return false
}

func nice2(s string) bool {
	// Var "l" will store the current letter when looping through "s"
	var l string
	for _, v := range s {
		if string(v) == l {
			return true
		}
		l = string(v)
	}
	return false
}

func nice3(s string) bool {
	// Checks if the combinations:
	// 	"ab", "cd", "pq", "xy"
	// are in the string.
	for i := 1; i < len(string(s)); i++ {
		c := string(s[i-1])

		switch string(c) {
		case "a":
			if string(s[i]) == "b" {
				return false
			}
		case "c":
			if string(s[i]) == "d" {
				return false
			}
		case "p":
			if string(s[i]) == "q" {
				return false
			}
		case "x":
			if string(s[i]) == "y" {
				return false
			}
		}
	}
	return true
}

func checkNice(s string) bool {
	if nice1(s) && nice2(s) && nice3(s) {
		theseAreNice++
		return true
	}
	return false
}
