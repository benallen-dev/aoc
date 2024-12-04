package main

import (
	"fmt"
)

func main() {
	input, err := readInput("input_example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(input)

	// For all strings in input, find 
	//   - How many XMAS
	//   - How many SAMX
	// Then transpose the input
	//   - do the same trick
	//   - add em together
}

