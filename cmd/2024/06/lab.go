package main

import (
	"fmt"
//	"time"
)

// [line, col]
type position [2]int

type direction int
const (
	DirUp = iota
	DirRight
	DirDown
	DirLeft
)
func (d direction) String() string {
	if d > 3 {
		return "bad direction"
	}

	return []string{"^", ">", "v", "<"}[d]
}

type guard struct {
	pos position
	dir direction
}

func (g guard) String() string {
	return fmt.Sprintf("Guard: %v facing %s", g.pos, g.dir.String())
}

type lab struct {
	width int
	height int
	obstacles map[position]bool
	guard guard
	guardInit guard
}

func (l lab) String() string {
	out := ""
	out += fmt.Sprintf("Lab %v x %v (h x w)\n", l.width, l.height)
	out += "Guard: " + l.guard.String() + "\n"
	out += fmt.Sprintf("Obstacles: %v\n", l.obstacles)

	return out
}

func (l lab) Draw() string {
	out := ""
	for i := 0; i < l.height; i++ {
		for j := 0; j < l.width; j++ {
			if l.guard.pos == [2]int{i, j} {
				
				out += l.guard.dir.String()
			} else if l.obstacles[position{i, j}] {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}

	return out
}

func (l *lab) Tick() (dirChange bool) {
	move := map[direction]func(position) position {
		DirUp: func(x position) position { return position{ x[0] -1, x[1] } },
		DirDown: func(x position) position { return position{ x[0] +1, x[1] } },
		DirLeft: func(x position) position { return position{ x[0], x[1] -1 } },
		DirRight: func(x position) position { return position{ x[0], x[1]+1 } },
	}
	
	next := move[l.guard.dir](l.guard.pos)

	// If obstacle, change direction
	if _, ok := l.obstacles[next]; ok {
		l.guard.dir = (l.guard.dir + 1) % 4
		return true
	} else { // If not, move guard
		l.guard.pos = next
		return false
	}
}

func (l lab) GuardPresent() bool {
	p := l.guard.pos

	if p[1] < 0 { // Guard is off to the left
		return false
	}
	if p[1] >= l.width { // Guard is off the right
		return false
	}
	if p[0] < 0 { // Guard is off above
		return false
	}
	if p[0] >= l.height { // Guard is off below
		return false
	}

	return true
}

func (l *lab) ResetGuard() {
	l.guard = l.guardInit
}

func (l lab) TestForLoop(newObstacle position) bool {
	l.ResetGuard()
	l.obstacles[newObstacle] = true

	guardPositions := map[guard]bool{}

	for l.GuardPresent() {
		dirChange := l.Tick()
	
		// redraw
		// fmt.Print("\033[H\033[2J")
		// fmt.Println(l.Draw())
		// time.Sleep(100 * time.Millisecond)



		if _, ok := guardPositions[l.guard] ; ok && !dirChange{
			// loop detected
			delete(l.obstacles, newObstacle)
			return true
		}

		guardPositions[l.guard] = true
	}


	delete(l.obstacles, newObstacle)
	return false
}

