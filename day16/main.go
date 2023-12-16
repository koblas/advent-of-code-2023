package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Dir int

const (
	NORTH Dir = iota
	EAST  Dir = iota
	SOUTH Dir = iota
	WEST  Dir = iota
)

func (d Dir) String() string {
	return []string{"N", "E", "S", "W"}[d]
}

type Pos struct {
	x, y int
}

type Beam struct {
	pos Pos
	dir Dir
}

type Visited map[Beam]bool

type Input struct {
	cells      map[Pos]rune
	maxX, maxY int
}

func (p Pos) Add(dir Pos) Pos {
	return Pos{
		x: p.x + dir.x,
		y: p.y + dir.y,
	}
}

func (b Beam) Move() Beam {
	bnew := Beam{dir: b.dir}
	switch b.dir {
	case NORTH:
		bnew.pos = b.pos.Add(Pos{0, -1})
	case EAST:
		bnew.pos = b.pos.Add(Pos{1, 0})
	case SOUTH:
		bnew.pos = b.pos.Add(Pos{0, 1})
	case WEST:
		bnew.pos = b.pos.Add(Pos{-1, 0})
	}
	return bnew
}

func stepBeams(input Input, beams []Beam, visited Visited) []Beam {
	var updated []Beam
	for _, b := range beams {
		if visited[b] {
			// 	fmt.Println("  Visited: ", b)
			continue
		}
		if _, found := input.cells[b.pos]; found {
			visited[b] = true
		}

		next := b.Move()

		act, found := input.cells[next.pos]
		if !found {
			// fmt.Println("  Off", next)
			continue
		}
		// fmt.Printf("AT %c: [%d,%d] %s\n", act, next.pos.x, next.pos.y, next.dir)
		switch act {
		case '.':
			// nothing
			updated = append(updated, next)
		case '|':
			switch next.dir {
			case NORTH, SOUTH:
				updated = append(updated, next)
			case EAST, WEST:
				updated = append(updated, Beam{
					pos: next.pos,
					dir: NORTH,
				}, Beam{
					pos: next.pos,
					dir: SOUTH,
				})
			}
		case '-':
			switch next.dir {
			case EAST, WEST:
				updated = append(updated, next)
			case NORTH, SOUTH:
				updated = append(updated, Beam{
					pos: next.pos,
					dir: EAST,
				}, Beam{
					pos: next.pos,
					dir: WEST,
				})
			}
		case '/':
			switch next.dir {
			case NORTH:
				next.dir = EAST
			case EAST:
				next.dir = NORTH
			case SOUTH:
				next.dir = WEST
			case WEST:
				next.dir = SOUTH
			}
			updated = append(updated, next)
		case '\\':
			switch next.dir {
			case NORTH:
				next.dir = WEST
			case EAST:
				next.dir = SOUTH
			case SOUTH:
				next.dir = EAST
			case WEST:
				next.dir = NORTH
			}
			updated = append(updated, next)
		}
	}

	return updated
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		cells: map[Pos]rune{},
	}

	if len(lines) == 0 {
		return input, fmt.Errorf("bad input")
	}

	for y, line := range lines {
		for x, ch := range line {
			input.cells[Pos{x, y}] = ch
		}
	}

	input.maxX = len(lines[0])
	input.maxY = len(lines)

	return input, nil
}

func runBeams(input Input, beams []Beam) int {
	visited := Visited{}
	for len(beams) != 0 {
		// fmt.Println("====")
		beams = stepBeams(input, beams, visited)
		// fmt.Println(beams)
	}

	bypos := map[Pos]bool{}
	for cell := range visited {
		bypos[cell.pos] = true
	}

	// for y := 0; y < 100; y += 1 {
	// 	if _, found := input.cells[Pos{0, y}]; !found {
	// 		break
	// 	}
	// 	for x := 0; x < 100; x += 1 {
	// 		p := Pos{x, y}
	// 		if _, found := input.cells[p]; !found {
	// 			break
	// 		}
	// 		if bypos[p] {
	// 			fmt.Printf("#")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }

	return len(bypos)
}

func PartOneSolution(input Input) (int, error) {
	beams := []Beam{
		{
			pos: Pos{-1, 0},
			dir: EAST,
		},
	}

	return runBeams(input, beams), nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PartTwoSolution(input Input) (int, error) {
	maxV := 0
	for y := 0; y < input.maxY; y++ {
		beams := []Beam{
			{
				pos: Pos{-1, y},
				dir: EAST,
			},
		}

		maxV = max(maxV, runBeams(input, beams))

		beams = []Beam{
			{
				pos: Pos{input.maxX, y},
				dir: WEST,
			},
		}

		maxV = max(maxV, runBeams(input, beams))
	}
	for x := 0; x < input.maxX; x++ {
		beams := []Beam{
			{
				pos: Pos{x, -1},
				dir: SOUTH,
			},
		}

		maxV = max(maxV, runBeams(input, beams))

		beams = []Beam{
			{
				pos: Pos{x, input.maxY},
				dir: NORTH,
			},
		}

		maxV = max(maxV, runBeams(input, beams))
	}

	return maxV, nil
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
