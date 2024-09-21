package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) ([]Round, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Ideally you'd know the input size but my unverified guess is
	// that reading the entire file into memory and then allocating
	// a slice of the correct size not going to be noticable at the
	// scale of this problem. As such, guess 100 and let the slice
	// grow as needed.
	scanner := bufio.NewScanner(file)
	rounds := []Round{}
	
	for scanner.Scan() {
		var us, them rune

		line := scanner.Text()

		_, err := fmt.Sscanf(line, "%c %c", &them, &us)
		if err != nil {
			return nil, err
		}

		usScore := int(us - 87)
		themScore := int(them - 64)

		rounds = append(rounds, Round{usScore, themScore})
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return rounds, nil
}

func readInputProperly(filename string) ([]Strat, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Ideally you'd know the input size but my unverified guess is
	// that reading the entire file into memory and then allocating
	// a slice of the correct size not going to be noticable at the
	// scale of this problem. As such, guess 100 and let the slice
	// grow as needed.
	scanner := bufio.NewScanner(file)
	strats := []Strat{}
	
	for scanner.Scan() {
		var them, outcome rune

		line := scanner.Text()

		_, err := fmt.Sscanf(line, "%c %c", &them, &outcome)
		if err != nil {
			return nil, err
		}

		strats = append(strats, Strat{
			Them: them,
			Outcome: outcome,
		})
	}

	scanerr := scanner.Err()
	if scanerr != nil {
		return nil, scanerr
	}

	return strats, nil
}
