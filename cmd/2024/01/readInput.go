package main

import (
	"bufio"
	"fmt"
	"os"
)

// Returns two slices containing the left and right lists of the input file
func readInput(filename string) (out [][]int, err error) {
	out = make([][]int, 2)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var l, r int

		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%d   %d", &l, &r)
		if err != nil {
			return nil, err
		}

		out[0] = append(out[0], l)
		out[1] = append(out[1], r)
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
