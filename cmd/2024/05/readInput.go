package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// first section: rules

// second section: updates

// plan: make a map of updates -> update->position
// Then, for each rule check if both numbers are in the map
// if they are, then we can easily check if map[page1] < map[page2]
// if no rules are violated, the update is valid

// ergo, we need to return a []rule, and a []update where
// - update is a map of int->int
// - rule is a tuple of int,int

// why not []int instead of a map? because this simplifies checking if a number is in the update

type rule [2]int

func (r rule) String() string {
	return fmt.Sprintf("%d|%d", r[0], r[1])
}

func readInput(filename string) (rules []rule, updates []update, err error) {
	rules = []rule{}
	updates = []update{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Scan until we hit a blank line to build rules
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		var l, r int

		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%d|%d", &l, &r)
		if err != nil {
			return nil, nil, err
		}

		rules = append(rules, rule{l, r})
	}

	// Scan until we hit EOF to build updates
	for scanner.Scan() {
		update := make(update)

		line := scanner.Text()
		parts := strings.Split(line, ",")
		for val, part := range parts {
			key, err := strconv.Atoi(part)
			if err != nil {
				return nil, nil, err
			}

			update[key] = val
		}

		updates = append(updates, update)
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, nil, scanerr
	}

	return rules, updates, nil
}
