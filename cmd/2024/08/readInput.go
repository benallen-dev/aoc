package main

import (
	"bufio"
	"fmt"
	"os"
)

type location [2]int
func (l location) Subtract(r location) location {
	return location{l[0] - r[0], l[1] - r[1]}
}
func (l location) Add(r location) location {
	return location{l[0] + r[0], l[1] + r[1]}
}
func (l location) String() string {
	return fmt.Sprintf("[%v, %v]", l[0], l[1])
}

type city struct {
	antennae map[rune][]location
	width int
	height int
}

func (c city) Draw() string {
	out := ""
	
	a := make(map[location]rune)
	for k, v := range c.antennae {
		for _, l := range v {
			a[l] = k
		}
	}

	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			if a[location{i, j}] != 0 {
				out += string(a[location{i, j}])
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	return out
}

func readInput(filename string) (city, error) {
	var width, height int
	antennae := make(map[rune][]location)

	file, err := os.Open(filename)
	if err != nil {
		return city{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	lc := 0
	for scanner.Scan() {
		line := scanner.Text()

		if lc == 0 {
			width = len(line)
		}

		for i, c := range line {
			if c != '.' {
				if _, ok := antennae[c]; !ok {				
					antennae[c] = []location{}
				}

				antennae[c] = append(antennae[c], location{lc, i})
			}
		}
		lc++
	}

	height = lc


	scanerr := scanner.Err()
	if scanerr != nil {
		return city{}, scanerr
	}

	return city{
		antennae: antennae,
		width: width,
		height: height,
	}, nil
}
