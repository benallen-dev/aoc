package main

import (
	"bufio"
	"os"
)


func readInput(filename string) ([]rope, error) {
	var out []rope

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, NewRope(scanner.Text()))
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
