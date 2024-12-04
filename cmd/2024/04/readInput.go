package main

import (
	"bufio"
//	"fmt"
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
		// var us, them rune

		// line := scanner.Text()
		// _, err := fmt.Sscanf(line, "%c %c", &them, &us)
		// if err != nil {
		// 	return nil, err
		// }

		out = append(out, scanner.Text())
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
