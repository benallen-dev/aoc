package main

import (
	"fmt"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	paper := 0
	ribbon := 0

	for _, b := range(input) {
		paper += b.PaperSurface()
		ribbon +=b.RibbonLength()
	}

	fmt.Println("Total area to order:", paper)
	fmt.Println("Total ribbon to order:", ribbon)

}

