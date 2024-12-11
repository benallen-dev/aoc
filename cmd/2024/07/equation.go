package main

type operator int
const (
	OpAdd = iota
	OpMul
)
func (o operator) String() string {
	if o > 1 {
		return "smooth operator"
	}

	return []string{"+", "*"}[o]
}


type equation struct {
	target int
	operands []int
}

type pathNode struct {
	remaining []int
	plus *pathNode
	mult *pathNode
}

func (p *pathNode) Process() {
	if p.plus != nil || p.mult != nil {
		panic("not a leaf node")
	}

	if len(p.remaining) == 1 {
		return
	}

	plusRemaining := []int{ p.remaining[0] + p.remaining[1] }
	plusRemaining = append(plusRemaining, p.remaining[2:]...)

	multRemaining := []int{ p.remaining[0] * p.remaining[1] }
	multRemaining = append(multRemaining, p.remaining[2:]...)

	p.plus = &pathNode{ remaining: plusRemaining }
	p.mult = &pathNode{ remaining: multRemaining }

	p.plus.Process()
	p.mult.Process()
}

func (p pathNode) GetLeafs() []*pathNode {
	// Symmetrical is fine here because this is a perfect tree
	if p.plus == nil && p.mult == nil {
		return []*pathNode{&p}
	}

	plusLeafs := p.plus.GetLeafs()
	multLeafs := p.mult.GetLeafs()

	return append(plusLeafs, multLeafs...)
}

func (e equation) Solvable() bool {
	root := pathNode{
		remaining: e.operands,
	}

	root.Process() // side effects let's gooo
	leafs := root.GetLeafs()

	solutions := 0
	for _, l := range leafs {
		if l.remaining[0] == e.target {
			solutions++
		}
	}

	return solutions > 0


}
