package main

import (
	"bufio"
	"os"
)

func readInput(filename string) ([]*Operation, error) {
	var out []*Operation

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		newOp, err := NewOperation(scanner.Text())
		if err != nil {
			panic(err)
		}

		out = append(out, newOp)
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
