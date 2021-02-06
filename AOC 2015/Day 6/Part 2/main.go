package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	file := strings.Split(string(bs), "\n")

	testing(file)
}

func testing(s []string) {
	var lights [][]int
	for i := 0; i < 1000; i++ {
		var row []int
		for j := 0; j < 1000; j++ {
			row = append(row, 0)
		}
		lights = append(lights, row)
	}

	for _, line := range s {
		coords := extractInts(line)
		rule := extractInstructionType(line)

		for k := coords[0][1]; k < coords[1][1]+1; k++ {
			for l := coords[0][0]; l < coords[1][0]+1; l++ {
				switch string(rule) {
				case "toggle":
					lights[k][l] += 2
				case "turn on":
					lights[k][l]++
				case "turn off":
					if lights[k][l] > 0 {
						lights[k][l]--
					}
				}
			}
		}
	}
	count := 0
	for _, v := range lights {
		for _, q := range v {
			count += q
		}
	}
	fmt.Println(count)
}
