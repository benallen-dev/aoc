package main

import (
	"fmt"
)

func part1(input [][]rune) int {
	xmasCount := 0

	height := len(input)
	width := len(input[0]) // if len(input) == 0 we want to panic anyway

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if input[i][j] != 'X' {
				continue
			}

			fitRight := j <= width-4
			fitLeft := j >= 3
			fitTop := i >= 3
			fitBottom := i <= height-4

			if fitRight {
				x := string([]rune{
					input[i][j],
					input[i][j+1],
					input[i][j+2],
					input[i][j+3],
				})

				if x == "XMAS" {
					xmasCount++
				}
			}
			if fitRight && fitBottom {
				x := string([]rune{
					input[i][j],
					input[i+1][j+1],
					input[i+2][j+2],
					input[i+3][j+3],
				})

				if x == "XMAS" {
					xmasCount++
				}
			}
			if fitBottom {
				x := string([]rune{
					input[i][j],
					input[i+1][j],
					input[i+2][j],
					input[i+3][j],
				})

				if x == "XMAS" {
					xmasCount++
				}
			}
			if fitBottom && fitLeft {
				x := string([]rune{
					input[i][j],
					input[i+1][j-1],
					input[i+2][j-2],
					input[i+3][j-3],
				})

				if x == "XMAS" {
					xmasCount++
				}
			}
			if fitLeft {
				x := string([]rune{
					input[i][j],
					input[i][j-1],
					input[i][j-2],
					input[i][j-3],
				})

				if x == "XMAS" {
					xmasCount++
				}
			}
			if fitLeft && fitTop {
				x := string([]rune{
					input[i][j],
					input[i-1][j-1],
					input[i-2][j-2],
					input[i-3][j-3],
				})

				if x == "XMAS" {
					xmasCount++
				}
			}
			if fitTop {
				x := string([]rune{
					input[i][j],
					input[i-1][j],
					input[i-2][j],
					input[i-3][j],
				})

				if x == "XMAS" {
					xmasCount++
				}
			}
			if fitTop && fitRight {
				x := string([]rune{
					input[i][j],
					input[i-1][j+1],
					input[i-2][j+2],
					input[i-3][j+3],
				})

				if x == "XMAS" {
					xmasCount++
				}
			}
		}
	}
	return xmasCount
}

func part2(input [][]rune) int {
	masCount := 0

	height := len(input)
	width := len(input[0]) // if len(input) == 0 we want to panic anyway

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if input[i][j] != 'A' {
				continue
			}

			diagonals := []rune{
				input[i-1][j-1],
				input[i-1][j+1],
				input[i+1][j-1],
				input[i+1][j+1],
			}

			// Valid diagonals are:
			// M M  |  M S  |  S M  |  S S
			// S S  |  M S  |  S M  |  M M

			// We can check if the diagonals are valid by casting the slice to
			// a string and checking if the string is MMSS, MSMS, SMSM, or SSMM
			diagonalString := string(diagonals)
			if diagonalString == "MMSS" || diagonalString == "MSMS" || diagonalString == "SMSM" || diagonalString == "SSMM" {
				masCount++
			}

		}
	}

	return masCount
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}
