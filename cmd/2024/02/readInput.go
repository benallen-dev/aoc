package main

import (
	"bufio"
	"strings"
	"strconv"
	"os"
)

func readInput(filename string) ([][]int, error) {
	var out [][]int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// var us, them rune

		line := scanner.Text()
		// _, err := fmt.Sscanf(line, "%c %c", &them, &us)
		// if err != nil {
		// 	return nil, err
		// }

		parts := strings.Split(line, " ")
		intParts := []int{}
		for _, x := range parts {
			c, err := strconv.Atoi(x)
			if err != nil {
				return nil, err
			}
			intParts = append(intParts, c)
		}

		out = append(out, intParts)



	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
