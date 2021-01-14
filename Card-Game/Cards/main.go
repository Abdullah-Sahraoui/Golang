package main

import "fmt"

func main() {
	cards := newDeck()
	cards.toString()
	fmt.Println(cards.toString())
}
