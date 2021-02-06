package main

import (
	"fmt"
	"strconv"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	var userName string
	var numOfPlayers string
	fmt.Println("Hi there! What is your name?")
	fmt.Scanln(&userName)

	fmt.Println("Hello, ", userName, ". Please enter the number of players:")
	fmt.Scanln(&numOfPlayers)

	i, err := strconv.Atoi(numOfPlayers)
	if err != nil {
		panic(err)
	}
	numToDeal := 52 / i
	userDeck, remaining := cards.deal(numToDeal)
	cards = remaining

	fmt.Println("Your cards are: ")

	userDeck.print()
}
