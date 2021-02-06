package main

import (
	"regexp"
	"strconv"
	"strings"
)

func extractInts(s string) [][]int {
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
