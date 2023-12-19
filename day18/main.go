package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
func (pos Pos) Min(b Pos) Pos {
	p2 := pos
	if p2.x > b.x {
		p2.x = b.x
	}
	if p2.y > b.y {
		p2.y = b.y
	}
	return p2
}
func (pos Pos) Max(b Pos) Pos {
	p2 := pos
	if p2.x < b.x {
		p2.x = b.x
	}
	if p2.y < b.y {
		p2.y = b.y
	}
	return p2
}

type InputStep struct {
	dirOne   Dir
	stepsOne int
	dirTwo   Dir
	stepsTwo int
	// color int
}

type Input struct {
	steps []InputStep
	cells map[Pos]int
}

var moveDir = [5]Pos{}

func init() {
	moveDir[NORTH] = Pos{0, -1}
	moveDir[EAST] = Pos{1, 0}
	moveDir[SOUTH] = Pos{0, 1}
	moveDir[WEST] = Pos{-1, 0}
}

func (pos Pos) Move(dir Dir, dist int) Pos {
	if dir == INVALID {
		panic("invalid dir")
	}
	return pos.Add(Pos{x: dist * moveDir[dir].x, y: dist * moveDir[dir].y})
}

type Visited map[Pos]bool

func Dump(minBound, maxBound Pos, visited Visited) {
	fmt.Println("=============", minBound, maxBound)

	for y := minBound.y; y < maxBound.y; y += 1 {
		for x := minBound.x; x < maxBound.x; x += 1 {
			pos := Pos{x: x, y: y}
			ch := '.'
			if _, found := visited[pos]; found {
				ch = '#'
			}
			fmt.Printf("%c", ch)
		}
		fmt.Printf("\n")
	}
}

func bounds(visited Visited) (Pos, Pos) {
	minBound := Pos{math.MaxInt64, math.MaxInt64}
	maxBound := Pos{math.MinInt64, math.MinInt64}

	for pos := range visited {
		minBound = minBound.Min(pos)
		maxBound = maxBound.Max(pos)
	}

	maxBound.x += 1
	maxBound.y += 1

	return minBound, maxBound
}

func fill(start Pos, visited Visited) {
	queue := []Pos{start}

	for len(queue) != 0 {
		pos := queue[0]
		queue = queue[1:]
		if _, found := visited[pos]; found {
			continue
		}

		visited[pos] = true
		queue = append(queue,
			pos.Move(NORTH, 1),
			pos.Move(EAST, 1),
			pos.Move(SOUTH, 1),
			pos.Move(WEST, 1),
		)
	}
}

func Abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

// https://en.wikipedia.org/wiki/Shoelace_formula
func shoelaceArea(points []Pos) int {
	prev := points[len(points)-1]
	sum := 0

	length := 0
	for _, point := range points {
		sum += (prev.x * point.y) - (prev.y * point.x)
		length += Abs(prev.x-point.x) + Abs(prev.y-point.y)
		prev = point
	}

	return (sum / 2) + length/2 + 1
}

func solveOne(input Input) int {
	visited := Visited{}
	pos := Pos{}

	points := []Pos{}
	for _, row := range input.steps {
		for i := 0; i < row.stepsOne; i++ {
			visited[pos] = true
			pos = pos.Move(row.dirOne, 1)
		}
		points = append(points, pos)
	}

	// minPos, maxPos := bounds(visited)

	// Dump(minPos, maxPos, visited)

	// Technically we should find an "enclosed point"
	fill(Pos{1, 1}, visited)
	a := shoelaceArea(points)

	if a != len(visited) {
		fmt.Println("FAIL")
	}

	// Dump(minPos, maxPos, visited)

	return len(visited)
}

func solveTwo(input Input) int {
	pos := Pos{}

	points := []Pos{}
	for _, row := range input.steps {
		pos = pos.Move(row.dirTwo, row.stepsTwo)
		points = append(points, pos)
	}

	return shoelaceArea(points)
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		cells: map[Pos]int{},
	}

	if len(lines) == 0 {
		return input, fmt.Errorf("bad input")
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			return input, fmt.Errorf("not enough parts")
		}
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			return input, err
		}
		row := InputStep{
			stepsOne: val,
		}
		switch parts[0] {
		case "U":
			row.dirOne = NORTH
		case "D":
			row.dirOne = SOUTH
		case "L":
			row.dirOne = WEST
		case "R":
			row.dirOne = EAST
		}

		cval, err := strconv.ParseInt(parts[2][2:len(parts[2])-1], 16, 64)
		if err != nil {
			return input, err
		}
		switch cval & 0xf {
		case 0:
			row.dirTwo = EAST
		case 1:
			row.dirTwo = SOUTH
		case 2:
			row.dirTwo = WEST
		case 3:
			row.dirTwo = NORTH
		}
		row.stepsTwo = int(cval / 16)

		input.steps = append(input.steps, row)
	}

	return input, nil
}

func PartOneSolution(input Input) (int, error) {
	// 902
	return solveOne(input), nil
}

func PartTwoSolution(input Input) (int, error) {
	// 1073
	return solveTwo(input), nil
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
