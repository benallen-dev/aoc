package main

import (
	"bufio"
//	"fmt"
	"os"

	"aoc/internal/stack"
)

// --------------------
//     [D]    
// [N] [C]    
// [Z] [M] [P]
//  1   2   3 

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2
// --------------------

func createStack () stack.Stack[rune] {
	foo := []rune{}
	return stack.New(foo)
}

func readInput(filename string) ([]string, error) {
	var out []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// var us, them rune

		// line := scanner.Text()
		// _, err := fmt.Sscanf(line, "%c %c", &them, &us)
		// if err != nil {
		// 	return nil, err
		// }

		out = append(out, scanner.Text())
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
