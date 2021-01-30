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

	content := strings.Split(string(file), "\n")

	var sliceString [][]string

	for _, value := range content {
		dimensions := strings.Split(value, "x")
		sliceString = append(sliceString, dimensions)
	}
	sliceInt := strToInt(sliceString)

	fmt.Println(calcSA(sliceInt))
	fmt.Println(calcRibbon(sliceInt))
}

func calcSA(s [][]int) int {
	var sa int
	var total int
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
				fmt.Println("ERROR WITH THE VALUE: ", value)
				panic(err)
			}
			temp = append(temp, conv)
		}
		slint = append(slint, temp)
		temp = nil

	}

	return slint
}

func calcRibbon(s [][]int) int {
	var p int
	var r int
	var total int
	for _, d := range s {
		sort.Ints(d)
		p = (2 * (d[0])) + (2 * (d[1]))
		r = (d[0]) * (d[1]) * (d[2])
		total = total + p + r
		p, r = 0, 0
	}

	return total
}
