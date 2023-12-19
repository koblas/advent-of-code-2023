package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Dir int

const (
	INVALID Dir = iota
	NORTH   Dir = iota
	EAST    Dir = iota
	SOUTH   Dir = iota
	WEST    Dir = iota
	dirmax  Dir = iota
)

func (d Dir) String() string {
	return []string{"*", "N", "E", "S", "W"}[d]
}
func (d Dir) IsReverse(a Dir) bool {
	switch d {
	case NORTH:
		return a == SOUTH
	case EAST:
		return a == WEST
	case SOUTH:
		return a == NORTH
	case WEST:
		return a == EAST
	}
	return false
}

type Pos struct {
	x, y int
}

func (pos Pos) Add(offset Pos) Pos {
	return Pos{
		x: pos.x + offset.x,
		y: pos.y + offset.y,
	}
}
func (pos Pos) Eq(b Pos) bool {
	return pos.x == b.x && pos.y == b.y
}

type Crucible struct {
	pos   Pos
	dir   Dir
	steps int
	loss  int
}

type Visited map[Crucible]int

type Input struct {
	cells      map[Pos]int
	maxX, maxY int
}

var moveDir = [5]Pos{}

func init() {
	moveDir[NORTH] = Pos{0, -1}
	moveDir[EAST] = Pos{1, 0}
	moveDir[SOUTH] = Pos{0, 1}
	moveDir[WEST] = Pos{-1, 0}
}

func (pos Pos) Move(dir Dir) Pos {
	if dir == INVALID {
		panic("invalid dir")
	}
	return pos.Add(moveDir[dir])
}

func (c Crucible) String() string {
	return fmt.Sprintf("{x:%d, y:%d} dir:%s loss:%d steps:%d", c.pos.x, c.pos.y, c.dir.String(), c.loss, c.steps)
}

type PriorityQueue []*Crucible

func Dump(input Input, visited Visited, pos *Pos) {
	fmt.Println("=============")
	simple := map[Pos]int{}

	for item, loss := range visited {
		v, found := simple[item.pos]
		if !found || loss < v {
			simple[item.pos] = loss
		}
	}

	for y := 0; y < input.maxY; y++ {
		for x := 0; x < input.maxX; x++ {
			p := Pos{x, y}
			val, found := simple[p]
			if pos != nil && p.Eq(*pos) {
				fmt.Printf("%4d", val)
				// fmt.Printf("X")
			} else if found {
				fmt.Printf("%4d", val)
				// fmt.Printf(":")
			} else {
				fmt.Printf("%4s", "")
				// fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

// Interface for container/heap
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].loss < pq[j].loss }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Crucible)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func solve(input Input, minSteps, maxSteps int) int {
	start := Pos{0, 0}
	end := Pos{x: input.maxX - 1, y: input.maxY - 1}

	queue := make(PriorityQueue, 0, 1_000_000)
	heap.Push(&queue, &Crucible{
		pos: start,
		dir: EAST,
	})
	heap.Push(&queue, &Crucible{
		pos: start,
		dir: SOUTH,
	})

	visited := Visited{}
	visited[Crucible{pos: start, dir: EAST}] = 0
	visited[Crucible{pos: start, dir: SOUTH}] = 0

	for queue.Len() != 0 {
		front := heap.Pop(&queue).(*Crucible)
		// fmt.Println(front.String())

		if front.pos.Eq(end) && front.steps >= minSteps {
			return front.loss
		}

		var allowedDirs []Dir
		switch front.dir {
		case NORTH, SOUTH:
			allowedDirs = []Dir{EAST, WEST, front.dir}
		case EAST, WEST:
			allowedDirs = []Dir{NORTH, SOUTH, front.dir}
		case INVALID:
			allowedDirs = []Dir{NORTH, EAST, SOUTH, WEST}
		}

		for _, dir := range allowedDirs {
			steps := 1
			if dir == front.dir {
				steps = front.steps + 1
				if steps > maxSteps {
					continue
				}
			} else if front.steps < minSteps {
				// Cannot turn until we've gotten far enough
				continue
			}

			next := front.pos.Move(dir)

			if loss, found := input.cells[next]; found {
				loss += front.loss
				vpos := Crucible{pos: next, dir: dir, steps: steps}
				if value, found := visited[vpos]; !found || loss < value {
					// does not hash loss
					visited[vpos] = loss
					vpos.loss = loss
					heap.Push(&queue, &vpos)
				}
			}
		}
	}

	return -1
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		cells: map[Pos]int{},
	}

	if len(lines) == 0 {
		return input, fmt.Errorf("bad input")
	}

	for y, line := range lines {
		for x, ch := range line {
			val, _ := strconv.Atoi(string(ch))
			input.cells[Pos{x, y}] = val
		}
	}

	input.maxX = len(lines[0])
	input.maxY = len(lines)

	return input, nil
}

func PartOneSolution(input Input) (int, error) {
	// 902
	return solve(input, 0, 3), nil
}

func PartTwoSolution(input Input) (int, error) {
	// 1073
	return solve(input, 4, 10), nil
}

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	timeStart := time.Now()
	input, err := ParseInput(lines)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Build input (%.2fms)\n", float64(time.Since(timeStart).Microseconds())/1000)

	timeStart = time.Now()
	values, err := PartOneSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values)

	timeStart = time.Now()
	values, err = PartTwoSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values)
}
