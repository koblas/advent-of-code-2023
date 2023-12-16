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
	grids []Grid
}

func ParseInput(lines []string, jokers bool) (Board, error) {
	result := Board{}

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	grid := Grid{}
	for _, line := range lines {
		if line == "" {
			result.grids = append(result.grids, grid)
			grid = Grid{}
			continue
		}

		grid = append(grid, line)
	}
	result.grids = append(result.grids, grid)

	return result, nil
}

func Dump(grid Grid) {
	fmt.Println("====")

	for _, line := range grid {
		fmt.Println(line)
	}
}

func transpose(grid Grid) Grid {
	// Dump(grid)

	lines := make([][]rune, len(grid[0]))
	for _, line := range grid {
		for idx, ch := range line {
			lines[idx] = append(lines[idx], ch)
		}
	}

	var output []string
	for _, row := range lines {
		output = append(output, string(row))
	}

	// Dump(output)

	return output
}

func checkMirror(row int, grid Grid) bool {
	maxrow := len(grid)
	top := row
	bottom := row + 1
	for top >= 0 && bottom < maxrow {
		if grid[top] != grid[bottom] {
			return false
		}
		top -= 1
		bottom += 1
	}
	return true
}

func matchGrid(grid Grid) *int {
	// Dump(grid)
	rows := len(grid)
	for row := 0; row < rows-1; row += 1 {
		if checkMirror(row, grid) {
			val := row + 1
			return &val
		}
	}
	return nil
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, grid := range input.grids {
		v1 := matchGrid(grid)
		if v1 != nil {
			sum += *v1 * 100
		} else {
			v2 := matchGrid(transpose(grid))
			if v2 != nil {
				sum += *v2
			} else {
				panic("FATAL")
			}
		}
	}

	// correct: 28651

	return sum, nil
}

func diffCountMirror(row int, grid Grid) [][]int {
	maxrow := len(grid)
	top := row
	bottom := row + 1

	pos := [][]int{}
	for top >= 0 && bottom < maxrow {
		for idx := 0; idx < len(grid[top]); idx += 1 {
			if grid[top][idx] != grid[bottom][idx] {
				pos = append(pos, []int{top, idx})
			}
		}
		top -= 1
		bottom += 1
	}
	return pos
}

func countMirror(grid Grid) [][][]int {
	result := [][][]int{}
	rows := len(grid)
	for row := 0; row < rows-1; row += 1 {
		cnt := diffCountMirror(row, grid)
		result = append(result, cnt)
	}
	return result
}

func runPartTwo(grid Grid) *int {
	cnts := countMirror(grid)

	for row, diffs := range cnts {
		if len(diffs) != 1 {
			continue
		}

		dup := append(Grid{}, grid...)
		diff := diffs[0]

		top := diff[0]
		offset := diff[1]

		chars := []rune(grid[top])
		if chars[offset] == '.' {
			chars[offset] = '#'
		} else {
			chars[offset] = '.'
		}
		dup[top] = string(chars)

		if checkMirror(row, dup) {
			v := row + 1
			return &v
		}
	}

	return nil
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum := 0

	for _, grid := range input.grids {
		v1 := runPartTwo(grid)
		if v1 != nil {
			sum += *v1 * 100
		} else {
			v2 := runPartTwo(transpose(grid))
			if v2 != nil {
				sum += *v2
			} else {
				panic("FATAL")
			}
		}
	}

	// sum := findMirror(input.grids[0])
	// sum += findMirror(transpose(input.grids[1])) * 100

	return sum, nil
}
