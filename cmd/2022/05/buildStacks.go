package main


import (
	"fmt"
	"strings"
	"strconv"

	"aoc/internal/stack"
)

func buildStacks(input []string) []stack.Stack[rune] {
	// find the index of the first line that is empty
	emptyLineNum := -1
	for i, v := range input {
		if v == "" {
			emptyLineNum = i
			break
		}
	}

	if emptyLineNum <0 {
		panic("Could not find empty line in input")
	}

	// Find all the labels, take the last one and convert to an int to 
	// find how many stacks there are.
	labels := strings.Fields(input[emptyLineNum-1])
	numberOfStacks, err := strconv.Atoi(labels[len(labels)-1])
	if err != nil {
		panic(err)
	}

	stacks := make([]stack.Stack[rune], numberOfStacks, numberOfStacks)
	// Apply a better formatter to each rune stack, but do it by exposing a
	// bunch of Stack internals because I didn't bother to make stacks really
	// good ok don't judge me I'll make it an option builder at some point
	for i := range stacks {
		stacks[i].Formatter = func(r rune) string {
			return fmt.Sprintf("%c", r)
		}
	}

	// Now that we have the stacks, time to populate them.
	// Line by line, starting at emptyLineNum - 2, find runes that 
	// are on stacks and push them.

	for i := emptyLineNum -2; i >= 0; i = i-1 {
		line := input[i]
		for j := range numberOfStacks {
			loc := (j*4) + 1 // position of character on the line, 1, 5, 9, 13, etc
			r := rune(line[loc])
			if r != ' ' {
				stacks[j].Push(r)
			}
		}
	}

	return stacks
}
