package main

import (
	"bufio"
	"strings"
	"os"
)

func readInput(filename string) (string, error) {
	var out []string

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return "", scanerr
	}
	
	return strings.Join(out, ""), nil
}
