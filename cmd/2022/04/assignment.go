package main

import (
	"fmt"
)

type Assignment struct {
	start int
	end   int
}

func (a Assignment) String() string {
	return fmt.Sprintf("%d-%d", a.start, a.end)
}

func (a Assignment) Contains(other Assignment) bool {
	return a.start <= other.start && a.end >= other.end
}
