package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) ([]box, error) {
	var out []box

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var w, l, h int

		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%dx%dx%d", &w, &l, &h)
		if err != nil {
			return nil, err
		}

		out = append(out, NewBox(w,l,h))
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
