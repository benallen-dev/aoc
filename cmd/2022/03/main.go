package main

import (
	"fmt"
)

// Runes and bytes are interchangable
// string[idx] returns a byte, not a rune
// therefore findShared returns a byte to avoid conversions all over the place
func findShared (line string) byte {
	var out byte

	// For the first half of the string build a map of rune counts		
	charmap := map[byte]bool{}
	i := 0

	// for the first half of the input, build the map
	for ; i<len(line)/2; i++ {
		c := line[i]
		charmap[c] = true
	}

	// for the second half, return when finding a value that exists in the first half
	for ;i<len(line); i++ {
		c := line[i]
		if _,ok := charmap[line[i]]; ok {
			return c
		}
	}

	return out
}

func runeToPrio(b byte) int {
	// a -> 97 -> 1
	// z -> 122 -> 26

	// A -> 65 -> 27
	// Z -> 90 -> 52

	if b < 97 { 
		// offset is 65 - 27
		return int(b - 38)
	} else {
		// offset is 97 - 1
		return int(b - 96)
	}
}

// again, bytes and runes are interchangeable
func findCommonInGroup(group [3]string) byte {
	maps := map[int]map[rune]bool{}
	
	// build the maps
	for g := range 3 {
		maps[g] = map[rune]bool{}
		for _, r := range group[g] {
			maps[g][r] = true
		}
	}

	// Probably not the most efficient way but oh well, it works
	for c := range maps[0] {
		_, ok2 := maps[1][c]
		_, ok3 := maps[2][c]

		if ok2 && ok3 {
			return byte(c)
		}
	}

	fmt.Println(group)
	panic("could not find common char in groups")
}


func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, line := range input {
		rune := findShared(line)

		total += runeToPrio(rune)
	}

	fmt.Printf("Part 1: %d\n", total)

	// combine inputs into groups of three
	// for each group, find that character which exists in all 3 strings
	// sum the priorities of said items

	// we could initialise the size of groups but I doubt for 300 items it will matter
	groups := [][3]string{}
	for i := 0; i<len(input); i = i+3 {
		groups = append(groups, [3]string{
			input[i],
			input[i+1],
			input[i+2],
		})
	}

	total = 0
	for _, g := range groups {
		c := findCommonInGroup(g)
		total += runeToPrio(c)
	}

	fmt.Printf("Part 2: %d\n", total)


}

