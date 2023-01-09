package main

import "fmt"

func main() {
	cards := newDeckFromFile("test.txt")

	fmt.Println(cards.toString())

}
