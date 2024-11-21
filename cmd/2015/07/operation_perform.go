package main

func performOperations(ops []*Operation) (map[string]uint16, error) {
	wires := make(map[string]uint16)
	wires["1"] = 1

	for _, op := range ops {
		var err error
		inputs := make([]uint16, len(op.Inputs), len(op.Inputs))
		for i, in := range op.Inputs {
			if val, ok := wires[in]; ok {
				inputs[i] = val
			}
		}
		wires[op.Output], err = op.Perform(inputs)
		if err != nil {
			panic(err)
		}
	}
	return wires, nil
}
