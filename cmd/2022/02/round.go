package main

import (
	"fmt"
)

type Round struct {
	us int
	them int
}

func (r Round) String() string {
	var us, them rune
	us = rune(87 + r.us) // 87 is the ASCII code for 'W', so 87 + 1 = 88 = 'X'
	them = rune(64 + r.them) // 64 is the ASCII code for 'A', so 64 + 1 = 65 = 'A'

	return fmt.Sprintf("%c %c", them, us)
}

func (r Round) Score() (score int) {
	score = 0
	score += r.us

	// A 01 Rock
	// B 01 Paper
	// C 01 Scissors

	// X 01 Rock
	// Y 01 Paper
	// Z 01 Scissors

	// Win conditions
	if (r.us == 1 && r.them == 3) || (r.us == 2 && r.them == 1) || (r.us == 3 && r.them == 2) {
		score += 6
	}

	// Draw
	if r.us == r.them {
		score += 3
	}

	return score
}
