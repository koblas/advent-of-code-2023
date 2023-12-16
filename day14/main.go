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

type Grid []string

type Board struct {
	grid Grid
}

func ParseInput(lines []string, jokers bool) (Board, error) {
	result := Board{}

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	result.grid = Grid{}
	for _, line := range lines {
		result.grid = append(result.grid, line)
	}

	result.grid = rotateCCW(result.grid)

	return result, nil
}

func Dump(grid Grid) {
	fmt.Println("====")

	for _, line := range grid {
		fmt.Println(line)
	}
}

func rotateCW(grid Grid) Grid {
	lines := make([][]rune, len(grid[0]))

	for idx := range grid {
		line := grid[len(grid)-idx-1]
		for idx, ch := range line {
			// y := len(grid) - idx - 1
			y := idx
			lines[y] = append(lines[y], ch)
		}
	}

	var output []string
	for _, row := range lines {
		output = append(output, string(row))
	}

	// Dump(output)

	return output
}

func rotateCCW(grid Grid) Grid {
	// Dump(grid)

	lines := make([][]rune, len(grid[0]))
	for _, line := range grid {
		for idx, ch := range line {
			y := len(grid) - idx - 1
			lines[y] = append(lines[y], ch)
		}
	}

	var output []string
	for _, row := range lines {
		output = append(output, string(row))
	}

	// Dump(output)

	return output
}

func moveRocks(input Grid) Grid {
	output := Grid{}
	for _, line := range input {
		row := make([]rune, len(line))
		rollTo := 0
		for idx, ch := range line {
			switch ch {
			case 'O':
				row[idx] = '.'
				row[rollTo] = 'O'
				rollTo = rollTo + 1
			case '#':
				row[idx] = '#'
				rollTo = idx + 1
			case '.':
				row[idx] = '.'
			}
		}
		output = append(output, string(row))
	}

	return output
}

func weight(grid Grid) int {
	sum := 0
	for _, line := range grid {
		for idx, ch := range line {
			if ch == 'O' {
				sum += len(line) - idx
			}
		}
	}
	return sum
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	result := moveRocks(input.grid)
	sum := weight(result)

	// correct: 106186

	return sum, nil
}

func runPartTwo(grid Grid) int {
	result := grid
	weights := []int{
		weight(result), // cycles start at 1, just a placeholder
	}
	for idx := 0; idx < 1000; idx += 1 {
		for i := 0; i < 4; i++ {
			result = moveRocks(result)
			result = rotateCW(result)
		}
		weights = append(weights, weight(result))
	}

	// fmt.Println(weights)

	rangeEqual := func(a, b []int) bool {
		if len(a) != len(b) {
			panic("BAD")
		}
		for idx, v := range a {
			if b[idx] != v {
				return false
			}
		}
		return true
	}

	size := 200

	// fmt.Println(weights)

	findRepeat := func() (int, int) {
		for idx := 0; idx < 2000; idx += 1 {
			for check := 2; check < 400; check += 1 {
				if rangeEqual(weights[idx:idx+size], weights[idx+check:idx+check+size]) {
					return idx, check
				}
			}
		}
		return -1, -1
	}

	offset, stride := findRepeat()

	// fmt.Println("OFF, STRIDE", offset, stride)

	pos := offset + ((1_000_000_000 - offset) % (stride))

	// fmt.Println("POS = ", pos, weights[offset], weights[offset+(1000%(stride+1))])

	// fmt.Println("REPEAT ", weights[offset:offset+stride])
	// fmt.Println("RANGE  ", weights[pos:pos+stride])

	return weights[pos]
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum := runPartTwo(input.grid)

	// 97993 too low
	// 106390 is good

	return sum, nil
}
