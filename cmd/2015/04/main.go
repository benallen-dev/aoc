package main

import (
	"fmt"
	"os"
	"crypto/md5"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// remove newline from input
	// I'm not particularly proud of this but also
	// don't care enough to do it properly
	input = []byte(strings.Split(string(input), "\n")[0])

	// The naive way:
	iter := 0
	fiveZeroes := 0
	sixZeroes := 0

	for {
		foo := append(input, []byte(fmt.Sprintf("%d", iter))...)

		hasher := md5.New()
		hasher.Write(foo)
		hexHash := fmt.Sprintf("%x", hasher.Sum(nil))

		if fiveZeroes == 0 && strings.HasPrefix(hexHash, "00000") {
			fiveZeroes = iter
		}

		if sixZeroes == 0 && strings.HasPrefix(hexHash, "000000") {
			sixZeroes = iter
		}

		if (fiveZeroes > 0 && sixZeroes > 0) {
			break
		}

		iter++
	}

	// Turns out the naive way is fast enough to answer the question
	fmt.Println("Five zeroes:", fiveZeroes)
	fmt.Println("Six zeroes:", sixZeroes)
}

