package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	content := strings.Split(string(file), "\n")

	// evaluateLights(extractInts(content[0]), extractInstructionType(content[0]))
	count := 0
	for _, line := range content {
		num := evaluateLights(extractInts(string(line)), extractInstructionType(string(line)))
		count += num
	}
	fmt.Println(count)
}

func extractInts(s string) [][]int {

	// Remember when using this function to assign "x" and "y" values to [x, y, x, y] respectively

	var si [][]int
	re := regexp.MustCompile(`(\d+,\w+)`)
	str := re.FindAllString(s, -1)
	for _, v := range str {
		split := strings.Split(string(v), ",")
		var convert []int
		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			convert = append(convert, i)
		}
		si = append(si, convert)
	}
	return si
}

func extractInstructionType(s string) string {
	re := regexp.MustCompile(`toggle|turn on|turn off`)
	ins := re.FindAllString(s, -1)
	return ins[0]
}

func evaluateLights(n [][]int, s string) int {
	startX := n[0][0]
	startY := n[0][1]
	stopX := n[1][0]
	stopY := n[1][1]

	var lightsOn int

	calcX := stopX - startX + 1
	calcY := stopY - startY + 1

	change := (calcX * calcY)

	switch s {
	case "turn on":
		if lightsOn >= change {
			break
		} else if lightsOn < change {
			diff := change - lightsOn
			lightsOn += diff
		}

	case "turn off":
		if lightsOn >= change {
			lightsOn -= change
		} else if lightsOn < change {
			break
		}
	// case "turn off":
	// 	lightsOn -= change

	case "toggle":
		if (1000000 - lightsOn) < change {
			lightsOn += change
		} else {
			lightsOn -= change
		}
	}
	return lightsOn
}
