package main

import (
	"fmt"
)

type Santa struct {
	position [2]int
}

func (s Santa) String() string {
	return fmt.Sprintf("(%d, %d)", s.position[0], s.position[1])
}

func (s *Santa) Move(char rune) {
	if char == '^' {
		s.position[1]++
	} else if char == 'v' {
		s.position[1]--
	} else if char == '>' {
		s.position[0]++
	} else if char == '<' {
		s.position[0]--
	}
}

func (s Santa) Pos() [2]int {
	return s.position
}
