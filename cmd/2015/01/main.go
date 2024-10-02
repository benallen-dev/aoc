package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pos := 0
	basement := false
	floor := 0

	reader := bufio.NewReader(file)
	for {
		pos = pos + 1

		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if char == '(' {
			floor++
		} else if char == ')'{
			floor--
		}

		if basement == false && floor == -1 {
			basement = true
			fmt.Printf("Entered basement at position: %d\n", pos)
		}
	}

	fmt.Printf("Instructions lead to floor %d\n", floor)
}

