package main

import (
	"fmt"
	"regexp"

	"strings"
)

func getMuls(input string) []string {
	mulRegex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return mulRegex.FindAllString(input, -1)
}

func performMul(input string) int {
	var a, b int
	fmt.Sscanf(input, "mul(%d,%d)", &a, &b)

	return a * b
}

func findDos(input string) [][]int {
	doRegex := regexp.MustCompile(`do\(\)`)
	return doRegex.FindAllStringIndex(input, -1)
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	var total int
	muls := getMuls(input)
	for _, mul := range muls {
		total += performMul(mul)
	}

	fmt.Printf("Part 1: %d\n", total)

	// Part 2:
	// 10 take input until "don't()", find mul()s
	// 20 find next do() unless EOF
	// 30 goto 10

	// split into parts that end with "don't()"
	parts := strings.Split(input, "don't()")

	var activeInput = make([]string, 0)
	for i, part := range parts {
		if i == 0 { // the first do() is implicit
			activeInput = append(activeInput, part)
			continue
		}

		// find the first do() - anything between this and \n is active
		dos := findDos(part)
		// Note: findDos can return nil for no match, if there's no
		// `do()` inbetween these `don't()`s we can safely move on.
		if len(dos) == 0 {
			continue
		}
		
		startat := dos[0][1] // include everything after 'do()'
		activeInput = append(activeInput, part[startat:])
	}

	var total2 int
	muls2 := getMuls(strings.Join(activeInput, ""))
	for _, mul := range muls2 {
		total2 += performMul(mul)
	}

	fmt.Printf("Part 2: %d\n", total2)
}

