package main

import (
	"bufio"
	"errors"
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

type Pos [2]int

type Grid map[Pos][]Pos
type Visited map[Pos]struct{}

type Board struct {
	start Pos
	grid  Grid
	orig  map[Pos]rune
	maxX  int
	maxY  int
}

var moveDir = map[rune][]Pos{
	'|': {{+0, +1}, {+0, -1}},
	'-': {{+1, +0}, {-1, +0}},
	'L': {{+0, -1}, {+1, +0}},
	'J': {{+0, -1}, {-1, +0}},
	'F': {{+0, +1}, {+1, +0}},
	'7': {{+0, +1}, {-1, +0}},
}

func (a Pos) Eq(b Pos) bool {
	return a[0] == b[0] && a[1] == b[1]
}
func (a Pos) Add(b Pos) Pos {
	return Pos{a[0] + b[0], a[1] + b[1]}
}
func (a Pos) EqL(b []Pos) bool {
	for _, item := range b {
		if a.Eq(item) {
			return true
		}
	}
	return false
}

func ParseInput(lines []string, jokers bool) (Board, error) {
	result := Board{
		grid: map[Pos][]Pos{},
		orig: map[Pos]rune{},
	}

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	result.maxX = len(lines[0])
	for y, line := range lines {
		for x, ch := range line {
			pos := Pos{x, y}
			result.orig[pos] = ch
			var move []Pos
			if ch == 'S' {
				result.start = pos
			}
			if dirs, found := moveDir[ch]; found {
				for _, dir := range dirs {
					move = append(move, pos.Add(dir))
				}
			}
			if move != nil {
				result.grid[pos] = move
			}
		}
		result.maxY += 1
	}

	// re-insert the start bar
	dirs := []Pos{}
	for dx := -1; dx <= 1; dx += 1 {
		for dy := -1; dy <= 1; dy += 1 {
			delta := Pos{dx, dy}
			check := result.start.Add(delta)
			for _, g := range result.grid[check] {
				if g[0] == result.start[0] && g[1] == result.start[1] {
					dirs = append(dirs, delta)
				}
			}
		}
	}
	if len(dirs) != 2 {
		return result, errors.New("Invalid start")
	}

	fmt.Println("START MOVES", dirs)
	for ch, check := range moveDir {
		if dirs[0].EqL(check) && dirs[1].EqL(check) {
			result.orig[result.start] = ch
			result.grid[result.start] = []Pos{
				result.start.Add(dirs[0]),
				result.start.Add(dirs[1]),
			}
			// fmt.Println("INSERTING ", string(ch))
			break
		}
	}

	return result, nil
}

func runSolutionOne(input Board) (int, Visited) {
	seen := Visited{
		input.start: struct{}{},
	}
	tokens := input.grid[input.start]

	steps := 0
	for len(tokens) != 0 {
		// fmt.Println(tokens)
		newToken := []Pos{}
		for _, pos := range tokens {
			for _, move := range input.grid[pos] {
				if _, found := seen[move]; !found {
					newToken = append(newToken, move)
				}
			}
		}
		tokens = newToken
		steps += 1
		for _, token := range tokens {
			seen[token] = struct{}{}
		}
	}

	return steps, seen
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum, _ := runSolutionOne(input)

	return sum, nil
}

func runSolutionTwo(input Board) int {
	_, seen := runSolutionOne(input)

	check := []Pos{}
	for y := 0; y < input.maxY; y++ {
		for x := 0; x < input.maxX; x++ {
			if _, found := seen[Pos{x, y}]; !found {
				check = append(check, Pos{x, y})
			}

		}
	}

	inside := map[Pos]struct{}{}

	count := 0
	for _, check := range check {
		cross := 0
		y := check[1]
		prev := '.'
		for x := check[0] + 1; x < input.maxX; x += 1 {
			cpos := input.orig[Pos{x, y}]
			if cpos == '-' {
				continue
			}
			if _, found := seen[Pos{x, y}]; !found {
				continue
			}
			if cpos == '|' {
				cross += 1
				prev = '.'
			} else if prev == 'F' && cpos == 'J' {
				cross += 1
				prev = '.'
			} else if prev == 'F' && cpos == '7' {
				prev = '.'
			} else if prev == 'L' && cpos == '7' {
				cross += 1
				prev = '.'
			} else if prev == 'L' && cpos == 'J' {
				prev = '.'
			} else {
				prev = cpos
			}
		}
		// fmt.Println("CROSS", cross)
		if cross%2 == 1 {
			inside[check] = struct{}{}
			// fmt.Println(check)
			count += 1
		}
	}

	//
	for y := 0; y < input.maxY; y++ {
		for x := 0; x < input.maxX; x++ {
			pos := Pos{x, y}
			if _, found := seen[pos]; found {
				fmt.Printf("*")
			} else if _, found := inside[pos]; found {
				fmt.Printf("I")
			} else {
				fmt.Printf("o")
			}

		}
		fmt.Printf("\n")
	}
	return count
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum := runSolutionTwo(input)

	return sum, nil
}
