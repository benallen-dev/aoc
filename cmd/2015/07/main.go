package main

import (
	"fmt"
)

// This question is more about parsing the input than applying the operations

// We know that wires only ever have one source driving them, so we can
// sort the input into a list of operations where we first apply hardcoded
// values, then apply operations to known variables, which will make it possible
// to apply operations to more variables, and so on until all operations have
// been performed. If this sort is not possible, the circuit is invalid anyway.

// Once sorted, apply operations in order to get all values.

func main() {
	ops, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	sortedOps := sortOperations(ops)

	wiresPartOne, err := performOperations(sortedOps)
	if err != nil {
		panic(err)
	}

	res, ok := wiresPartOne["a"]
	if !ok {
		panic("no value for wire 'a'")
	}
	fmt.Println(res)

	// For part 2, we take the result of part 1 and set it as the value of
	// wire 'b', while removing both all other SET operations, and the
	// operation that has 'b' as an output.
	for i, op := range sortedOps {
		if op.Output == "b" {
			sortedOps[i] = &Operation{
				Opcode: "SET",
				Inputs: []string{},
				Output: "b",
				Perform: func(inputs []uint16) (uint16, error) {
					return res, nil
				},
			}
		}

		if op.Opcode == "SET" && op.Output != "b" {
			sortedOps = append(sortedOps[:i], sortedOps[i+1:]...)
		}
	}

	// You might need to re-sort, but in my input, b is already being set by
	// a SET operation, so it's already in the right order.

	wiresPartTwo, err := performOperations(sortedOps)

	res2, ok := wiresPartTwo["a"]
	if !ok {
		panic("no value for wire 'a'")
	}
	fmt.Println(res2)
}
