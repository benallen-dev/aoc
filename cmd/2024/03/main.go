package main

import (
	"fmt"
	"regexp"
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
	doRegex := regexp.MustCompile(`do\(\d+\)`)
	return doRegex.FindAllStringIndex(input, -1)
}

func findDonts(input string) [][]int {
	dontRegex := regexp.MustCompile(`don't\(\)`)
	return dontRegex.FindAllStringIndex(input, -1)
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

	fmt.Println(total)

	// Ha! Serves me right for using regexes
	// 10 take input until "don't()", find mul()s
	// 20 find next do() unless EOF
	// 30 goto 10

	// doIdxs := findDos(input)
	// dontIdxs := findDonts(input)

	
		

}

