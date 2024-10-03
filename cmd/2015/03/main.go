package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	santa := &Santa{[2]int{0, 0}}
	roboSanta := &Santa{[2]int{0, 0}}

	var counts map[[2]int]int = make(map[[2]int]int) // the fact this works is pretty neat
	
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var i int = 0
	reader := bufio.NewReader(file)
	for {

		// alternate between santa and robo-santa
		var currentSanta *Santa
		if i % 2 == 0 {
			currentSanta = santa
		} else {
			currentSanta = roboSanta
		}
		i++ // comment out this line for the answer to part 1

		// increment the count for the current position
		counts[currentSanta.Pos()]++

		// Move to next position
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		currentSanta.Move(char)
	}

	fmt.Printf("Visited %d houses\n", len(counts))
}
