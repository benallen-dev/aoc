package main

import (
	"bufio"
	"os"
)

func readInput(filename string) ([][]rune, error) {

	var out [][]rune

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, []rune(line))
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
