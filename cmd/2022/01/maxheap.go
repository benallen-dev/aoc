package main

// See https://pkg.go.dev/container/heap

type Elf struct {
	id int // The number of the elf
	calories []int // The energy they are carrying
	total int // The sum of the total calories
	index int // used to keep track of position
}

// Implements the heap interface
type MaxHeap []*Elf

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i].total > mh[j].total // we want highest first, so less is more
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}

func (mh *MaxHeap) Push(x any) {
	n := len(*mh)
	item := x.(*Elf)
	item.index = n
	*mh = append(*mh, item)
}

func (mh *MaxHeap) Pop() any {
	old := *mh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // Some GC magic apparently
	item.index = -1 // "for safety"
	*mh = old[0 : n-1]
	return item
}

func (mh *MaxHeap) update(item *Elf, calories []int) {
	item.calories = calories
	sum := 0
	for _, c := range calories {
		sum += c
	}
	item.total = sum
}

