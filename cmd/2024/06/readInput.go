package main

import (
	"bufio"
	"os"
)


func readInput(filename string) (lab, error) {
	out := lab{
		obstacles: map[position]bool{},
	}
	
	file, err := os.Open(filename)
	if err != nil {
		return lab{}, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lc := 0 // line count

	for scanner.Scan() {
		l := scanner.Text()
		if lc == 0 { // This is the first line, set width
			out.width = len(l)
		}

		for i, c := range l {
			// find all `#` and store as lab.obstacles[lc][idx]
			if c == '#' {
				out.obstacles[[2]int{lc,i}] = true
			}

			// find `^` and store as guard state
			if c == '^' {
				out.guard = guard{
					pos: [2]int{lc, i},
					dir: DirUp,
				}

				out.guardInit = out.guard
			}
		}

		lc++
	}

	out.height = lc

	if err := scanner.Err(); err != nil {
		return lab{}, err
	}

	return out, nil
}
