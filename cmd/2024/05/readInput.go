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

type update map[int]int

func (u update) IntSlice() []int {

	ints := make([]int, len(u), len(u))
	for k, v := range u {
		ints[v] = k
	}

	return ints
}

func (u update) String() string {
	return fmt.Sprint(u.IntSlice())
}

func (u update) Valid(rules []rule) bool {
	for _, rule := range rules {
		_, okleft := u[rule[0]]  // left exists in update
		_, okright := u[rule[1]] // right exists in update

		// if both left and right exist and position of left > position of right, it's invalid
		if okleft && okright && u[rule[0]] > u[rule[1]] {
			return false
		}
	}

	return true
}
func (u update) ValidString(rules []rule) string {
	if u.Valid(rules) {
		return "valid"
	}
	return "invalid"
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
