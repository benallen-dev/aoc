package main

import (
	"aoc/internal/set"
	"fmt"
)

func insideBounds(c city, l location) bool {
	outTop := l[0] < 0
	outBottom := l[0] >= c.height
	outLeft := l[1] < 0
	outRight := l[1] >= c.width

	return !(outTop || outBottom || outLeft || outRight)
}

func getAntinodes(c city) []location {
	out := []location{}

	// f for frequency
	for _, f := range c.antennae {
		// if there are more than one antennae in the same location
		if len(f) > 1 {
			for i := 0; i < len(f); i++ {
				for j := 0; j < len(f); j++ {
					if i == j {
						continue
					}

					// This antenna is also an antinode
					out = append(out, f[i])

					// calculate the difference between the two antennae
					diff := f[i].Subtract(f[j])

					newLoc := f[i].Add(diff)
					for insideBounds(c, newLoc) {
						out = append(out, newLoc)
						newLoc = newLoc.Add(diff)
					}
				}
			}
		}
	}
	return out
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	antiSet := set.Set[location]{}
	for _, a := range getAntinodes(input) {
		// Check inside bounds
		if insideBounds(input, a) {
			antiSet.Add(a)
		}
	}

 	fmt.Println("Part 1:", antiSet.Length())
}

