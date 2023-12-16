package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	value, err := PartOneSolution(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1: ", value)

	values, err := PartTwoSolution(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2: ", values)
}

type Pos struct {
	x, y int
}

func (a Pos) Add(b Pos) Pos {
	return Pos{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

type Grid map[Pos]rune

type Pairs [][2]Pos

type Board struct {
	orig       Grid
	maxX, maxY int
	usedX      map[int]bool
	usedY      map[int]bool
}

func ParseInput(lines []string, jokers bool) (Board, error) {
	result := Board{
		orig:  Grid{},
		usedX: map[int]bool{},
		usedY: map[int]bool{},
	}

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	result.maxX = len(lines[0])
	for y, line := range lines {
		for x, ch := range line {
			pos := Pos{x, y}
			if ch == '#' {
				result.usedX[x] = true
				result.usedY[y] = true
				result.orig[pos] = ch
			}
		}
		result.maxY += 1
	}

	return result, nil
}

func Grow(board Board, scale int) Grid {
	growX := make([]int, board.maxX)
	growY := make([]int, board.maxY)

	offset := 0
	for pos := 0; pos < board.maxX; pos++ {
		if !board.usedX[pos] {
			offset += scale
		}
		growX[pos] = offset
	}
	offset = 0
	for pos := 0; pos < board.maxY; pos++ {
		if !board.usedY[pos] {
			offset += scale
		}
		growY[pos] = offset
	}

	grid := Grid{}
	for pos, ch := range board.orig {
		grow := Pos{growX[pos.x], growY[pos.y]}
		grid[pos.Add(grow)] = ch
	}

	return grid
}

func computePairs(grid Grid) Pairs {
	pairs := Pairs{}
	starts := []Pos{}
	for pos := range grid {
		starts = append(starts, pos)
	}

	for pos, start := range starts {
		for _, cur := range starts[pos+1:] {
			pairs = append(pairs, [2]Pos{start, cur})
		}
	}

	return pairs
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func computeSum(pairs Pairs) int {
	sum := 0
	for _, pair := range pairs {
		sum += abs(pair[0].x - pair[1].x)
		sum += abs(pair[0].y - pair[1].y)
	}
	return sum
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	grid := Grow(input, 1)
	pairs := computePairs(grid)
	sum := computeSum(pairs)

	// expect 9724940
	return sum, nil
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	grid := Grow(input, 1_000_000-1)
	pairs := computePairs(grid)
	sum := computeSum(pairs)

	return sum, nil
}
