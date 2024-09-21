package main

import (
	"strconv"
	"log"
	"os"
	"bufio"
)

func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Panic("Cannot read input file from disk", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result [][]int
	var currentGroup []int

	for	scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// If currentGroup has data, append to result and reset
			if len(currentGroup) > 0 {
				result = append(result, currentGroup)
				currentGroup = []int{}
			}
		} else {
			// convert line to int and append to currentGroup
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Panic("Cannot convert line to int", err)
			}

			currentGroup = append(currentGroup, num)
		}
	}

	// Append the last group
	if len(currentGroup) > 0 {
		result = append(result, currentGroup)
	}

	return result
}
