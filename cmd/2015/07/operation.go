package main

import (
	"fmt"
	"strings"
)

// Normally you would just use `int` instead of `uint16`, but the problem
// specifically states unsigned 16 bit integers, and if you use a signed
// architecture-width int you won't get the same over/underflow behaviour. As
// such every numerical operation is done on uint16s.

type Operation struct {
	// The operation that `Perform` will apply
	Opcode string
	// The input wires to the operation
	Inputs []string
	// The output wire of the operation
	Output string
	// The function to apply
	// Operations have 0, 1, or 2 inputs. We could use [2]uint16 but this way we can check the length and verify we're not being passed garbage
	Perform func(inputs []uint16) (uint16, error)
}

func NewOperation(input string) (*Operation, error) {
	switch {
	case strings.Contains(input, "NOT"):
		return NewNOT(input)
	case strings.Contains(input, "AND"):
		return NewAND(input)
	case strings.Contains(input, "OR"):
		return NewOR(input)
	case strings.Contains(input, "LSHIFT"):
		return NewLSHIFT(input)
	case strings.Contains(input, "RSHIFT"):
		return NewRSHIFT(input)
	default:
		return NewASSIGN(input)
	}
}

// Ideally the below functions wouldn't be exported but this is all in the main package so, meh

// Assigns are a pain in the bum because they are (wire | const) -> wire
func NewASSIGN(input string) (*Operation, error) {
	output := &Operation{
		Opcode: "ASSIGN",
		Inputs: make([]string, 0, 1), // Length 0 or 1, so slice of length 0 with capacity 1
	}

	// This janky try-and-maybe-it-works method is a hint that strings.Contains wasn't the best choice here but it's too late now
	var value uint16 
	_, numErr := fmt.Sscanf(input, "%d -> %s", &value, &output.Output)
	if numErr == nil { // this feels wrong but ok
		// This is a constant value, so it becomes a "SET"
		output.Opcode = "SET"
		output.Perform = func(inputs []uint16) (uint16, error) {
			return value, nil
		}

		return output, nil
	}

	// If we get here, decoding as a number didn't work, so let's try as a string
	output.Inputs = append(output.Inputs, "") // increase size of Inputs to 1
	_, wireErr := fmt.Sscanf(input, "%s -> %s", &output.Inputs[0], &output.Output)
	if wireErr != nil {
		return nil, fmt.Errorf("Could not decode %s: %s", input, wireErr.Error())
	}

	output.Perform = func(inputs []uint16) (uint16, error) {
		if len(inputs) != 1 {
			return 0, fmt.Errorf("ASSIGN wire expects input of length 1, got %v", inputs)
		}
		return inputs[0], nil
	}

	return output, nil
}

func NewNOT(input string) (*Operation, error) {
	output := &Operation{
		Opcode: "NOT",
		Inputs: make([]string, 1, 1),
	}
	_, err := fmt.Sscanf(input, "NOT %s -> %s", &output.Inputs[0], &output.Output)
	if err != nil {
		return nil, fmt.Errorf("Could not decode %s", input)
	}

	output.Perform = func(inputs []uint16) (uint16, error) {
		if len(inputs) != 1 {
			return 0, fmt.Errorf("NOT expects input of length 1, got %v", inputs)
		}

		return ^inputs[0], nil
	}

	return output, nil
}

func NewAND(input string) (*Operation, error) {
	output := &Operation{
		Opcode: "AND",
		Inputs: make([]string, 2, 2),
	}
	_, err := fmt.Sscanf(input, "%s AND %s -> %s", &output.Inputs[0], &output.Inputs[1], &output.Output)
	if err != nil {
		return nil, fmt.Errorf("Could not decode %s", input)
	}

	output.Perform = func(inputs []uint16) (uint16, error) {
		if len(inputs) != 2 {
			return 0, fmt.Errorf("AND expects input of length 2, got %v", inputs)
		}

		return inputs[0] & inputs[1], nil
	}

	return output, nil
}

func NewOR(input string) (*Operation, error) {
	output := &Operation{
		Opcode: "OR",
		Inputs: make([]string, 2, 2),
	}
	_, err := fmt.Sscanf(input, "%s OR %s -> %s", &output.Inputs[0], &output.Inputs[1], &output.Output)
	if err != nil {
		return nil, fmt.Errorf("Could not decode %s", input)
	}

	output.Perform = func(inputs []uint16) (uint16, error) {
		if len(inputs) != 2 {
			return 0, fmt.Errorf("OR expects input of length 2, got %v", inputs)
		}

		return inputs[0] | inputs[1], nil
	}

	return output, nil
}

func NewLSHIFT(input string) (*Operation, error) {
	output := &Operation{
		Opcode: "LSHIFT",
		Inputs: make([]string, 1, 1),
	}

	var amount uint16
	_, err := fmt.Sscanf(input, "%s LSHIFT %d -> %s", &output.Inputs[0], &amount, &output.Output)
	if err != nil {
		return nil, fmt.Errorf("Could not decode %s", input)
	}

	output.Perform = func(inputs []uint16) (uint16, error) {
		if len(inputs) != 1 {
			return 0, fmt.Errorf("OR expects input of length 1, got %v", inputs)
		}

		return inputs[0] << amount, nil
	}

	return output, nil
}

func NewRSHIFT(input string) (*Operation, error) {
	output := &Operation{
		Opcode: "RSHIFT",
		Inputs: make([]string, 1, 1),
	}

	var amount uint16
	_, err := fmt.Sscanf(input, "%s RSHIFT %d -> %s", &output.Inputs[0], &amount, &output.Output)
	if err != nil {
		return nil, fmt.Errorf("Could not decode %s", input)
	}

	output.Perform = func(inputs []uint16) (uint16, error) {
		if len(inputs) != 1 {
			return 0, fmt.Errorf("RSHIFT expects input of length 1, got %v", inputs)
		}

		return inputs[0] >> amount, nil
	}

	return output, nil
}
