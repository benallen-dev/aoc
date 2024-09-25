package main

import (
	"fmt"
	"errors"
	// "strings"
	// "strconv"

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

func getMessage(stacks []stack.Stack[rune]) string {
	// peek each stack to build the message
	msg := ""
	for _, s := range stacks {
		r, err := s.Peek()
		if err != nil {
			panic(err)
		}

		msg += string(r)
	}

	return msg
}


func findDivider(input []string) (int, error) {
	// find the index of the first line that is empty
	emptyLineNum := -1
	for i, v := range input {
		if v == "" {
			emptyLineNum = i
			break
		}
	}

	if emptyLineNum <0 {
		return -1, errors.New("Could not find empty line")
	}

	return emptyLineNum, nil
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	stacks := buildStacks(input)

	div, err := findDivider(input)

	ops := input[div:]
	for _, op := range ops {
		// parse instruction
		var count, source, target int
		fmt.Sscanf(op, "move %d from %d to %d", &count, &source, &target)
	
		// these are 1-indexed in instructions but 0-indexed in the program
		source--
		target--

		for i := 0; i <count; i++ {
			v, err := stacks[source].Pop()
			if err != nil {
				panic(err)
			}
			stacks[target].Push(v)
		}
	}

	fmt.Println("Part 1:", getMessage(stacks))

	// Guess what, the crane moves slices of stacks!
	// Let's start again, but this time just jam another stack between
	// pop and push - that'll reverse the order back to what it should be.

	stacks = buildStacks(input)
	for _, op := range ops {
		// parse instruction
		var count, source, target int
		fmt.Sscanf(op, "move %d from %d to %d", &count, &source, &target)
	
		// these are 1-indexed in instructions but 0-indexed in the program
		source--
		target--

		tmp := stack.New([]rune{})

		for i := 0; i <count; i++ {
			v, err := stacks[source].Pop()
			if err != nil {
				panic(err)
			}
			tmp.Push(v)
		}

		for ;; {
			v, err := tmp.Pop()
			if err != nil {
				// tmp is empty
				break
			}
			stacks[target].Push(v)
		}
	}

	fmt.Println("Part 2:", getMessage(stacks))

}

