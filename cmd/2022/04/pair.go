package main

import (
	"fmt"
)

type Pair struct {
	Ass1 Assignment
	Ass2 Assignment
}

func NewPair(start1, end1, start2, end2 int) Pair {
	return Pair{
		Ass1: Assignment{start: start1, end: end1},
		Ass2: Assignment{start: start2, end: end2},
	}
}

func (p Pair) String() string {
	return fmt.Sprintf("%s, %s", p.Ass1.String(), p.Ass2.String())
}

// FullOverlap returns true when one of the Pair's assignments is fully contained in the other
func (p Pair) FullOverlap() bool {
	return p.Ass1.Contains(p.Ass2) || p.Ass2.Contains(p.Ass1)
}

// Oh man I got so lucky choosing the right structure, nice
func (p Pair) PartialOverlap() bool {
	return

}
