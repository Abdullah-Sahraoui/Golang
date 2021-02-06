package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	file := strings.Split(string(bs), "\n")

	testing(file)
}

func testing(s []string) {
	var lights [][]bool
	for i := 0; i < 1000; i++ {
		var row []bool
		for j := 0; j < 1000; j++ {
			row = append(row, false)
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
					lights[k][l] = !lights[k][l]
				case "turn on":
					lights[k][l] = true
				case "turn off":
					lights[k][l] = false
				}
			}
		}
	}
	count := 0
	for _, v := range lights {
		for _, q := range v {
			if q {
				count++
			}
		}
	}
	fmt.Println(count)
}
