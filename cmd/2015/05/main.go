package main

import (
	"fmt"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	nice := 0
	nicer := 0
	for _, r := range input {
		if r.Nice() {
			nice++
		}

		if r.Nicer() {
			nicer++
		}
	}

	fmt.Println(nice, "nice strings")
	fmt.Println(nicer, "nicer strings")
}

