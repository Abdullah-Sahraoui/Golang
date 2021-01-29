package main

import (
	"fmt"
	"io/ioutil"
)

// "da" stands for "delivered at"
var da [][2]int

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	content := string(file)

	// Remember to add starting point as a location.
	da = append(da, [2]int{0, 0})

	// "cl" stands for "current location"
	cl := &[2]int{0, 0}

	for _, v := range content {
		switch string(v) {
		case "^":
			cl[1]++
			checkLocation(cl)
		case "v":
			cl[1]--
			checkLocation(cl)
		case ">":
			cl[0]++
			checkLocation(cl)

		case "<":
			cl[0]--
			checkLocation(cl)
		}
	}

	fmt.Println(len(da))
}

func checkLocation(s *[2]int) {
	for i := 0; i < len(da); i++ {
		if *s == da[i] {
			return
		}
	}

	da = append(da, *s)
}
