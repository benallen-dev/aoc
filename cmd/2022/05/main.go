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
}

