package main

import (
	"fmt"
)

type Strat struct {
	Them rune
	Outcome rune
}

func (s Strat) String() string {
	return fmt.Sprintf("%c %c", s.Them, s.Outcome)
}

// A -> Rock	 = 1 pt
// B -> Paper	 = 2 pts
// C -> Scissors = 3 pts

// X -> Lose = 0 pts
// Y -> Draw = 3 pts
// Z -> Win  = 6 pts

var scoreMatrix = map[rune]map[rune]int {
	'A': { // Opponent chooses ROCK
		'X': 3, // Lose (0) -> scissors (3)
		'Y': 4, // Draw (3) -> rock (1)
		'Z': 8, // Win (6) -> paper (2)
	},
	'B': { // Opponent chooses PAPER
		'X': 1, // Lose (0) -> rock (1)
		'Y': 5, // Draw (3) -> paper (2)
		'Z': 9, // Win (6) -> scissors (3)
	},
	'C': { // Oppponent chooses SCISSORS
		'X': 2, // Lose (0) -> paper (2)
		'Y': 6, // Draw (3) -> scissors (3)
		'Z': 7, // Win (6) -> rock (1)
	},
}

// We're making a HUGE assumption that there's no bad data, btw
func (s Strat) Score() int {
	return scoreMatrix[s.Them][s.Outcome]
}

