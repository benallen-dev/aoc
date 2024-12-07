package main

import (
	"aoc/internal/color"
	"fmt"
)

type update map[int]int

func (u update) IntSlice() []int {

	ints := make([]int, len(u), len(u))
	for k, v := range u {
		ints[v] = k
	}

	return ints
}

func (u update) String() string {
	c := len(u) / 2

	out := "["
	for i, v := range u.IntSlice() {
		if i == c {
			out += color.Bold + fmt.Sprint(v) + color.Reset
		} else {
			out += fmt.Sprint(v)
		}

		if i < len(u)-1 {
			out += ", "
		}
	}
	out += "]"

	return out
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

// Mutates the update to fix it. Note that this mutates u even without
// a pointer receiver because maps are reference types.
func (u update) Fix(rules []rule) {
	for !u.Valid(rules) {
		for _, rule := range rules {
			l, okleft := u[rule[0]]  // left exists in update
			r, okright := u[rule[1]] // right exists in update

			// if both left and right exist and position of left > position of right, it's invalid
			if okleft && okright && u[rule[0]] > u[rule[1]] {
				// swap
				u[rule[0]] = r
				u[rule[1]] = l
			}
		}
	}
}
