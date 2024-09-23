package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) ([]Pair, error) {
	var out []Pair

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var start1, start2, end1, end2 int

		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &start1, &end1, &start2, &end2)
		if err != nil {
			return nil, err
		}

		out = append(out, NewPair(start1, end1, start2, end2))
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
