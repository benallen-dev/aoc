package main

import (
	"slices"
)

type box struct {
	width  int
	length int
	height int
}

func NewBox(w, l, h int) box {
	return box{
		width: w,
		height: h,
		length: l,
	}
}

func (b box) areas() []int {
	return []int{
		b.length * b.width,
		b.width * b.height,
		b.height * b.length,
	}
}

func (b box) Surface() int {
	total := 0

	for _, a := range b.areas() {
		total += 2*a
	}

	return total
}

func (b box) SmallestArea() int {
	areas := b.areas()
	slices.Sort(areas) // gotta love me some in-place sorting

	return areas[0]
}

func (b box) PaperSurface() int {
	return b.Surface() + b.SmallestArea()
}

func (b box) Volume() int {
	return b.width*b.height*b.length
}

func (b box) SmallestPerimeter() int {
	sides := []int{b.width, b.height, b.length}
	slices.Sort(sides)

	// Because areas always returns a slice of length 3 this is ok
	return sides[0] + sides[0] + sides[1] + sides[1]
}

func (b box) RibbonLength() int {
	return b.Volume()+b.SmallestPerimeter()
}
