package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	file := strings.Split(string(bs), "\r\n")

	c1, c2 := counting(file)
	fmt.Println("c1 - c2: ", c1-c2)
}

func reg1(s string) int {
	re := regexp.MustCompile(`(\\x)[0-9a-z][0-9a-z]`)
	m := re.FindAllString(s, -1)
	return len(m)
}

func reg2(s string) int {
	re := regexp.MustCompile(`\\"`)
	m := re.FindAllString(s, -1)
	return len(m)
}

func reg3(s string) int {
	re := regexp.MustCompile(`\\\\`)
	m := re.FindAllString(s, -1)
	return len(m)
}

func counting(s []string) (int, int) {
	c1 := 0
	c2 := 0
	for _, str := range s {
		c1 += len(str)
		c2 += len(str) - 2

		r1 := (reg1(str)) * 3
		r2 := reg2(str)
		r3 := reg3(str)

		c2 -= (r1 + r2 + r3)
	}

	return c1, c2
}
