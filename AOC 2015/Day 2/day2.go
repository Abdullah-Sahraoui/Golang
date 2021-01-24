package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	content := strings.Split(string(file), "\r\n")

	var sliceString [][]string

	for _, value := range content {
		dimensions := strings.Split(value, "x")
		sliceString = append(sliceString, dimensions)
	}
	sliceInt := strToInt(sliceString)

	fmt.Println(calcSA(sliceInt))
}

func calcSA(s [][]int) int {
	var sa int
	var total int`
	for _, child := range s {
		// For the SURFACE AREA:
		sa = (2 * (child[0]) * (child[1])) + (2 * (child[1]) * (child[2])) + (2 * (child[2]) * (child[0]))

		// Sorting to calculate slack:
		sort.Ints(child)

		// TOTAL INCLUDING SURFACE AREA AND SLACK
		total = total + sa + (child[0] * (child[1]))
		sa = 0
	}
	return total
}

// 1. Split input into []string
// 2. Loop over each string and split at "x"
// Loop over resulting splice:
// 3. Convert strings into integers
// 4. Loop over integers: surface area = (2 * l * w) + (2 * w * h) + (2 * h * l)
// 5. Calculate amount of slack needed =
// 5. Update total to equal total + above calculation

func strToInt(s [][]string) [][]int {
	var slint [][]int
	var temp []int

	for _, sl := range s {
		for _, value := range sl {
			conv, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("ERROR WITH THE value: ", value)
				panic(err)
			}
			temp = append(temp, conv)
		}
		slint = append(slint, temp)
		temp = nil

	}

	return slint
}
