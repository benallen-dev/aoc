package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]equation, error) {
	out := []equation{}

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

		in := scanner.Text()
		// split on : to get target and operands
		x := strings.Split(in, ": ")

		target, err := strconv.Atoi(x[0])
		if err != nil {
			return nil, err
		}

		operands := []int{}
		for _, v := range strings.Split(x[1], " ") {
			op, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			operands = append(operands, op)
		}

		out = append(out, equation{
			target:   target,
			operands: operands,
		})
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return out, nil
}
