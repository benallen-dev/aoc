package main

// I could sort in place using the sort package but I'm lazy and don't want to
// figure out how to shoehorn the stateful sort (eg which signals are known by
// that point in the sort) in there and the input is not that big.
//
// Also if you like functional programming the amount of mutating of 'in' is
// going to make you cringe, but this is Go and 'in' is a slice of pointers, not
// a pointer in itself so we're getting a copy of the slice of pointers anyway.
func sortOperations(in []*Operation) (out []*Operation) {
	inputlength := len(in)

	// Because AND sometimes uses a '1' as an input, make a special known signal with a value of 1
	knownSignals := map[string]bool{
		"1": true,
	}

	// First, the SET operations
	for i := 0; i < len(in); i++ {
		op := in[i]
		if op.Opcode == "SET" {
			knownSignals[op.Output] = true
			out = append(out, op)            // push onto output
			in = append(in[:i], in[i+1:]...) // remove from input
		}

	}

	// Now we have our starting point
	idx := 0

outer:
	for len(out) < inputlength {
		if idx >= len(in) {
			//	something's gone wrong
			panic("got to the end of the input but we're not done sorting??")
		}

		op := in[idx]

		// If any input signal is not in the map, skip this operation
		for _, input := range op.Inputs {
			if _, ok := knownSignals[input]; !ok {
				idx++
				continue outer
			}
		}

		knownSignals[op.Output] = true
		out = append(out, op)
		in = append(in[:idx], in[idx+1:]...) // remove from input
		idx = 0                              // start over
	}

	return out
}
