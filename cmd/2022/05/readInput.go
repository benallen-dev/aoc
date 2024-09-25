package main

import (
	"bufio"
	"os"
)

func readInput(filename string) ([]string, error) {
	var out []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	// cool, so out is a slice of strings now.
	// You could just use ioutil.ReadFile but this is what I have in the template.
	//
	// You could also parse these as they come in but the initial input is only
	// 500 lines and so I'm not too bothered about memory performance of reading
	// the input, when we could be focusing on the interesting part of parsing
	// the initial state and doing some stack operations.
	return out, nil
}
