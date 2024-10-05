package main

import (
	"fmt"
	"strings"
)

type rope string

func NewRope(s string) rope {
	return rope(s)
}

func (r rope) String() string {
	out := ""

	out += fmt.Sprintf("Rope: %s\n", string(r))
	// out += fmt.Sprintf("    Three vowels:              %t\n", r.ThreeVowels())
	// out += fmt.Sprintf("    Has double letter:         %t\n", r.HasDoubleLetter())
	// out += fmt.Sprintf("    Has no naughty strings:    %t\n", r.NoNaughtyString())
	out += fmt.Sprintf("    Has non-overlapping pairs: %t\n", r.NonOverlappingDoublePair())
	out += fmt.Sprintf("    Has xyx pattern:           %t\n", r.RepeatedWithLetterBetween())

	return out
}

func (r rope) LastIdx() int {
	return len(r) - 1
}

func (r rope) ThreeVowels() bool {
	// scan rune by rune
	// if vowel increase count
	// if count > 2 return true
	// return false if done with iteration

	c := 0
	for _, r := range r {

		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			c++
		}

		if c > 2 {
			return true
		}
	}

	return false
}

func (r rope) HasDoubleLetter() bool {
	for i := 0; i < r.LastIdx(); i++ {
		if r[i] == r[i+1] {
			return true
		}
	}

	return false
}


func (r rope) RepeatedWithLetterBetween () bool {
	for i := 0; i < r.LastIdx() -1; i++ {
		if r[i] == r[i+2] {
			return true
		}
	}

	return false
}

func (r rope) NoNaughtyString() bool {
	// Check for 'ab', 'cd', 'pq', or 'xy'
	// if any of these are found, return false

	if strings.Contains(string(r), "ab") ||
		strings.Contains(string(r), "cd") ||
		strings.Contains(string(r), "pq") ||
		strings.Contains(string(r), "xy") {
		return false
	}

	return true
}

func (r rope) NonOverlappingDoublePair() bool {
	for i := 0; i < r.LastIdx() - 2; i++ {
		candidate := string(r[i:i+2])
		remainder := string(r[i+2:])

		if strings.Contains(remainder, candidate) {
			return true
		}
	}
	return false
}

func (r rope) Nice() bool {
	return r.ThreeVowels() && r.HasDoubleLetter() && r.NoNaughtyString()
}

func (r rope) Nicer() bool {
	return r.NonOverlappingDoublePair() && r.RepeatedWithLetterBetween()
}
